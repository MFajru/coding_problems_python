def primNumbers(num):
    count = 0
    if num == 1:
        return "1 is not prime number"
    elif num > 1:
        for i in range(1, num + 1):
            if num % i == 0:
                count += 1
                if count > 2:
                    return str(num) + " is not a prime number"
        return str(num) + " is a prime number"
    else:
        return "Invalid input"


num = int(input("Input number: "))
print(primNumbers(num))
