package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	fmt.Println("[+] Let the scan begin... ")
	ports := 65535
	for i := 1; i <= ports; i++ {
		waitGroup.Add(1)
		go func(j int) {
			defer waitGroup.Done()
			address := fmt.Sprintf("127.0.0.1:%d", j)
			checkPort, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("[-] either port %d is not open, filtered, or an error occured. Err: (%s).\n", j, err)
				return
			}
			checkPort.Close()
			fmt.Printf("%d is open.\n", j)
		}(i)
	}
	waitGroup.Wait()
}
