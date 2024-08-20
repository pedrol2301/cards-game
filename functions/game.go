package functions

import (
	"cli-test/models"
	"fmt"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"golang.org/x/exp/rand"
)

func CreateDeck() []models.Card {
	var deck []models.Card
	suits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10} // Considerando 10 para J, Q, K

	for _, naipe := range suits {
		for _, valor := range values {
			deck = append(deck, models.Card{Value: valor, Suit: naipe})
		}
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}

func CalcularPontos(hand []models.Card) int {
	total := 0
	numAs := 0

	for _, carta := range hand {
		if carta.Value == 1 {
			numAs++
			total += 11
		} else {
			total += carta.Value
		}
	}

	for total > 21 && numAs > 0 {
		total -= 10
		numAs--
	}

	return total
}

func DetermineWinner(game *models.Game) {
	if game.Player1.Points > 21 {
		fmt.Printf("%s estourou! %s ganha!\n", game.Player1.Name, game.Player2.Name)
		return
	}

	if game.Player2.Points > 21 {
		fmt.Printf("%s estourou! %s ganha!\n", game.Player2.Name, game.Player1.Name)
		return
	}

	if game.Player1.Points > game.Player2.Points {
		fmt.Printf("%s ganha!\n", game.Player1.Name)
	} else if game.Player2.Points > game.Player1.Points {
		fmt.Printf("%s ganha!\n", game.Player2.Name)
	} else {
		fmt.Printf("Empate!\n")
	}
}

func PlayerTurn(player *models.Player, deck *[]models.Card) {
	for {

		_ = spinner.New().Title("Shuffle cards...").Run()

		time.Sleep(2 * time.Second)

		var choice bool
		var cardsString string
		for _, cardVal := range player.Hand {
			cardsString += fmt.Sprintf("%d de %s\n", cardVal.Value, cardVal.Suit)
		}
		// fmt.Printf("%s, suas cartas: %+v\n", player.Name, player.Hand)

		form := huh.NewForm(huh.NewGroup(
			huh.NewNote().Description(fmt.Sprintf("Suas cartas são:\n%s", cardsString)),

			huh.NewConfirm().Title(fmt.Sprintf("Pontos: %d", player.Points)).
				Affirmative("Pedir").
				Negative("Parar").
				Value(&choice),
		))

		form.Run()

		if choice {
			*deck = append(*deck, (*deck)[0])
			player.Hand = append(player.Hand, (*deck)[0])
			*deck = (*deck)[1:]

			player.Points = CalcularPontos(player.Hand)
			if player.Points >= 21 {
				break
			}
		} else {
			break
		}
	}
}

func DealerTurn(game *models.Game) {
	// fmt.Printf("%+v\n", game.Player2.Hand)
	var dealerHandString string
	for _, cards := range game.Player2.Hand {
		dealerHandString += fmt.Sprintf("%d de %s\n", cards.Value, cards.Suit)
	}
	huh.NewForm(
		huh.NewGroup(huh.NewNote().Title("Dealer revela sua mão:\n").Description(dealerHandString)),
	).Run()

	for game.Player2.Points < 17 {
		game.Player2.Hand = append(game.Player2.Hand, game.Deck[0])
		game.Deck = game.Deck[1:]

		game.Player2.Points = CalcularPontos(game.Player2.Hand)
	}
}

func StartGame(game *models.Game) {
	game.Player1.Hand = append(game.Player1.Hand, game.Deck[0], game.Deck[1])
	game.Player2.Hand = append(game.Player2.Hand, game.Deck[2], game.Deck[3])

	game.Player1.Points = CalcularPontos(game.Player1.Hand)
	game.Player2.Points = CalcularPontos(game.Player2.Hand)

	// Mostra a carta inicial do dealer
	if game.IsDealer {

		dealerHandString := fmt.Sprintf("%d de %s", game.Player2.Hand[0].Value, game.Player2.Hand[0].Suit)
		huh.NewForm(
			huh.NewGroup(huh.NewNote().Title("Dealer mostra sua mão:\n").Description(dealerHandString)),
		).Run()
	}

	PlayerTurn(&game.Player1, &game.Deck)

	if game.SoloGame {
		DealerTurn(game)
	} else {
		PlayerTurn(&game.Player2, &game.Deck)
	}

	DetermineWinner(game)
}
