# Quest07 - Slices, chaînes et parsing 🧩

## Objectifs
- Comprendre le fonctionnement des slices en Go.
- Manipuler des chaînes de caractères, les découper, les concaténer.
- Utiliser des strucutres simples et l'indexation.
- Développer les fonctions qui imisent ou remplacent des fonctions standars (```make```, ```append```).

## Contenu
[appendrange.go](./appendrange.go) : Crée dynamiquement une slice d'entiers entre deux bornes sans utiliser ```make```.<br>
[makerange.go](./makerange.go) : Crée une slice d'entiers entre deux bornes sans utiliser  ```append```.<br>
[concatparams.go](./concatparams.go) : Concatène tous les arguments reçus dans un seul string, séparés par des sauts de ligne.<br>
[splitwhitespaces.go](./splitwhitespaces.go) : Sépare une chaîne de caractères en mots en supprimant les espaces, tabulations, etc.<br>
[printwordstables.go](./printwordstables.go) : Affiche une slice de chaînes de caractères ligne par ligne.<br>

## Ce que j'ai appris
- Créer et manipuler des slices sans utliser ```make``` ni ```append```.
- Traiter des chaînes de caractères pour les découper ou les assembler.
- Approfondir l'utilisation de boucles ```for```, du mot-clé ```range```, et de la gestion des index.
- Gérer les entrées utilisateur passées en ligne de commande.
- Simuler des comportements similaires à ```strings.Split```, ```strings.Join```, etc.

## Comment exécuter
Assurez-vous d'être dans le dossier 'quest07' du dépôt cloné :<br>

```cd quest07```

Pour lancer un programme :<br>
```go run nom_du_fichier.go```<br>

**Ou** Place toi dans le dossier du programme :<br>
```cd nom_du_dossier```<br>
```go run main.go```<br>

Pour compiler et exécuter :<br>
```go build nom_du_fichier.go```<br>
```./nom_du_fichier```