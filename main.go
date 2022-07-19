package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/phantommachine/go-pmod-tc1/pmodtc1"
)

func CtoF(celsius float32) float32 {
	return celsius*(9.0/5.0) + 32.0
}

func main() {
	tc1 := pmodtc1.New("/dev/spidev1.0")
	close, err := tc1.Open()
	if err != nil {
		panic(err)
	}
	defer close()

	fmt.Printf("Temp %f\n", CtoF(tc1.ReadTemp()))

	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		var respBody TemperatureResponseBody
		respBody.Temperature = CtoF(tc1.ReadTemp())
		respBody.Scale = "fahrenheit"

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(respBody)
	})

	fmt.Println("serving @ 127.0.0.1:9991")
	http.ListenAndServe(":9991", nil)
}

type TemperatureResponseBody struct {
	Temperature float32 `json:"temperature"`
	Scale       string  `json:"scale"`
}
