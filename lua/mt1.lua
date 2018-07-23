A = {b=100}

function A.w(self,v)
    self.b = self.b - v
    -- body
end


function A.w1(v)
    A.b = A.b - v
    -- body
end


a = A
-- A = nil

a.w1(1)
a.w(a,1)
print(a.b)
return