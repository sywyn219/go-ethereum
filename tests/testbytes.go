package main


import (
	"fmt"
	"encoding/hex"
)

func main(){
	// 506c6564676
	// 52656465656d
	fmt.Println(hex.EncodeToString([]byte("redeem")))
	
}