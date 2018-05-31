"""
Group of characters which are enclosed by two parenthesis
Tuples are immutable objects.so we cannot alter elements which are inside the tuple.
"""

# count
"""
used to find repeating or occurrences of particular element in the tuple
takes one parameter
"""
a = (10, 0, 3, 0)
print a.count(0)

# index
"""
used to get index position of particular element
"""
print a.index(0, 2)
