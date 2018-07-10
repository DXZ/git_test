local cjson = require "cjson"

string.split = function(s, p)
    local rt= {}
    string.gsub(s, '[^'..p..']+', function(w) table.insert(rt, w) end )
    return rt
end

local ret = {"[\"2013732_0.07\", \"2013733_0.07\"]","[\"1_1.07\", \"2_2.07\"]"}
local thrs = 1
local input_list = string.split("3 4"," ")
local result = {}

for i,data in ipairs(ret)
do 
    local data_list = cjson.decode(data)
    print(i,data,input_list[i])
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
result_json = cjson.encode(result)
print(result_json)
return result_json