class A(object):
    def __init__(self):
        self.a = 10
        self.b = 50


class B(A):
    def __init__(self):
        A.__init__(self)
        self.c = 11
        self.d = 20

print "a", B().a

print "c", B().c
