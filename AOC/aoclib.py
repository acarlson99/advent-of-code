def doDFS(g,start,expandF,isEndF):
    q=[]
    visited = set()
    q.append((start,1))
    while len(q)>0:
        pos,dist = q.pop(0)
        if isEndF(g,pos):
            return dist
        if pos in visited:
            continue
        visited.add(pos)
        for newPos in expandF(g,pos):
            if newPos not in visited:
                if isEndF(g,newPos):
                    return dist
                q.append((newPos,dist+1))
    return -1

def readFile(name):
    f = open(name, 'r')
    contents = f.read().strip()
    f.close()
    return contents
