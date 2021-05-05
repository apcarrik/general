

def URLify_Brute(strg, len):
    words = []
    tmpword = ""
    for ch in strg:
        if(ch == " "):
            if tmpword:
                words.append(tmpword)
            tmpword = ""
        else:
            tmpword+=ch
    
    #Overwrite string with new string
    new_str = []
    for word in words:
        if(new_str):
            new_str+="%20"
            new_str+=word
        else:
            new_str = word
    return new_str

print(URLify_Brute("Mr John Smith    ",13))