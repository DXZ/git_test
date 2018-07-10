--定义元表m
m = {}
--定义元表的__index的元方法
--对任何找不到的键，都会返回"undefined"
m.__index = function ( table, key )
    if key == "z"
    then
        return "zzzzz"
    end
    return "undefined"
end   
 
--表pos
pos = {x=1, y=2}
--初始没有元表，所以没有定义找不到的行为
--因为z不在pos中，所以直接返回nil
print(pos.z) -- nil
--将pos的元表设为m
setmetatable(pos, m)
--这是虽然pos里仍然找不到z，但是因为pos有元表，
--而且元表有__index属性，所以执行其对应的元方法，返回“undefined”
print(pos.x) -- undefined

print(pos.z) -- undefined

print(pos.qqq) -- undefined