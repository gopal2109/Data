a = [1, 3, 5, 6, 7, 8]
for i in range(0, len(a)):
    if a[i] != a[i-1]+1:
        print a[i-1]+1, "is missing"
