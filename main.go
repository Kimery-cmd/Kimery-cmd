package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// Structure pour stocker les données du jeu
type HangManData struct {
	MaskedWord       string          // Mot masqué affiché au joueur
	SecretWord       string          // Mot secret à deviner
	Attempts         int             // Nombre de tentatives
	HangmanStages    [11]string      // Représentations ASCII des étapes du pendu
	RemainingLives   int             // Vies restantes
	GuessedLetters   map[string]bool // Lettres déjà proposées par le joueur
	IncorrectGuesses []string        // Lettres/mots incorrects proposés
}

// Fonction pour choisir la difficulté
func chooseDifficulty() string {
	for {
		var input string
		fmt.Println("Choisissez la difficulté :")
		fmt.Println("1 - Facile")
		fmt.Println("2 - Moyen")
		fmt.Println("3 - Difficile")
		fmt.Print("Entrez votre choix (1, 2 ou 3) : ")
		fmt.Scanln(&input) // Lecture de l'entrée utilisateur

		// Vérifie si le choix est valide
		if input == "1" || input == "2" || input == "3" {
			return input
		} else {
			fmt.Println("Erreur : Nombre non valide") // Message d'erreur si entrée incorrecte
		}
	}
}

// Fonction pour charger les étapes du pendu depuis un fichier texte
func loadHangmanStages() [11]string {
	var stages [11]string
	file, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		log.Fatal(err) // Arrête le programme en cas d'erreur
	}

	parts := strings.Split(string(file), "\n\n") // Sépare les parties du fichier
	for i, part := range parts {
		if i < 11 { // Assure-toi que 11 éléments sont chargés
			stages[i] = part
		}
	}
	return stages
}

// Affiche le mot masqué, chaque lettre séparée par un espace
func displayMaskedWord(maskedWord string) {
	for _, char := range maskedWord {
		fmt.Printf("%c ", char) // Affichage de chaque caractère
	}
	fmt.Println()
}

// Fonction pour révéler un nombre de lettres initiales dans le mot masqué
func revealInitialLetters(secretWord string) string {
	rand.Seed(time.Now().UnixNano()) // Initialise le générateur aléatoire
	wordLength := len(secretWord)
	numLettersToReveal := wordLength/2 - 1 // Nombre de lettres à révéler

	revealedWord := make([]rune, wordLength) // Tableau pour le mot masqué
	for i := range revealedWord {
		revealedWord[i] = '_' // Initialise chaque caractère comme un underscore
	}

	lettersRevealed := make(map[rune]bool) // Lettres déjà révélées
	countRevealed := 0

	// Révèle les lettres aléatoires
	for countRevealed < numLettersToReveal {
		randomIndex := rand.Intn(wordLength)          // Choix d'un mot aléatoire
		chosenLetter := rune(secretWord[randomIndex]) // Lettre choisie

		if lettersRevealed[chosenLetter] { // Ignore si déjà révéler
			continue
		}

		countOccurrences := 0
		// Compte les répétitions de la lettre dans le mot
		for _, char := range secretWord {
			if char == chosenLetter {
				countOccurrences++
			}
		}

		// Vérifie si on peut révéler la lettre ou pas
		if countOccurrences > numLettersToReveal-countRevealed {
			continue
		}

		// Mettre à jour le mot masqué avec la lettre révélée
		for i, char := range secretWord {
			if char == chosenLetter {
				revealedWord[i] = char
			}
		}

		lettersRevealed[chosenLetter] = true // Marque la lettre comme révélé
		countRevealed += countOccurrences    // Mettre à jour le nombre de lettres révélées
	}

	return string(revealedWord) // Retourne le mot masqué avec lettres révélé
}

