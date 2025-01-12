def wordAbbrv(strArr):
    for word in strArr:
        lenWord = len(word)
        lenBetween = lenWord-2
        if (lenWord > 10):
            print(word[0]+str(lenBetween)+word[lenWord-1])
            continue
        print(word)

numOfWord = input()
words = []
for i in range(int(numOfWord)):
    word = input()
    words.append(word)

wordAbbrv(words)