#!/usr/bin/env python
# -*- coding: utf-8 -*-
from aizuoye.v1_1 import get_ocr_result_aizuoye
from xiaoyuan.v1_1 import get_ocr_result_xiaoyuan

if __name__ == '__main__':
    get_ocr_result_xiaoyuan(open('./aizuoye/1.jpg','rb'))