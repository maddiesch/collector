package data

import (
	"compress/bzip2"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/maddiesch/collector/internal/task"
)

type DownloadInput struct {
	Task task.Task

	// The URL the file will be downloaded from.
	FromURL string

	// The writer the file will be written to.
	Dest io.Writer
}

// Download loads a file from the network and copies it to the destination writer.
func Download(ctx context.Context, in DownloadInput) error {
	defer in.Task.MarkAsDone()

	req, err := http.NewRequestWithContext(ctx, "GET", in.FromURL, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("invalid http status code: %d", res.StatusCode)
	}

	in.Task.SetValue(res.ContentLength)

	dest := io.MultiWriter(in.Dest, &task.TaskWriter{Task: in.Task})

	_, err = io.Copy(dest, res.Body)

	return err
}

type InflateCompressedFileInput struct {
	Task task.Task

	In  io.Reader
	Out io.Writer
}

// InflateCompressedFile reads a bzip2 compressed file, uncompressed is and writes the data to the writer.
func InflateCompressedFile(ctx context.Context, in InflateCompressedFileInput) error {
	defer in.Task.MarkAsDone()

	r, w := io.Pipe()
	defer r.Close()
	defer w.Close()

	doneChan := make(chan error, 1)

	go func() {
		defer close(doneChan)

		if _, err := io.Copy(in.Out, bzip2.NewReader(in.In)); err != nil {
			doneChan <- err
		}
	}()

	return <-doneChan
}
