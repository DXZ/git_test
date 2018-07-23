local mt = require('mt')
mt.setId(0,1)
print(mt.getId)
print(mt.getId())
mt.setId(2)
print(mt.getId())



mt1 = mt.new(1)
print(mt.getId(),mt1.getId())


print("------m2------")

local mtgood = require("mt_pro")
local mtgood1 = mtgood:new()
local mtgood2 = mtgood:new()


print(mtgood1:getId(),mtgood2:getId(),mtgood1.global,mtgood2.global,mtgood1.global1[1],mtgood2.global1[2])
mtgood1:setId(100)
mtgood2:setId(200)
mtgood1:setId(100)
mtgood2:setG(2000)
mtgood1:setG(1000)
mtgood2:setG1(2,2000)
mtgood1:setG1(1,1000)
-- mtgood1:setG1(2,1000)
print(mtgood1.id,mtgood2.id,mtgood1.global,mtgood2.global,mtgood1.global1[1],mtgood2.global1[2])



print("------good------")
local mtgood = require("mt_good")
local mtgood1 = mtgood:new()
local mtgood2 = mtgood:new()


print(mtgood1:getId(),mtgood2:getId(),mtgood1.global,mtgood2.global,mtgood1.global1[1],mtgood2.global1[2])
mtgood1:setId(100)
mtgood2:setId(200)
mtgood2:setG(2000)
mtgood1:setG(1000)
mtgood2:setG1(2,2000)
mtgood1:setG1(1,1000)
-- mtgood1:setG1(2,1000)
print(mtgood1.id,mtgood2.id,mtgood1.global,mtgood2.global,mtgood1.global1[1],mtgood2.global1[2])

return