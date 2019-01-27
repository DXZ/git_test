#!/usr/bin/env python
# -*- coding: utf-8 -*-
import json
import requests
import urllib2
import time

def get_ocr_result_xiaoyuan(file):
    proxy = {"https":"http://localhost:8888"}
    url = "https://xyks.yuanfudao.com/leo-oral/android/oral-evaluations?_productId=611&platform=android22&version=1.0.0&vendor=qihu&av=5&sign=9f3f7f9534ac20835c40fd39d04eeeea"
    s = requests.Session()
    cookies = {
    'ks_sess':'Zx2E8whArMI1VZhaukc56Uv2afVWiSy+rF3J3aux8h8WPqVjtzrMi9ZN83LVJDUs',
    'ks_persistent':'7Vxly03YT406MZbSMDqqWz71OFBdSrkKrPBYR25/pxV7TfaTrRT9SlGP7OSz4fcTKqZaLEuzx41Y44lk+lMy9c9nmHVatpReGq35N17G5Hc=',
    'ks_deviceid':'565943',
    }

    header = {
        'Connection' : 'Keep-Alive',
        'User-Agent': 'Leo/1.0.0 (OPPOOPPO A37m; Android 5.1; Scale/2.00)',
        'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
        'Accept-Language': 'en-US,en;q=0.5',
        'Connection': 'keep-alive',
        'Accept-Encoding': 'gzip',
        'Host':'xyks.yuanfudao.com',
    }
    s.headers = header
    files = {'image': file}
    # r = s.post(url,files = files, verify="./ca/gdroot-g2.crt")
    # r = s.post(url,files = files, proxies=proxy)
    r = s.post(url,files=files, cookies=cookies)
    
    print r.text,r.status_code,r.content
    if r.status_code != 200:
        print "error post!!!"
        return "error post"
    return r.content

