#!/usr/bin/env python
# -*- coding: utf-8 -*-
import json
import requests
import urllib2
import time

def spider_17():
s = requests.Session()
cookies = {
    "lupld":"1",
    "_ga":"GA1.2.744943285.1525785534",
    "voxauth":"CpTc9B7fcgFFyQkdWmSFY6/Z8g+rsRoswzerbIfCX2mK35OUNRSnngwkLFfKbq/MjrQL+IoMCwy4lzpcmaQI2w",
    "va_sess":"CpTc9B7fcgFFyQkdWmSFY6/Z8g+rsRoswzerbIfCX2mK35OUNRSnngwkLFfKbq/MjrQL+IoMCwy4lzpcmaQI2w",
    "uid":"13224717",
    "ugcxxAty":"1",
}
header = {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.117 Safari/537.36',
    'Accept': '*/*',
    'Accept-Language': 'zh-CN,zh;q=0.9,en;q=0.8',
    'Connection': 'keep-alive',
    'Accept-Encoding': 'gzip, deflate, br',
    'Host':'www.17zuoye.com',
    'Referer':'https://www.17zuoye.com/teacher/invite/activateteacher.vpage',
    'X-Requested-With':'XMLHttpRequest',
}
s.headers = header
# s.cookies = cookies
unixtime_str = str(int(round(time.time(), 3)*1000))
base_url = "https://www.17zuoye.com/teacher/mentor/unauthteacherlist.api?&pageNum=2&_=%s"%unixtime_str
result = s.get(base_url,cookies=cookies)



if __name__ == '__main__':
    spider_17()