awk -F',' 'BEGIN{for(n=0;n<256;n++)ord[sprintf("%c",n)]=n}{for(i=1;i<=NF;i++){n=0;split($i,chars,"");for(j=1;j<=length(chars);j++){n+=ord[chars[j]];n*=17;n=n%256};print n}}' | awk '{n+=$0}END{print n}'
