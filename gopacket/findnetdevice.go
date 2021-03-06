/* 
   Refer : Packet Capture, Injection, and Analysis with Gopacket
   https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
*/
package main

import (
	"fmt"
	"log"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Find all devices including network devices and USB devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Devices found : ")
	for _, device := range devices {
		fmt.Println("\nName : ", device.Name)
		fmt.Println("Description : ", device.Description)
		fmt.Println("Devices addresses : ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP address : ", address.IP)
			fmt.Println("- Subnet mask : ", address.Netmask)
		}
	}
}
