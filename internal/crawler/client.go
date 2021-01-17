package crawler

import (
	"errors"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/buger/jsonparser"
	"github.com/tupyy/stock/internal/config"
	"github.com/tupyy/stock/models"
)

func createUrl(baseUrl, company string) string {
	return fmt.Sprintf("%s?symbol=1rP%s&period=-1", baseUrl, company)
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

func getStock(client *http.Client, company string) (models.StockValue, error) {

	baseUrl := config.GetServerBaseUrl()
	if len(baseUrl) == 0 {
		return models.StockValue{}, errors.New("server base url not  found")
	}

	data, err := doRequest(client, createUrl(baseUrl, company))
	if err != nil {
		return models.StockValue{}, err
	}

	s := models.StockValue{Label: company}
	if val, err := jsonparser.GetFloat(data, "d", "[0]", "c"); err == nil {
		s.Value = val
	}

	if val, err := jsonparser.GetFloat(data, "d", "[0]", "var"); err == nil {
		s.Variation = val
	}

	return s, nil
}
