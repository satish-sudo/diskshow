package main

import (
	"fmt"
	"github.com/jaypipes/ghw"
	"strings"
)

func main() {
	block, err := ghw.Block()
	if err != nil {
		fmt.Printf("Error getting block storage info: %v", err)
	}

	//fmt.Printf("Block: %v\n", block)

	var vmdisks []string
	vgdisks := []string {"c20a22d0-5169-47b0-86e4-9c44a5c6ae9f", "869f3bdc-2579-4d45-acdc-076932b7c138", "741315f1-9c57-4899-bb20-a6fb789efbc5"}

    fmt.Println("Printing disk")
	
	for _, disk := range block.Disks {
		fmt.Printf("Disk: %v \n", disk)
	}
	
	for i, vgdisk := range vgdisks {
		fmt.Printf("Outer: Checking with disk[%d]: %s \n", i, vgdisk)
		vgdisk = strings.ReplaceAll(vgdisk, "-", "_") 
		for j, disk := range block.Disks {		
			fmt.Printf("  Inner: Checking with disk[%d]: %s \n", j, disk.SerialNumber)
			if strings.Contains(disk.SerialNumber, vgdisk) {
				fmt.Println("Found match, appending...")
				vmdisks = append(vmdisks, disk.Name)
				break;
			}
		}
	}
	
	fmt.Println("Printing sorted disks")
	for _, disk := range vmdisks {
		fmt.Printf("%v\n", disk)
	}
}