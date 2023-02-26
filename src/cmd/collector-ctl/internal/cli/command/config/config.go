package config

import (
	"context"

	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/color"
	"github.com/maddiesch/collector/internal/repositories/database"
)

type Config struct {
	Dir string
	DB  *database.Database
}

func (c Config) Shutdown(ctx context.Context) error {
	if err := c.DB.Close(); err != nil {
		color.Warn.Printf("Close Database Error: %s\n", err)
	}

	return nil
}
