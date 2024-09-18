/*
Copyright © 2024 MARIA NIRANJAN <s19835@sci.pdn.ac.lk>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace IP address details using the ipinfo API.",
	Long:  `The "trace" command allows you to trace and retrieve detailed information about an IP address using the ipinfo API. This command provides key insights such as the IP's geographical location, including country, city, and coordinates, as well as the Internet Service Provider (ISP) and organization details. It supports both IPv4 and IPv6 addresses. By leveraging the ipinfo API, the "trace" command ensures accurate and up-to-date information, making it a powerful tool for network diagnostics, security analysis, and general IP tracking.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trace called")
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

	// Here you will define your flags and configuration settings.
}

//	{
//	    "ip": "8.8.8.8",
//	    "hostname": "dns.google",
//	    "anycast": true,
//	    "city": "Mountain View",
//	    "region": "California",
//	    "country": "US",
//	    "loc": "37.4056,-122.0775",
//	    "org": "AS15169 Google LLC",
//	    "postal": "94043",
//	    "timezone": "America/Los_Angeles"
//	}

type IP struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func showData() {
	url := "https://ipinfo.io/8.8.8.8/json?token=d39218ee15f924"
	responseByte := getData(url)

	data := IP{}

	json.Unmarshal(responseByte, &data)
}

func getData(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		log.Println("Response not found")
	}

	responseByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println("unable to read the response")
	}

	return responseByte
}
