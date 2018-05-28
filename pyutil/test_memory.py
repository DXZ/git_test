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

if __name__ == '__main__':
    test(1)