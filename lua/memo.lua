local _M = {}

local bucket_write = {}

local bucket_read = {}

function _M.set(key,value)
    bucket_write[key] = value
end

function _M.refresh()
    bucket_read = bucket_write
end

function _M.get(key)
    return bucket_read[key]
end

return _M