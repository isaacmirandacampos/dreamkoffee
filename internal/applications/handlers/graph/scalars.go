package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/isaacmirandacampos/dreamkoffee/pkg/scalar"
	"github.com/shopspring/decimal"
	"github.com/vektah/gqlparser/v2/ast"
)

func (ec *executionContext) unmarshalInputDecimal(_ context.Context, v interface{}) (decimal.Decimal, error) {
	return scalar.UnmarshalDecimal(v)
}

func (ec *executionContext) _Decimal(_ context.Context, _ ast.SelectionSet, v *decimal.Decimal) graphql.Marshaler {
	return graphql.MarshalString(scalar.MarshalDecimal(*v))
}
