package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
)

const (
	api = "https://www.boursorama.com/bourse/action/graph/ws/UpdateCharts"
)

type StockValue struct {
	Value     float64
	Variation float64
}

func createUrl(company string) string {
	return fmt.Sprintf("%s?symbol=1rP%s&period=-1", api, company)
}

func doRequest(client *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func getStock(client *http.Client, company string) (StockValue, error) {
	data, err := doRequest(client, createUrl(company))
	if err != nil {
		return StockValue{}, err
	}

	s := StockValue{}
	if val, err := jsonparser.GetFloat(data, "d", "[0]", "c"); err == nil {
		s.Value = val
	}

	if val, err := jsonparser.GetFloat(data, "d", "[0]", "var"); err == nil {
		s.Variation = val
	}

	return s, nil
}

func main() {

	client := &http.Client{}
	s, _ := getStock(client, "RNO")
	fmt.Printf("%2.2f,%1.2f", s.Value, s.Variation)

}
