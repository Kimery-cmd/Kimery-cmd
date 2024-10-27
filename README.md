Jeu du Pendu en Go
Ce projet est une implémentation du jeu du pendu en langage Go. Le joueur doit deviner un mot caché en entrant des lettres, tout en ayant un nombre limité de vies. Le jeu affiche progressivement un motif de pendu en fonction des vies restantes.

Fichiers du projet
main.go : Le fichier principal contenant la logique du jeu.
words1.txt : Liste de mots faciles pour le niveau de difficulté 1.
words2.txt : Liste de mots moyens pour le niveau de difficulté 2.
words3.txt : Liste de mots difficiles pour le niveau de difficulté 3.
hangman.txt : Fichier contenant les différentes étapes du pendu en ASCII art, affichées en fonction des vies restantes.
Prérequis
Go : Assurez-vous d'avoir Go installé pour exécuter le code. Installer Go.
Installation
Clonez le dépôt :
bash
Copier le code
git clone https://github.com/votre-utilisateur/pendu-go.git
cd pendu-go
Assurez-vous que tous les fichiers (words1.txt, words2.txt, words3.txt, hangman.txt) sont dans le même dossier que main.go.
Utilisation
Lancez le jeu avec la commande suivante :
bash
Copier le code
go run main.go
Choisissez un niveau de difficulté en entrant 1, 2, ou 3 pour respectivement facile, moyen, ou difficile.
Devinez le mot en entrant des lettres ou un mot complet. Vous pouvez aussi entrer exit pour quitter la partie.
Le jeu affiche vos vies restantes et le motif du pendu. Une fois le mot deviné ou les vies épuisées, le jeu se termine.
Fonctionnalités
Niveaux de difficulté : Trois niveaux de difficulté avec des mots de longueurs variées.
Révélation de lettres : Au début du jeu, certaines lettres du mot sont révélées aléatoirement.
Gestion des vies : Le joueur commence avec 10 vies et perd une vie pour chaque lettre incorrecte. Une pénalité de 2 vies est appliquée pour un mot incorrect.
Affichage dynamique : Le motif du pendu est affiché en fonction des vies restantes, n'apparaissant qu'après 9 vies.