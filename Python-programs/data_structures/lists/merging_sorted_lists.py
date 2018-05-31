a = [3, 4, 6, 10, 11, 18]
b = [1, 5, 7, 12, 13, 19, 21]


def merge_two_sorted_lists(a, b):
    return sorted(a + b)

print merge_two_sorted_lists(a, b)
