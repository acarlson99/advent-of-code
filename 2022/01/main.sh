SORTED_TOTALS=`awk '/[0-9]+/ { n += $1 } /^$/ { print n; n=0; }' input.txt | sort -ruh`

echo $SORTED_TOTALS | tr ' ' '\n' | head -1
echo $SORTED_TOTALS | tr ' ' '\n' | head -3 | awk '{ n+=$1; } END { print n; }'
