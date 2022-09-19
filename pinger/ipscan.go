package pinger

import (
	"fmt"
	"strings"
	"sync"
)

type ScanIP struct {
	idx  int
	addr string
	wg   *sync.WaitGroup
}

func NewScanIP(idx int, addr string) *ScanIP {
	scp := &ScanIP{idx: idx, addr: addr}
	scp.wg = &sync.WaitGroup{}
	return scp
}

func (s ScanIP) Index() int {
	return s.idx
}

func (s ScanIP) Addr() string {
	return s.addr
}

func (s *ScanIP) Next() {
	ip := s.FormIPAddr()
	s.wg.Add(1)
	go SinglePing(ip, s)
}

func (s *ScanIP) Update() {
	s.idx += 1
}

func (s *ScanIP) Finish() {
	s.wg.Wait()
}

func (s *ScanIP) FormIPAddr() string {
	sample := strings.Split(s.addr, ".")
	sample = sample[:len(sample)-1]
	addr := fmt.Sprintf("%s.%d", strings.Join(sample, "."), s.idx)
	return addr
}
