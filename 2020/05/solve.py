import fileinput


def toRow(s):
	return eval('0b' + s.replace('F', '0').replace('B', '1'))

def toCol(s):
	return eval('0b' + s.replace('R', '1').replace('L', '0'))

def seatID(s):
	return toRow(s[:-3]) * 8 + toCol(s[7:])

ids = []

for line in fileinput.input():
	ids.append(seatID(line.rstrip()))

ids.sort()

print(ids[-1])

lst = None

for i in ids:
	if lst and lst + 1 != i:
		print(lst+1)
		break
	lst = i
