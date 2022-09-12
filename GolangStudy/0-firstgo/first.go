package main

import(
	"fmt"
	
)
func jisheng(n int)(result int){
	if(n>1){
	result=n*jisheng(n-1)
		return result
	}
	return 1
}

func main() {  
	b:=make([]byte,20)
    a:=jisheng(6)
	fmt.Printf("格式%T a=%v\n",a,a)
	
	b[0]=63
	fmt.Printf("格式%T a=%v\n",b,b)
}