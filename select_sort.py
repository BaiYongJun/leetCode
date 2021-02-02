A = [64, 25, 12, 22, 11]

for i in range(len(A)):
    mid_index = i
    for j in range(i+1, len(A)):
        if A[mid_index] > A[j]:
            mid_index = j
    A[i], A[mid_index] = A[mid_index], A[i]

print(A)
