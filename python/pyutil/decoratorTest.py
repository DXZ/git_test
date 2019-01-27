#!/usr/bin/env Python
#coding=utf-8

from functools import wraps
import logging
import traceback
from flask import Flask


app = Flask(__name__)


def memo(func):
    cache={}
    # print "in cache:",cache
    @wraps(func)
    def wrap(*args,**kwargs):
        if kwargs.get("c"):
            print "aaa",kwargs.get("c")
        else:
            print "none"
        print "in wrap:",cache,args,kwargs
        if args not in cache:
            cache[args]=func(*args,**kwargs)
        return cache[args]
    return wrap

def testDe(func):
    print "testDe"
    return func



def exception(func):
    @wraps(func)
    def wrap(*args,**kargs):
        result = None
        try:
            result = func(*args,**kargs)
            return result
        except Exception as e:
            msg = traceback.format_exc()
            logging.error(msg)
            return msg
            # raise e
    return wrap

@testDe
@memo
def test(a,b):
    print "a*b=",a,b
    return a*b

@testDe
@memo
def test1(a,b):
    print "a+b=",a,b
    return a+b

# @exception
@memo
def test3(a,b,**kwargs):
    print "test3"
    return a/b


@app.route("/hello")
@exception
def hello():
    return 1/0


if __name__ == '__main__':
    # print test(1,3)
    # print test(2,3)
    # print test(1,3)
    # print test1(3,3)
    # print test1(3,3)

    print test3(a=3,b=1,c="asaaa")
    print test3(3,1,c="asd")
    # print test3(3,0)
    print "end"
    # app.run(debug=False, host='0.0.0.0')