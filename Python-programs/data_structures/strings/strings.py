"""
- Group of characters which are enclosed by two single, double and triple codes
- single and double codes are same but triple code strings are used to define or declare multi-line strings

***syntax***
    variable_name  = string
"""

# 1. Capitalize
"""
It is used to make first letter has capital
It won't take any parameter
"""
a = "hello"
print a.capitalize()

# 2. Count
"""
It is used to find repeatings or occurrence of particular element in the string
It takes one parameter
"""

print a.count('l')
print a.count('l', 0, 3)

# 3. startswith & endswith
"""
It used to check whether string startswith or endswith given element or not
It will true if condition satisfy
It takes one parameter has input
"""
print a.startswith('h')
print a.startswith('H')

print a.endswith('o')
print a.endswith('l')

# 4. find
"""
it is used to get the index position of element
it will take one or two parameters
"""

print a.find('l')
print a.find('l', 3)

# 5. index
"""
used to get index position
it will take one parameter
"""
print a.index('e')

# 6. format
"""
used to replace any variable with given string
takes one or more parameters
"""
print "hello {}".format('world')

# 7. encode & decode

"""
Encode method will provide the security to the data by converting it into non human redable language.
Decoder used to convert encripted string into human redable string.
"""
print a.encode('base64')
print a.decode()

# 8. lower & upper

"""
used convert sting to lower and upper cases
"""
print a.lower()
print a.islower()
print a.upper()
print a.isupper()
print a.isalnum()
print a.isalpha()
print a.isdigit()
print a.isspace()
print a.istitle()

c = a.replace('l', 'k')
print c

# split
"""
It is used  to convert a string into list.
If we are not mentioning the element with which we want to split the string then by default  python compiler split the complete string  by using space.
It will take one parameter.
"""

print c.split()
print c.splitlines()
print c.rsplit()

# strip

d = " Hai "
print d.strip()
print d.rstrip()
print d.lstrip()

print d.join(a)
