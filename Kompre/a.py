n = int(input("input amount of data: "))
upperRange = int(input("input upper range: "))
bottomRange = int(input("input bottom range: "))
lsUpper = []
lsMiddle = []
lsBottom = []
for i in range(n):
    a = int(input("input data: "))
    if a > upperRange:
        lsUpper.append(a)
    elif a > bottomRange:
        lsMiddle.append(a)
    else:
        lsBottom.append(a)
print("Above " + str(upperRange) + ":", lsUpper)
print("In the middle: ", lsMiddle)
print("Below " + str(bottomRange) + ":", lsBottom)
