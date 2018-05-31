"""
It is concise way(easy way)to create a list based on the requirement.
This is advanced methodology for map,filter,loops and conditions.
The output of list comprehension is always a list.
syntax:-
    [expression for i in variable if(condition)]
"""

print [i for i in range(0, 10) if i % 2 == 0]
print [i for i in range(0, 10) if i % 3 == 1]
