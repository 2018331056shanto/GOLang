package main
import "fmt"

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a Artist) int { //passing an Artist by value
	a.Songs++
	return a.Songs
	//pass by value
}

func newRelease1(a *Artist) int{//we are receiving the address so anything we make change will change the actual value
	//pass by reference

	a.Songs++
	return a.Songs
}
func main() {
	me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	me1:=&Artist{Name: "Matt", Genre: "Electro", Songs: 42}//we are passing the address 
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs\n", me.Name, me.Songs)
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease1(me1))
	fmt.Printf("%s has a total of %d songs\n", me.Name, me1.Songs)
}