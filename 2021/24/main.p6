#! /usr/bin/env raku

my %fns := {
	'*' => sub { @_[0] * @_[1]; },
	'/' => sub { @_[0] / @_[1]; },
	'+' => sub { @_[0] + @_[1]; },
	'%' => sub { @_[0] % @_[1]; },
	'=' => sub { @_[0] == @_[1]; }
};

my %env := {};

sub calc {
	# my %env = shift @_;
	my $lvar = shift @_;
	my $lhs = %env{$lvar};
	# $lhs = -1 unless $lhs!='';
	if ($lhs=='') { $lhs = 0 };
	my $rhs = shift @_;
	$rhs = %env{$rhs} unless $rhs ~~ /\d/;
	# $rhs = -1 unless $rhs!='';
	if ($rhs=='') { $rhs = 0 };
	my $op = shift @_;
	# say "<$lhs> <$op> <$rhs>";
	# say %env{$lhs};
	%env{$lvar} = (%fns{$op})($lhs, $rhs);
	# say %env{$lhs};
	# %env{$lhs} = $lhs * $rhs;
}

sub MAIN($file) {
	# my @num = (1,3,5,7,9,2,4,6,8,9,9,9,9,9);
	loop (my $n=99999999999999; $n>1; $n--) {
		%env := {};
		next if $n.contains(0);
		my @num = "$n".split('',:v);
		shift @num;
		say @num;
		# my @num = (1,3,5,7,9,2,4,6,8,9,9,9,9,9);
		for $file.IO.lines -> $line {
			given $line {
				when rx:s/inp (.)/ {
					# %env{$/[0]} = $*IN.lines[0];
					%env{$/[0]} = shift @num;
				}
				when rx:s/add (.) (.)/ {
					# calc %env, $/[0], $/[1], '+';
					calc $/[0], $/[1], '+';
				}
				when rx:s/mul (.) (.)/ {
					# calc %env, $/[0], $/[1], '*';
					calc $/[0], $/[1], '*';
				}
				when rx:s/div (.) (.)/ {
					# calc %env, $/[0], $/[1], '/';
					calc $/[0], $/[1], '/';
				}
				when rx:s/mod (.) (.)/ {
					# calc %env, $/[0], $/[1], '%';
					calc $/[0], $/[1], '%';
				}
				when rx:s/eql (.) (.)/ {
					# calc %env, $/[0], $/[1], '=';
					calc $/[0], $/[1], '=';
				}
			}
		}
		# say %env;
		say %env{'z'};
		if (%env{'z'}==0) { say $n; return; }
	}
}
