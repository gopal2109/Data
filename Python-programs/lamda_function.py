"""
- Lambda functions are one line in code and these are not reusable so the scope of lambda functions is very very less in python.
- Lambda keywords is used to create a lambda functions.
Syntax:-
    Function_name=lambda var1,var2=var1*var2
"""

lam = lambda x: x**2 + x

print lam(2)


fs = [(lambda n, i=i: i + n) for i in range(10)]

print fs[3](4)
