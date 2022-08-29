def checkSubstring(s, p):
    s = s.split(" ")
    p = p.split(" ")
    # print(s, p)
    for i in range(len(p)):
        for j in range(len(s)):
            if p[i] == s[j]:
                s[j] = ""
    # print(s)
    a = " ".join([s[k] for k in range (len(s)-1, -1, -1) if s[k] != ""])
    # print(a)
    return a[::-1]
   

print(checkSubstring('This is java test', 'is java'))
print(checkSubstring('This is a very nice', 'is a very'))