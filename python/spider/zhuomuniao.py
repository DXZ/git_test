#!/usr/bin/env python
# -*- coding: utf-8 -*-

import re



url = "https://www.zhuomuniaochacha.com/maths/ocr/oralcalcevaluate"


'''
# POST /maths/ocr/oralcalcevaluate HTTP/1.1
# Cookie: cuid=9F54B3988ED026490E0DD601180C921F%7C038014630971668
# X-Wap-Proxy-Cookie: none
# Content-Type: multipart/form-data; boundary=2fd57IjisEwMJcahPBoSw_3q4sdPH_Ct0oM
# User-Agent: Dalvik/2.1.0 (Linux; U; Android 5.1; OPPO A37m Build/LMY47I)
# Host: www.zhuomuniaochacha.com
# Connection: Keep-Alive
# Accept-Encoding: gzip
# Content-Length: 277578
'''
def get_header():
    return {
        'User-Agent': 'Dalvik/2.1.0 (Linux; U; Android 5.1; OPPO A37m Build/LMY47I)',
        'Connection': 'keep-Alive',
        'Accept-Encoding': 'gzip',
        'X-Wap-Proxy-Cookie':'none',
        'Cookie':'cuid=9F54B3988ED026490E0DD601180C921F%7C038014630971668',
        'Host':'www.zhuomuniaochacha.com',
        'Content-Length': 277578,
    }



if __name__ == '__main__':
    img_path = ''
    s = requests.Session()
    s.headers = get_header()

    my_data   = {
    'key1' : 'value1',
    'key2' : 'value2'
    }
    s = s.post(url,data = my_data)
    print s.status,'-----',s.content

