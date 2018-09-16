/* 
   Refer : Packet Capture, Injection, and Analysis with Gopacket
   https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
*/


package main

import (
	"github.com/google/gopacket/pcap"
	"log"
	"github.com/google/gopacket"
	"fmt"
	"time"
)

func main() {
	var (
		device       string = "enp4s0"
		snapshotLen int32  = 1024
		promiscuous  bool   = false
		err          error
		timeout      time.Duration = 5 * time.Second
		handle       *pcap.Handle
	)
	// Open device for live capture
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSrc := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSrc.Packets() {
		fmt.Println(packet)
	}
}
