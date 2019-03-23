package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// map JSON keys

type Sip struct {
	AddressOfRecord string `json:"addressOfRecord"`
	URI             string `json:"uri"`
	Content         string `json:"content"`
	TenantID        string `json:"tenantid"`
	Source          string `json:"source"`
	Target          string `json:"target"`
	UserAgent       string `json:"useragent"`
	RawUserAgent    string `json:"rawuseragent"`
	Created         string `json:"created"`
	LineID          string `json:"lineid"`
}

type SipDumps []Sip

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse and load file in memory
		http.ServeFile(w, r, "../AOR/serverContent/SIP_Stack.json")

	})

	myRouter.HandleFunc("/sipRegistrations/{SipDumps}", returnOneSip).Methods("GET")

	myRouter.HandleFunc("/sipRegistrations", returnAll).Methods("GET")

	fmt.Println("Server Running!")
	log.Fatal(http.ListenAndServe(":3500", myRouter))

}

func returnOneSip(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	value := variables["AddressOfRecord"]

	fmt.Fprintf(w, "Value: "+value)

}

func returnAll(w http.ResponseWriter, r *http.Request) {
	sips := SipDumps{
		Sip{AddressOfRecord: "0142e2fa3543cb32bf000100620002", TenantID: "0127d974-f9f3-0704-2dee-000100420001", URI: "sip:0142e2fa3543cb32bf000100620002@109.149.135.172;jbcuser=cpe70",
			Source: "29.211.204.173:19622", Target: "60.124.57.147:5061", UserAgent: "polycom.vvx.600", RawUserAgent: "PolycomVVX-VVX_600-UA/132.244.41.145", Created: "2016-12-12T22:40:40.764Z", LineID: "013db2ba-2175-6d29-6157-000100620002"},

		Sip{AddressOfRecord: "0146a51532d4fdb52d000100620002", TenantID: "0127d974-f9f3-0704-2dee-000100420001", URI: "sip:0146a51532d4fdb52d000100620002@33.186.185.199;jbcuser=cpe70",
			Source: "96.171.8.24:19622", Target: "121.236.199.226:5061", RawUserAgent: "PolycomVVX-VVX_500-UA/152.17.17.73", Created: "2016-12-12T22:42:03.538Z", LineID: "0146a514-e123-9bee-452d-000100620002"},
	}
	fmt.Println("This Endpoint returns all SIPs")

	json.NewEncoder(w).Encode(sips)
}
