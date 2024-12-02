def parte1(lines):
	cont = 0

	for i in enumerate(lines) :
		stringa = "Game " + str(i[0] + 1) + ": "
		line = str(i[1])
		line = line.replace(stringa, "")
		comb = line.split(";")
		
		flag = True
		for element in comb:
			pesca = element.split(",")
			for e in pesca:
				lista = e.strip().split(" ")
				if "red" in lista[1]: 
					if int(int(lista[0]) > 12):
						flag = False 
						break
				if "green" in lista[1]:
					if int(int(lista[0]) > 13):
						flag = False 
						break
				if "blue" in lista[1]:
					if int(int(lista[0]) > 14):
						flag = False 
						break
			if not flag:
				break
		if flag:
			cont += i[0] + 1
		
	print(cont)

def parte2(lines):
    potenze = []
	
    for i in enumerate(lines): 
        verde, rosso, blu = [], [], []
        stringa = "Game " + str(i[0] + 1) + ": "
        line = str(i[1])
        line = line.replace(stringa, "")
        comb = line.split(";")
        for element in comb:
            pesca = element.split(",")
            for e in pesca:
                lista = e.strip().split(" ")
                if "red" in lista[1]: 
                    rosso.append(int(lista[0]))
                if "green" in lista[1]:
                    verde.append(int(lista[0]))
                if "blue" in lista[1]:
                    blu.append(int(lista[0]))
        rosso.sort()
        verde.sort()
        blu.sort()

        potenze.append(rosso[len(rosso)-1] * verde[len(verde)-1] * blu[len(blu)- 1])
    som = 0
    for n in potenze:
        som += n
    print(som)

lines = open("input.txt", "r")
parte1(lines)
parte2(lines)