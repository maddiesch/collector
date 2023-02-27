package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReleaseDate(t *testing.T) {
	t.Run("Value", func(t *testing.T) {
		d := ReleaseDate{
			Year:  2023,
			Month: time.February,
			Day:   3,
		}

		v, err := d.Value()

		require.NoError(t, err)

		assert.Equal(t, "2023-02-03", v)
	})

	t.Run("Scan", func(t *testing.T) {
		r := new(ReleaseDate)

		err := r.Scan("2023-02-03")

		require.NoError(t, err)

		assert.Equal(t, 2023, r.Year)
		assert.Equal(t, time.February, r.Month)
		assert.Equal(t, 3, r.Day)
	})
}
