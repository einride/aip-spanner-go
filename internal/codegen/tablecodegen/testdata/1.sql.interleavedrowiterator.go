// +build testdata.1.sql.interleavedrowiterator

package testdata

// Code generated by TestInterleavedRowIteratorCodeGenerator_GenerateCode/interleavedrowiterator/testdata/1.sql. DO NOT EDIT.

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
)

type SingersRow struct {
	SingerId   int64              `spanner:"SingerId"`
	FirstName  spanner.NullString `spanner:"FirstName"`
	LastName   spanner.NullString `spanner:"LastName"`
	SingerInfo []uint8            `spanner:"SingerInfo"`
}

func (*SingersRow) ColumnNames() []string {
	return []string{
		"SingerId",
		"FirstName",
		"LastName",
		"SingerInfo",
	}
}

func (*SingersRow) ColumnIDs() []spansql.ID {
	return []spansql.ID{
		"SingerId",
		"FirstName",
		"LastName",
		"SingerInfo",
	}
}

func (*SingersRow) ColumnExprs() []spansql.Expr {
	return []spansql.Expr{
		spansql.ID("SingerId"),
		spansql.ID("FirstName"),
		spansql.ID("LastName"),
		spansql.ID("SingerInfo"),
	}
}

func (r *SingersRow) Validate() error {
	if !r.FirstName.IsNull() && len(r.FirstName.StringVal) > 1024 {
		return fmt.Errorf("column FirstName length > 1024")
	}
	if !r.LastName.IsNull() && len(r.LastName.StringVal) > 1024 {
		return fmt.Errorf("column LastName length > 1024")
	}
	return nil
}

func (r *SingersRow) UnmarshalSpannerRow(row *spanner.Row) error {
	for i := 0; i < row.Size(); i++ {
		switch row.ColumnName(i) {
		case "SingerId":
			if err := row.Column(i, &r.SingerId); err != nil {
				return fmt.Errorf("unmarshal Singers row: SingerId column: %w", err)
			}
		case "FirstName":
			if err := row.Column(i, &r.FirstName); err != nil {
				return fmt.Errorf("unmarshal Singers row: FirstName column: %w", err)
			}
		case "LastName":
			if err := row.Column(i, &r.LastName); err != nil {
				return fmt.Errorf("unmarshal Singers row: LastName column: %w", err)
			}
		case "SingerInfo":
			if err := row.Column(i, &r.SingerInfo); err != nil {
				return fmt.Errorf("unmarshal Singers row: SingerInfo column: %w", err)
			}
		default:
			return fmt.Errorf("unmarshal Singers row: unhandled column: %s", row.ColumnName(i))
		}
	}
	return nil
}

func (r *SingersRow) MutationForColumns(columns []string) (string, []string, []interface{}) {
	var values []interface{}
	return "Singers", columns, values
}

func (r *SingersRow) Mutation() (string, []string, []interface{}) {
	return r.MutationForColumns(r.ColumnNames())
}

func (r *SingersRow) Insert() *spanner.Mutation {
	return spanner.Insert(r.Mutation())
}

func (r *SingersRow) InsertOrUpdate() *spanner.Mutation {
	return spanner.InsertOrUpdate(r.Mutation())
}

func (r *SingersRow) Update() *spanner.Mutation {
	return spanner.Update(r.Mutation())
}

func (r *SingersRow) InsertColumns(columns []string) *spanner.Mutation {
	return spanner.Insert(r.MutationForColumns(columns))
}

func (r *SingersRow) InsertOrUpdateColumns(columns []string) *spanner.Mutation {
	return spanner.InsertOrUpdate(r.MutationForColumns(columns))
}

func (r *SingersRow) UpdateColumns(columns []string) *spanner.Mutation {
	return spanner.Update(r.MutationForColumns(columns))
}

func (r *SingersRow) PrimaryKey() SingersPrimaryKey {
	return SingersPrimaryKey{
		SingerId: r.SingerId,
	}
}

type SingersPrimaryKey struct {
	SingerId int64
}

func (k SingersPrimaryKey) SpannerKey() spanner.Key {
	return spanner.Key{
		k.SingerId,
	}
}

func (k SingersPrimaryKey) SpannerKeySet() spanner.KeySet {
	return k.SpannerKey()
}

func (k SingersPrimaryKey) BoolExpr() spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("SingerId"),
		RHS: spansql.StringLiteral(k.SingerId),
	})
	return spansql.Paren{Expr: b}
}

func (k SingersPrimaryKey) QualifiedBoolExpr(prefix spansql.PathExp) spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: append(prefix, spansql.ID("SingerId")),
		RHS: spansql.StringLiteral(k.SingerId),
	})
	return spansql.Paren{Expr: b}
}

type ListQuery struct {
	Where  spansql.BoolExpr
	Order  []spansql.Order
	Limit  int64
	Offset int64
}

type SpannerReadTransaction interface {
	Read(ctx context.Context, table string, keys spanner.KeySet, columns []string) *spanner.RowIterator
	ReadRow(ctx context.Context, table string, key spanner.Key, columns []string) (*spanner.Row, error)
	Query(ctx context.Context, statement spanner.Statement) *spanner.RowIterator
}
