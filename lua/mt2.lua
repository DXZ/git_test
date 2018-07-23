local _M  = { _VERSION = "7.7.7" }


local mt    = { __index = _M }


_M.global = 100
_M.global1 = {name=2,test=3}


function _M.new(self,o)
    o = o or {}
    setmetatable(o,self)
    self.__index = self
    return o
end




function _M:new(o)
    o = o or {}
    setmetatable(o,self)
    self.__index = self
    return o
end


function _M.getId(self)
    return self.id
end


function _M.setId(self,idtmp)
    self.id = idtmp
end

function _M.setG(self,g)
    self.global = g
    -- body
end

function _M.setG1(self,k,g)
    self.global1[k] = g
    -- body
end

return _M