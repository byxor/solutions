char*M="no hiss",m[31],i,s=0;main(){gets(m);for(;s<2&&m[i];)s=m[i++]^'s'?0:s+1;puts(M+(s>1)*3);}
