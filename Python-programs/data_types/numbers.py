"""
Python supports four different numerical types
- int (signed integers) − They are often called just integers or ints, are positive or negative whole numbers with no decimal point
- long (long integers ) − Also called longs, they are integers of unlimited size, written like integers and followed by an uppercase or lowercase L.
- float (floating point real values) − Also called floats, they represent real numbers and are written with a decimal
    point dividing the integer and fractional parts. Floats may also be in scientific notation, with E or e indicating
    the power of 10 (2.5e2 = 2.5 x 102 = 250).
- complex (complex numbers) − are of the form a + bJ, where a and b are floats and J (or j) represents the square root
     of -1 (which is an imaginary number). The real part of the number is a, and the imaginary part is b. Complex
     numbers are not used much in Python programming.
"""
a = 12
b = 12.5
c = 5e+10j

print type(a)
print type(b)
print type(c)
