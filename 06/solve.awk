BEGIN {
	FS = ")"
}

{
	parents[$2] = $1
	if (children[$1]) {
		children[$1] = children[$1]","
	}
	children[$1] = children[$1]$2
}

# return number of parents of node
func eval_elem(str) {
	if (str == "COM") {
		return (0)
	}
	return eval_elem(parents[str]) + 1
}

# return 1 if found santa, else 0
# append expansion of node in tosearch, inc elems
func expand(str) {
	split(children[str], kids, ",")
	kids[0] = parents[str]

	for (elem in kids) {
		if (searched[kids[elem]] != 1) {
			chain[kids[elem]] = chain[str] + 1
			if (kids[elem] == "SAN")
				return 1
			searched[kids[elem]] = 1
			tosearch[elems] = kids[elem]
			elems++
		}
	}
	return 0
}

# BFS for SAN node from YOU node
# store length from YOU in chain["SAN"]
func find_san() {
	# chain = arr of distance from YOU
	chain["YOU"] = 0

	# list of nodes to search
	elems = 0
	tosearch[elems++] = "YOU"

	# BFS
	ii = 0
	while (1 == 1) {
		if (tosearch[ii] == "SAN")
			break
		if (expand(tosearch[ii]))
			break
		ii++
	}
}

END {
	# count parents
	for (elem in parents) {
		total += eval_elem(elem)
	}
	print "Part one:", total
	find_san()
	# -2 because counting length from YOU to SAN rather than len from parent nodes
	print "Part two:", chain["SAN"] - 2
}
