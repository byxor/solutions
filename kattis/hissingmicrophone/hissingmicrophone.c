char*M="no hiss\n";char m[31];main(){scanf("%s",m);int i,s=0;for(;i<31;){if(m[i]==0||s>=2)break;if(m[i++]=='s')s++;else s=0;}printf(s>=2?M+3:M);}
