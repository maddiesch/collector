package conditional

import (
	"github.com/maddiesch/collector/internal/db/statement/generator"
)

type Conditional interface {
	Generate(generator.ArgumentNameProvider) (string, []any, error)
}
