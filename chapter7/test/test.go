package main

import "fmt"

func main() {
	d := make(map[int]string)
	d[0] = "i"
	d[1] = "do"
	d[2] = "house"
	d[3] = "with"
	d[4] = "mouse"
	d[5] = "in"
	d[6] = "not"
	d[7] = "like"
	d[8] = "them"
	d[9] = "ham"
	d[10] = "a"
	d[11] = "anywhere"
	d[12] = "green"
	d[13] = "eggs"
	d[14] = "and"
	d[15] = "here"
	d[16] = "or"
	d[17] = "there"
	d[18] = "sam"
	d[19] = "am"

	fmt.Println(d[18], d[0], d[19])
}
