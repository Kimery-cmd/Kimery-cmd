# Hangman

Le hangman est un jeu en Go où les joueurs doivent deviner un mot caché en proposant des lettres.
C'est le jeu du pendu

## Installation

Assurez-vous d'avoir Go installé, puis clonez le dépôt et exécutez le fichier principal.


```bash
git clone https://github.com/Kimery-cmd/hangman-classic.git
cd hangman-classic
go run main.go
```

## Utilisation

Le jeu propose un mot caché que le joueur doit deviner. Voici un exemple de fonctionnement :

```go
// Lancement du jeu
main()

// Choix d'une lettre
var guess string
fmt.Print("Proposez une lettre, un mot ou tapez 'exit' pour quitter : ")
fmt.Scanln(&guess) // Lecture de l'entrée utilisateur
```
