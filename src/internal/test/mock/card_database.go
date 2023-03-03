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

func TempDir() string {
	_, currentFilePath, _, _ := runtime.Caller(0)

	return filepath.Join(filepath.Dir(currentFilePath), "../../../tmp")
}

func CopyCardDatabase(to string) error {
	compressedPath := CompressedCardDatabasePath()

	in, err := os.Open(compressedPath)
	if err != nil {
		return err
	}
	defer in.Close()

	if _, err := os.Stat(to); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(filepath.Dir(to), 0755); err != nil {
			return err
		}

		out, err := os.OpenFile(to, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer out.Close()

		if _, err := in.Seek(0, io.SeekStart); err != nil {
			return err
		}

		if _, err := io.Copy(out, bzip2.NewReader(in)); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func CardDatabaseLocation() string {
	dbPath := filepath.Join(TempDir(), "test_card_database.sqlite")

	CopyCardDatabase(dbPath)

	return dbPath
}
