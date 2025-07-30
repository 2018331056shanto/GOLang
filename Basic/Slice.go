
// Slices wrap arrays to give a more general,
//  powerful, and convenient interface to sequences of data.
// Slices hold references to an underlying array, and 
// if you assign one slice to another, both refer to the same array
//  Syntax s[lo:hi]
package main

import "fmt"

func main() {
// 	mySlice := []int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(mySlice)

// 	fmt.Println(mySlice[1:4])
// 	// [3 5 7]

// 	// missing low index implies 0
// 	fmt.Println(mySlice[:3])
// 	// [2 3 5]

// 	// missing high index implies len(s)
// 	fmt.Println(mySlice[4:])
// 	// [11 13]

// 	names := [4]string{
// 		"John",
// 		"Paul",
// 		"George",
// 		"Ringo",
// 	}
// 	fmt.Println(names)

// 	a := names[0:2]    //slice a 
// 	b := names[1:3]      //slice b
// 	fmt.Println(a, b)

// 	// Slices hold references to an underlying array, and 

// 	b[0] = "XXX"      // value at zeroth index of slice b changed
// 	// as b starts with names array's 1st index(0 indexing) so if we change the value in b's 0th index 
// 	// as it is referencing to the actual array so the original array's 1st index will also change
// 	// all the slices which have the 1st index will also change 

// 	fmt.Println(a, b)
// 	// In the code above two slices, a and b are made. With a containing 
// 	// elements at the index 0 and 1 of the array and b containing the elements at index 1 and 2 of the array.
// 	fmt.Println(names)

// // Besides creating slices by passing the values right away (slice literal), you can also use make. You create an empty
// // slice of a specific length and then populate each entry:
// 	cities := make([]string,6 )
// 	cities[0] = "Santa Monica"
// 	cities[1] = "Venice"
// 	cities[2] = "Los Angeles"
// 	fmt.Printf("%q", cities)
// 	// Appending to slice
// 	cities[3]="Dhaka"


// 	// To specify a capacity, pass a third argument to make:

// 	bold := make([]int, 10, 20) // len(b)=0, cap(b)=5
// 	fmt.Println(bold)


// 	//  a slice is seating on top of an array, in this case,
// 	//  the array is empty and the slice can’t set a value in the referred array. There is a way to do that though, and that is by using the append function:
// 	division:=[]string{}
// 	division=append(division, "Dhaka","Sylhet","Chottogram")
// 	fmt.Printf("%q\n",division)
// // append a slice to another using ellipsis
// 	merge:=[]string{}
// 	// You obviously can’t append a slice of type []int to another slice of type []string.
// 	// merge=append(merge,cities...)
// 	// fmt.Printf("%q\n",merge)
// 	// merge=append(merge, division...)
// 	merge=append(division,cities...)
// 	fmt.Printf("%q\n",merge)
// 	fmt.Printf("%q\n",merge[1:5])


// 	s := []struct {
// 		i int
// 		b bool
// 	}{
// 		{2, true},
// 		{3, false},
// 		{5, true},
// 		{7, true},
// 		{11, false},
// 		{13, true},
// 	}
// 	fmt.Println(s)
	


// 	// Length
// 	fmt.Println(len(merge))

// 	xx:=[]int {1,2,4,5}
// 	fmt.Println(xx)
// 	yy:=xx[0:2]
// 	fmt.Println(yy)
// 	zz:=xx[2:]
// 	fmt.Println(zz)
// 	kk:=append(yy,zz...)
// 	fmt.Println(kk)
// 	kk[0]=123
// 	fmt.Println(kk)
// 	fmt.Println(xx)

//Slice 

var slice0 []int // Declare a slice without initializing it
slice0=append(slice0, 1) // Append elements to the slice
fmt.Println(slice0) // Print the slice (will be empty since it was not initialized)


var slice1 []int= []int{1, 2, 3, 4, 5}

fmt.Println(len(slice1)) // Length of slice
fmt.Println(cap(slice1)) // Capacity of slice
fmt.Println(slice1)      // Print the slice


slice2:=[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

fmt.Println(slice2) // Print the slice
fmt.Println(cap(slice2)) // Capacity of slice
fmt.Println(len(slice2)) // Length of slice


var arr =[10] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}	


slice3:=arr[2:5] // Slice from index 2 to 4 (5 is exclusive)
fmt.Println(slice3) // Print the slice
fmt.Println(len(slice3)) // Length of slice
fmt.Println(cap(slice3)) // Capacity of slice\


slice4:=slice3[1:3] // Slice from index 1 to 2 (3 is exclusive)
fmt.Println(slice4) // Print the slice
fmt.Println(len(slice4)) // Length of slice
fmt.Println(cap(slice4)) // Capacity of slice

/*

Slice ={

	Pointer to the first element of the array,
	Length of the slice,
	Capacity of the slice (the maximum number of elements it can hold without reallocating).
}
*/

slice5 := make([]int, 0, 5) // Create a slice with length 0 and capacity 5
fmt.Println(slice5) // Print the slice
fmt.Println(len(slice5)) // Length of slice
fmt.Println(cap(slice5)) // Capacity of slice

slice5 = append(slice5, 1, 2, 3) // Append elements to the slice
fmt.Println(slice5) // Print the slice after appending
fmt.Println(len(slice5)) // Length of slice after appending
fmt.Println(cap(slice5)) // Capacity of slice after appending

}