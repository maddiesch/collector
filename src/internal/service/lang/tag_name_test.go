package lang_test

import (
	"testing"

	"github.com/maddiesch/collector/internal/service/lang"
	"github.com/stretchr/testify/assert"
)

func TestDisplayNameForTagString(t *testing.T) {
	cases := map[string]string{
		"en":  "English",
		"ja":  "Japanese",
		"es":  "Spanish",
		"it":  "Italian",
		"ph":  "Phyrexian",
		"zhs": "zhs",
		"sa":  "Sanskrit",
		"zht": "zht",
		"he":  "Hebrew",
		"de":  "German",
		"ko":  "Korean",
		"ru":  "Russian",
		"ar":  "Arabic",
		"grc": "Ancient Greek",
		"pt":  "Portuguese",
		"la":  "Latin",
		"fr":  "French",
	}

	for tag, expected := range cases {
		t.Run(tag, func(t *testing.T) {
			value := lang.DisplayNameForTagString(tag)

			assert.Equal(t, expected, value)
		})
	}
}
