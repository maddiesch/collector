package mock

import (
	"compress/bzip2"
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func CardDatabaseLocation() string {
	_, file, _, _ := runtime.Caller(0)

	in, err := os.Open(filepath.Join(filepath.Dir(file), "card_datatabase.sqlite.bz2"))
	if err != nil {
		panic(err)
	}
	defer in.Close()

	dbPath := filepath.Join(filepath.Dir(file), "../../../tmp", "test_card_database.sqlite")
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
