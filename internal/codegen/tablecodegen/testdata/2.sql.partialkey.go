// +build testdata.2.sql.partialkey

package testdata

// Code generated by TestPartialKeyCodeGenerator_GenerateCode/partialkey/testdata/2.sql. DO NOT EDIT.

import (
	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
)

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

type AlbumsPartialKey struct {
	SingerId     int64
	AlbumId      int64
	ValidAlbumId bool
}

func (k AlbumsPartialKey) SpannerKey() spanner.Key {
	n := 1
	if k.ValidAlbumId {
		n++
	}
	result := make(spanner.Key, 0, n)
	result = append(result, k.SingerId)
	if k.ValidAlbumId {
		result = append(result, k.AlbumId)
	}
	return result
}

func (k AlbumsPartialKey) SpannerKeySet() spanner.KeySet {
	return k.SpannerKey()
}

func (k AlbumsPartialKey) BoolExpr() spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("SingerId"),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	if k.ValidAlbumId {
		b = spansql.LogicalOp{
			Op:  spansql.And,
			LHS: b,
			RHS: spansql.ComparisonOp{
				Op:  spansql.Eq,
				LHS: spansql.ID("AlbumId"),
				RHS: spansql.IntegerLiteral(k.AlbumId),
			},
		}
	}
	return spansql.Paren{Expr: b}
}

func (k AlbumsPartialKey) QualifiedBoolExpr(prefix spansql.PathExp) spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: append(prefix, spansql.ID("SingerId")),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	if k.ValidAlbumId {
		b = spansql.LogicalOp{
			Op:  spansql.And,
			LHS: b,
			RHS: spansql.ComparisonOp{
				Op:  spansql.Eq,
				LHS: append(prefix, spansql.ID("AlbumId")),
				RHS: spansql.IntegerLiteral(k.AlbumId),
			},
		}
	}
	return spansql.Paren{Expr: b}
}
