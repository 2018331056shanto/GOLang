package main

import (
	"fmt"
	"strings"
)

// A map maps​ keys to values.
func main(){

	// map[key data type] valueDataType
	cities:=map[string] int{
		"Dhaka":10,
		"Rajshahi":6,
		"Mymensingh":5,
	}
	for k,v:=range cities{
		fmt.Println(k,v)
	}
	// map must be created using "make" key word
	mp:=make(map[string]int)
	mp["Tufail"]=10
	mp["Titu"]=32
	mp["Shanto"]=56
	for k,v:=range mp{
		fmt.Println(k,v)
	}
	// delete an element from map
	// delete(mapName, Key)
	delete(mp,"Titu")
	for k,v:=range mp{
		fmt.Println(k,v)
	}
	elem,ok:=mp["Titu"]
	fmt.Println(ok,elem)
	s:="i am shanto "
	// it turns a string into  a slice and all the white space are count as separator

	vv:=strings.Fields(s)
	fmt.Println(vv)
	
	for x:=range vv{
		fmt.Println(vv[x])
	}
}