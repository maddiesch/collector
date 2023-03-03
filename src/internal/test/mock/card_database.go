package mock

import (
	"compress/bzip2"
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func CompressedCardDatabasePath() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Join(filepath.Dir(file), "card_database.sqlite.bz2")
}

func CardDatabaseLocation() string {
	compressedPath := CompressedCardDatabasePath()

	in, err := os.Open(compressedPath)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	dbPath := filepath.Join(filepath.Dir(compressedPath), "../../../tmp", "test_card_database.sqlite")
	if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
			panic(err)
		}

		out, err := os.OpenFile(dbPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		if _, err := in.Seek(0, io.SeekStart); err != nil {
			panic(err)
		}

		if _, err := io.Copy(out, bzip2.NewReader(in)); err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	return dbPath
}
