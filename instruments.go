package angelone

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Instruments -

// Instrument -
type Instrument struct {
	Token          string `json:"token"`
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Expiry         string `json:"expiry"`
	Strike         string `json:"strike"`
	Lotsize        string `json:"lotsize"`
	Instrumenttype string `json:"instrumenttype"`
	ExchSeg        string `json:"exch_seg"`
	TickSize       string `json:"tick_size"`
}

var dateFormat = "02-01-2006"

// GetInstruments - get instruments map
func (c *Client) GetInstruments() (map[string]Instrument, error) {
	today := time.Now()
	todayStr := today.Format(dateFormat)
	if todayStr == c.date && len(c.instruments) > 0 {
		fmt.Println("cached instruments")
		return c.instruments, nil
	}
	res, err := c.httpClient.Do(http.MethodGet, URIInstrument, nil, nil)
	if err != nil {
		return nil, err
	}
	if res.Response.StatusCode != 200 {
		return nil, errors.New("response is not success")
	}
	var instruments []Instrument
	err = json.Unmarshal(res.Body, &instruments)
	if err != nil {
		return nil, err
	}
	m := map[string]Instrument{}
	for _, v := range instruments {
		m[v.Symbol] = v
	}
	c.instruments = m
	c.date = todayStr
	return m, nil
}

// GetInstrument - get instrument by symbol
func (c *Client) GetInstrument(symbol string) (*Instrument, error) {
	res, err := c.GetInstruments()
	if err != nil {
		return nil, err
	}
	instrument := res[symbol]
	return &instrument, nil
}
