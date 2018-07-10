local cjson = require "cjson"

string.split = function(s, p)
    local rt= {}
    string.gsub(s, '[^'..p..']+', function(w) table.insert(rt, w) end )
    return rt
end

local ret = {"{\"821765\": 0.300482, \"821769\": 0.327004, \"821770\": 0.352931, \"821771\": 0.333025, \"821772\": 0.372128, \"818703\": 0.304051, \"818704\": 0.337808, \"818706\": 0.348868, \"818709\": 0.489781, \"821783\": 0.419626, \"818713\": 0.352092, \"818715\": 0.301591, \"818717\": 0.489781, \"818719\": 0.370042, \"821798\": 0.341957, \"821802\": 0.422317, \"821809\": 0.326729, \"821811\": 0.369528, \"821814\": 0.301378, \"821816\": 0.35111, \"822331\": 0.471012, \"822333\": 0.37681, \"2050651\": 0.318892, \"2050655\": 0.376117, \"2050658\": 0.454341, \"2050660\": 0.428053, \"2050661\": 0.411369, \"2050663\": 0.364413, \"2050664\": 0.374699, \"2050670\": 0.439381, \"822384\": 0.352742, \"818803\": 0.317399, \"2050682\": 0.353641, \"818711\": 0.370042, \"818800\": 0.449658, \"816803\": 0.401783, \"2012838\": 0.321469, \"2012840\": 0.336448, \"2012841\": 0.331355, \"2012845\": 0.428258, \"822451\": 0.490647, \"821791\": 0.511232, \"819392\": 0.324172, \"819394\": 0.368924, \"818883\": 0.330837, \"819407\": 0.348713, \"819408\": 0.370508, \"822486\": 0.422317, \"823535\": 0.356598, \"823547\": 0.347909, \"822031\": 0.386409, \"822032\": 0.373303, \"822033\": 0.406214, \"822040\": 0.365752, \"822041\": 0.350184, \"822042\": 0.301651, \"822044\": 0.313268, \"822557\": 0.331125, \"2052894\": 0.455101, \"822047\": 0.431688, \"822048\": 0.436837, \"822562\": 0.331885, \"822052\": 0.382483, \"822054\": 0.443698, \"822568\": 0.382655, \"849201\": 0.333182, \"2011448\": 0.308845, \"870203\": 0.319582, \"821728\": 0.383062, \"890707\": 0.379384, \"2013551\": 0.358681, \"2050672\": 0.346576, \"2013567\": 0.436113, \"2013568\": 0.431493, \"2001334\": 0.359342, \"2001336\": 0.354854, \"2001338\": 0.386847, \"2001340\": 0.342506, \"821702\": 0.500231, \"821704\": 0.471456, \"821714\": 0.406214, \"2011104\": 0.43215, \"2011107\": 0.462826, \"2001381\": 0.338512, \"2011304\": 0.37426, \"821747\": 0.372544, \"821755\": 0.303448, \"821758\": 0.373773, \"821759\": 0.325472}"}
local thrs = 0.3
local input_list = string.split("3 4"," ")
local result = {}


for i,data in ipairs(ret)
do
    print(i,data)
    repeat
        if(type(data) ~= "string")
        then
            break
        end
    
    local data_list = cjson.decode(data)
    local input_ques_id = input_list[i]
    result[input_ques_id] = {}
    print (data,data_list)
    for ques_id,item in pairs(data_list)
    do
        print("1111",ques_id,item)
        local ques_thrs = tonumber(item)
        if(ques_thrs > thrs)
        then
            result[input_ques_id][ques_id] = ques_thrs
        end
    end
    until true
end
local result_json = cjson.encode(result)
print(result_json)
return result_json