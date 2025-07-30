package main

import "fmt"

func main(){

var n int

var nums [] int

fmt.Scan(&n)

for i:=0;i<n;i++{

	var x int
	fmt.Scan(&x)

	nums=append(nums,x)

}

mx:=-1
mn:=10000000
sum:=0
for _,n:=range nums{

	sum+=n
	mx=max(mx,n)
	mn=min(mn,n)
}

fmt.Println(sum-mx-mn) 
}
