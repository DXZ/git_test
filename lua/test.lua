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
return "asadas"
