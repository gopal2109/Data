class A(object):
    def __init__(self):
        self.age = 25
        self.name = "test"


class B(A):
    def __init__(self):
        A.__init__(self)
        self.company = "accion"
        self.address = "bangalore"


class C(B):
    def __init__(self):
        B.__init__(self)
        self.car = "volvo"


print "name : ", C().name

print "company : ", C().company