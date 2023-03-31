package angelone

import (
	"crypto/tls"
	"net/http"
	"testing"
)

func TestClient_Instruments(t *testing.T) {
	t.Parallel()
	c := &Client{}
	c.SetHTTPClient(&http.Client{
		Timeout:   requestTimeout,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	})
	res, err := c.GetInstruments()
	if err != nil {
		t.Errorf("Error while fetching instruments. %v", err)
	}

	if len(res) == 0 {
		t.Errorf("instruments length is zero")
	}
}

func TestClient_Instrument(t *testing.T) {
	t.Parallel()
	c := &Client{}
	c.SetHTTPClient(&http.Client{
		Timeout:   requestTimeout,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	})
	res, err := c.GetInstrument("BANKNIFTY20APR2338700PE")
	if err != nil {
		t.Errorf("Error while fetching instrument. %v", err)
	}

	if res != nil && res.Token != "" {
		t.Errorf("instrument not found")
	}
}
