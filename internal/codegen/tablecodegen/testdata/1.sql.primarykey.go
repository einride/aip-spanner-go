// +build testdata.1.sql.primarykey

package testdata

// Code generated by TestPrimaryKeyCodeGenerator_GenerateCode/primarykey/testdata/1.sql. DO NOT EDIT.

import (
	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
)

type SingersKey struct {
	SingerId int64
}

func (k SingersKey) SpannerKey() spanner.Key {
	return spanner.Key{
		k.SingerId,
	}
}

func (k SingersKey) SpannerKeySet() spanner.KeySet {
	return k.SpannerKey()
}

func (k SingersKey) Delete() *spanner.Mutation {
	return spanner.Delete("Singers", k.SpannerKey())
}

func (k SingersKey) BoolExpr() spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("SingerId"),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	return spansql.Paren{Expr: b}
}

func (k SingersKey) QualifiedBoolExpr(prefix spansql.PathExp) spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: append(prefix, spansql.ID("SingerId")),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	return spansql.Paren{Expr: b}
}
