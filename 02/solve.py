# 509871

import fileinput

def eval_program(arr, noun, verb):
    arr[1] = noun
    arr[2] = verb
    i = 0
    while True:
        if arr[i] == 1:
            a = arr[i + 1]
            b = arr[i + 2]
            arr[arr[i + 3]] = arr[a] + arr[b]
        elif arr[i] == 2:
            a = arr[i + 1]
            b = arr[i + 2]
            arr[arr[i + 3]] = arr[a] * arr[b]
        elif arr[i] == 99:
            break
        else:
            print(arr[i])
            print("YIKES")
        i += 4

for line in fileinput.input():
    copy = [int(n) for n in line.split(",")]

    arr = [n for n in copy]
    eval_program(arr, 12, 2)
    print(arr[0])

    for noun in range(100):
        for verb in range(100):
            arr = [n for n in copy]
            eval_program(arr, noun, verb)
            if arr[0] == 19690720:
                print("N: ", noun, "V: ", verb)
