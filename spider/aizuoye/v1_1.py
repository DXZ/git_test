#!/usr/bin/env python
# -*- coding: utf-8 -*-
import json
import requests
import urllib2
import time
from io import BytesIO
# import MysqlDB

# def get_image(image_url):
#     # r = requests.get(image_url)
#     r = urllib2.urlopen(image_url)
#     # f = open('./1.jpg', 'rb')
#     r.seek(0)
#     b = BytesIO(r.read())
#     return b

def get_ocr_result_aizuoye(file):
    url = "http://api.zuoye.ai/v1_1/homework/submit"

    s = requests.Session()
    my_data   = {
    '__user_id' : 4516608,
    '__version' : '1.0',
    '__access_token' : '13134303895ae1e1c4beb088.93966333',
    '__device_type' : 1,
    'batch' : 0,
    '__device_id' : 'b692e78c-654c-3a74-ab44-4a6d25e677c1',
    'demo' : 0,
    }

    files = {'img': file}
    r = s.post(url,data=my_data,files = files)
    # print r.status_code,r.content
    if r.status_code != 200:
        print "error post!!!"
        return "error post"
    # data = json.loads(r.content)
#     print data
#     '''
# # {"result":1,"response":{"valid_areas":[{"x":0,"y":319,"width":571,"height":775}],"topics":[{"area":{"x":451,"y":329,"width":2,"height":151},"origin_text":"","topic_valid":true,"confidence":1,"neat_text":"","type":0,"question_text":"","empty":false,"answers":[{"text":"2","correct_text":"2","correct":true}],"feedback":false},{"area":{"x":506,"y":554,"width":2,"height":201},"origin_text":"","topic_valid":true,"confidence":1,"neat_text":"","type":0,"question_text":"","empty":false,"answers":[{"text":"4","correct_text":"4","correct":true}],"feedback":false},{"area":{"x":43,"y":768,"width":513,"height":261},"origin_text":"","topic_valid":true,"confidence":1,"neat_text":"3+5=$7$","type":0,"question_text":"","empty":false,"answers":[{"text":"7","correct_text":"8","correct":false}],"feedback":false}],"history_id":193891055,"history_batch_id":-1}}
# # 0
#     '''
#     if data['result'] != 1:
#         print "result bad!"
#         return

    return r.content


if __name__ == '__main__':
    # product_conn=MySQLdb.connect(host='10.10.48.120',user='lyj',passwd='lyj123',port=3308,unix_socket='/tmp/mysql.sock',charset="utf8")
    # product_conn.select_db('abi')
    # product_cur = product_conn.cursor()

    # product_cur.execute("SELECT url FROM ocr_image where status=0")
    # data = product_cur.fetchall()

    # for i in data:
        # get_image(i[0])
        # time.sleep(5)

    f = get_image("https://susuanqiniu.knowbox.cn/20180314/5825b234bfb1ed134ddc2ddcb0aeb35b")
    run(f)
