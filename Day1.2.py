import re

def main():
    text = open('Input1.txt','r')
    sum = 0
    for line in text:
        print("\n")
        print("\n")
        linetotal=findnumber(line)
        print("linetotal:",linetotal)
        sum+=linetotal
    print(findnumber("eight82cppmnpvkvthreesevenseven8h"))
    print(sum)

        
    


def findnumber(texto:str):
    print(texto)
    FirstNumberindex = 99
    LastNumberindex = -1
    FirstNumber = 0
    LastNumber=0
    elements = [0,1,2,3,4,5,6,7,8,9]
    elements2 = ["one","two","three","four","five","six","seven","eight","nine"]
    for number in elements:
        finder =re.finditer(str(number),texto)
        match = list(finder)
        print(match)
        try:
                matchpos = match[0].span()[0]
                if match != -1:
                # print("tried to find:",number,"\n Found it there in position",(match+1))
                    if matchpos < FirstNumberindex:
                        # print("OldFirstNumber:",FirstNumber,"newFirstNumber",number)
                        # print("OldFirstNumberindex:",FirstNumberindex,"newFirstNumberindex",match+1)
                        FirstNumberindex = match[0].span()[0]
                        FirstNumber = number
                    if match[-1].span()[0]> LastNumberindex:
                        # print("OldLastNumber:",LastNumber,"newLastNumber",number)
                        # print("OldLastNumberindex:",FirstNumberindex,"newLastNumberindex",match+1)
                        LastNumberindex = match[-1].span()[0]
                        LastNumber = number
        except:
            print('a')
    for count, number in enumerate(elements2, start=1):
        finder =re.finditer(str(number),texto)
        match = list(finder)
        try:
                if match != -1:
                # print("tried to find:",number,"\nFound it there in position",match+1)
                    if match[0].span()[0] < FirstNumberindex:
                        # print("OldFirstNumber:",FirstNumber,"newFirstNumber",count)
                        # print("OldFirstNumberindex:",FirstNumberindex,"newFirstNumberindex",match+1)
                        # print("\n")
                        FirstNumberindex = match[0].span()[0]
                        FirstNumber = count
                    if match[-1].span()[0] > LastNumberindex:
                        # print("OldLastNumber:",LastNumber,"newLastNumber",count)
                        # print("OldLastNumberindex:",LastNumberindex,"newLastNumberindex",match+1)
                        # print("\n")           
                        LastNumberindex = match[-1].span()[0]
                        LastNumber = count
        except:
             print('')
                
    result = str(elements[FirstNumber])+''+str(elements[LastNumber])
    print("this line number is:",result)
    return int(result)




main()