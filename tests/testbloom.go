package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	bloomfilter "github.com/holiman/bloomfilter/v2"
)


type stateBloom struct {
	bloom *bloomfilter.Filter
}

type stateBloomHasher []byte

func (f stateBloomHasher) Write(p []byte) (n int, err error) { panic("not implemented") }
func (f stateBloomHasher) Sum(b []byte) []byte               { panic("not implemented") }
func (f stateBloomHasher) Reset()                            { panic("not implemented") }
func (f stateBloomHasher) BlockSize() int                    { panic("not implemented") }
func (f stateBloomHasher) Size() int                         { return 8 }
func (f stateBloomHasher) Sum64() uint64                     { return binary.BigEndian.Uint64(f) }

func main(){
	stateBloom,_:=newStateBloomWithSize(1)

	// for i:=0;i<=5000000;i++{
	// 	fmt.Println(i)
	// 	hash:=sha256.Sum256([]byte(strconv.Itoa(i)))
	// 	err:=stateBloom.Put(hash[:],nil)
	// 	if err!=nil{
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }

	
    var hash0 []byte

    hash:=sha256.Sum256([]byte(strconv.Itoa(25000)))
	fmt.Println(hex.EncodeToString(hash[:]))
	hash2:=sha256.Sum256([]byte(strconv.Itoa(25001)))
	fmt.Println(hex.EncodeToString(hash2[:]))
	fmt.Println(len(hex.EncodeToString(hash2[:]))/64)

	fmt.Println(hex.EncodeToString([]byte("pledge")))
     hash0=append([]byte("pledge"),hash[:]...)
     hash0=append(hash0,hash2[:]...)
	
	fmt.Println(hex.EncodeToString(hash0))

	fmt.Println(hex.EncodeToString(hash0[:6]))


		err:=stateBloom.Put(hash[:],nil)
		if err!=nil{
			fmt.Println(err)
			return
		}

	key,err:=hex.DecodeString("6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b")
	booll,err:=stateBloom.Contain(key)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(booll)
	}

	key,err=hex.DecodeString("0812a4ef4ea9e800e2cb87a311317fdb06f80ab8700c7185b35588b3d6953739")
	booll,err=stateBloom.Contain(key)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(booll)
	}
	
}

func newStateBloomWithSize(size uint64)(*stateBloom, error) {
	bloom, err := bloomfilter.New(8,4)
	if err != nil {
		return nil, err
	}
	return &stateBloom{bloom: bloom}, nil
}

var HashLength = 32

func (bloom *stateBloom) Put(key []byte, value []byte) error {
	if len(key) != HashLength {
		return errors.New("error")
	}
	bloom.bloom.Add(stateBloomHasher(key))
	return nil
}


func (bloom *stateBloom) Contain(key []byte) (bool, error) {
	return bloom.bloom.Contains(stateBloomHasher(key)), nil
}
