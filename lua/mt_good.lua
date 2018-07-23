local _M  = { _VERSION = "7.7.7" }


local mt    = { __index = _M }


_M.global = 100
_M.global1 = {name=2,test=3}


_M.new = function(self,o)
    o = o or {}
    setmetatable(o,self)
    self.__index = self
    return o
end


_M.getId = function(self)
    return self.id
end


_M.setId = function(self,idtmp)
    self.id = idtmp
end

_M.setG = function(self,g)
    self.global = g
    -- body
end

_M.setG1 = function(self,k,g)
    self.global1[k] = g
    -- body
end

return _M