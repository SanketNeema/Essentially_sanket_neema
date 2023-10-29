package main

import (
	"fmt"
)

type Response struct {
	firstSubPackage   int16
	secondSubPackage  string
	thirdSubPackage   int16
	fourthSubPackage  string
	fifthSubPackage   int16
	sixthSubPackage   string
	seventhSubPackage int32
}

func DecodePackage(apackage []byte) Response {

	// 1. First 2 bytes represent a short datatype
	// 2. The next 12 bytes represent 12 characters
	// 3. The next 1 byte represent a single byte
	// 4. The next 8 bytes represent 8 characters
	// 5. The next 2 bytes represent a short datatype
	// 6. The next 15 bytes represent 15 characters
	// 7. The next 4 bytes represent a long datatype

	firstchan := make(chan int16)
	secondchan := make(chan string)
	thirdchan := make(chan int16)
	fourthchan := make(chan string)
	fifthchan := make(chan int16)
	sixthchan := make(chan string)
	seventhchan := make(chan int32)

	go func() {
		firstchan <- int16(apackage[0])<<8 | int16(apackage[1])
	}()

	go func() {
		secondchan <- string(apackage[2:14])
	}()

	go func() {
		thirdchan <- int16(apackage[14])
	}()

	go func() {
		fourthchan <- string(apackage[15:23])
	}()

	go func() {
		fifthchan <- int16(apackage[23])<<8 | int16(apackage[24])
	}()

	go func() {
		sixthchan <- string(apackage[25:40])
	}()

	go func() {
		seventhchan <- int32(apackage[40])<<24 | int32(apackage[41])<<16 | int32(apackage[42])<<8 | int32(apackage[43])
	}()

	return Response{
		firstSubPackage:   <-firstchan,
		secondSubPackage:  <-secondchan,
		thirdSubPackage:   <-thirdchan,
		fourthSubPackage:  <-fourthchan,
		fifthSubPackage:   <-fifthchan,
		sixthSubPackage:   <-sixthchan,
		seventhSubPackage: <-seventhchan,
	}

}

func main() {
	// Given byte slice
	byteSlice := []byte{'\x04', '\xD2', '\x6B', '\x65', '\x65', '\x70', '\x64', '\x65', '\x63', '\x6F', '\x64', '\x69', '\x6E', '\x67', '\x38', '\x64', '\x6F', '\x6E', '\x74', '\x73', '\x74', '\x6F', '\x70', '\x03', '\x15', '\x63', '\x6F', '\x6E', '\x67', '\x72', '\x61', '\x74', '\x75', '\x6C', '\x61', '\x74', '\x69', '\x6F', '\x6E', '\x73', '\x07', '\x5B', '\xCD', '\x15'}

	// Convert byte slice to string
	str := DecodePackage(byteSlice)

	// Print the resulting string
	fmt.Println(str)
}
