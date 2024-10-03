package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quinta-nails/bots/internal/config"
	"github.com/quinta-nails/bots/internal/helpers"
	pb "github.com/quinta-nails/protobuf/gen/go/bots"
)

func (s *Service) SetupBot(ctx context.Context, in *pb.SetupBotRequest) (*pb.SetupBotResponse, error) {
	frontendConfig := config.FrontendConfig{}
	if err := env.Parse(&frontendConfig); err != nil {
		return nil, err
	}

	err := s.validator.Validate(in)
	if err != nil {
		return nil, err
	}

	botModel, err := s.db.GetBotById(ctx, in.Id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New(`bot not exist`)
	}

	b, err := helpers.NewTelegramBot(botModel.Token)
	if err != nil {
		return nil, err
	}

	_, err = b.SetChatMenuButton(ctx, &bot.SetChatMenuButtonParams{
		MenuButton: &models.MenuButtonWebApp{
			Type: `web_app`,
			Text: `Записаться`,
			WebApp: models.WebAppInfo{
				URL: fmt.Sprintf("%s?bot=%d", frontendConfig.URL, botModel.ID),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.SetupBotResponse{}, nil
}