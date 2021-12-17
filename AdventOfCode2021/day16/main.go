package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Packet struct {
	value        int
	version      int
	operatorType int
	lengthType   int
	length       int

	subPackets []*Packet
}

func (p *Packet) getVersionSum() (versionSum int) {
	versionSum += p.version
	for _, subPacket := range p.subPackets {
		versionSum += subPacket.getVersionSum()
	}
	return versionSum
}

func createPacket(packetString string) *Packet {
	var packet Packet

	versionString := packetString[:3]
	packetString = packetString[3:]
	versionNumber, _ := strconv.ParseUint(versionString, 2, 32)
	packet.version = int(versionNumber)

	operatorTypeString := packetString[:3]
	packetString = packetString[3:]
	operatorTypeNumber, _ := strconv.ParseUint(operatorTypeString, 2, 32)
	packet.operatorType = int(operatorTypeNumber)

	if packet.operatorType != 4 {

		lengthTypeString := packetString[:1]
		packetString = packetString[1:]
		lengthTypeNumber, _ := strconv.ParseUint(lengthTypeString, 2, 32)
		packet.lengthType = int(lengthTypeNumber)

		var lengthString string
		if lengthTypeNumber == 0 {
			lengthString = packetString[:15]
			packetString = packetString[15:]
		} else {
			lengthString = packetString[:11]
			packetString = packetString[11:]
		}
		lengthNumber, _ := strconv.ParseUint(lengthString, 2, 32)
		packet.length = int(lengthNumber)

		// Current packet has a subpacket at this point
		if lengthTypeNumber == 0 {
			subpacketAString := packetString[:11]
			packetString = packetString[11:]
			packet.subPackets = append(packet.subPackets, createPacket(subpacketAString))

			if lengthNumber > 11 {
				subpacketBString := packetString[:lengthNumber-11]
				packet.subPackets = append(packet.subPackets, createPacket(subpacketBString))
			}
		} else {
			if lengthNumber == 1 {
				packet.subPackets = append(packet.subPackets, createPacket(packetString))
			} else {
				for i := 0; i < int(lengthNumber); i++ {
					subPacketString := packetString[:11]
					packetString = packetString[11:]
					packet.subPackets = append(packet.subPackets, createPacket(subPacketString))
				}
			}
		}

	} else {
		var valueString string
		for {
			groupString := packetString[:5]
			packetString = packetString[5:]

			prefix := groupString[:1]
			groupString = groupString[1:]
			valueString += groupString
			if prefix == "0" {
				break
			}
		}
		valuenumber, _ := strconv.ParseUint(valueString, 2, 32)
		packet.value = int(valuenumber)
	}
	return &packet
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var hexString string
	var binaryString string
	for scanner.Scan() {
		line := scanner.Text()
		hexString = line

		for _, char := range strings.Split(hexString, "") {

			data, _ := strconv.ParseUint(char, 16, 32)
			bin := fmt.Sprintf("%004b", data)
			binaryString += bin
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	packet := createPacket(binaryString)

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("packets numbers for part 1: ", packet.getVersionSum())

}

func doSubPacket(packet string) {
	subPacket := packet

	subPacketVersionString := packet[:3]
	packet = packet[3:]
	subPacketVersionNumber, _ := strconv.ParseUint(subPacketVersionString, 2, 32)
	fmt.Println("Subpacket version: ", subPacketVersionNumber)

	subPacketTypeIdString := packet[:3]
	packet = packet[3:]
	subPacketTypeIdNumber, _ := strconv.ParseUint(subPacketTypeIdString, 2, 32)

	// Literal value
	if subPacketTypeIdNumber == 4 {

	} else {
		subPacketLengthTypeIdString := packet[:1]
		packet = packet[1:]
		subPacketLengthTypeIdNumber, _ := strconv.ParseUint(subPacketLengthTypeIdString, 2, 32)

		var subPacketLengthString string

		if subPacketLengthTypeIdNumber == 1 {
			subPacketLengthString = packet[:11]
			packet = packet[11:]
		} else {
			subPacketLengthString = packet[:15]
			packet = packet[15:]
		}

		subPacketLengthNumber, _ := strconv.ParseUint(subPacketLengthString, 2, 32)
		if subPacketLengthNumber != 0 {
			if subPacketLengthTypeIdNumber == 1 {
				fmt.Println(len(packet))
				// subPacketLength := (len(packet) / int(subPacketLengthNumber))
				for i := 0; i < int(subPacketLengthNumber); i++ {
					subPacketString := packet[:11]
					packet = packet[11:]
					doSubPacket(subPacketString)
				}
			} else {
				subPacketString := packet[:subPacketLengthNumber]
				packet = packet[subPacketLengthNumber:]

				subPacketAString := subPacketString[:11]
				subPacketA := subPacketAString
				subPacketString = subPacketString[11:]

				subPacketAVersionString := subPacketAString[:3]
				subPacketAString = subPacketAString[3:]
				subPacketAVersionNumber, _ := strconv.ParseUint(subPacketAVersionString, 2, 32)

				subPacketATypeIdString := subPacketAString[:3]
				subPacketAString = subPacketAString[3:]
				subPacketATypeIdNumber, _ := strconv.ParseUint(subPacketATypeIdString, 2, 32)

				subPacketALiteralString := subPacketAString
				subPacketALiteralNumber, _ := strconv.ParseUint(subPacketALiteralString, 2, 32)
				fmt.Println("Subpacket A for part 1: ", subPacketA, " Version: ", subPacketAVersionNumber, " ", subPacketATypeIdNumber, " ", subPacketALiteralNumber)

				if subPacketLengthNumber != 11 {
					subPacketBString := subPacketString[:subPacketLengthNumber-11]
					subPacketB := subPacketBString

					subPacketBVersionString := subPacketAString[:3]
					subPacketBString = subPacketBString[3:]
					subPacketBVersionNumber, _ := strconv.ParseUint(subPacketBVersionString, 2, 32)

					subPacketBTypeIdString := subPacketBString[:3]
					subPacketBString = subPacketBString[3:]
					subPacketBTypeIdNumber, _ := strconv.ParseUint(subPacketBTypeIdString, 2, 32)

					subPacketBLiteralString := subPacketBString
					subPacketBLiteralNumber, _ := strconv.ParseUint(subPacketBLiteralString, 2, 32)
					fmt.Println("Subpacket B for part 1: ", subPacketB, " ", subPacketBVersionNumber, " ", subPacketBTypeIdNumber, " ", subPacketBLiteralNumber)
				}
			}
		}
	}

	fmt.Println("Subpacket for part 1: ", subPacket, " ", subPacketVersionNumber, " ", subPacketTypeIdNumber)
}
