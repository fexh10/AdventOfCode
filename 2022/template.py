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

if __name__ == "__main__":
    start_time = time.time()
    main()
    print("--- {} ---".format(approx(time.time() - start_time)))