package main

import (
	"fmt"
	"os"

	"strings"

	"errors"

	"github.com/vishvananda/netlink"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}
	action := os.Args[1]
	interfaceName := os.Args[2]

	link, err := netlink.LinkByName(interfaceName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	switch strings.ToLower(action) {
	case "add":
		err = add(link)
	case "del":
		err = del(link)
	case "show":
		err = show(link)
	case "replace":
		err = replace(link)
	default:
		printUsage()
		fmt.Printf("\n\nInvalid action %q\n", action)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("go-iptools [action] [network interface] [address to add]")
	fmt.Println("example: go-iptools add eth0 10.64.244.1/16")
	fmt.Println("available actions: add del show")
}

func add(link netlink.Link) error {
	if len(os.Args) != 4 {
		return errors.New("Incorrect number of arguments")
	}
	stringAddress := os.Args[3]
	addr, err := netlink.ParseAddr(stringAddress)
	if err != nil {
		return err
	}
	return netlink.AddrAdd(link, addr)
}

func del(link netlink.Link) error {
	if len(os.Args) != 4 {
		return errors.New("Incorrect number of arguments")
	}
	stringAddress := os.Args[3]
	addr, err := netlink.ParseAddr(stringAddress)
	if err != nil {
		return err
	}
	return netlink.AddrDel(link, addr)
}

func replace(link netlink.Link) error {
	if len(os.Args) != 4 {
		return errors.New("Incorrect number of arguments")
	}
	stringAddress := os.Args[3]
	addr, err := netlink.ParseAddr(stringAddress)
	if err != nil {
		return err
	}
	return netlink.AddrReplace(link, addr)
}

func show(link netlink.Link) error {
	addrs, err := netlink.AddrList(link, netlink.FAMILY_ALL)
	if err != nil {
		return err
	}
	for _, addr := range addrs {
		fmt.Println(addr)
	}
	return nil
}
