package main

import (
	"net"
	"net/http"
)

func summary(w http.ResponseWriter, r *http.Request){
	r.Header.Write(w)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		w.Write([]byte(""))
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				w.Write([]byte(ipnet.IP.String()))
				w.Write([]byte("\n"))
			}
		}
	}
	w.Write([]byte(""))
}

func main() {
	http.HandleFunc("/", summary)
	http.ListenAndServe(":8080", nil)
}
