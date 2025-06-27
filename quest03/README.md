# Quest03 - Pointeurs, chaînes et conversions de types 🧠🪢

## Objectifs
- Comprendre le fonctionnement des pointeurs en Go : manipulation directe de la mémoire.
- Apprendre à modifier des viaraibles via des pointeurs simples ou multiples.
- Savoir parcourir et manipuler des chaînes de caractères rune par rune.
- Réaliser des conversions entre chaînes et entiers (```string``` → ```int```) de façon manuelle.
- Approfondir la logique et la rigueur dans la manipulation de types simples.

## Contenu
[pointone.go](./pointone.go) : Modifie une variable via un pointeur.<br>
[ultimatepointone.go](./ultimatepointone.go) : Modifie une variable via trois niveaux de pointeurs.<br>
[divmod.go](./divmod.go) : Division avec reste en utilisant des pointeurs.<br>
[ultimatedivmod.go](./ultimatedivmod.go) : Division avec reste en utilisant des pointeurs.<br>
[printstr.go](./printstr.go) : Affiche les caractères d'une chaîne un par un.<br>
[strlen.go](./strlen.go) : Renvoie le nombre de runes dans une chaîne.<br>
[swap.go](./swap.go) : Echange deux entiers via leurs pointeurs.<br>
[strrev.go](./strrev.go) : Inverse une chaîne de caractères.<br>
[basicatoi.go](./basicatoi.go) : Convertit une chaîne valide de chiffres en int.<br>
[basicatoi2.go](./basicatoi2.go) : Convertit une chaîne en int, ignore les caratères non numériques.<br>
[atoi.go](./atoi.go) : Convertit une chaîne en int avec gestion des signes + et -.<br>

## Ce que j'ai appris
- Manipulation de pointeurs simples et multiples en Go (```*int```, ```**int```, etc.).
- Gestion manuelle des chaînes de caractères, rune par rune, sans ```fmt```.
- Construction de fonctions utilitaires pour lire ou écrire du texte sans bibliothèque externe.
- Conversion manuelles de ```string``` en ```int``` sans utiliser ```strconv.Atoi```.
- Rigueur dans la gestion des cas limites.
- Approche algorithmique bas niveau pour mieux comprendre le fonctionnement du langage.

## Comment exécuter
Assurez-vous d'être dans le dossier 'quest03' du dépôt cloné :<br>

```cd quest03```

Pour lancer un programme :<br>
```go run nom_du_fichier.go```<br>

**Ou** Place toi dans le dossier du programme :<br>
```cd nom_du_dossier```<br>
```go run main.go```<br>

Pour compiler et exécuter :<br>
```go build nom_du_fichier.go```<br>
```./nom_du_fichier```