// Fonction principale du programme
func main() {
	var wordList []string // Liste des mots selon la difficulté
	var difficulty string // Niveau de difficulté choisi
	var secretWord string // Mot choisi aléatoirement
	var hangmanData HangManData
	hangmanData.RemainingLives = 10
	hangmanData.HangmanStages = loadHangmanStages()
	hangmanData.GuessedLetters = make(map[string]bool)

	// Choix de la difficulté et chargement des mots
	difficulty = chooseDifficulty()
	switch difficulty {
	case "1":
		fmt.Printf("Vous avez choisi la difficulté facile\n")
		file, err := ioutil.ReadFile("words1.txt") // Charge le fichier des mots faciles
		if err != nil {
			log.Fatal(err) // Arrête le programme si le fichier est introuvable
		}
		wordList = strings.Split(string(file), "\n") // Sépare chaque mot
	case "2":
		fmt.Printf("Vous avez choisi la difficulté moyen\n")
		file, err := ioutil.ReadFile("words2.txt") // Charge le fichier des mots moyens
		if err != nil {
			log.Fatal(err)
		}
		wordList = strings.Split(string(file), "\n")
	case "3":
		fmt.Printf("Vous avez choisi la difficulté difficile\n")
		file, err := ioutil.ReadFile("words3.txt") // Charge le fichier des mots difficiles
		if err != nil {
			log.Fatal(err)
		}
		wordList = strings.Split(string(file), "\n")
	}

	// Sélection aléatoire d'un mot de la liste
	rand.Seed(time.Now().UnixNano())
	max := len(wordList)
	randomIndex := rand.Intn(max)                             // Génère un index aléatoire
	secretWord = strings.TrimSpace(wordList[randomIndex])     // Mot à deviner sans espaces
	hangmanData.SecretWord = secretWord                       // Enregistre le mot à trouver
	hangmanData.MaskedWord = revealInitialLetters(secretWord) // Initialise le mot masqué avec lettres révélées

	validInput := regexp.MustCompile(`^[a-zA-Z]+$`) // Expression régulière pour valider l'entrée

	// Boucle principale du jeu
	for hangmanData.RemainingLives > 0 && hangmanData.MaskedWord != hangmanData.SecretWord {
		// Affiche le motif du pendu uniquement si le joueur a moins de 10 vies
		if hangmanData.RemainingLives < 10 {
			fmt.Println(hangmanData.HangmanStages[10-hangmanData.RemainingLives]) // Affiche le motif du pendu
		}

		fmt.Printf("Vies restantes : %d\n", hangmanData.RemainingLives) // Affiche les vies restantes
		displayMaskedWord(hangmanData.MaskedWord)                       // Affiche le mot masquer actuel

		// Affiche les propositions incorrectes
		if len(hangmanData.IncorrectGuesses) > 0 {
			fmt.Println("Propositions incorrectes :", strings.Join(hangmanData.IncorrectGuesses, ", "))
		}

		fmt.Println() // Saute une ligne entre chaque essais

		var guess string
		fmt.Print("Proposez une lettre, un mot ou tapez 'exit' pour quitter : ")
		fmt.Scanln(&guess)             // Lecture de l'utilisateur
		guess = strings.ToLower(guess) // Convertit en minuscule

		if guess == "exit" { // Quitte le jeu si on tape "exit"
			fmt.Println("Vous avez quitté le jeu. À bientôt !")
			return
		}

		// Vérifie si l'entrée est valide (lettres uniquement)
		if !validInput.MatchString(guess) {
			fmt.Println("Erreur : Veuillez entrer uniquement des lettres.")
			continue
		}

		// Vérifie si la lettre ou le mot a déjà été proposer
		if hangmanData.GuessedLetters[guess] {
			fmt.Println("Vous avez déjà proposé cette lettre ou ce mot.")
			continue
		}

		// Vérifie si le joueur propose un mot complet ou une lettre unique
		if len(guess) > 1 {
			// Cas du mot complet
			if guess == hangmanData.SecretWord { // Si le mot est correct
				hangmanData.MaskedWord = hangmanData.SecretWord // Mettre à jour le mot masqué
				break
			} else {
				hangmanData.RemainingLives -= 2                                            // Pénalité de deux vies pour un mot incorrect
				hangmanData.IncorrectGuesses = append(hangmanData.IncorrectGuesses, guess) // Ajoute à la liste des propositions incorrectes
			}
		} else {
			// Cas d'une lettre seule
			if strings.Contains(hangmanData.SecretWord, guess) { // Lettre correcte
				for i := 0; i < len(hangmanData.SecretWord); i++ {
					if string(hangmanData.SecretWord[i]) == guess {
						hangmanData.MaskedWord = hangmanData.MaskedWord[:i] + guess + hangmanData.MaskedWord[i+1:] // Mettre à jour le mot masquer
					}
				}
			} else {
				hangmanData.RemainingLives--                                               // Perte d'une vie pour une lettre incorrecte
				hangmanData.IncorrectGuesses = append(hangmanData.IncorrectGuesses, guess) // Ajoute à la liste des proposition incorrectes
			}
		}
		hangmanData.GuessedLetters[guess] = true // Marque la lettre/mot comme proposer
	}

	// Vérifie si le joueur a gagné ou perdu
	if hangmanData.MaskedWord == hangmanData.SecretWord {
		fmt.Println("Félicitations, vous avez trouvé le mot :", hangmanData.SecretWord) // Gagné
	} else {
		if hangmanData.RemainingLives == 0 {
			fmt.Println(hangmanData.HangmanStages[10])                                      // Affiche le dernier motif du pendu
			fmt.Println("Dommage, vous avez perdu. Le mot était :", hangmanData.SecretWord) // Message de perte
		}
	}
	time.Sleep(5 * time.Second) // Pause de 5 secondes avant de fermer
}
