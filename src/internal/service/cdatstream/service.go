package cdatstream

import (
	"context"
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
)

func New() ports.CardDataStreamer {
	return &service{}
}

type service struct {
}

func (s *service) StreamCardData(ctx context.Context, endpoint string) (<-chan any, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	streamChan := make(chan any, 250)

	go func() {
		defer close(streamChan)
		defer response.Body.Close()

		haltAndCatchFire := func(err error) {
			streamChan <- err
			runtime.Goexit()
		}

		decoder := json.NewDecoder(response.Body)

		// Consume Open Bracket
		if _, err := decoder.Token(); err != nil {
			haltAndCatchFire(err)
		}

		for decoder.More() {
			var card domain.CardData

			if err := ctx.Err(); err != nil {
				runtime.Goexit()
			}

			if err := decoder.Decode(&card); err != nil {
				haltAndCatchFire(err)
			}

			streamChan <- card
		}

		// Consume Close Bracket
		if _, err := decoder.Token(); err != nil {
			haltAndCatchFire(err)
		}
	}()

	return streamChan, nil
}
