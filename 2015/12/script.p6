#!/usr/bin/env raku

use strict;

sub MAIN() {
	my @stack;
	my $inObj = False;
	my $objIgnore = False;
	my $total = 0;
	for $*IN.lines -> $line {
		# say $line;
		given $line {
			when /'{'/ {
				my $inO = $inObj;
				my $t = $total+0;
				my $ign = $objIgnore;
				# say 'PUSH {', " $inO, $t, $ign";
				@stack.push(($inO,$t,$ign));
				$total = 0;
				$objIgnore = False;
				$inObj = True;
			}
			when /'['/ {
				my $inO = $inObj;
				my $t = $total+0;
				my $ign = $objIgnore;
				# say 'PUSH [', " $inO, $t, $ign";
				@stack.push(($inO,$t,$ign));
				$total = 0;
				$objIgnore = False;
				$inObj = False;
			}
			when /'": "red"'/ {
				# say "ignoring";
				$objIgnore = True;
			}
			when /('-'*\d+)/ {
				$total += 0+$/;
				# say "total is $total";
			}
			when /'}'/ {
				# say 'POP }';
				my $t = $total;
				if ($objIgnore and $inObj) { $t=0; }
				($inObj,$total,$objIgnore) = @stack.pop();
				$total += $t;
			}
			when /']'/ {
				# say 'POP ]';
				my $t = $total;
				($inObj,$total,$objIgnore) = @stack.pop();
				$total += $t;
			}
		}
		# say "$total ", @stack;
		# say "inObj = $inObj; objIgnore = $objIgnore; total = $total";
	}
	say $total;
}
