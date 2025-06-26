# Quest01 - Premiers pas avec Bash 🐚
Ces exercices ont été réalisés le premier jour de la piscine en Juin 2023, en totale autonomie et collaboration avec d'autres candidats.

## Objectifs
Créer des scripts Bash pour se familiariser avec :
- Les commandes de bases du shell ('ls', 'cat', 'echo', etc.)
- L'utilisation de 'curl' et 'jq' pour manipuler des données JSON
- La recherche et manipulation de fichier via la ligne de commande
- Les bases de l'automisation et de la logique conditionnelle

## Contenu
[to-git-or-not-to-git.sh](./to-git-or-not-to-git.sh) : affiche le nom, pouvoir et genre d'un super-héros avec l'ID 170 depuis un fichier JSON distant.
<br>
[who-are-you.sh](./who-are-you.sh) : identifie le super-héro avec l'ID 70 grâce à 'curl' et 'jq'.
<br>
[mastertheLS](./mastertheLS) : liste les fichiers de manière triée, sans fichiers cachés, séparés par des virgules.
<br>
[r](./r) : affiche simplement 'R' quand on exécute 'cat r'.
<br>
[look](./look) : recherche les fichiers commençant par 'a', finissant par 'z', ou commençant par 'z' et finissant par 'a!'.
<br>
[myfamily.sh](./myfamily.sh) : affiche les membres de la famille du super-héro identifié par l'ID contenu dans la variable d'environnement 'HERO_ID'.
<br>
[lookagain.sh](./lookagain.sh) : recherche les finissant par '.sh', affiche leur nom sans extension, par ordre décroissant.
<br>
[countfiles](./countfiles.sh) : affiche le nombre total de fichiers et dossiers.
<br>
[skip.sh](./skip.sh) : affiche un 'ls -l' en sautant une ligne sur deux à partir de la première.
<br>
[my_answer.sh](./my_answer.sh) : extrait une réponse à partir d'un test cloné, en combinant 'head' et 'tail'.
<br>
[explain.sh](./explain.sh) : extrait plusieurs informations clés d'un fichier d'enquête (témoin, voiture suspecte, noms de suspects non arrêtés).

## Ce que j'ai appris
- Syntaxe de base des scripts Bash
- Utilisation de commandes telles que ```curl```, ```jq```, ```ls```, ```find```, etc.
- Manipulation de fichiers et de données textuelles via le terminal.
- Rigueur dans l'écriture de scripts.
- Gestion des variables d'environnement.

## Comment exécuter
Assurez-vous d'être dans le dossier 'quest01' du dépôt cloné :
```cd quest01```

Pour lancer un script :
En rendant le script exécutable 
```chmod +x nom_du_script.sh```
<br>

```./nom_du_script.sh'```

**OU** en appelant Bash directement :
```bash nom_du_script.sh```
