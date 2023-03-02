package data

import (
	"compress/bzip2"
	"context"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type DownloadInput struct {
	*zap.Logger

	// The URL the file will be downloaded from.
	FromURL string

	// The writer the file will be written to.
	Dest io.Writer
}

// Download loads a file from the network and copies it to the destination writer.
func Download(ctx context.Context, in DownloadInput) error {
	req, err := http.NewRequestWithContext(ctx, "GET", in.FromURL, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(in.Dest, res.Body)

	return err
}

// InflateCompressedFile reads a bzip2 compressed file, uncompressed is and writes the data to the writer.
func InflateCompressedFile(ctx context.Context, in io.Reader, out io.Writer) error {
	r, w := io.Pipe()
	defer r.Close()
	defer w.Close()

	doneChan := make(chan error, 1)

	go func() {
		defer close(doneChan)

		if _, err := io.Copy(out, bzip2.NewReader(in)); err != nil {
			doneChan <- err
		}
	}()

	return <-doneChan
}
