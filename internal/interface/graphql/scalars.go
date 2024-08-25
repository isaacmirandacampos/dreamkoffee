package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/isaacmirandacampos/finkoffee/internal/utils"
	"github.com/shopspring/decimal"
	"github.com/vektah/gqlparser/v2/ast"
)

func (ec *executionContext) unmarshalInputDecimal(_ context.Context, v interface{}) (decimal.Decimal, error) {
	return utils.UnmarshalDecimal(v)
}

func (ec *executionContext) _Decimal(_ context.Context, _ ast.SelectionSet, v *decimal.Decimal) graphql.Marshaler {
	return utils.MarshalDecimal(*v)
}
