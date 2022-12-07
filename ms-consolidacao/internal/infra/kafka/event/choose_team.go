package event

import (
	"context"
	"encoding/json"

	"github.com/augustoyuudi/ms-consolidacao/internal/usecase"
	"github.com/augustoyuudi/ms-consolidacao/pkg/uow"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ProcessChooseTeam struct{}

func (p ProcessChooseTeam) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.MyTeamChoosePlayersInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	addNewMyTeamUsecase := usecase.NewMyTeamChoosePlayersUseCase(uow)
	err = addNewMyTeamUsecase.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
