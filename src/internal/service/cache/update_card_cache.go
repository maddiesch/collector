package cache

import (
	"context"
	"time"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type UpdateCardCacheService struct {
	ports.BulkDataEndpointRepository
	ports.MetadataRepository
	ports.CardDataRepository
	ports.CardDataStreamer
	ports.ProgressReporter
}

func (s *UpdateCardCacheService) Call(ctx context.Context) error {
	endpoints, err := s.BulkDataEndpointRepository.GetBulkDataEndpoint(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to load bulk data endpoints")
	}

	endpoint, ok := lo.Find(endpoints, func(e domain.BulkDataEndpoint) bool {
		return e.Kind == "default_cards"
	})
	if !ok {
		return errors.New("unable to locate an endpoint for default_cards")
	}

	stream, err := s.CardDataStreamer.StreamCardData(ctx, endpoint.DownloadURL)
	if err != nil {
		return errors.Wrap(err, "failed create card data streamer")
	}

	if err := s.CardDataRepository.FlushCardData(ctx); err != nil {
		return errors.Wrap(err, "failed to flush card data")
	}

	var count int64

	for event := range stream {
		switch value := event.(type) {
		case error:
			return errors.Wrap(value, "stream encountered an error")
		case domain.CardData:
			count++

			if err := s.CardDataRepository.InsertCardData(ctx, value); err != nil {
				return errors.Wrap(err, "failed to insert card data")
			}

			s.ProgressReporter.ReportProgress(1.0, "")
		}
	}

	return s.MetadataRepository.SetMetadata(ctx, "cache.default_cards.last_updated", time.Now())
}
