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
return "asadas"
