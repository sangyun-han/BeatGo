package main

// example
// 2019/07/25 14:49:10 95.76 MBit/s
// 2019/07/25 14:49:10 Invalid (too small) IP header length (3 < 5)
// 2019/07/25 14:49:11 95.76 MBit/s
// 2019/07/25 14:49:12 95.75 MBit/s
// 2019/07/25 14:49:12 Invalid (too small) IP header length (3 < 5)

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/afpacket"
	"github.com/google/gopacket/layers"
)

var rx int64

func printRxStats() {
	t := time.NewTicker(1 * time.Second)

	lastTime := time.Now()
	var lastValue int64 = 0
	for _ = range t.C {
		now := time.Now()
		tmp := rx
		delta := now.Sub(lastTime).Seconds()
		if delta == 0 {
			continue
		}

		d := float64(tmp-lastValue) / delta
		log.Printf("%.2f MBit/s", 8*d/(1024*1024))

		lastValue = tmp
		lastTime = now
	}
}

func main() {
	inputAFpacket, err := afpacket.NewTPacket(afpacket.OptInterface("ens10"))
	if err != nil {
		log.Fatal(err)
	}
	defer inputAFpacket.Close()

	go printRxStats()

	for {
		pkt, _, err := inputAFpacket.ZeroCopyReadPacketData()
		if err != nil {
			log.Fatal(err)
		}

		atomic.AddInt64(&rx, int64(len(pkt)))

		ether := layers.Ethernet{}
		if err = ether.DecodeFromBytes(pkt, gopacket.NilDecodeFeedback); err != nil {
			log.Println(err)
		}
		if ether.EthernetType != layers.EthernetTypeIPv4 {
			continue
		}
		ipv44 := layers.IPv4{}
		if err = ipv44.DecodeFromBytes(ether.Payload, gopacket.NilDecodeFeedback); err != nil {
			log.Println(err)
			continue
		}
		udpp := layers.UDP{}
		if err = udpp.DecodeFromBytes(ipv44.Payload, gopacket.NilDecodeFeedback); err != nil {
			log.Println(err)
		}
	}
}
