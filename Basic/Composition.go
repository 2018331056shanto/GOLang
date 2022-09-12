package main

// go does not support Inheritence or extending a class t does have the ability to “borrow”
// pieces of an implementation by embedding types within a struct or interface.
import (
	"fmt"
	"time"
)
type Engine struct{
	Model string
	Engine_Number int
	Litter int
}
type Car struct{
	// car has also engine so we can composite it
	Engine
	Car_Model string
	Color string
}
func (e Engine) key(k string) string{

	return k+"i am don"
}
  
type purchase interface {
    sell()
}
  
type display interface {
    show()
}
type salesman interface {
    purchase
    display
}
type game struct {
    name, price    string
    gameCollection []string
}
func (t game) sell() {
    fmt.Println("--------------------------------------")
    fmt.Println("Name:", t.name)
    fmt.Println("Price:", t.price)
    fmt.Println("--------------------------------------")
}
func (t game) show() {
    fmt.Println("The Games available are: ")
    for _, name := range t.gameCollection {
        fmt.Println(name)
    }
    fmt.Println("--------------------------------------")
}
func (t game) Vanue(s string) string{

	time:=time.Now()
	return s+" " +time.String()
}
func shop(s salesman) {
    fmt.Println(s)
    s.sell()
    s.show()
	x:=s.(game).Vanue("Dhaka")
	fmt.Println(x)
}
func main(){

	p:=Car{
		Engine{Model: "Bugati",Engine_Number: 2345,Litter: 8},"Chiron","Blue"}

	fmt.Println(p.Engine,p.Engine_Number,p.Litter,p.Car_Model,p.Color)
	fmt.Println(p.key("skdlasdal"))

	collection := []string{"XYZ", 
	"Trial by Code", "Sea of Rubies"}
	game1 := game{"ABC", "$125", collection}
	shop(game1)

}