package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Card struct {
	Suit  string
	Value string
}

type Player struct {
	Name  string
	Cards []Card
	Score int
}

var suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
var values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var deck []Card
var player, dealer Player

func main() {
	rand.Seed(time.Now().UnixNano())

	initializeGame()

	http.HandleFunc("/", gameHandler)
	http.HandleFunc("/hit", hitHandler)
	http.HandleFunc("/stand", standHandler)
	http.HandleFunc("/reset", resetGameHandler)

	fmt.Println("Server is running on port 8081...")
	http.ListenAndServe(":8081", nil)
}

func initializeGame() {
	deck = createDeck()
	shuffleDeck()

	// Clear player and dealer cards and reset their scores
	player = Player{Name: "Player", Cards: []Card{}, Score: 0}
	dealer = Player{Name: "Dealer", Cards: []Card{}, Score: 0}

	// Deal initial cards
	player.Cards = append(player.Cards, drawCard(), drawCard())
	dealer.Cards = append(dealer.Cards, drawCard(), drawCard())

	updateScores()
}

func createDeck() []Card {
	var deck []Card
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}
	return deck
}

func shuffleDeck() {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

func drawCard() Card {
	card := deck[0]
	deck = deck[1:]
	return card
}

func updateScores() {
	player.Score = calculateScore(player.Cards)
	dealer.Score = calculateScore(dealer.Cards)
}

func calculateScore(cards []Card) int {
	score := 0
	aceCount := 0

	for _, card := range cards {
		switch card.Value {
		case "A":
			aceCount++
			score += 11
		case "K", "Q", "J":
			score += 10
		default:
			value, _ := strconv.Atoi(card.Value)
			score += value
		}
	}

	for score > 21 && aceCount > 0 {
		score -= 10
		aceCount--
	}

	return score
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	gameState := `
    <div>
        <h2>Player</h2>
        <ul>`
	for _, card := range player.Cards {
		gameState += fmt.Sprintf("<li>%s of %s</li>", card.Value, card.Suit)
	}
	gameState += fmt.Sprintf("</ul><p>Score: %d</p>", player.Score)

	gameState += `
        <h2>Dealer</h2>
        <ul>`
	for i, card := range dealer.Cards {
		if i == 1 && r.URL.Path != "/stand" {
			gameState += "<li>Hidden</li>"
		} else {
			gameState += fmt.Sprintf("<li>%s of %s</li>", card.Value, card.Suit)
		}
	}
	gameState += fmt.Sprintf("</ul><p>Score: %d</p>", dealer.Score)

	if r.URL.Path != "/stand" && player.Score <= 21 {
		gameState += `
        <form action="/hit" method="post">
            <button type="submit">Hit</button>
        </form>
        <form action="/stand" method="post">
            <button type="submit">Stand</button>
        </form>`
	} else if r.URL.Path == "/stand" || player.Score > 21 {
		if player.Score > 21 || dealer.Score > 21 || player.Score == dealer.Score {
			gameState += "<p>It's a tie! Restarting in 10 seconds...</p>"
		} else if player.Score > dealer.Score {
			gameState += "<p>Player wins! Restarting in 10 seconds...</p>"
		} else {
			gameState += "<p>Dealer wins! Restarting in 10 seconds...</p>"
		}

		gameState += `<script>
            setTimeout(function() {
                window.location.href = "/reset";
            }, 10000);
            </script>`
	}

	gameState += `</div>`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, gameState)
}

func hitHandler(w http.ResponseWriter, r *http.Request) {
	player.Cards = append(player.Cards, drawCard())
	updateScores()

	if player.Score > 21 {
		http.Redirect(w, r, "/stand", http.StatusSeeOther)
		return
	}

	gameHandler(w, r)
}

func standHandler(w http.ResponseWriter, r *http.Request) {
	for dealer.Score < 17 {
		dealer.Cards = append(dealer.Cards, drawCard())
		updateScores()
	}

	gameHandler(w, r)
}

func resetGameHandler(w http.ResponseWriter, r *http.Request) {
	initializeGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
