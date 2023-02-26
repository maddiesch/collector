package conditional

import (
	"github.com/maddiesch/collector/internal/db/statement/generator"
)

type Conditional interface {
	Generate(generator.ArgumentNameProvider) (string, []any, error)
}

/**

type Conditional interface {
	Generate() (string, []any, error)
}

func Eq(column string, value any) Conditional {
	return &EqConditional{Column: column, Value: value}
}

type EqConditional struct {
	Column string
	Value  any
}

func (e *EqConditional) Generate() (string, []any, error) {
	return fmt.Sprintf(`"%s" = ?`, e.Column), []any{e.Value}, nil
}

func Like(column string, value string, leading bool) Conditional {
	return &LikeConditional{Column: column, Value: value, Leading: leading}
}

type LikeConditional struct {
	Column  string
	Value   string
	Leading bool
}

func (c *LikeConditional) Generate() (string, []any, error) {
	var value string

	if c.Leading {
		value = `%` + c.Value
	} else {
		value = c.Value + `%`
	}

	return fmt.Sprintf(`"%s" LIKE ?`, c.Column), []any{value}, nil
}

func And(left, right Conditional) Conditional {
	return &JoinConditional{Left: left, Right: right, Instruction: "AND"}
}

func Or(left, right Conditional) Conditional {
	return &JoinConditional{Left: left, Right: right, Instruction: "OR"}
}

type JoinConditional struct {
	Instruction string
	Left        Conditional
	Right       Conditional
}

func (a *JoinConditional) Generate() (string, []any, error) {
	lq, la, err := a.Left.Generate()
	if err != nil {
		return "", nil, err
	}
	rq, ra, err := a.Right.Generate()
	if err != nil {
		return "", nil, err
	}

	return fmt.Sprintf("(%s) %s (%s)", lq, a.Instruction, rq), append(la, ra...), nil
}

// Synthetic

func StringContains(column, value string) Conditional {
	return Or(
		Like(column, value, true),
		Like(column, value, false),
	)
}

*/
