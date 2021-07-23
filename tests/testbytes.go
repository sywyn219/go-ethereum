package main


import (
	"fmt"
	"encoding/hex"
)

func main(){
	// 506c6564676
	// 52656465656d
	fmt.Println(hex.EncodeToString([]byte("pledge")))
	fmt.Println(hex.EncodeToString([]byte("redeem")))
	fmt.Println(hex.EncodeToString([]byte("unlockReward")))
   str:="0x0812a4ef4ea9e800e2cb87a311317fdb06f80ab8700c7185b35588b3d6953739"
	fmt.Println(str[2:])

	
}