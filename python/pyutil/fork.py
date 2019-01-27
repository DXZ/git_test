import os
print "starting"
j = 1

print "doing"

import importT

def test():
    print "test"
    print "1111",os.fork()
    # from importT import 
    # import importT
    print "00000",id(importT.T),importT.T
    importT.T["1"].append("123")
    print "1111",id(importT.T),importT.T
    # print "gggg",os.getgid(),j
    print "end test"
# os.fork()
print "aaa"
test()
print j,importT.T

