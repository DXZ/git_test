local inspect = require "conf.lua.lib.inspect"
local stablesort = require "conf.lua.lib.stablesort"
local cjson = require "cjson"
local memo = require "conf.lua.memo"
local redis = require "resty.redis"

function table.deepcopy(object)
    local lookup_table = {}
    local function _copy(object)
        if type(object) ~= "table" then
            return object
        elseif lookup_table[object] then
            return lookup_table[object]
        end
        local new_table = {}
        lookup_table[object] = new_table
        for index, value in pairs(object) do
            new_table[_copy(index)] = _copy(value)
        end
        return setmetatable(new_table, getmetatable(object))
    end
    return _copy(object)
end


local conn = redis.new()
conn.connect(conn, '10.10.213.219', '6379')
conn:select(5)


local request_method = ngx.var.request_method
local args = nil
local q_lists = nil
local oq_list = nil
local thrs = nil
if "GET" == request_method then
    args = ngx.req.get_uri_args()
    q_lists = args["questionList"]
    oq_list  = args["avoidList"]
    thrs  = args["thrs"]
elseif "POST" == request_method then
    ngx.req.read_body()
    args = cjson.decode(ngx.req.get_body_data())
    q_lists = args["questionList"]
    oq_list  = args["avoidList"]
    thrs  = args["thrs"]
end

function get_candidate_similarity(q_list, oq_list, thrs)
    local all_list = table.deepcopy(q_list)
    for k,v in pairs(oq_list)
    do
        table.insert(all_list, v)
    end
    local all_list_ret = conn:mget(unpack(all_list))

    local intersect_list = {}
    for k1,v1 in pairs(q_list) do
        for k2,v2 in pairs(oq_list) do
            if (v1 == v2) then
                intersect_list[tostring(v1)] = 1
            end
        end
    end

    local s_info = {}
    local qids = {}
    local oqids = {}
    local df_sq = {}
    local input_ques_id = nil
    for i,key in pairs(q_list)
    do
        qids[tostring(key)] = 1
        s_info[tostring(key)] = {}
        df_sq[tostring(key)] = 0
    end

    for i,key in pairs(oq_list)
    do
        oqids[tostring(key)] = 1
    end

    for i,data in pairs(all_list_ret)
    do
        repeat
            if(type(data) ~= "string")
            then
                break
            end

            local data_list = cjson.decode(data)
            if (i<=#q_list)
            then
                input_ques_id = tostring(q_list[i])
                for ques_id,item in pairs(data_list)
                do
                    repeat
                        if (qids[ques_id] == nil)
                        then
                            break
                        end
                        local ques_thrs = tonumber(item)
                        if(ques_thrs >= thrs)
                        then
                            s_info[input_ques_id][ques_id] = ques_thrs
                            if(input_ques_id > ques_id)
                            then
                                df_sq[input_ques_id] = df_sq[input_ques_id] + ques_thrs
                            end
                            if(intersect_list[input_ques_id]~=nil)
                            then
                                df_sq[input_ques_id] = df_sq[input_ques_id] + ques_thrs
                            end
                        end
                    until true
                end
            else
                local input_ques_id = tostring(oq_list[i-#q_list])
                for ques_id,item in pairs(data_list)
                do
                    repeat
                        if (qids[ques_id] == nil)
                        then
                            break
                        end

                        local ques_thrs = tonumber(item)
                        if(ques_thrs >= thrs and intersect_list[ques_id]==nil)
                        then
                            s_info[ques_id][input_ques_id] = ques_thrs
                            df_sq[ques_id] = df_sq[ques_id] + ques_thrs
                        end
                    until true
                end

            end
        until true
    end

    for k,v in pairs(intersect_list)
    do
        s_info[tostring(k)][k] = 1
        df_sq[tostring(k)] = df_sq[tostring(k)] + 1
    end


    local tmpsim = {}
    for k,v in pairs(q_list) do
        local index = #tmpsim + 1
        tmpsim[index] = {}
        tmpsim[index].qid = v
        tmpsim[index].sim = df_sq[tostring(v)]
    end
    --table.sort(tmpsim, function(a,b) return a.sim < b.sim end)
    stablesort.stable_sort(tmpsim, function(a,b) return a.sim < b.sim end)
    newSimVal = {}
    for k,v in pairs(tmpsim) do
        table.insert(newSimVal, {penal=v.sim, s_info=s_info[tostring(v.qid)], q_id=v.qid })
    end
    return newSimVal
end



function get_memo_data(q_list)
    for kk,q_item in pairs(q_list)
    do
        table.insert(result, memo.get(q_item))
    end


function get_candidate_similarity2(q_list, oq_list, thrs)
    local all_list_ret = conn:mget(q_list)

    for i,data in pairs(all_list_ret)
    do
        if(type(data) == "string")
        then
            cjson.decode(data)
        end
    end
    return {}
end

local result = {}
for k,q_list in pairs(q_lists)
do
    for kk,q_item in pairs(q_list)
    do

        if (memo.get(q_item) ~= nil)
        then
            table.insert(result, memo.get(q_item))
        else
            ret = get_candidate_similarity2(q_item, oq_list, thrs)
            memo.add(q_item,ret)
            table.insert(result,ret)
        end
    end
end
ngx.say(cjson.encode({code=200, data=result}))
