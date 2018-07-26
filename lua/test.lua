function revtab(tab)
    local revtab = {}
    for k,v in pairs(tab) do
        revtab[v] = true
    end
    return revtab
end
print("asd")

local a = {123,345}
print(table.unpack(a).."aaaa")

local i = {}
i[1]  = 1
i[3]  = 1
i[2]  = 1

local ccc = {asd={1,3,4,4}}
print("ccc=",ccc)


i[100]  = 1
local c = {}
c[i] = 1

print("len(i)=",#i,c[i],i)
-- print(table.maxn(i)) 启用了
for k,v in pairs(c)
do
    print(v)
end
i = {}
c[i] = 2
print("len(i)=",#i,c[i],i)
for k,v in pairs(c)
do
    print(v)
end

print("teeeeeeee------------")
local cjson = require "cjson"
local a = {}
a[1] = 1
a[3] = 3
a[7] = 7
print(cjson.encode(a))



function make_array(list,max_length)
    local result = {}
    local count = 1
    for i=1,max_length do
        if list[i] then
            result[count] = list[i]
            count = count + 1
        end
    end
    return result
end

local b = make_array(a,5)

print(cjson.encode(b))

local MithOrder = {}

MithOrder["数感"] = 1
MithOrder["符号意识"] = 2
MithOrder["空间观念"] = 3
MithOrder["几何直观"] = 4
MithOrder["数据分析观念"] = 5
MithOrder["运算能力"] = 6
MithOrder["推理能力"] = 7
MithOrder["模型思想"] = 8
MithOrder["应用意识"] = 9
MithOrder["创新意识"] = 10

print(#MithOrder)

return "asadas"
