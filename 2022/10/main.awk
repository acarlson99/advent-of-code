function f() {
	im = i%40;
	i+=1;
	if (n-1==im||n==im||n+1==im) {
		s = s"#"
	} else {
		s = s"."
	}
	# print i,n,s
	if (i==20||i==60||i==100||i==140||i==180||i==220) {
		# print n*i;
		sum+=n*i
	}
	if (i%40==0) {
		print s
		s = ""
	}
}

BEGIN {
	print "p2 V"
	n=1;i=0;
}

/addx/ {
	f();
	f();
	n += $2;
}

/noop/ { f() }

END {
	print "p1 >", sum
}
