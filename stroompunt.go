package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"

	"github.com/micmonay/keybd_event"
)

var (
	//go:embed static/*
	static embed.FS
)

func handleKeyLeft(w http.ResponseWriter, r *http.Request) {
	handleKeyInput(w, r, "left")
}

func handleKeyRight(w http.ResponseWriter, r *http.Request) {
	handleKeyInput(w, r, "right")
}

func handleKeyInput(w http.ResponseWriter, r *http.Request, direction string) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	if direction == "left" {
		kb.SetKeys(keybd_event.VK_LEFT)
	} else {
		kb.SetKeys(keybd_event.VK_RIGHT)
	}

	err = kb.Launching()
	if err != nil {
		panic(err)
	}
}

func getIpAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}

	log.Fatal("IP address could not be detected. Make sure you are connected to a network and try again.")
	return ""
}

func main() {
	fmt.Println("------------------------------------------")
	fmt.Println("                STROOMPUNT                ")
	fmt.Println("------------------------------------------")
	fmt.Println("Connect your phone to the same network as")
	fmt.Println("this computer and open the following URL:")
	fmt.Println(getIpAddress() + ":1234")

	fSys, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}

	http.Handle("/", http.FileServer(http.FS(fSys)))
	http.HandleFunc("/left", handleKeyLeft)
	http.HandleFunc("/right", handleKeyRight)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
