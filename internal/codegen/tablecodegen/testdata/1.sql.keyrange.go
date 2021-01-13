// +build testdata.1.sql.keyrange

package testdata

// Code generated by TestKeyRangeCodeGenerator_GenerateCode/keyrange/testdata/1.sql. DO NOT EDIT.

import (
	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
)

type SingersKeyRange struct {
	Start SingersPartialKey
	End   SingersPartialKey
	Kind  spanner.KeyRangeKind
}

func (k SingersKeyRange) SpannerKeySet() spanner.KeySet {
	return spanner.KeyRange{
		Start: k.Start.SpannerKey(),
		End:   k.End.SpannerKey(),
		Kind:  k.Kind,
	}
}

type SingersPartialKey struct {
	SingerId int64
}

func (k SingersPartialKey) SpannerKey() spanner.Key {
	return spanner.Key{k.SingerId}
}

func (k SingersPartialKey) BoolExpr() spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("SingerId"),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	return spansql.Paren{Expr: b}
}

func (k SingersPartialKey) QualifiedBoolExpr(prefix spansql.PathExp) spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: append(prefix, spansql.ID("SingerId")),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	return spansql.Paren{Expr: b}
}
