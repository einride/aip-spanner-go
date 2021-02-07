// +build testdata.5.sql.database

package testdata

// Code generated by TestDatabaseCodeGenerator_GenerateCode/database/testdata/5.sql. DO NOT EDIT.

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
)

type UserAccessLogRow struct {
	UserId     int64     `spanner:"UserId"`
	LastAccess time.Time `spanner:"LastAccess"`
}

func (*UserAccessLogRow) ColumnNames() []string {
	return []string{
		"UserId",
		"LastAccess",
	}
}

func (*UserAccessLogRow) ColumnIDs() []spansql.ID {
	return []spansql.ID{
		"UserId",
		"LastAccess",
	}
}

func (*UserAccessLogRow) ColumnExprs() []spansql.Expr {
	return []spansql.Expr{
		spansql.ID("UserId"),
		spansql.ID("LastAccess"),
	}
}

func (r *UserAccessLogRow) Validate() error {
	return nil
}

func (r *UserAccessLogRow) UnmarshalSpannerRow(row *spanner.Row) error {
	for i := 0; i < row.Size(); i++ {
		switch row.ColumnName(i) {
		case "UserId":
			if err := row.Column(i, &r.UserId); err != nil {
				return fmt.Errorf("unmarshal UserAccessLog row: UserId column: %w", err)
			}
		case "LastAccess":
			if err := row.Column(i, &r.LastAccess); err != nil {
				return fmt.Errorf("unmarshal UserAccessLog row: LastAccess column: %w", err)
			}
		default:
			return fmt.Errorf("unmarshal UserAccessLog row: unhandled column: %s", row.ColumnName(i))
		}
	}
	return nil
}

func (r *UserAccessLogRow) Mutate() (string, []string, []interface{}) {
	return "UserAccessLog", r.ColumnNames(), []interface{}{
		r.UserId,
		r.LastAccess,
	}
}

func (r *UserAccessLogRow) MutateColumns(columns []string) (string, []string, []interface{}) {
	if len(columns) == 0 {
		columns = r.ColumnNames()
	}
	values := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		switch column {
		case "UserId":
			values = append(values, r.UserId)
		case "LastAccess":
			values = append(values, r.LastAccess)
		default:
			panic(fmt.Errorf("table UserAccessLog does not have column %s", column))
		}
	}
	return "UserAccessLog", columns, values
}

func (r *UserAccessLogRow) Key() UserAccessLogKey {
	return UserAccessLogKey{
		UserId:     r.UserId,
		LastAccess: r.LastAccess,
	}
}

type UserAccessLogKey struct {
	UserId     int64
	LastAccess time.Time
}

func (k UserAccessLogKey) SpannerKey() spanner.Key {
	return spanner.Key{
		k.UserId,
		k.LastAccess,
	}
}

func (k UserAccessLogKey) SpannerKeySet() spanner.KeySet {
	return k.SpannerKey()
}

func (k UserAccessLogKey) Delete() *spanner.Mutation {
	return spanner.Delete("UserAccessLog", k.SpannerKey())
}

func (UserAccessLogKey) Order() []spansql.Order {
	return []spansql.Order{
		{Expr: spansql.ID("UserId"), Desc: false},
		{Expr: spansql.ID("LastAccess"), Desc: false},
	}
}

func (k UserAccessLogKey) BoolExpr() spansql.BoolExpr {
	cmp0 := spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("UserId"),
		RHS: spansql.IntegerLiteral(k.UserId),
	}
	cmp1 := spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("LastAccess"),
		RHS: spansql.TimestampLiteral(k.LastAccess),
	}
	b := spansql.BoolExpr(cmp0)
	b = spansql.LogicalOp{
		Op:  spansql.And,
		LHS: b,
		RHS: cmp1,
	}
	return spansql.Paren{Expr: b}
}

type UserAccessLogRowIterator struct {
	*spanner.RowIterator
}

