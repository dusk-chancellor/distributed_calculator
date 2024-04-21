package manager

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/dusk-chancellor/distributed_calculator/internal/grpc/orchestrator"
	"github.com/dusk-chancellor/distributed_calculator/internal/storage"
)

type ExpressionUpdater interface {
	UpdateExpression(ctx context.Context, answer, status string, id int64,) error
	SelectAllExpressions(ctx context.Context) ([]storage.Expression, error)
}

// all this states must be somewhere else than here
var (
	done = "done"
	trouble = "error"
	null = "null"
)

func RunManager(ctx context.Context, expressionUpdater ExpressionUpdater) {
	
	log.Println("running Orchestrator manager")
	for {
		go func() {
			storedExpressions, err := expressionUpdater.SelectAllExpressions(ctx)
			if err != nil {
				log.Printf("could not SelectExpressions() from database: %v", err)
			}

			for _, expression := range storedExpressions {
				if expression.Status == done || expression.Status == trouble {
					continue
				} else {
					ans, err := orchestrator.Calculate(ctx, expression.Expression)
					if err != nil {
						log.Printf("could not Calculate(): %v", err)
						expressionUpdater.UpdateExpression(
							ctx, null, trouble, expression.ID,
						)
						continue
					}

					res := strconv.FormatFloat(ans, 'g', -1, 64)

					if err = expressionUpdater.UpdateExpression(
						ctx, res, done, expression.ID,
					); err != nil {
						log.Printf("could not UpdateExpression(): %v", err)
					}
				}
			}
		}()

		time.Sleep(7 * time.Second)
	}
}

