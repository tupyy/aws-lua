local twt = require("twt")

local o = { test = "test" }
local tag, err = twt.create(o)
print(tag)

local verr = twt.verify(o, tag)
if verr == nil then
    print("token valid")
else
    print(verr)
end
