# -*- coding: UTF-8 -*-

lis = [10, 8, 4, 7, 5]
for i in range(len(lis)-1):
    for j in range(len(lis)-1-i):
        if lis[j] > lis[j+1]:
            lis[j], lis[j+1] = lis[j+1], lis[j]
print(lis)
