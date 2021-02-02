def merge_sort(ori_list):
    n = len(ori_list)
    if n <= 1:
        return ori_list
    pivot_index = n // 2
    first_list = merge_sort(ori_list[:pivot_index])
    second_list = merge_sort(ori_list[pivot_index:])

    first_index = 0
    second_index = 0
    res = []
    while first_index < len(first_list) and second_index < len(second_list):
        if first_list[first_index] < second_list[second_index]:
            res.append(first_list[first_index])
            first_index += 1
        else:
            res.append(second_list[second_index])
            second_index += 1
    res += first_list[first_index:]
    res += second_list[second_index:]
    return res


sorted_list = merge_sort([3, 7, 5, 9, 1, 2, 7, 4])
print(sorted_list)
