package types

import (
	"context"

	"github.com/AlexTerra21/gophkeeper/pb"
)

// Текущее состояние клинта
type Condition struct {
	Ctx       context.Context
	ISAuth    bool
	Client    pb.GophkeeperClient
	AuthToken string
}
