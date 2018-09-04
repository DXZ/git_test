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




function split(str, split_char)      
    local sub_str_tab = {}
    while true do          
        local pos = string.find(str, split_char) 
        if not pos then              
            table.insert(sub_str_tab,str)
            break
        end  
        local sub_str = string.sub(str, 1, pos - 1)              
        table.insert(sub_str_tab,sub_str)
        str = string.sub(str, pos + 1, string.len(str))
    end      
    return sub_str_tab
end

local c = split("asd___asd123_asd___12312","___")
print("--A-----")
for i =1 ,#c do
    print(c[i])
end

--在LUA中，没有类似python的str.split()函数，如果要将一个字符串按照指定的符号分割，并存到table中。
--如将a="1,2,3"转换为{1,2,3}

a="1,2,3,4,"    --或者 a="1,2,3,4" 不论最后一个是否有","
t={}

for w in string.gmatch("asd___asd123m_asd___1231,2","([^_]+)") do     --按照“,”分割字符串
    table.insert(t,w) 
end

for k,v in ipairs(t) do 
    print(k..":"..v) 
end


-- table.cancat()



local a = {}
a.bbb = "123"
print(a.bbb,a.aa)
-- function make_array(list,max_length)
--     local result = {}
--     local count = 1
--     for i=1,max_length do
--         if list[i] then
--             result[count] = list[i]
--             count = count + 1
--         end
--     end
--     return result
-- end

-- local b = make_array(a,5)

-- print(cjson.encode(b))

-- local MithOrder = {}

-- MithOrder["数感"] = 1
-- MithOrder["符号意识"] = 2
-- MithOrder["空间观念"] = 3
-- MithOrder["几何直观"] = 4
-- MithOrder["数据分析观念"] = 5
-- MithOrder["运算能力"] = 6
-- MithOrder["推理能力"] = 7
-- MithOrder["模型思想"] = 8
-- MithOrder["应用意识"] = 9
-- MithOrder["创新意识"] = 10

-- print(#MithOrder)





-- local config = {}
-- config.ProblemOrder = {}
-- config.ProblemOrder["阅读理解"] = 1
-- config.ProblemOrder["推理分析"] = 2
-- config.ProblemOrder["列式运算"] = 3
-- config.ProblemOrderMaxLength = 4


-- for k,v in pairs(config.ProblemOrder) do
--     print(1111);print(k,v);print("11111")
-- end
-- print("start")
-- function test( ... )
--     for i, v in ipairs{...} do
--         if type(v) == "table" then
--             for k,vv in pairs(v) do 
--                 print("111111",k,vv)
--             end
--         else
--             print("22222",k,v)
--         end
--     end
--     return false
-- end


-- local testaaaa = {a=1,b=2,c=3}

-- test(testaaaa,1)

-- b = {testaaaa=testaaaa}
-- print(b.testaaaa.a)
-- local a =10
-- local a= 100
-- print(a)

-- -- socket.sleep(0.2)
-- return 
-- '''