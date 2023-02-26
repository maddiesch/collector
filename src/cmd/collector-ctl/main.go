package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/maddiesch/collector/internal/db"
	"github.com/maddiesch/collector/internal/repositories/database"
)

func main() {
	ctx := context.Background()

	dir := mustLoadWorkingDir()

	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal(err)
	}

	config := config.Config{
		Dir: dir,
		DB: database.New(
			mustCreateDatabaseConn(ctx, filepath.Join(dir, "database.sqlite")),
		),
	}

	cli.Execute(ctx, config)
}

func mustLoadWorkingDir() string {
	user, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(user, ".collector")
}

func mustCreateDatabaseConn(ctx context.Context, path string) *db.Conn {
	conn, err := db.NewConn(ctx, db.NewConnInput{
		FilePath: path,
	})
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