func (i *UserAccessLogRowIterator) Next() (*UserAccessLogRow, error) {
	spannerRow, err := i.RowIterator.Next()
	if err != nil {
		return nil, err
	}
	var row UserAccessLogRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

func (i *UserAccessLogRowIterator) Do(f func(row *UserAccessLogRow) error) error {
	return i.RowIterator.Do(func(spannerRow *spanner.Row) error {
		var row UserAccessLogRow
		if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
			return err
		}
		return f(&row)
	})
}

type ReadTransaction struct {
	Tx SpannerReadTransaction
}

func Query(tx SpannerReadTransaction) ReadTransaction {
	return ReadTransaction{Tx: tx}
}

func (t ReadTransaction) ReadUserAccessLogRows(
	ctx context.Context,
	keySet spanner.KeySet,
) *UserAccessLogRowIterator {
	return &UserAccessLogRowIterator{
		RowIterator: t.Tx.Read(
			ctx,
			"UserAccessLog",
			keySet,
			((*UserAccessLogRow)(nil)).ColumnNames(),
		),
	}
}

type GetUserAccessLogRowQuery struct {
	Key UserAccessLogKey
}

func (t ReadTransaction) GetUserAccessLogRow(
	ctx context.Context,
	query GetUserAccessLogRowQuery,
) (*UserAccessLogRow, error) {
	spannerRow, err := t.Tx.ReadRow(
		ctx,
		"UserAccessLog",
		query.Key.SpannerKey(),
		((*UserAccessLogRow)(nil)).ColumnNames(),
	)
	if err != nil {
		return nil, err
	}
	var row UserAccessLogRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

type BatchGetUserAccessLogRowsQuery struct {
	Keys []UserAccessLogKey
}

func (t ReadTransaction) BatchGetUserAccessLogRows(
	ctx context.Context,
	query BatchGetUserAccessLogRowsQuery,
) (map[UserAccessLogKey]*UserAccessLogRow, error) {
	spannerKeys := make([]spanner.KeySet, 0, len(query.Keys))
	for _, key := range query.Keys {
		spannerKeys = append(spannerKeys, key.SpannerKey())
	}
	foundRows := make(map[UserAccessLogKey]*UserAccessLogRow, len(query.Keys))
	if err := t.ReadUserAccessLogRows(ctx, spanner.KeySets(spannerKeys...)).Do(func(row *UserAccessLogRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

type ListUserAccessLogRowsQuery struct {
	Where  spansql.BoolExpr
	Order  []spansql.Order
	Limit  int32
	Offset int64
	Params map[string]interface{}
}

func (t ReadTransaction) ListUserAccessLogRows(
	ctx context.Context,
	query ListUserAccessLogRowsQuery,
) *UserAccessLogRowIterator {
	if len(query.Order) == 0 {
		query.Order = UserAccessLogKey{}.Order()
	}
	params := map[string]interface{}{
		"__limit":  int64(query.Limit),
		"__offset": query.Offset,
	}
	for param, value := range query.Params {
		if _, ok := params[param]; ok {
			panic(fmt.Errorf("invalid param: %s", param))
		}
		params[param] = value
	}
	if query.Where == nil {
		query.Where = spansql.True
	}
	stmt := spanner.Statement{
		SQL: spansql.Query{
			Select: spansql.Select{
				List: ((*UserAccessLogRow)(nil)).ColumnExprs(),
				From: []spansql.SelectFrom{
					spansql.SelectFromTable{Table: "UserAccessLog"},
				},
				Where: query.Where,
			},
			Order:  query.Order,
			Limit:  spansql.Param("__limit"),
			Offset: spansql.Param("__offset"),
		}.SQL(),
		Params: params,
	}
	return &UserAccessLogRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

type SpannerReadTransaction interface {
	Read(ctx context.Context, table string, keys spanner.KeySet, columns []string) *spanner.RowIterator
	ReadRow(ctx context.Context, table string, key spanner.Key, columns []string) (*spanner.Row, error)
	Query(ctx context.Context, statement spanner.Statement) *spanner.RowIterator
}
