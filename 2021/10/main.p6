#! /usr/bin/env raku

use strict;

sub MAIN($file) {
	my %matchList := {
		'{' => '}',
		# '}' => '{',

		'[' => ']',
		# ']' => '[',

		'(' => ')',
		# ')' => '(',

		'<' => '>',
		# '>' => '<',
	};
	my $total = 0;
	my @total2;
	for $file.IO.lines -> $line {
		my ($a,$b) = sub {
			my @stack = [];
			for $line.comb {
				# say @stack;
				if %matchList.keys.contains($_) {
					@stack.push($_);
				}
				else {
					my $c = @stack.pop;
					return {
						')' => (3,0),
						']' => (57,0),
						'}' => (1197,0),
						'>' => (25137,0),
					}{$_} unless (%matchList{$c} === $_);

				}
			}
			# calculate for incomplete
			my @l.push(%matchList{$_}) for @stack.reverse;
			# say @l;
			my $t=0;
			$t = (($t*5) + { ')'=>1, ']'=>2, '}'=>3, '>'=>4 }{$_}) for @l;
			# say $t;
			return (0,$t);
		}();
		$total += $a;
		@total2.push($b) unless ($b==0);
	}
	say $total;
	say @total2.sort[@total2.elems/2];
	# say @total2;
}
