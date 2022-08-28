def repeatedCharacters(string):
    count = 99
    char = ""
    for i in range (len(string)-1):
        for j in range(i+1, len(string)):
            if string[i] == string[j]:
                if count >= j-i:
                    count = j-i
                    char = string[i]
    return char

print(repeatedCharacters("Programming is important"))