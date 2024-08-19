package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/charmbracelet/huh"
// 	"github.com/charmbracelet/huh/spinner"
// 	"github.com/charmbracelet/lipgloss"
// 	xstrings "github.com/charmbracelet/x/exp/strings"
// )

// type Spice int

// const (
// 	Mild Spice = iota + 1
// 	Medium
// 	Hot
// )

// func (s Spice) String() string {
// 	switch s {
// 	case Mild:
// 		return "Mild "
// 	case Medium:
// 		return "Medium-Spicy "
// 	case Hot:
// 		return "Spicy-Hot "
// 	default:
// 		return ""
// 	}
// }

// type Order struct {
// 	Burger       Burger
// 	Side         string
// 	Name         string
// 	Instructions string
// 	Discount     bool
// }

// type Burger struct {
// 	Type     string
// 	Toppings []string
// 	Spice    Spice
// }

// func main() {
// 	var burger Burger
// 	var order = Order{Burger: burger}

// 	// Should we run in accessible mode?
// 	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

// 	form := huh.NewForm(
// 		huh.NewGroup(huh.NewNote().
// 			Title("Charmburger").
// 			Description("Welcome to _Charmburger™_.\n\nHow may we take your order?\n\n").
// 			Next(true).
// 			NextLabel("Next"),
// 		),

// 		// Choose a burger.
// 		// We'll need to know what topping to add too.
// 		huh.NewGroup(
// 			huh.NewSelect[string]().
// 				Options(huh.NewOptions("Charmburger Classic", "Chickwich", "Fishburger", "Charmpossible™ Burger")...).
// 				Title("Choose your burger").
// 				Description("At Charm we truly have a burger for everyone.").
// 				Validate(func(t string) error {
// 					if t == "Fishburger" {
// 						return fmt.Errorf("no fish today, sorry")
// 					}
// 					return nil
// 				}).
// 				Value(&order.Burger.Type),

// 			huh.NewMultiSelect[string]().
// 				Title("Toppings").
// 				Description("Choose up to 4.").
// 				Options(
// 					huh.NewOption("Lettuce", "Lettuce").Selected(true),
// 					huh.NewOption("Tomatoes", "Tomatoes").Selected(true),
// 					huh.NewOption("Charm Sauce", "Charm Sauce"),
// 					huh.NewOption("Jalapeños", "Jalapeños"),
// 					huh.NewOption("Cheese", "Cheese"),
// 					huh.NewOption("Vegan Cheese", "Vegan Cheese"),
// 					huh.NewOption("Nutella", "Nutella"),
// 				).
// 				Validate(func(t []string) error {
// 					if len(t) <= 0 {
// 						return fmt.Errorf("at least one topping is required")
// 					}
// 					return nil
// 				}).
// 				Value(&order.Burger.Toppings).
// 				Filterable(true).
// 				Limit(4),
// 		),

// 		// Prompt for toppings and special instructions.
// 		// The customer can ask for up to 4 toppings.
// 		huh.NewGroup(
// 			huh.NewSelect[Spice]().
// 				Title("Spice level").
// 				Options(
// 					huh.NewOption("Mild", Mild).Selected(true),
// 					huh.NewOption("Medium", Medium),
// 					huh.NewOption("Hot", Hot),
// 				).
// 				Value(&order.Burger.Spice),

// 			huh.NewSelect[string]().
// 				Options(huh.NewOptions("Fries", "Disco Fries", "R&B Fries", "Carrots")...).
// 				Value(&order.Side).
// 				Title("Sides").
// 				Description("You get one free side with this order."),
// 		),

// 		// Gather final details for the order.
// 		huh.NewGroup(
// 			huh.NewInput().
// 				Value(&order.Name).
// 				Title("What's your name?").
// 				Placeholder("Margaret Thatcher").
// 				Validate(func(s string) error {
// 					if s == "Frank" {
// 						return errors.New("no franks, sorry")
// 					}
// 					return nil
// 				}).
// 				Description("For when your order is ready."),

