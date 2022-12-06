package main

import (
	"bytes"
	"fmt"
)

func main() {
	data, _ := ReadInput()

	fmt.Println("Part 1 start: ", demarc([]byte(data[0]), 4))
	fmt.Println("Part 2 start: ", demarc([]byte(data[0]), 14))
}

func demarc(stream []byte, marker int) int {
	currentIndex := 0
	startOfPacketSize := marker

	for index, character := range stream {
		foundStartOfPacket := false
		if index < startOfPacketSize {
			if bytes.Contains(stream[:index], []byte{character}) {
				currentIndex = index
			}
			continue
		} // skip past initial few characters

		currentPacket := stream[index-startOfPacketSize : index]

		foundCharacter := map[byte]bool{}
		for packetIndex, packetCharacter := range currentPacket {
			if foundCharacter[packetCharacter] {
				break
			} else if !foundCharacter[packetCharacter] {
				foundCharacter[packetCharacter] = true
			}

			if packetIndex >= startOfPacketSize-1 {
				currentIndex = index
				foundStartOfPacket = true
			}
		}

		currentIndex = index

		if foundStartOfPacket { // If we found our packet, then don't check the rest of the byte stream
			break
		}
	}

	return currentIndex
}
