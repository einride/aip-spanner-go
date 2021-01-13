// +build testdata.2.sql.interleavedrowiterator

package testdata

// Code generated by TestInterleavedRowIteratorCodeGenerator_GenerateCode/interleavedrowiterator/testdata/2.sql. DO NOT EDIT.

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
)

type SingersAndAlbumsRowIterator struct {
	*spanner.RowIterator
}

func (i *SingersAndAlbumsRowIterator) Next() (*SingersAndAlbumsRow, error) {
	spannerRow, err := i.RowIterator.Next()
	if err != nil {
		return nil, err
	}
	var row SingersAndAlbumsRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

func (i *SingersAndAlbumsRowIterator) Do(f func(row *SingersAndAlbumsRow) error) error {
	return i.RowIterator.Do(func(spannerRow *spanner.Row) error {
		var row SingersAndAlbumsRow
		if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
			return err
		}
		return f(&row)
	})
}

type SingersAndAlbumsRow struct {
	SingerId   int64              `spanner:"SingerId"`
	FirstName  spanner.NullString `spanner:"FirstName"`
	LastName   spanner.NullString `spanner:"LastName"`
	SingerInfo []uint8            `spanner:"SingerInfo"`
	Albums     []*AlbumsRow       `spanner:"Albums"`
}

func (r *SingersAndAlbumsRow) PrimaryKey() SingersPrimaryKey {
	return SingersPrimaryKey{
		SingerId: r.SingerId,
	}
}

func (r *SingersAndAlbumsRow) UnmarshalSpannerRow(row *spanner.Row) error {
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
		case "Albums":
			if err := row.Column(i, &r.Albums); err != nil {
				return fmt.Errorf("unmarshal Singers interleaved row: Albums column: %w", err)
			}
		default:
			return fmt.Errorf("unmarshal Singers row: unhandled column: %s", row.ColumnName(i))
		}
	}
	return nil
}

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
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	return spansql.Paren{Expr: b}
}

func (k SingersPrimaryKey) QualifiedBoolExpr(prefix spansql.PathExp) spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: append(prefix, spansql.ID("SingerId")),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	return spansql.Paren{Expr: b}
}

type AlbumsRow struct {
	SingerId   int64              `spanner:"SingerId"`
	AlbumId    int64              `spanner:"AlbumId"`
	AlbumTitle spanner.NullString `spanner:"AlbumTitle"`
}

func (*AlbumsRow) ColumnNames() []string {
	return []string{
		"SingerId",
		"AlbumId",
		"AlbumTitle",
	}
}

func (*AlbumsRow) ColumnIDs() []spansql.ID {
	return []spansql.ID{
		"SingerId",
		"AlbumId",
		"AlbumTitle",
	}
}

func (*AlbumsRow) ColumnExprs() []spansql.Expr {
	return []spansql.Expr{
		spansql.ID("SingerId"),
		spansql.ID("AlbumId"),
		spansql.ID("AlbumTitle"),
	}
}

func (r *AlbumsRow) Validate() error {
	return nil
}

func (r *AlbumsRow) UnmarshalSpannerRow(row *spanner.Row) error {
	for i := 0; i < row.Size(); i++ {
		switch row.ColumnName(i) {
		case "SingerId":
			if err := row.Column(i, &r.SingerId); err != nil {
				return fmt.Errorf("unmarshal Albums row: SingerId column: %w", err)
			}
		case "AlbumId":
			if err := row.Column(i, &r.AlbumId); err != nil {
				return fmt.Errorf("unmarshal Albums row: AlbumId column: %w", err)
			}
		case "AlbumTitle":
			if err := row.Column(i, &r.AlbumTitle); err != nil {
				return fmt.Errorf("unmarshal Albums row: AlbumTitle column: %w", err)
			}
		default:
			return fmt.Errorf("unmarshal Albums row: unhandled column: %s", row.ColumnName(i))
		}
	}
	return nil
}

func (r *AlbumsRow) MutationForColumns(columns []string) (string, []string, []interface{}) {
	var values []interface{}
	return "Albums", columns, values
}

func (r *AlbumsRow) Mutation() (string, []string, []interface{}) {
	return r.MutationForColumns(r.ColumnNames())
}

func (r *AlbumsRow) Insert() *spanner.Mutation {
	return spanner.Insert(r.Mutation())
}

func (r *AlbumsRow) InsertOrUpdate() *spanner.Mutation {
	return spanner.InsertOrUpdate(r.Mutation())
}

func (r *AlbumsRow) Update() *spanner.Mutation {
	return spanner.Update(r.Mutation())
}

func (r *AlbumsRow) InsertColumns(columns []string) *spanner.Mutation {
	return spanner.Insert(r.MutationForColumns(columns))
}

func (r *AlbumsRow) InsertOrUpdateColumns(columns []string) *spanner.Mutation {
	return spanner.InsertOrUpdate(r.MutationForColumns(columns))
}

func (r *AlbumsRow) UpdateColumns(columns []string) *spanner.Mutation {
	return spanner.Update(r.MutationForColumns(columns))
}

func (r *AlbumsRow) PrimaryKey() AlbumsPrimaryKey {
	return AlbumsPrimaryKey{
		SingerId: r.SingerId,
		AlbumId:  r.AlbumId,
	}
}

type AlbumsPrimaryKey struct {
	SingerId int64
	AlbumId  int64
}

func (k AlbumsPrimaryKey) SpannerKey() spanner.Key {
	return spanner.Key{
		k.SingerId,
		k.AlbumId,
	}
}

func (k AlbumsPrimaryKey) SpannerKeySet() spanner.KeySet {
	return k.SpannerKey()
}

func (k AlbumsPrimaryKey) BoolExpr() spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("SingerId"),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	b = spansql.LogicalOp{
		Op:  spansql.And,
		LHS: b,
		RHS: spansql.ComparisonOp{
			Op:  spansql.Eq,
			LHS: spansql.ID("AlbumId"),
			RHS: spansql.IntegerLiteral(k.AlbumId),
		},
	}
	return spansql.Paren{Expr: b}
}

func (k AlbumsPrimaryKey) QualifiedBoolExpr(prefix spansql.PathExp) spansql.BoolExpr {
	b := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: append(prefix, spansql.ID("SingerId")),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	b = spansql.LogicalOp{
		Op:  spansql.And,
		LHS: b,
		RHS: spansql.ComparisonOp{
			Op:  spansql.Eq,
			LHS: append(prefix, spansql.ID("AlbumId")),
			RHS: spansql.IntegerLiteral(k.AlbumId),
		},
	}
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
