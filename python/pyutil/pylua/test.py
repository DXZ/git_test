import lupa  
lua = lupa.LuaRuntime()  
  
LIBS = [  
    "./test.lua",  
]  
llibs = {}  
  
def get_file_name(filename):  
    import os  
    (_, tmp) = os.path.split(filename) 
    # print tmp 
    (f_name, ext) = os.path.splitext(tmp)  
    return f_name  

def load_libs():
    global LIBS, llibs  
    for lib_p in LIBS:  
        f = open(lib_p, 'r')  
        code_str = f.readlines()  
        filename = get_file_name(lib_p)  
        print filename
        llibs[filename] = lua.execute('\n'.join(code_str))  

if __name__ == '__main__':  
    load_libs()
    print llibs
    print llibs['test'].sayhi()
    print llibs['test'].callback(100, 200, 300, 400) 