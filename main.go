package main

import (
	"cli-test/functions"
	"cli-test/models"
	"errors"

	"github.com/charmbracelet/huh"
)

func main() {
	var game models.Game

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

	game.Deck = functions.CreateDeck()
	functions.StartGame(&game)
}
