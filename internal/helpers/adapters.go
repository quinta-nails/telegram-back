package helpers

import (
	pb "github.com/quinta-nails/protobuf/gen/go/telegram_backend"
	"github.com/quinta-nails/telegram-backend/internal/db"
)

func NewPbBotFromBotRow(row *db.Bot) *pb.Bot {
	return &pb.Bot{
		Id:        row.ID,
		FirstName: row.FirstName,
		Username:  row.Username,
	}
}