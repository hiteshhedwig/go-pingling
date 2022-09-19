package pinger

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/go-ping/ping"
)

func IsValidateAddr(addr string) bool {
	strs := strings.Split(addr, ".")
	if len(strs) == 4 {
		return true
	}
	return false
}

func GetGateway(addr string) string {
	strs := strings.Split(addr, ".")
	strs[len(strs)-1] = "1"
	return strings.Join(strs, ".")
}

func PingInit(addr string) error {

	if !IsValidateAddr(addr) {
		return fmt.Errorf("invalid addr: %s", addr)
	}

	pinger, err := ping.NewPinger(addr)
	if err != nil {
		fmt.Printf("error creating ping: %v", err)
		return err
	}

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%vms\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt.Milliseconds())
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)

	}

	pinger.Count = 10
	pinger.Run() // blocks until finished
	return nil
}

func SinglePing(addr string, sip *ScanIP) {

	defer sip.wg.Done()

	if sip.idx >= 255 {
		return
	}

	sip.idx += 1

	//fmt.Println("addrr ", addr)
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		fmt.Printf("error creating ping: %v", err)
	}

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%vms \n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt.Milliseconds())
	}
	pinger.Count = 1
	pinger.Timeout = time.Duration(1100 * time.Millisecond)

	pinger.Run()

}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func PingSearch(addr string) error {
	if !IsValidateAddr(addr) {
		return fmt.Errorf("invalid addr: %s", addr)
	}

	ip := GetGateway(addr)
	newscan := NewScanIP(0, ip)
	for newscan.Index() < 255 {
		go newscan.Next()
	}

	newscan.Finish()
	return nil
}
