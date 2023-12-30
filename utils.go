package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"log"
	"strconv"
)

// IntToHex converts an int64 to a byte array

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// ReverseBytes
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// code by Shon
func IntToHex2(num int64) []byte {
	hexstr := strconv.FormatInt(num, 16)
	byteArr, _ := hex.DecodeString(hexstr)
	return byteArr
}
