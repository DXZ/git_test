from functools import wraps
from guppy import hpy

def checkout_memory(f):
    @wraps(f)
    def func(*arg,**kw):
        hp = hpy()
        result = f(*arg,**kw)
        print hp.heap()
        return result
    return func

@checkout_memory
def test(a):
    return a*a

@checkout_memory
def testB():
    import pandas as pd
    return pd.DataFrame(np.random.randint(0,100,size=(10000, 4)), columns=list('ABCD'))

if __name__ == '__main__':
    test(1)