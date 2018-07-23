local _M  = { _VERSION = "0.0.1" }

_M.id = 0

_M.new = function(idtmp)
    _M.id = idtmp
    return _M
    -- body
end

_M.getId = function()
    return _M.id
end


_M.setId = function(idtmp)
    _M.id = idtmp
end

return _M