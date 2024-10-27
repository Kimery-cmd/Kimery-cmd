cd hangman-classic && echo "# Pendu

Pendu est un jeu en Go où les joueurs doivent deviner un mot caché en proposant des lettres.

## Installation

Assurez-vous d'avoir [Go](https://golang.org/doc/install) installé, puis clonez le dépôt et exécutez le fichier principal.

```bash
git clone https://github.com/Kimery-cmd/hangman-classic.git
cd hangman-classic
go run main.go
\`\`\`

## Utilisation

Le jeu propose un mot caché que le joueur doit deviner. Voici un exemple de fonctionnement :

\`\`\`go
// Lancement du jeu
main()

// Choix d'une lettre
var guess string
fmt.Print(\"Proposez une lettre, un mot ou tapez 'exit' pour quitter : \")
fmt.Scanln(&guess) // Lecture de l'entrée utilisateur
\`\`\`

Le joueur peut entrer une lettre ou un mot. Si le mot est deviné, le jeu se termine, sinon le joueur perd des vies jusqu'à épuisement.

## Contributing

Les pull requests sont les bienvenues. Pour des changements majeurs, veuillez d'abord ouvrir une issue pour discuter de ce que vous aimeriez modifier.

Assurez-vous de mettre à jour les tests si nécessaire.

## License

[MIT](https://choosealicense.com/licenses/mit/)" > README.md
