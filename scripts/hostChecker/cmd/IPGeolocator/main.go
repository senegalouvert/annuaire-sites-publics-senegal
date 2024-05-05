package main

import (
	"encoding/csv"
	"fmt"
	"hostChecker/pkg/geoloc"
	"os"
)

const (
	InputFilePath  = "data/non_teste.csv"
	OutputFilePath = "out/ip_non_senegal.csv"
	Country        = "Senegal"
)

type GeoIP struct {
	Country string `json:"country"`
}

func main() {
	// Compteurs pour le nombre total d'IPs et le nombre d'IPs non hébergées au Sénégal
	totalIPs := 0
	nonSenegalIPs := 0

	// Ouvrir le fichier CSV
	f, err := os.Open(InputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Lire le fichier CSV
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Créer un fichier CSV pour les IPs non hébergées au Sénégal
	outfile, err := os.Create(OutputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outfile.Close()
	writer := csv.NewWriter(outfile)
	defer writer.Flush()

	// Parcourir chaque ligne (chaque IP)
	for i, line := range lines {
		if i == 0 {
			// Ignorer la première ligne (ligne d'en-tête)
			continue
		}

		// Obtenir l'IP à partir de la deuxième colonne du fichier CSV
		ip := line[1]

		// Vérifier si l'IP est vide
		if ip == "" {
			fmt.Printf("L'IP à la ligne %d est vide. Ignorée.\n", i+1)
			continue
		}

		geo := geoloc.GeoIP{}
		err = geoloc.GetLocation(ip, &geo)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Si l'IP n'est pas hébergée au Sénégal, écrire l'IP dans le fichier CSV
		if geo.Country != Country {
			writer.Write([]string{line[0], ip, geo.Country})
			nonSenegalIPs++
		}
		totalIPs++
	}

	// À la fin du script, imprimer les statistiques
	fmt.Printf("Nombre total d'IPs testées : %d\n", totalIPs)
	fmt.Printf("Nombre d'IPs non hébergées au %s : %d\n", Country, nonSenegalIPs)
	if totalIPs > 0 {
		fmt.Printf("Pourcentage d'IPs non hébergées au %s : %.2f%%\n", Country, float64(nonSenegalIPs)/float64(totalIPs)*100)
	}
}
