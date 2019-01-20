char*M="no hiss",m[31],i,s=0;
main(){
 gets(m);
 for(;m[i]&&s<2;){
  s=m[i++]^'s'?0:s+1;
 }
 puts(M+(s>1)*3);
}
