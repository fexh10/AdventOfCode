import time

def approx(seconds: float):
    notation = ["s", "ms", "µs"]
    i = 0
    while seconds < 1:
        seconds *= 1000
        i += 1
    return str(seconds)+notation[i]

def main():
    file = open("input.txt", "r")
    lett_comuni = []
    #parte 1
    """
    for line in file:
        scomp1, scomp2 = line[:len(line) // 2], line[len(line) // 2:]
        for i in range(len(scomp1)):
            if scomp1[i] in scomp2:
                if scomp1[i].isupper():
                    lett_comuni.append(ord(scomp1[i]) - 38)
                    break
                else:
                    lett_comuni.append(ord(scomp1[i]) % 32)
                    break
    print("parte1:", sum(lett_comuni))
    """
    #parte 2
    righe = file.readlines()
    i = 0
    while(i < len(righe) - 2):
        temp = righe[i]
        for j in range(len(temp)):
            if temp[j] in righe[i + 1] and temp[j] in righe[i + 2]:
                if temp[j].isupper():
                    lett_comuni.append(ord(temp[j]) - 38)
                    i += 3
                    break
                else:
                    lett_comuni.append(ord(temp[j]) % 32) 
                    i += 3
                    break
    print("parte 2:", sum(lett_comuni))

if __name__ == "__main__":
    start_time = time.time()
    main()
    print("--- {} ---".format(approx(time.time() - start_time)))

