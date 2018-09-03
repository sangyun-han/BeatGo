/* 
   Refer : Packet Capture, Injection, and Analysis with Gopacket
   https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
*/

package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

//var (
////	device       string = "enp4s0"
////	snapshot_len int32  = 1024
////	promiscuous  bool   = false
////	err          error
////	timeout      time.Duration = 5 * time.Second
////	handle       *pcap.Handle
////)

func main() {
	// Open device for live capture
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	var filter string = "tcp and port 80"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only capturing TCP port 80 packets.")

	// Use the handle as a packet source to process all packets
	packetSrc := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSrc.Packets() {
		fmt.Println(packet)
	}
}
