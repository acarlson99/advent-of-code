#! /usr/bin/env raku

use strict;

sub MAIN($file) {
	my $score;
	my $score2;
	for $file.IO.lines -> $line {
		$line ~~ /(.)' '(.)/;
		my $a = Str("{$0}");
		my $b = Str("{$1}");
		sub getOutcome($a,$b) {
			if (	($a~~'A' and $b~~'Y') or
				($a~~'B' and $b~~'Z') or
				($a~~'C' and $b~~'X')) {
				return 6;
			} 
			elsif (	($a~~'A' and $b~~'X') or
				($a~~'B' and $b~~'Y') or
				($a~~'C' and $b~~'Z')) {
				return 3;
			} 
			else { return 0; }
		}
		my $outcome = getOutcome($a,$b);
		sub score($b) {
			if ($b~~'X') { return 1; }
			if ($b~~'Y') { return 2; }
			if ($b~~'Z') { return 3; }
		}
		my $s = score($b);
		$score += $s + $outcome;

		# p2
		# X = lose
		# Y = draw
		# Z = win
		my $ob;
		# lose
		if ($b~~'X') {
			if ($a~~'A') { $ob = 'Z'; }
			if ($a~~'B') { $ob = 'X'; }
			if ($a~~'C') { $ob = 'Y'; }
		}
		# draw
		if ($b~~'Y') {
			if ($a~~'A') { $ob = 'X'; }
			if ($a~~'B') { $ob = 'Y'; }
			if ($a~~'C') { $ob = 'Z'; }
		}
		# win
		if ($b~~'Z') {
			if ($a~~'A') { $ob = 'Y'; }
			if ($a~~'B') { $ob = 'Z'; }
			if ($a~~'C') { $ob = 'X'; }
		}
		my $out = getOutcome($a, $ob);
		my $sc2 = score($ob);
		$score2+=$out+$sc2;
	}
	say $score;
	say $score2;
}
