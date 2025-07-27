package main

import "fmt"

func main() {
	var (
		i  int     = 42
		f  float64 = 42.5
		u  uint    = 42
		b  byte    = 65
		s  string  = "123"
		r  rune    = 'A'
		tf float32 = 3.14
		ti int32   = 100
	)

	// int to float64
	fmt.Println(float64(i))
	// float64 to int
	fmt.Println(int(f))
	// int to uint
	fmt.Println(uint(i))
	// uint to int
	fmt.Println(int(u))
	// byte to string
	fmt.Println(string(b))
	// string to byte slice
	fmt.Println([]byte(s))
	// rune to string
	fmt.Println(string(r))
	// float32 to int32
	fmt.Println(int32(tf))
	// int32 to float32
	fmt.Println(float32(ti))
	// string to int (using strconv)
	// import "strconv" for this
	// n, _ := strconv.Atoi(s)
	// fmt.Println(n)
}
