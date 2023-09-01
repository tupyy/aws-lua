print("hello from lua")
local aws = require("aws")

local function print_table(t)
    if t == nil then
        return
    end
    for k, v in pairs(t) do
        if type(v) == "table" and v ~= nil then
            print(k .. print_table(v))
        else
            print(string.format("%s: %s", k, v))
        end
    end
end

-- get AZ
local ret, err = aws.list("aws_availability_zones", {})
if err ~= nil then
    print(err)
    os.exit(1)
end
print("AZ in region:")
for k, v in pairs(ret.AvailabilityZones) do
    print("ZoneID:" .. v.ZoneName)
end

-- Look for a VPC with the tag myvpc=true.
-- If not found create one with cidr=10.0.0.0/16

-- create vpc
local function create_vpc(cidr, tags)
    vpc = {
        cidr = cidr,
    }
    if tags ~= nil and type(tags) == "table" then
        vpc["tags"] = tags
    end
    return aws.create("aws_vpc", vpc)
end

-- Get VPCs
local vpcs, err = aws.list("aws_vpc", {})
if err ~= nil then
    print(err)
    os.exit(1)
end

local vpc = { found = false }
if vpcs ~= nil then
    for _, _vpc in pairs(vpcs.Vpcs) do
        if _vpc.Tags ~= nil then
            for _, v in pairs(_vpc.Tags) do
                if v.Key == "myvpc" and v.Value == "true" then
                    vpc.found = true
                    vpc.id = _vpc.VpcId
                end
            end
        end
    end
end

local vpc_id
if vpc.found then
    print("myvpc VPC found. ID: " .. vpc.id)
    vpc_id = vpc.id
else
    local vpc, err = create_vpc("10.0.0.0/16", { myvpc = "true" })
    if err ~= nil then
        print("vpc creation failed: " .. err)
        os.exit(1)
    end
    if vpc ~= nil then
        print("vpc created: " .. vpc.Vpc.VpcId)
        vpc_id = vpc.Vpc.VpcId
    end
end

------------
--
-- SUBNETS
--
-- ---------
-- List subnets
local ss, err = aws.list("aws_subnet", {
    filters = {
        { Name = "vpc-id", Values = { vpc_id } },
    }
}
)

local subnets = {}
function subnets:find(cidr)
    for _, v in ipairs(self) do
        if v.cidr == cidr then return true end
    end
    return false
end

if err ~= nil then
    print("listing subnets failed: " .. err)
else
    if ss.Subnets == nil then
        print("no subnets found for vpc " .. vpc_id)
    else
        for k, v in pairs(ss.Subnets) do
            print(string.format("found subnet cidr=%s, id=%s", v.CidrBlock, v.SubnetId))
            table.insert(subnets, { id = v.SubnetId, cidr = v.CidrBlock })
        end
    end
end

print("create one private and one public subnet in each az...")
print("private subnets cidr starts at 10.0.1.0/24")
print("public subnets cidr starts at 10.0.100.0/24")

local tags = {
    myvpc = "true"
}

local i = 1
for _, az in pairs(ret.AvailabilityZones) do
    -- for each az create one subnet
    for _, j in ipairs({ i, i + 100 }) do
        local next_cidr = "10.0." .. j .. ".0/24"
        if not subnets:find(next_cidr) then
            local s, err = aws.create("aws_subnet",{
                vpc_id = vpc_id,
                cidr = next_cidr,
                availability_zone_id = az.ZoneID,
                tags = tags,
            })
            if err ~= nil then
                print("error creating subnet in az " .. err)
            else
                print(string.format("subnet %s created in az %s", next_cidr, az.ZoneID))
            end
        end
        i = i + 1
    end
end

-------------------
---
--- NATS
---
-------------------
