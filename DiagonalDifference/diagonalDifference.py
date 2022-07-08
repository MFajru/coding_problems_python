sizeArr = int(input("enter size of array: "))
matrix = []
# input matrix value
for i in range(sizeArr):
    row = []
    inp = input("enter number in row " +
                str(i + 1) + " (separated by space): ")
    inp2 = inp.split(" ")
    for j in range(sizeArr):
        row.append(int(inp2[j]))
        # if (i == j):
        #     total1 += int(inp2[j])
    matrix.append(row)

diagonal1 = sum(matrix[i][i] for i in range(sizeArr))
diagonal2 = sum(matrix[i][sizeArr - i - 1] for i in range(sizeArr))

print(matrix)
print(abs(diagonal1 - diagonal2))
