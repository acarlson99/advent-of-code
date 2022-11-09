#!/bin/bash

# ./get.sh 2020 01 "0123456789OAuthCookie9876543210"

OUTPATH=$(printf "%s/%d/%02d/input.txt" $(dirname $0) $1 $2)

linkFile() {
	fname=$(dirname $0)/input.txt
	echo "linking $fname"
	ln -f -s `readlink $OUTPATH` $fname
	echo "linking /tmp/input.txt"
	ln -f -s `readlink $OUTPATH` /tmp/input.txt
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

TMPFILE=`mktemp`
curl -s $LOC -b "session=$3" -o $TMPFILE

if [ $? != 0 ]
then
	echo "CURL ERROR"
elif [[ `cat $TMPFILE` =~ "404 Not Found" ]]
then
	echo "NOT FOUND: $LOC"
	cat $TMPFILE
elif [[ `cat $TMPFILE` =~ "Please don't repeatedly request this endpoint before it unlocks" ]]
then
	echo "Not yet available: $LOC"
	cat $TMPFILE
elif [[ `cat $TMPFILE` =~ "Puzzle inputs differ by user.  Please log in to get your puzzle input." ]]
then
	echo "Missing session cookie"
	cat $TMPFILE
else
	echo $OUTPATH
	mkdir -p $(dirname $OUTPATH)
	cp $TMPFILE $OUTPATH
	linkFile
fi
rm $TMPFILE
