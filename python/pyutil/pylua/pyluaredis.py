#!/usr/bin/env Python
#coding=utf-8

import time
import json
 
import redis
 
pool = redis.ConnectionPool(host='localhost', port=6379, db=0)
r = redis.Redis(connection_pool=pool)



lua1 = """
   redis.call("select", ARGV[1])
   return redis.call("get",KEYS[1])
"""
script1 = r.register_script(lua1)



lua2 = """
   redis.call("select", ARGV[1])
   local ret = redis.call("get",KEYS[1])
   redis.call("select", ARGV[2])
   return ret
"""

script2 = r.register_script(lua2)

# print r.get("mykey")
# print script2(keys=["mykey"], args = [0,1] )
# print r.get("mykey"), "ok"

print r.mget(["mykey","mykey1"])
print "script2",script2( keys=["mykey"], args = [1,0] )
print script2( keys=["mykey"], args=[0,1] )
print "=========="
luasplit = """
    string.split = function(s, p)
        local rt= {}
        string.gsub(s, '[^'..p..']+', function(w) table.insert(rt, w) end )
        return rt
    end

    redis.call("select",ARGV[1])
    local ret = redis.call("MGET",KEYS[1])
    local thrs = tonumber(KEYS[2])
    local input_list = string.split(KEYS[1]," ")
    local result = {}

    for i,data in ipairs(ret)
    do 
        local data_list = cjson.decode(data)
        local input_ques_id = input_list[i]
        result[input_ques_id] = {}
        for j,item in ipairs(data_list)
        do
            local list = string.split(item, '_')
            local ques_id = list[1]
            local ques_thrs = tonumber(list[2])
            if(ques_thrs > thrs)
            then
                result[input_ques_id][ques_id] = ques_thrs
            end
        end
    end
    local result_json = cjson.encode(result)
    return result_json
"""

luasplitscript = r.register_script(luasplit)
print "11"
print luasplitscript(keys=["1 2 3",0.08],args=[0,1],)
