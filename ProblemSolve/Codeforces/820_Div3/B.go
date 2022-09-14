package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var t int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for x := 0; x < t; x++ {
		var n int
		fmt.Fscan(in,&n)
		var s string
		fmt.Fscan(in,&s)
		start:=n-1
		var ans string 
		const st int =96
		for start>=0{
			if s[start]!='0'{
				// fmt.Println("hello")
				c,_:=strconv.Atoi(string(s[start]))
				val:=string(st+c)
				ans+=val
				start--
			}else{


				c,_:=strconv.Atoi(string(s[start-2])+string(s[start-1]))
				val:=string(st+c)
				ans +=val
				start-=3
			}
		}
		for i:=len(ans)-1;i>=0;i--{
			fmt.Print(string(ans[i]))
		}
		fmt.Println()
	}

}
