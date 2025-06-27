# Quest08 - Structs, fichiers et adaptation de code 📦📄

## Objectifs
- Comprendre et manipuler les structs en Go.
- Corriger et adapter des programmes existants pour qu'ils compilent et respectent les consignes.
- Gérer les arguments de la ligne de commande.
- Lire le contenu d'un fichier passé en paramètre.
- Approfondir l'usage de pointeurs, de types personnalisés et de l'affichage sans fmt.

## Contenu
Ces exercices sont des programmes indépendants dans leur propre dossier contenant un fichier ```main.go```:<br>
[boolean](./boolean/main.go) : Corriger un programme utilisant une structure booléenne personnalisée pour qu'il affiche si le nombre d'arguments est pair ou impair.<br>
[point](./point/main.go) : Corriger une structure ```Point``` avec deux champs ```int```, l'instancier avec des coordinnées et l'afficher au bon format.<br>
[displayfile](./displayfile/main.go) : Lire et afficher le contenu d'un fichier passé en argument en gérant les erreurs courantes: aucun argument, trop d'arguments, fichier introuvable.<br>

## Ce que j'ai appris
- Déclarer et utiliser une structure personnalisée.
- Manipuler des pointeurs vers des strucs pour modifier leurs champs.
- Lire des fichiers avec ```os.ReadFile```, et traiter les erreurs de manière propre.
- Gérer les arguments de la ligne de commmande et vérifier leur nombre.

## Comment exécuter
Assurez-vous d'être dans le dossier 'quest08' du dépôt cloné et dans le dossier du programme souhaité :<br>

```cd quest08```<br>
```cd nom_du_dossier```<br>

Pour lancer un programme :<br>
```go run main.go```<br>

Pour **displayfile** :<br>
```go run main.go fichier.txt```<br>

**Ou** pour compiler et exécuter :<br>
```go build nom_du_fichier.go```<br>
```./nom_du_fichier```

