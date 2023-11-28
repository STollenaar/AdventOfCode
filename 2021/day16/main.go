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

// Simple version sum for part 1
func (p *Packet) getVersionSum() (versionSum int) {
	versionSum += p.version
	for _, subPacket := range p.subPackets {
		versionSum += subPacket.getVersionSum()
	}
	return versionSum
}

// Get the value of the packet, if operation packet keep going
// 	down till a literal packet has been reached
func (p *Packet) getValue() int {
	if p.operatorType == 4 {
		return p.value
	} else {
		return p.doOperation()
	}
}

// Hardest part of this problem, parsing the subpackets correctly
func createPacket(packetString string) (*Packet, string) {
	var packet Packet

	versionString := packetString[:3]
	packetString = packetString[3:]
	versionNumber, _ := strconv.ParseUint(versionString, 2, 64)
	packet.version = int(versionNumber)

	operatorTypeString := packetString[:3]
	packetString = packetString[3:]
	operatorTypeNumber, _ := strconv.ParseUint(operatorTypeString, 2, 64)
	packet.operatorType = int(operatorTypeNumber)

	if packet.operatorType != 4 {

		lengthTypeString := packetString[:1]
		packetString = packetString[1:]
		lengthTypeNumber, _ := strconv.ParseUint(lengthTypeString, 2, 64)
		packet.lengthType = int(lengthTypeNumber)

		var lengthString string
		if lengthTypeNumber == 0 {
			lengthString = packetString[:15]
			packetString = packetString[15:]
		} else {
			lengthString = packetString[:11]
			packetString = packetString[11:]
		}
		lengthNumber, _ := strconv.ParseUint(lengthString, 2, 64)
		packet.length = int(lengthNumber)

		// Current packet has a subpacket at this point
		if lengthTypeNumber == 0 {
			subPackets := packetString[:lengthNumber]
			packetString = packetString[lengthNumber:]

			for subPackets != strings.Repeat("0", len(subPackets)) {
				subPacket, nextSubPacket := createPacket(subPackets)
				packet.subPackets = append(packet.subPackets, subPacket)
				subPackets = nextSubPacket
			}
		} else {
			for i := 0; i < int(lengthNumber); i++ {
				subPacket, returnedPacket := createPacket(packetString)
				packet.subPackets = append(packet.subPackets, subPacket)
				packetString = returnedPacket
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
		valuenumber, _ := strconv.ParseUint(valueString, 2, 64)
		packet.value = int(valuenumber)
	}
	return &packet, packetString
}

func (p *Packet) doOperation() (value int) {
	switch p.operatorType {
	case 0:
		for _, subPacket := range p.subPackets {
			value += subPacket.getValue()
		}
	case 1:
		value = 1
		for _, subPacket := range p.subPackets {
			value *= subPacket.getValue()
		}
	case 2:
		for i, subPacket := range p.subPackets {
			subPacketValue := subPacket.getValue()
			if i == 0 || subPacketValue < value {
				value = subPacketValue
			}
		}
	case 3:
		for i, subPacket := range p.subPackets {
			subPacketValue := subPacket.getValue()
			if i == 0 || subPacketValue > value {
				value = subPacketValue
			}
		}
	case 5:
		subPacket1, subPacket2 := p.subPackets[0].getValue(), p.subPackets[1].getValue()
		if subPacket1 > subPacket2 {
			value = 1
		} else {
			value = 0
		}
	case 6:
		subPacket1, subPacket2 := p.subPackets[0].getValue(), p.subPackets[1].getValue()
		if subPacket1 < subPacket2 {
			value = 1
		} else {
			value = 0
		}
	case 7:
		subPacket1, subPacket2 := p.subPackets[0].getValue(), p.subPackets[1].getValue()
		if subPacket1 == subPacket2 {
			value = 1
		} else {
			value = 0
		}
	default:
		// How did we get here? did the matrix break?
		log.Fatal("Unknown operatorType: ", p.operatorType)
	}
	return value
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

			data, _ := strconv.ParseUint(char, 16, 64)
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
	packet, _ := createPacket(binaryString)

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Version sum for part 1: ", packet.getVersionSum())

	start = time.Now()
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Operation result for part 2: ", packet.doOperation())
}
