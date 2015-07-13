package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/soniah/gosnmp"
	"net"
	"os"
	"path/filepath"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("   %s [-community=<community>] host\n", filepath.Base(os.Args[0]))
		fmt.Printf("     host      - the host to scan\n\n")
		flag.PrintDefaults()
	}

	var community string
	flag.StringVar(&community, "community", "public", "the community string for device")

	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	target := flag.Args()[0]

	gosnmp.Default.Target = target
	gosnmp.Default.Community = community
	gosnmp.Default.Timeout = time.Duration(10 * time.Second) // Timeout better suited to walking
	err := gosnmp.Default.Connect()
	if err != nil {
		fmt.Printf("Connect err: %v\n", err)
		os.Exit(1)
	}
	defer gosnmp.Default.Conn.Close()

	oid := "1.3.6.1.2.1.4.22.1.2"
	err = gosnmp.Default.BulkWalk(oid, printValue)
	if err != nil {
		fmt.Printf("Walk Error: %v\n", err)
		os.Exit(1)
	}
}

func printValue(pdu gosnmp.SnmpPDU) error {
	fmt.Printf("IP: %s MAC:", pdu.Name[24:])
	names, _ := net.LookupAddr(pdu.Name[24:])
	b := hex.EncodeToString(pdu.Value.([]byte))
	mac := b[:2] + ":" + b[2:4] + ":" + b[4:6] + ":" + b[6:8] + ":" + b[8:10] + ":" + b[10:12]
	fmt.Printf("%s\n", mac)
	fmt.Println(names)
	return nil
}
