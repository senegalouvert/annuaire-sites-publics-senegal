package geoloc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type GeoIP struct {
	Country string `json:"country"`
}

func LookupIP(host string) (string, error) {
	addrs, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String(), nil
		}
	}
	return "", fmt.Errorf("no A (IPv4) record found for host: %s", host)
}

func GetLocation(ip string, geo *GeoIP) error {
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &geo)
	if err != nil {
		return err
	}

	return nil
}
