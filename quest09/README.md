# Quest09 - Fonctions comme valeurs (fonctionnelles) 🔁📦

## Objectifs
- Apprendre à utiliser des fonctions comme paramètres.
- Appliquer des concepts de programmation fonctionnelle en Go.
- Maîtriser le traitement de slices à l'aide de fonctions comme ```Map```, ```ForEach```, ```Any```, etc.
- Implémenter des fonctions génériques adaptées à divers cas logiques.

## Contenu
[foreach](./foreach.go) : applique une fonction ```f(int)``` à chaque élément d'une slice d'entiers.<br>
[map.go](./map.go) : Retourne une nouvelle slice de booléens résultant de l'application d'une fonction ```f(int) bool``` à chaque élément d'une slice de int.<br>
[any.go](./any.go) : Renvoie ```true``` si au moins un élément d'une slice de string satisfait une condition donnée.<br>
[countif.go](./countif.go) : Compte le nombre de string dans une slice qui remplissent une certaine condition.<br>
[issorted.go](./issorted.go) : Vérifie si une slice d'entiers est triée, selon une fonction de comparaison personnalisée.<br>


## Ce que j'ai appris
- Passer des fonctions en tant qu'arguments, et à les utiliser dans des boucles.
- Construire des slices de retour dynamiquement avec ```append```.
- Implémenter des fonctions de traitement générique (type ```map```, ```filter```, ```any```, etc.).
- Détecter un ordre de tri via une fonction de comparaison.

## Comment exécuter
Assurez-vous d'être dans le dossier 'quest09' du dépôt cloné :<br>

```cd quest09```

Pour lancer un programme :<br>
```go run nom_du_fichier.go```<br>

**Ou** Place toi dans le dossier du programme :<br>
```cd nom_du_dossier```<br>
```go run main.go```<br>

Pour compiler et exécuter :<br>
```go build nom_du_fichier.go```<br>
```./nom_du_fichier```