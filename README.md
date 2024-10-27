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
Choisissez la difficulté :
1 - Facile
2 - Moyen
3 - Difficile
Entrez votre choix (1, 2 ou 3) : 1
Vous avez choisi la difficulté facile

Vies restantes : 10
Mot masqué : _ _ _ _ _

Proposez une lettre, un mot ou tapez 'exit' pour quitter : e
Vies restantes : 10
Mot masqué : _ _ _ _ _

Propositions incorrectes : 

Proposez une lettre, un mot ou tapez 'exit' pour quitter : a
Vies restantes : 9
Mot masqué : _ _ _ _ _

Propositions incorrectes : a

Proposez une lettre, un mot ou tapez 'exit' pour quitter : t
Vies restantes : 8
Mot masqué : t _ _ _ _

Propositions incorrectes : a

Proposez une lettre, un mot ou tapez 'exit' pour quitter : exit
Vous avez quitté le jeu. À bientôt !
```
