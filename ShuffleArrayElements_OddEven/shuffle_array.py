def shuffleArray(inputString):
    arr = inputString.split(" ")
    arrInt = [int(i) for i in arr]
    arrInt.sort()
    arrEven = []
    arrOdd = []
    for i in arrInt:
        if i % 2 != 0:
            arrOdd.append(i)
        elif i % 2 == 0:
            arrEven.append(i)

    arrResult = arrOdd + arrEven
    for i in arrResult:
        print(str(i), end=" ")


inputString = "0 0 0 0 0 1 1 1 1"
shuffleArray(inputString)