// 			huh.NewText().
// 				Value(&order.Instructions).
// 				Placeholder("Just put it in the mailbox please").
// 				Title("Special Instructions").
// 				Description("Anything we should know?").
// 				CharLimit(400).
// 				Lines(5),

// 			huh.NewConfirm().
// 				Title("Would you like 15% off?").
// 				Value(&order.Discount).
// 				Affirmative("Yes!").
// 				Negative("No."),
// 		),
// 	).WithAccessible(accessible)

// 	err := form.Run()

// 	if err != nil {
// 		fmt.Println("Uh oh:", err)
// 		os.Exit(1)
// 	}

// 	prepareBurger := func() {
// 		time.Sleep(2 * time.Second)
// 	}

// 	_ = spinner.New().Title("Preparing your burger...").Accessible(accessible).Action(prepareBurger).Run()

// 	// Print order summary.
// 	{
// 		var sb strings.Builder
// 		keyword := func(s string) string {
// 			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
// 		}
// 		fmt.Fprintf(&sb,
// 			"%s\n\nOne %s%s, topped with %s with %s on the side.",
// 			lipgloss.NewStyle().Bold(true).Render("BURGER RECEIPT"),
// 			keyword(order.Burger.Spice.String()),
// 			keyword(order.Burger.Type),
// 			keyword(xstrings.EnglishJoin(order.Burger.Toppings, true)),
// 			keyword(order.Side),
// 		)

// 		name := order.Name
// 		if name != "" {
// 			name = ", " + name
// 		}
// 		fmt.Fprintf(&sb, "\n\nThanks for your order%s!", name)

// 		if order.Discount {
// 			fmt.Fprint(&sb, "\n\nEnjoy 15% off.")
// 		}

// 		fmt.Println(
// 			lipgloss.NewStyle().
// 				Width(40).
// 				BorderStyle(lipgloss.RoundedBorder()).
// 				BorderForeground(lipgloss.Color("63")).
// 				Padding(1, 2).
// 				Render(sb.String()),
// 		)
// 	}
// }

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"golang.org/x/exp/rand"
)

type Game struct {
	SoloGame bool
	IsDealer bool
	Player1  Player
	Player2  Player
	Deck     []Card
}

type Card struct {
	Value int
	Suit  string
}

type Player struct {
	Name   string
	Hand   []Card
	Points int
}

func CreateDeck() []Card {
	var deck []Card
	suits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10} // Considerando 10 para J, Q, K

	for _, naipe := range suits {
		for _, valor := range values {
			deck = append(deck, Card{Value: valor, Suit: naipe})
		}
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}

func CalcularPontos(hand []Card) int {
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

func DetermineWinner(game *Game) {
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

func PlayerTurn(player *Player, deck *[]Card) {
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

func DealerTurn(game *Game) {
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

func StartGame(game *Game) {
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

func main() {
	var game Game

	accessible := false

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Blackjack").
				Description("Bem-vindo ao Blackjack.\n\n").
				Next(true).
				NextLabel("Next"),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Nome do jogador 1").
				Placeholder("Nome").
				Value(&game.Player1.Name).Validate(func(s string) error {
				if s == "" {
					return errors.New("nome é obrigatório")
				}
				return nil
			}),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Jogar sozinho?").
				Value(&game.SoloGame).
				Affirmative("Sim").
				Negative("Não"),
		),
	).WithAccessible(accessible)

	err := form.Run()
	if err != nil {
		println("Erro ao rodar form")
		return
	}

	if game.SoloGame {
		game.Player2.Name = "Dealer"
		game.IsDealer = true
	} else {
		secondPlayerForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Nome do jogador 2").
					Placeholder("Nome").
					Value(&game.Player2.Name).Validate(func(s string) error {
					if s == "" {
						return errors.New("nome é obrigatório")
					}
					return nil
				}),
			),
		).WithAccessible(accessible)

		err := secondPlayerForm.Run()
		if err != nil {
			println("Erro ao rodar form")
			return
		}

		game.IsDealer = false
	}

	game.Deck = CreateDeck()
	StartGame(&game)
}
