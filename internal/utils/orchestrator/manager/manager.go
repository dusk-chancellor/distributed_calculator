package manager

import (
	"context"
	"log"
	"strconv"

	"github.com/dusk-chancellor/distributed_calculator/internal/grpc/orchestrator"
	"github.com/dusk-chancellor/distributed_calculator/internal/storage"
)

// all this status must be somewhere else than here
var (
	trouble = "error"
	done    = "done"
	null    = "null"
)

func Manage(ctx context.Context, expressionOperator storage.ExpressionInteractor, agentAddr string) {

	go func() {
		storedExpressions, err := expressionOperator.SelectAllExpressions(ctx)
		if err != nil {
			log.Printf("could not SelectExpressions() from database: %v", err)
		}

		for _, expression := range storedExpressions {
			if expression.Status == done || expression.Status == trouble {
				continue
			} else {
				ans, err := orchestrator.Calculate(ctx, expression.Expression, agentAddr)
				if err != nil {
					expressionOperator.UpdateExpression(
						ctx, null, trouble, expression.ID,
					)
					continue
				}

				res := strconv.FormatFloat(ans, 'g', -1, 64)

				if err = expressionOperator.UpdateExpression(
					ctx, res, done, expression.ID,
				); err != nil {
					log.Printf("could not UpdateExpression(): %v", err)
				}
			}
		}
	}()
}

