local cjson = require "cjson"
local aaa={}

aaa["asd"]="123"
aaa["bbb"]="222"

local aaastr = cjson.encode(aaa)

local bbb = {}
bbb["1"] = aaastr
bbb["2"] = aaastr

local bbbstr = cjson.encode(bbb)
local bbbdic = cjson.decode(bbbstr)
print(type(bbbdic["1"]),bbbstr,bbb,aaastr)