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
	"time"

	"github.com/charmbracelet/huh"
	"golang.org/x/exp/rand"
)

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

func main() {

	var Player1 Player
	var Player2 Player

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
				Value(&Player1.Name).Validate(func(s string) error {
				if s == "" {
					return errors.New("nome é obrigatório")
				}
				return nil
			}))).WithAccessible(accessible)

	err := form.Run()
	if err != nil {
		println("Erro ao rodar form")
	}

	deck := CreateDeck()

	Player1.Hand = append(Player1.Hand, deck[0])
	Player1.Hand = append(Player1.Hand, deck[1])

	Player2.Hand = append(Player2.Hand, deck[2])
	Player2.Hand = append(Player2.Hand, deck[3])

	Player1.Points = CalcularPontos(Player1.Hand)
	Player2.Points = CalcularPontos(Player2.Hand)

	for Player1.Points < 21 && Player2.Points < 21 {
		Player1.Hand = append(Player1.Hand, deck[0])
		Player1.Points = CalcularPontos(Player1.Hand)

		Player2.Hand = append(Player2.Hand, deck[1])
		Player2.Points = CalcularPontos(Player2.Hand)
	}

	if Player1.Points == 21 {
		println("Player 1 ganhou")
		return
	}

	if Player2.Points == 21 {
		println("Player 2 ganhou")
		return
	}

	if Player1.Points > 21 {
		println("Player 2 ganhou")
		return
	}

	if Player2.Points > 21 {
		println("Player 1 ganhou")
		return
	}

	if Player1.Points == Player2.Points {
		println("Empate")
		return
	}

}
