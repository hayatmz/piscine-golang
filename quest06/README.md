# Quest06 - Arguments de ligne de commande en Go 🧭

## Objectifs
- Manipuler les arguments passés au programme via la ligne de commande.
- Utiliser le package ```os``` et sa fonction ```os.Args```.
- Comprendre comment Go interprète et récupère les entrées utilisateur à l'exécution.
- Travailler la logique de parcours de tableaux (```[]string```), dans l'ordre ou en sens inverse.

## Contenu
Ces exercices sont des programmes indépendants dans leur propre dossier contenant un fichier ```main.go```:<br>
[printprogramname](./printprogramname/main.go) : Affiche le nom du programme exécuté.<br>
[printparams/](./printparams/main.go) : Affiche les arguments de la ligne de commande dans l'ordre.<br>
[revparams](./revparams/main.go) : Affiche les arguments de la ligne de commande en ordre inverse.<br>

## Ce que j'ai appris
- Accéder aux arguments via ```os.Args```.
- Comprendre que ```os.Args[0]``` contient toujours le nom du programme.
- Parcourir un tableau de chaînes (```[]string```) de différentes manières (ordre, ordre inverse).
- Gérer l'absence ou la présence d'arguments, sans faire planter le programme.
- Ecrire des programmes robustes prenant en compte différents cas d'utilisation.

## Comment exécuter
Assurez-vous d'être dans le dossier 'quest06' du dépôt cloné et dans le dossier du programme souhaité :<br>

```cd quest06```<br>
```cd nom_du_dossier```<br>

Pour lancer un programme :<br>
```go run main.go```<br>

**Ou** pour compiler et exécuter :<br>
```go build nom_du_fichier.go```<br>
```./nom_du_fichier```