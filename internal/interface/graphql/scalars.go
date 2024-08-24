package graphql

import (
	"context"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
	"github.com/vektah/gqlparser/v2/ast"
)

func MarshalDecimal(d decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf("\"%s\"", d.String()))
	})
}

func UnmarshalDecimal(v interface{}) (decimal.Decimal, error) {
	switch v := v.(type) {
	case string:
		return decimal.NewFromString(v)
	case float64:
		return decimal.NewFromFloat(v), nil
	case int:
		return decimal.NewFromInt(int64(v)), nil
	default:
		return decimal.Decimal{}, fmt.Errorf("unexpected type %T for Decimal", v)
	}
}

func (ec *executionContext) unmarshalInputDecimal(_ context.Context, v interface{}) (decimal.Decimal, error) {
	return UnmarshalDecimal(v)
}

func (ec *executionContext) _Decimal(_ context.Context, _ ast.SelectionSet, v *decimal.Decimal) graphql.Marshaler {
	return MarshalDecimal(*v)
}
