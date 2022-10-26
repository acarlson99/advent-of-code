#! /usr/bin/env raku

use strict;

sub go(%map, $x1, $x2, $y1, $y2) {
	if ($x1 > $x2 || $y1 > $y2) {
		return go(%map, $x2, $x1, $y2, $y1);
	}
	if ($y1 == $y2) {
		for $x1..$x2 {
			%map{$y1}{$_}++;
		}
	}
	if ($x1 == $x2) {
		for $y1..$y2 {
			%map{$_}{$x1}++;
		}
	}
}

sub go2(%map, $x1, $x2, $y1, $y2) {
	if ($y1 > $y2) {
		return go2(%map, $x2, $x1, $y2, $y1);
	}
	my $dir = 1;
	if $x1 > $x2 { $dir = -1; }
	for $y1..$y2 {
		my $y = $_;
		my $x = $x1 + $dir * ($y - $y1);
		%map{$y}{$x}++;
	}
}

sub countem(%arr) {
	my $cnt=0;
	for %arr.keys {
		my %d = %arr{$_};
		my $k = $_;
		for %d.keys {
			if (%d{$_} >= 2) {
				$cnt++;
			}
		}
	}
	return $cnt
}

sub MAIN($file) {
	my %arr;
	my %arr2;
	for $file.IO.lines -> $line {
		$line ~~ /(\d+)\,(\d+)....(\d+)\,(\d+)/;
		my $x1 = +$0;
		my $y1 = +$1;
		my $x2 = +$2;
		my $y2 = +$3;
		if ($x1==$x2 || $y1==$y2) {
			go(%arr, $x1, $x2, $y1, $y2);
			go(%arr2, $x1, $x2, $y1, $y2);
		} else {
			go2(%arr2, $x1, $x2, $y1, $y2);
		}
	}
	say(countem(%arr));
	say(countem(%arr2));
}
