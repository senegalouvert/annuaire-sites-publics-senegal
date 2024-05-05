package main

import (
	"hostChecker/pkg/geoloc"
	"testing"
)

func TestLookupIP(t *testing.T) {
	// Test avec un site Web connu
	host := "google.com"
	ip, err := geoloc.LookupIP(host)
	if err != nil {
		t.Errorf("Erreur lors de la recherche de l'IP pour %s: %v", host, err)
	}
	if ip == "" {
		t.Errorf("Aucune IP trouvée pour %s", host)
	}
}

func TestGetLocation(t *testing.T) {
	// Test avec une adresse IP connue
	ip := "8.8.8.8" // Adresse IP de Google DNS
	geo := geoloc.GeoIP{}
	err := geoloc.GetLocation(ip, &geo)
	if err != nil {
		t.Errorf("Erreur lors de la récupération de l'emplacement pour %s: %v", ip, err)
	}
	if geo.Country == "" {
		t.Errorf("Aucun pays trouvé pour %s", ip)
	}
}
