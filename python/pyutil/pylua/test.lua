libfoo = {}  
  
function libfoo.sayhi()  
    return "hi from lupa"  
end  
  
function libfoo.callback(a, b, c, d)  
    return a * b + c - d  
end  
  
return libfoo 