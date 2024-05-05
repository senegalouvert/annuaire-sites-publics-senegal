package main

import (
	"encoding/csv"
	"fmt"
	"hostChecker/pkg/geoloc"
	"os"
)

const (
	InputFilePath  = "../../data/annuaire.csv"
	OutputFilePath = "out/host_non_senegal.csv"
	FailedFilePath = "out/host_non_teste.csv"
	Country        = "Senegal"
)

func main() {
	// Compteurs pour le nombre total de sites, le nombre de sites non hébergés au Sénégal et le nombre de sites non testés
	totalSites := 0
	nonSenegalSites := 0
	untestedSites := 0

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

	// Créer un fichier CSV pour les sites Web non hébergés au Sénégal
	outfile, err := os.Create(OutputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outfile.Close()
	writer := csv.NewWriter(outfile)
	defer writer.Flush()

	// Créer un fichier CSV pour les sites Web non testés
	failedfile, err := os.Create(FailedFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer failedfile.Close()
	failedWriter := csv.NewWriter(failedfile)
	defer failedWriter.Flush()

	// Parcourir chaque ligne (chaque site Web)
	for _, line := range lines {
		// Obtenir l'URL du site Web à partir de la deuxième colonne du fichier CSV
		host := line[1]

		ip, err := geoloc.LookupIP(host)
		if err != nil {
			fmt.Println(err)
			failedWriter.Write([]string{host, ""})
			untestedSites++
			continue
		}

		geo := geoloc.GeoIP{}
		err = geoloc.GetLocation(ip, &geo)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Si le site Web n'est pas hébergé au Sénégal, écrire l'URL dans le fichier CSV
		if geo.Country != Country {
			writer.Write([]string{host, geo.Country})
			nonSenegalSites++
		}
		totalSites++
	}

	// À la fin du script, imprimer les résultats
	fmt.Printf("Nombre total de sites testés : %d\n", totalSites)
	fmt.Printf("Nombre de sites non hébergés au %s : %d\n", Country, nonSenegalSites)
	fmt.Printf("Nombre de sites non testés : %d\n", untestedSites)
	if totalSites > 0 {
		fmt.Printf("Pourcentage de sites non hébergés au %s : %.2f%%\n", Country, float64(nonSenegalSites)/float64(totalSites)*100)
	}
}
