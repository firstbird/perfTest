package main

import (
	"hash/crc32"
	"fmt"
)

func main() {
	fmt.Printf("go-crc32: %d \n ",crc32.Checksum([]byte{179}, crc32.MakeTable(crc32.Castagnoli)))
}
