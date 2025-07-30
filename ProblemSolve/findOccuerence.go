package main
import "fmt"


func main(){

	
	var n int 

	fmt.Scan(&n)

	var nums [] int

	for i:=0;i<n;i++{


		var x int
		fmt.Scan(&x)


		nums=append(nums,x)

	}


	mp:=make(map[int]int)

	for _,ele:=range nums {


		mp[ele]++
	}

	for key,val:=range mp{

		fmt.Println(key,val)
	}

}
