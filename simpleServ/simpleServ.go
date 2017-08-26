package main

import (
	"fmt"
	"net/http"
	"net"
)

func handler(w http.ResponseWriter, r *http.Request) {
	l, err := net.Interfaces()
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
	fmt.Println(w)
	for _, f := range l {
		fmt.Fprintf(w, "Name %s\n", f.Name)
		fmt.Fprintf(w, "MTU: %v\n", f.MTU)
		fmt.Fprintf(w, "HardwareAddr: %v\n", f.HardwareAddr)
		fmt.Fprintf(w, "Flags: %v\n\n", f.Flags)
	}
}

func main() {

	http.HandleFunc("/NetInterfaces/", handler)
	http.ListenAndServe(":2000", nil)

}
