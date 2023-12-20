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

// code by Shon
func IntToHex2(num int64) []byte {
	hexstr := strconv.FormatInt(num, 16)
	byteArr, _ := hex.DecodeString(hexstr)
	return byteArr
}
