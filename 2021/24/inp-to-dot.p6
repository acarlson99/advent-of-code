#! /usr/bin/env raku

my %nums := { 'w'=>0, 'x'=>0, 'y'=>0, 'z'=>0, 'inp'=>0, 'add'=>0, 'mul'=>0, 'div'=>0, 'mod'=>0, 'eql'=>0 };

sub f($opName, $v0, $v1) {
	# add W X
	my $opID = ++%nums{$opName};
	my $varID = ++%nums{$v0};
	say $opName,'_',$opID,' [label=',$opName,']';
	say Str($v0),'_',$varID,' [label=',Str($v0),']';

	# add_4 -> W_4
	say $opName,'_',$opID,' -> ',Str($v0),'_',$varID;
	# W_3 -> add_4
	say Str($v0),'_',$varID-1,' -> ',$opName,'_',$opID;
	# X_0 -> add_4
	my $v2 = Str($v1);
	if ('wxyz'.contains($v2)) {
		say Str($v1),'_',%nums{$v1},' -> ',$opName,'_',$opID;
	}
	else {
		my $nname = Str(rand);
		say $nname,' [label=',$v2,']';
		say $nname,' -> ',$opName,'_',$opID;
	}
}

sub MAIN($file) {
	say '# https://dreampuf.github.io/GraphvizOnline';
	say 'strict digraph {';
	for 'w'..'z' -> $z {
		say $z,'_0',' [label=',$z,']';
	}
	for $file.IO.lines -> $line {
		given $line {
			when rx:s/inp (.)/ {
				my $opID = ++%nums{'inp'};
				my $varID = ++%nums{$/[0]};
				say 'inp_',$opID,' [label=inp]';
				say Str($/[0]),'_',$varID,' [label=',Str($/[0]),']';
				say 'inp_',$opID,' -> ',Str($/[0]),'_',$varID;
			}
			when rx:s/add (.) (\-?.)/ {
				f('add', Str($/[0]), Str($/[1]));
			}
			when rx:s/mul (.) (\-?.)/ {
				f('mul', Str($/[0]), Str($/[1]));
			}
			when rx:s/div (.) (\-?.)/ {
				f('div', Str($/[0]), Str($/[1]));
			}
			when rx:s/mod (.) (\-?.)/ {
				f('mod', Str($/[0]), Str($/[1]));
			}
			when rx:s/eql (.) (\-?.)/ {
				f('eql', Str($/[0]), Str($/[1]));
			}
		}
	}
	say '}';
}
