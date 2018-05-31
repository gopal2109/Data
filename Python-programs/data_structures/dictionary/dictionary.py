"""
- Dictionaries consist group of items here each item is key, value pair
- Each item separated by comma
- Key is immutable object. we can use strings, integers, tuples has keys.
- value is mutable object we can use strings, tuples, dictionaries, again dict
- dictionary is unordered data type
- we can access values by using keys not using index values
"""

# copy
"""
used to copy one dict to another
"""
a = {'a': 10, 'b': 11}
b = a.copy()
print b

# keys
"""
used to get list of keys
"""
print b.keys()
print b.viewkeys()
print b.iterkeys()

# values
"""
used to get list of values
"""
print b.values()
print b.viewvalues()
print b.itervalues()

# items
"""
used to get list of items
"""
print b.items()
print b.viewitems()
print b.iteritems()

# has_key
"""
used to check key exist or not
"""

print b.has_key('b')

# pop & popitem

# used to read and remove first item from dict
print b.popitem()

# used to read any key value from dict and remove it

print b.pop('b')

print b.setdefault('a', 10)

print b.get('a')

b.update({'b': 25})

print b

seq = ('name', 'age')
d = b.fromkeys(seq, 10)
print str(d)

b.clear()
print b