if __name__ == '__main__':

    somma = []
    file = open("input.txt", "r")
    #parte 1 
    temp = 0
    for line in file:
        if line != "\n":
            temp += int(line)
        else:
            somma.append(temp)
            temp = 0
    somma.sort(reverse = True)
    print("parte 1:", somma[0])

    #parte 2
    print("parte 2", somma[0] + somma[1] + somma[2])
