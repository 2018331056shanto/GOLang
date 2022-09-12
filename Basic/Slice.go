
// Slices wrap arrays to give a more general,
//  powerful, and convenient interface to sequences of data.
// Slices hold references to an underlying array, and 
// if you assign one slice to another, both refer to the same array
//  Syntax s[lo:hi]
package main

import "fmt"

func main() {
	mySlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(mySlice)

	fmt.Println(mySlice[1:4])
	// [3 5 7]

	// missing low index implies 0
	fmt.Println(mySlice[:3])
	// [2 3 5]

	// missing high index implies len(s)
	fmt.Println(mySlice[4:])
	// [11 13]

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]    //slice a 
	b := names[1:3]      //slice b
	fmt.Println(a, b)

	// Slices hold references to an underlying array, and 

	b[0] = "XXX"      // value at zeroth index of slice b changed
	// as b starts with names array's 1st index(0 indexing) so if we change the value in b's 0th index 
	// as it is referencing to the actual array so the original array's 1st index will also change
	// all the slices which have the 1st index will also change 

	fmt.Println(a, b)
	// In the code above two slices, a and b are made. With a containing 
	// elements at the index 0 and 1 of the array and b containing the elements at index 1 and 2 of the array.
	fmt.Println(names)

// Besides creating slices by passing the values right away (slice literal), you can also use make. You create an empty
// slice of a specific length and then populate each entry:
	cities := make([]string,6 )
	cities[0] = "Santa Monica"
	cities[1] = "Venice"
	cities[2] = "Los Angeles"
	fmt.Printf("%q", cities)
	// Appending to slice
	cities[3]="Dhaka"


	// To specify a capacity, pass a third argument to make:

	bold := make([]int, 10, 20) // len(b)=0, cap(b)=5
	fmt.Println(bold)


	//  a slice is seating on top of an array, in this case,
	//  the array is empty and the slice can’t set a value in the referred array. There is a way to do that though, and that is by using the append function:
	division:=[]string{}
	division=append(division, "Dhaka","Sylhet","Chottogram")
	fmt.Printf("%q\n",division)
// append a slice to another using ellipsis
	merge:=[]string{}
	// You obviously can’t append a slice of type []int to another slice of type []string.
	// merge=append(merge,cities...)
	// fmt.Printf("%q\n",merge)
	// merge=append(merge, division...)
	merge=append(division,cities...)
	fmt.Printf("%q\n",merge)
	fmt.Printf("%q\n",merge[1:5])


	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	


	// Length
	fmt.Println(len(merge))

	xx:=[]int {1,2,4,5}
	fmt.Println(xx)
	yy:=xx[0:2]
	fmt.Println(yy)
	zz:=xx[2:]
	fmt.Println(zz)
	kk:=append(yy,zz...)
	fmt.Println(kk)
	kk[0]=123
	fmt.Println(kk)
	fmt.Println(xx)
}