#!/usr/bin/env python
# -*- coding: utf-8 -*-
import json
import requests
import urllib2
import time
from io import BytesIO
import MySQLdb
from aizuoye.v1_1 import get_ocr_result_aizuoye
from xiaoyuan.v1_1 import get_ocr_result_xiaoyuan
import random,datetime


def get_image(image_url):
    # r = requests.get(image_url)
    r = urllib2.urlopen(image_url)
    # f = open('./1.jpg', 'rb')
    # r.seek(0)
    b = BytesIO(r.read())
    return b


if __name__ == '__main__':
    # product_conn=MySQLdb.connect(host='10.10.48.120',user='lyj',passwd='lyj123',port=3308,unix_socket='/tmp/mysql.sock',charset="utf8")
    product_conn=MySQLdb.connect(host='127.0.0.1',user='lyj',passwd='lyj123',port=3338,unix_socket='/tmp/mysql.sock',charset="utf8")
    product_conn.select_db('abi')
    product_cur = product_conn.cursor()

    product_cur.execute("SELECT url,image_id FROM ocr_image where state=0")
    data = product_cur.fetchall()

    for i in data:
        print i
        date = datetime.datetime.now()
        f = get_image(i[0])
        xy_result = get_ocr_result_xiaoyuan(f)
        # xy_result = get_ocr_result_xiaoyuan(f)
        product_cur.execute("INSERT INTO ocr_competitor_result(`image_id`,`xiaoyuan`,`addtime`) Values(%s,%s,%s)",(i[1],xy_result,date))
        product_conn.commit()
        time.sleep(random.randint(15,60))
        print "done"

    product_cur.close()
    product_conn.close()