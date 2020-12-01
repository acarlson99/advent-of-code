#!/bin/bash

# ./get.sh 2020 01 "0123456789OAuthCookie9876543210"

LOC="https://adventofcode.com/$1/day/$(printf '%d' $2)/input"

curl -s $LOC -b "session=$3" -o /tmp/input.txt

if [ $? != 0 ]
then
	echo "CURL ERROR"
elif [[ $INPUT =~ "404 Not Found" ]]
then
	echo "ERROR: $LOC"
else
	OUTPATH=$(printf "%s/%d/%02d/input.txt" $(dirname $0) $1 $2)
	echo $OUTPATH
	mkdir -p $(dirname $OUTPATH)
	cp /tmp/input.txt $OUTPATH
fi
