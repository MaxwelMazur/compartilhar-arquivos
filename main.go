package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	ip, err := showIP()
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(os.Stdout)
	log.Println("Servindo arquivos na na url " + ip + ":8080")
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("files"))))
	http.ListenAndServe(":8080", nil)
}

func showIP() (string, error) {

	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {

		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()

		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()

			if ip == nil {
				continue // [nao ha endereco no ipv4]
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("Voce esta conectado a internet?")
}
