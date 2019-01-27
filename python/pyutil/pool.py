from multiprocessing import Pool
from time import sleep
 
# def f(x):
#     for i in range(10):
#         print '%s --- %s ' % (i, x)
#         sleep(1)
 
 
# def main():
#     pool = Pool(processes=3)    # set the processes max number 3
#     for i in range(11,20):
#         result = pool.apply_async(f, (i,))
#     pool.close()
#     pool.join()
#     if result.successful():
#         print 'successful'
 
 
# if __name__ == "__main__":
#     main()

# from multiprocessing import Pool
# import time

# pool = Pool(processes=10)

# # l =


def AAAAA(x,y):
    print("in pool_test",x,y)
    sleep(1)
    print("end pool_test",x,y)
    return x+y,x-y


def BBBBB(x,y):
    print("in BBBBB",x,y)
    sleep(1)
    print("end BBBBB",x,y)
    return x-y,x*y



def main():
    
    x = 123
    # AAAAA(1,2)

    # print eval("x")
    res = []
    pool = Pool(processes=10)
    for i in xrange(10):
        res.append(pool.apply_async(AAAAA, (i,eval("x"),)))

    # for i in xrange(10):
    #     res.append(pool.apply_async(BBBBB, (i,eval("x"),)))

    print("end apply_async",res)
    # pool.close()
    # pool.join()
    a = 1
    print("-------22222--------")
    for i in res:
        xx,yyy = i.get()
    print("---------------")
    for i in xrange(10):
        res.append(pool.apply_async(BBBBB, (i,eval("x"),)))
    for i in res:
        xx,yyy = i.get()
    pool.apply_async(AAAAA, [11,12])
    pool.close()
    pool.join()
    print eval("xx"),eval("yyy")

if __name__ == "__main__":
    
    main()
    
# [x.get() for x in [pool.apply_async(pool_test, (x,)) for x in gen_list(l)]]