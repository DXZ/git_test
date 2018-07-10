local cjson = require "cjson"
local memo = require "conf.lua.memo"
local redis = require "resty.redis"



local conn = redis.new()
conn.connect(conn, '10.10.213.219', '6379')
conn:select(5)



function update()
    local all_list_key = conn:keys("*")
    for i=1,#all_list_key,5000 
    do 
        local all_list_ret = conn:mget(unpack(all_list_key,i,math.min(i+4999, #all_list_key)))
        for ii, v in ipairs(all_list_ret)
        do
            local key = all_list_key[i+ii-1];
            ngx.say(cjson.encode({keys=key,value=v}))
            if(type(v) ~= "string")
            then
                break
            end
            data = cjson.decode(v)
            memo.set(k,data)
        end
    end 
end


function update2()
    local memo_cache = ngx.shared.memo_cache
    local all_list_key = conn:keys("*")
    for i=1,#all_list_key,5000
    do
        local all_list_ret = conn:mget(unpack(all_list_key,i,math.min(i+4999, #all_list_key)))
        for ii, v in ipairs(all_list_ret)
        do
            local key = all_list_key[i+ii-1];
            ngx.say(cjson.encode({keys=key,value=v}))
            if(type(v) ~= "string")
            then
                break
            end
            data = cjson.decode(v)
            memo_cache:set(key,data)
        end
    end
end

function refresh()
    memo.refresh()
end

ngx.say(cjson.encode({code=200, msg="success"}))