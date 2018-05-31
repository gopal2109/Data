"""
Group of characters which are enclosed by two square brackets i.e []
lists are mutable objects.We can change elements inside list
"""
a = [1, 2, 3, 4, 5, 1, 2, 8]
a[0] = 11
print a

# append
"""
used to insert an element at the end of list
takes one parameter
"""
a.append([0])
print a

# extend
"""
used to insert element at the end of list
"""
a.extend([12, 14])
print a

# count
"""
get repeating elements in the list
"""
print a.count(2)

# index
"""
get index position
"""
print a.index(12)

# insert
"""
used to insert element any where in the list
takes two parameters
"""
a.insert(1, 55)
print a

# pop
"""
used remove last element from the list
"""
print a.pop()

# remove
"""
used to remove any element from the list
"""
a.remove(2)
print a

# reverse
"""
used to reverse the element in the list
"""
a.reverse()
print a

# sort
"""
used to sort the elements in the list
"""
a.sort()
print a