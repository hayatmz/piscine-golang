# Quest02 - Premiers pas en Golang : affichage et conditions 🖨️🔁
Deuxième journée de la piscine Golang à Zone01 Rouen en juin 2023. Ils permettent de se familiariser avec les bases de Go, avec un focus sur l'affichage via runes, les boucles, conditions simples et fonctions.

## Objectifs
- Comprendre la syntaxe de base de Go (boucles, conditions, fonctions)
- Manipuler les runes pour afficher des caractères sans ```fmt```
- Travailler avec des conditions simples et des boucles ```for```
- Se familiariser avec la structure de projets Go avec un ```main.go```

## Contenu
Certains exercices sont des programmes indépendants dans leur propre dossier contenant un fichier ```main.go```:<br>
[printalphabet](./printalphabet) : affiche l'alphabet de ```a``` à ```z``` rune par rune.<br>
[printdigits](./printdigits) : affiches les chiffres de ```0``` à ```9``` rune par rune.<br>
[printreversealphabet](./printreversealphabet) : affiche l'alphabet en sens inverse (```z``` à ```a```). <br>
[isnegative.go](./isnegative.go) : fonction qui retourne ```true``` si un nombre est négatif.<br>
[printcomb.go](./printcomb.go) : affiche toutes les combinaisons possibles de trois chiffres en ordre croissant, séparés par des virgules, sans doubons et en utilisant uniquement ```z01.PrintRune```.

## Ce que j'ai appris
- Syntaxe de base de Go
- Affichage de runes
- Implémentation de fonctions simples et festion des valeurs booléennes
- Gestion des sorties formatées sans ```fmt```

## Comment exécuter
Assurez-vous d'être dans le dossier 'quest02' du dépôt cloné :<br>

```cd quest02```

Pour lancer un programme :<br>
```go run nom_du_fichier.go```<br>

**Ou** Place toi dans le dossier du programme :<br>
```cd nom_du_dossier```<br>
```go run main.go```<br>

Pour compiler et exécuter :<br>
```go build nom_du_fichier.go```<br>
```./nom_du_fichier```