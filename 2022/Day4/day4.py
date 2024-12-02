import time

def approx(seconds: float):
    notation = ["s", "ms", "µs"]
    i = 0
    while seconds < 1:
        seconds *= 1000
        i += 1
    return str(seconds)+notation[i]

def main():
    #parte 1
    file = open("input.txt").read().strip()
    lines = [x.strip() for x in file.split('\n')]
    
    cont = 0
    cont2 = 0
    for line in lines:
        one,two = line.split(',')
        s1, e1= one.split('-')
        s2, e2= two.split('-')
        s1,e1,s2,e2 = [int(x) for x in [s1,e1,s2,e2]]
        if s1<=s2 and e2<=e1 or s2<=s1 and e1<=e2:
            cont += 1
        #parte2
        if not (e1 < s2 or s1 > e2):
            cont2 += 1
    print(cont2)
        
        



if __name__ == "__main__":
    start_time = time.time()
    main()
    print("--- {} ---".format(approx(time.time() - start_time)))