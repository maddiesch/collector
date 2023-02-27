package conditional

import (
	"fmt"

	"github.com/maddiesch/collector/internal/db/statement/generator"
)

func And(left, right Conditional) Conditional {
	return &logicalInfixConditional{left, right, "AND"}
}

func Or(left, right Conditional) Conditional {
	return &logicalInfixConditional{left, right, "OR"}
}

type logicalInfixConditional struct {
	left     Conditional
	right    Conditional
	operator string
}

func (c *logicalInfixConditional) Generate(provider generator.ArgumentNameProvider) (string, []any, error) {
	var args []any

	left, lArgs, err := c.left.Generate(provider)
	if err != nil {
		return "", nil, err
	}
	args = append(args, lArgs...)

	right, rArgs, err := c.right.Generate(provider)
	if err != nil {
		return "", nil, err
	}
	args = append(args, rArgs...)

	return fmt.Sprintf("(%s %s %s)", left, c.operator, right), args, nil
}
