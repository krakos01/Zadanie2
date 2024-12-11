package main

import (
	"fmt"
  	"log"
  	"net/http"
  	"time"
  	"net"
)


// Declare global variable to hold current date and time
var currentDatetime string 


// Custom HTTP request handler
func handler(w http.ResponseWriter, r *http.Request) {
	ip := getServerIP()  // Get IP
	currentDatetime = time.Now().Format("02.01.2006 15:04:05")  // Set current date and time and format

	// Print text on a webpage when handled
	fmt.Fprintf(w, "Adres: %s!", ip)
	fmt.Fprintf(w, "\n%s", currentDatetime)
}


// Returns server IP
func getServerIP() string {

	// Get all interfaces
	interfaces, _ := net.Interfaces()
	
	for _, iface := range interfaces {
		// Skip disabled interfaces and loopbacks
		if iface.Flags&net.FlagUp == 0  || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		
		// Get all addresses of given interface
		addresses, _ := iface.Addrs()

		// Find not empty ipv4 address
		var ip net.IP
		for _, addr := range addresses {

			// Cover both cases: with and without mask
			switch v := addr.(type) {
				case *net.IPNet: ip = v.IP
				case *net.IPAddr: ip = v.IP
			}

			// Return first not null ipv4 address
			if ip != nil && ip.To4() != nil{
				return ip.String()
			}
		}
	}
	
	// Return empty string if address not found
	return ""
}



func main() {
	port := 8080
	address := fmt.Sprintf(":%d", port)  // Formatuje port na string ":8080" 
	author := "Dawid Krajewski"

	// Set current date and time
	currentDatetime = time.Now().Format("02.01.2006 15:04:05")

	// Log date, author and port
	fmt.Println("Data uruchomienia:", currentDatetime)
	fmt.Println("Autor:", author)
	fmt.Println("Port:", port)

	// Handle all request to root ("/")
	http.HandleFunc("/", handler)

	// Listen on port and log when unexpected error occurs
	log.Fatal(http.ListenAndServe(address, nil))
}

