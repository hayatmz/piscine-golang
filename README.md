# Piscine Golang 🏊‍♂️

Ce dépôt contient l'ensemble des projets réalisés durant ma "piscine Golang" à Zone01 Rouen en juin 2023. Une immersion intensive de plusieurs semaines pour apprendre les bases de la programmation avec le langage Golang.

> 📌 **Important :** Tous les exercices et consignes étaient en anglais, ce qui a également renforcé mes compétences en compréhension technique de l'anglais écrit.

## Objectifs

-   Apprendre les fondamentaux de la programmation (logique, conditions, boucles, fonctions, etc.)
-   Se familiariser avec Go.
-   Développer l'autonomie, la rigueur et la capacité à collaborer.

## Table des matières

-   [Quest01 - Premiers pas avec Bash](./quest01)
-   [Quest02 - Premiers pas en Go: affichage et conditions](./quest02)
-   [Quest03 - Pointeurs, chaînes et conversions de types en Go](./quest03)
-   [Quest04 - Fonctions récursives et calculs mathématiques en Go](./quest04)
-   [Quest05 - Manipulations avancées de chaînes de caractères](./quest05)
-   [Quete06 - Arguments de ligne de commande en Go](./quest06)
-   [Quest07 - Slices, chaînes et parsing](./quest07)
-   [Quest08 - Structs, fichiers et adaptation de code](./quest08)
-   [Quest09 - Fonctions comme valeurs (fonctionnelles)](./quest09)

## Ce que j'y ai appris

- Le travail en autonomie, la résolution de problèmes et la collaboration entre pairs
- L'utilisation de Bash et des commandes Unix pour automatiser des tâches
-   La syntaxe et les spécificités du langage Go (typage, pointeurs, slices, structs, etc.)
-   Les bases de l'algorithmie (conditions, boucles, récursivité, recherche, tri)
-   La manipulation de chaîne, de fichiers, et d'arguments en ligne de commande
- La logique de programmation fonctionnelle avec les fonctions d'ordre supérieur

## Installation

Pour explorer ou tester les projets localement :

1. Assure toi d'avoir Go installé sur ta machine. Tu peux vérifier avec :
```
go version
```

Si besoin, [installe Golang](https://go.dev/doc/install).

2. **Cloner le dépôt** :

```
git clone https/..github.com/hayatmz/piscine-go.git
```
Et rends toi dans le dossier.
```
cd piscine-go
```

> Tous les projets utilisent uniquement la bibliothèque standard de Go, **à l'exception de la bibliothèque pédagogique** [github.com/01-edu/z01](./github.com/01-edu/z01) utilisée pour afficher des caractères rune par rune, nous interdisant d'utiliser ```fmt```.

Pour l'installer, exécute cette commande dans le terminal :

```
go get github.com/01-edu/z01
```

Elle sera automatiquement ajoutée au fichier ```go.mod``` si le module a été initialisé.

💡 Si ton dossier n'a pas encore de module Go, commence par :<br>
```
go mod init piscine
```

3. **Naviguer vers la quête souhaitée** :

```
cd quest02
```

4. **Exécuter un fichier Go** :

```
go run nomdufichier.go
```

**ou, pour compiler** :

```
go build nomdufichier.go
```
```
./nomdufichier
```


## À propos de Zone01

Zone01 est une école de code innovante basée sur l'apprentissage par projets et l'évaluation par les pairs. Aucun professeur, 100% pratique, 100% autonomie.
