local memo = require "conf.lua.memo"
local cjson = require "cjson"


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



local result = {}
for k,q_list in pairs(q_lists)
do
    for kk,q_item in pairs(q_list)
    do
        table.insert(result, memo.get(q_item))
    end
end
ngx.say(cjson.encode({code=200, data=result}))
