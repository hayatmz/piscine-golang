# Quest05 - Manipulation avancée de chaînes de caractères. 🔤🧪

## Objectifs
- Approfondir la festion des chaînes de caractères en Go.
- S'exercer à manipuler les runes (caractères Unicode) uniquement.
- Implémenter des fonctions utilitaires sur les chaînes sans utiliser les fonctions de la bibliothèque standard.
- Travailler la logique de parsing, de validation et de transformation textuelle.

## Contenu
[firstrune.go](./firstrune.go) : Renvoie la première rune d'une chaîne de caractères.<br>
[lastrune.go](./lastrune.go) : Renvoie la dernière rune d'une chaîne de caractères.<br>
[nrune.go](./nrune.go) : Renvoie la rune à l'index donné.<br>
[compare.go](./compare.go) : Compare deux chaînes lexicographiquement.<br>
[alphacount.go](./alphacount.go) : Compte le nombre de lettres latines dans une chaîne.<br>
[index.go](./index.go) : Renvoie la position d'une sous-chaîne dans une chaîne.<br>
[concat.go](./concat.go) : Concatène deux chaînes.<br>
[isupper.go](./isupper.go) : Vérifie si une chaîne ne contient que des majuscules.<br>
[islower.go](./islower.go) : Vérifie si une chaîne ne contient que des minuscules.<br>
[isalpha.go](./isalpha.go) : Vérifie si une chaîne est uniquement alphabétique.<br>
[isnumeric.go](./isnumeric.go) : Vérifie si une chaîne contient uniquement des chiffres.<br>
[isprintable.go](./isprintable.go) : Vérifie si tous les caractères sont imprimables.<br>
[toupper.go](./toupper.go) : Convertit une chaîne en majuscules.<br>
[tolower.go](./tolower.go) : Convertit une chaîne en minuscules.<br>
[printnbrinorder.go](./printnbrinorder.go) : Affiche les chiffres d'un entier dans l'ordre croissant.<br>
[trimatoi.go](./trimatoi.go) : Extrait un nombre entier d'une chaîne en ignorant les caractères inutiles.<br>
[capitalize.go](./capitalize.go) : Met une majuscule la première lettre de chaque mot, les autres en minuscules.<br>

## Ce que j'ai appris
- Manipulation de chaînes et runes en Go.
- Gestion des cas limites (chaînes vides, caractères spéciaux, etc.).
- Implémentation de comportements proches de fonctions de la bibliothèque standard.
- Utilisation des boucles, conditions et manipulation de types ```rune``` et ```byte```.
- Compréhension des conversions entre types (```string```, ```[]rune```, etc.).
- Approche manuelle du parsing de texte, sans dépendances externes.

## Comment exécuter
Assurez-vous d'être dans le dossier 'quest05' du dépôt cloné :<br>

```cd quest05```

Pour lancer un programme :<br>
```go run nom_du_fichier.go```<br>

**Ou** Place toi dans le dossier du programme :<br>
```cd nom_du_dossier```<br>
```go run main.go```<br>

Pour compiler et exécuter :<br>
```go build nom_du_fichier.go```<br>
```./nom_du_fichier```