# Projet de vérification de l'emplacement d'hébergement des sites Web

Ce projet contient deux scripts en Go qui vérifient l'emplacement d'hébergement de différents sites Web à partir d'un fichier CSV. Les scripts utilisent une API de géolocalisation d'adresse IP pour déterminer l'emplacement d'hébergement de chaque site Web.

## Structure du projet

- `pkg/geoloc` : Contient les fonctions `LookupIP` et `GetLocation` qui sont utilisées par les deux scripts.
- `cmd/WebHostGeolocator` : Le premier script qui utilise `LookupIP` et `GetLocation` pour vérifier l'emplacement d'hébergement des sites Web à partir de leur URL.
- `cmd/IPGeolocator` : Le deuxième script qui utilise seulement `GetLocation` pour vérifier l'emplacement d'hébergement des sites Web à partir de leur adresse IP.
- `out/` : Le dossier avec les fichiers CSV généré par les scripts.

## Comment exécuter les scripts

1. Assurez-vous d'avoir Go installé sur votre machine.
2. Ouvrez un terminal et naviguez jusqu'au répertoire du projet.

3. Pour exécuter le premier script, utilisez la commande suivante :

```bash
go run cmd/WebHostGeolocator/main.go
```

4. Pour exécuter le deuxième, utilisez la commande suivante :

```bash
go run cmd/IPGeolocator/main.go
```

Ces commandes doivent être exécutées à partir du répertoire racine. Elles compileront et exécuteront le script spécifié.
