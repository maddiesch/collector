package conditional

import (
	"github.com/maddiesch/collector/internal/raptor/statement/generator"
)

type Conditional interface {
	Generate(generator.ArgumentNameProvider) (string, []any)
}
