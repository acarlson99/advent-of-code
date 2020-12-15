#!/bin/bash

# ./get.sh 2020 01 "0123456789OAuthCookie9876543210"

OUTPATH=$(printf "%s/%d/%02d/input.txt" $(dirname $0) $1 $2)

linkFile() {
	echo "linking $(dirname $0)/input.txt"
	ln -f -s $OUTPATH $(dirname $0)/input.txt
}

if [ -f $OUTPATH ]
then
	echo "$OUTPATH file exists.  [c]ontinue [s]kip [l]ink"
	read USRIN
	case $USRIN in
		"c")
			echo "Continuing..."
			;;
		"s")
			echo "Exiting"
			exit 0
			;;
		"l")
			linkFile
			exit 0
			;;
		*)
			echo "Bad input"
			exit 1
	esac
fi

LOC="https://adventofcode.com/$1/day/$(printf '%d' $2)/input"

curl -s $LOC -b "session=$3" -o /tmp/input.txt

if [ $? != 0 ]
then
	echo "CURL ERROR"
elif [[ $INPUT =~ "404 Not Found" ]]
then
	echo "ERROR: $LOC"
else
	echo $OUTPATH
	mkdir -p $(dirname $OUTPATH)
	cp /tmp/input.txt $OUTPATH
	linkFile
fi
