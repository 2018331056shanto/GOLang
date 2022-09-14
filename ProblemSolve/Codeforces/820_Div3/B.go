package main

import (
	"bufio"
	"fmt"
	"os"
	
)
var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)
func In(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}
func Out(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
func solve(){
	var n int
	In("%d\n",&n)
	var s string
	In("%s\n",&s)

	start:=n-1
		var ans string 
		const st int =96
		for start>=0{
			if s[start]!='0'{
				// fmt.Println("hello")
				s:=int(s[start])-48
				val:=string(st+s);
				ans+=val
				start--
			}else{


				z:=int(s[start-2])-48
				y:=int(s[start-1])-48
				val:=string(st+z*10+y)
				ans +=val
				start-=3
			}
		}
		for i:=len(ans)-1;i>=0;i--{
			fmt.Print(string(ans[i]))
		}
		fmt.Println()

}

func main() {
	var t int
	 In("%d\n",&t)
	 for t>0{
		solve()
		t--
	 }
	
	

}
