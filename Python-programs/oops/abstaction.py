from abc import ABCMeta, abstractmethod

class Animal(object):
    __metaclass__ = ABCMeta

    @abstractmethod
    def say_someting(self):
        return "I am an animal"


class Cat(Animal):

    def say_someting(self):
        s = super(Cat, self).say_someting()

        return "{} {}".format(s, "maiuu")

if __name__ == "__main__":
    c = Cat()
    print c.say_someting()