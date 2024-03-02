
def main():
    text = open('Input1.txt','r')
    f = open('result.txt','w')
    sum = 0
    for line in text:
        FirstNumber = None
        LastNumber = None
        for L in line:
            if L.isnumeric() and not FirstNumber:
                FirstNumber = L
                f.write(FirstNumber)
            if L.isnumeric() and FirstNumber:
                LastNumber = L
                f.write(LastNumber)
        result = FirstNumber+''+LastNumber
        sum+= int(result)
    print(sum)
    



main()