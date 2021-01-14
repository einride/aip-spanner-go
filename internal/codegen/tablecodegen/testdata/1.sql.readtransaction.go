// +build testdata.1.sql.readtransaction

package testdata

// Code generated by TestReadTransactionCodeGenerator_GenerateCode/readtransaction/testdata/1.sql. DO NOT EDIT.

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SingersReadTransaction struct {
	Tx SpannerReadTransaction
}

func Singers(tx SpannerReadTransaction) SingersReadTransaction {
	return SingersReadTransaction{Tx: tx}
}

func (t SingersReadTransaction) Read(
	ctx context.Context,
	keySet spanner.KeySet,
) *SingersRowIterator {
	return &SingersRowIterator{
		RowIterator: t.Tx.Read(
			ctx,
			"Singers",
			keySet,
			((*SingersRow)(nil)).ColumnNames(),
		),
	}
}

func (t SingersReadTransaction) Get(
	ctx context.Context,
	key SingersPrimaryKey,
) (*SingersRow, error) {
	spannerRow, err := t.Tx.ReadRow(
		ctx,
		"Singers",
		key.SpannerKey(),
		((*SingersRow)(nil)).ColumnNames(),
	)
	if err != nil {
		return nil, err
	}
	var row SingersRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

func (t SingersReadTransaction) BatchGet(
	ctx context.Context,
	keys []SingersPrimaryKey,
) ([]*SingersRow, error) {
	spannerKeys := make([]spanner.KeySet, 0, len(keys))
	for _, key := range keys {
		spannerKeys = append(spannerKeys, key.SpannerKey())
	}
	it := t.Read(ctx, spanner.KeySets(spannerKeys...))
	defer it.Stop()
	foundRows := make(map[SingersPrimaryKey]*SingersRow, len(keys))
	if err := it.Do(func(row *SingersRow) error {
		foundRows[row.PrimaryKey()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	rows := make([]*SingersRow, 0, len(keys))
	for _, key := range keys {
		row, ok := foundRows[key]
		if !ok {
			return nil, status.Errorf(codes.NotFound, "not found: %v", key)
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func (t SingersReadTransaction) List(
	ctx context.Context,
	query ListQuery,
) *SingersRowIterator {
	stmt := spanner.Statement{
		SQL: spansql.Query{
			Select: spansql.Select{
				List: ((*SingersRow)(nil)).ColumnExprs(),
				From: []spansql.SelectFrom{
					spansql.SelectFromTable{Table: "Singers"},
				},
				Where: query.Where,
			},
			Order:  query.Order,
			Limit:  spansql.Param("limit"),
			Offset: spansql.Param("offset"),
		}.SQL(),
		Params: map[string]interface{}{
			"limit":  query.Limit,
			"offset": query.Offset,
		},
	}
	return &SingersRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

type SingersRowIterator struct {
	*spanner.RowIterator
}

func (i *SingersRowIterator) Next() (*SingersRow, error) {
	spannerRow, err := i.RowIterator.Next()
	if err != nil {
		return nil, err
	}
	var row SingersRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

func (i *SingersRowIterator) Do(f func(row *SingersRow) error) error {
	return i.RowIterator.Do(func(spannerRow *spanner.Row) error {
		var row SingersRow
		if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
			return err
		}
		return f(&row)
	})
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

func (r *SingersRow) Mutation() (string, []string, []interface{}) {
	return "Singers", r.ColumnNames(), []interface{}{
		r.SingerId,
		r.FirstName,
		r.LastName,
		r.SingerInfo,
	}
}

func (r *SingersRow) MutationForColumns(columns []string) (string, []string, []interface{}) {
	if len(columns) == 0 {
		columns = r.ColumnNames()
	}
	values := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		switch column {
		case "SingerId":
			values = append(values, r.SingerId)
		case "FirstName":
			values = append(values, r.FirstName)
		case "LastName":
			values = append(values, r.LastName)
		case "SingerInfo":
			values = append(values, r.SingerInfo)
		default:
			panic(fmt.Errorf("table Singers does not have column %s", column))
		}
	}
	return "Singers", columns, values
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

func (k SingersPrimaryKey) Delete() *spanner.Mutation {
	return spanner.Delete("Singers", k.SpannerKey())
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
