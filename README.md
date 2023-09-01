## aws-lua

This small poc creates aws resources from lua.

For example, the following script list all vpc and look for a vpc with the tag `myvpc=true`. 
If such a vpc is not found, it will create one:
```lua
local aws = require("aws")
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
```

### Use
```shell
make build
bin/aws-lua -f path_to_lua_script --aws-access-key <access-key> --aws-secret-key <secret-key> --aws-region <aws-region>
```
### Current supported AWS API:

- DescribeAvailabilityZones
- CreateVpc
- DescribeVpcs
- CreateSubnet
- DescribeSubnets
- CreateUser
- ListUsers
- CreateAccessKey
- ListAccessKeys
