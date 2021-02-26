// +build testdata.6.sql.database

package testdata

// Code generated by TestDatabaseCodeGenerator_GenerateCode/database/testdata/6.sql. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShippersRow struct {
	ShipperId  string           `spanner:"shipper_id"`
	CreateTime time.Time        `spanner:"create_time"`
	UpdateTime time.Time        `spanner:"update_time"`
	DeleteTime spanner.NullTime `spanner:"delete_time"`
	Shipments  []*ShipmentsRow  `spanner:"shipments"`
}

func (*ShippersRow) ColumnNames() []string {
	return []string{
		"shipper_id",
		"create_time",
		"update_time",
		"delete_time",
	}
}

func (*ShippersRow) ColumnIDs() []spansql.ID {
	return []spansql.ID{
		"shipper_id",
		"create_time",
		"update_time",
		"delete_time",
	}
}

func (*ShippersRow) ColumnExprs() []spansql.Expr {
	return []spansql.Expr{
		spansql.ID("shipper_id"),
		spansql.ID("create_time"),
		spansql.ID("update_time"),
		spansql.ID("delete_time"),
	}
}

func (r *ShippersRow) Validate() error {
	if len(r.ShipperId) > 63 {
		return fmt.Errorf("column shipper_id length > 63")
	}
	return nil
}

func (r *ShippersRow) UnmarshalSpannerRow(row *spanner.Row) error {
	for i := 0; i < row.Size(); i++ {
		switch row.ColumnName(i) {
		case "shipper_id":
			if err := row.Column(i, &r.ShipperId); err != nil {
				return fmt.Errorf("unmarshal shippers row: shipper_id column: %w", err)
			}
		case "create_time":
			if err := row.Column(i, &r.CreateTime); err != nil {
				return fmt.Errorf("unmarshal shippers row: create_time column: %w", err)
			}
		case "update_time":
			if err := row.Column(i, &r.UpdateTime); err != nil {
				return fmt.Errorf("unmarshal shippers row: update_time column: %w", err)
			}
		case "delete_time":
			if err := row.Column(i, &r.DeleteTime); err != nil {
				return fmt.Errorf("unmarshal shippers row: delete_time column: %w", err)
			}
		case "shipments":
			if err := row.Column(i, &r.Shipments); err != nil {
				return fmt.Errorf("unmarshal shippers interleaved row: shipments column: %w", err)
			}
		default:
			return fmt.Errorf("unmarshal shippers row: unhandled column: %s", row.ColumnName(i))
		}
	}
	return nil
}

func (r *ShippersRow) Mutate() (string, []string, []interface{}) {
	return "shippers", r.ColumnNames(), []interface{}{
		r.ShipperId,
		r.CreateTime,
		r.UpdateTime,
		r.DeleteTime,
	}
}

func (r *ShippersRow) MutateColumns(columns []string) (string, []string, []interface{}) {
	if len(columns) == 0 {
		columns = r.ColumnNames()
	}
	values := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		switch column {
		case "shipper_id":
			values = append(values, r.ShipperId)
		case "create_time":
			values = append(values, r.CreateTime)
		case "update_time":
			values = append(values, r.UpdateTime)
		case "delete_time":
			values = append(values, r.DeleteTime)
		default:
			panic(fmt.Errorf("table shippers does not have column %s", column))
		}
	}
	return "shippers", columns, values
}

func (r *ShippersRow) Key() ShippersKey {
	return ShippersKey{
		ShipperId: r.ShipperId,
	}
}

type ShipmentsRow struct {
	ShipperId  string           `spanner:"shipper_id"`
	ShipmentId string           `spanner:"shipment_id"`
	CreateTime time.Time        `spanner:"create_time"`
	UpdateTime time.Time        `spanner:"update_time"`
	DeleteTime spanner.NullTime `spanner:"delete_time"`
}

func (*ShipmentsRow) ColumnNames() []string {
	return []string{
		"shipper_id",
		"shipment_id",
		"create_time",
		"update_time",
		"delete_time",
	}
}

func (*ShipmentsRow) ColumnIDs() []spansql.ID {
	return []spansql.ID{
		"shipper_id",
		"shipment_id",
		"create_time",
		"update_time",
		"delete_time",
	}
}

func (*ShipmentsRow) ColumnExprs() []spansql.Expr {
	return []spansql.Expr{
		spansql.ID("shipper_id"),
		spansql.ID("shipment_id"),
		spansql.ID("create_time"),
		spansql.ID("update_time"),
		spansql.ID("delete_time"),
	}
}

func (r *ShipmentsRow) Validate() error {
	if len(r.ShipperId) > 63 {
		return fmt.Errorf("column shipper_id length > 63")
	}
	if len(r.ShipmentId) > 63 {
		return fmt.Errorf("column shipment_id length > 63")
	}
	return nil
}

func (r *ShipmentsRow) UnmarshalSpannerRow(row *spanner.Row) error {
	for i := 0; i < row.Size(); i++ {
		switch row.ColumnName(i) {
		case "shipper_id":
			if err := row.Column(i, &r.ShipperId); err != nil {
				return fmt.Errorf("unmarshal shipments row: shipper_id column: %w", err)
			}
		case "shipment_id":
			if err := row.Column(i, &r.ShipmentId); err != nil {
				return fmt.Errorf("unmarshal shipments row: shipment_id column: %w", err)
			}
		case "create_time":
			if err := row.Column(i, &r.CreateTime); err != nil {
				return fmt.Errorf("unmarshal shipments row: create_time column: %w", err)
			}
		case "update_time":
			if err := row.Column(i, &r.UpdateTime); err != nil {
				return fmt.Errorf("unmarshal shipments row: update_time column: %w", err)
			}
		case "delete_time":
			if err := row.Column(i, &r.DeleteTime); err != nil {
				return fmt.Errorf("unmarshal shipments row: delete_time column: %w", err)
			}
		default:
			return fmt.Errorf("unmarshal shipments row: unhandled column: %s", row.ColumnName(i))
		}
	}
	return nil
}

func (r *ShipmentsRow) Mutate() (string, []string, []interface{}) {
	return "shipments", r.ColumnNames(), []interface{}{
		r.ShipperId,
		r.ShipmentId,
		r.CreateTime,
		r.UpdateTime,
		r.DeleteTime,
	}
}

func (r *ShipmentsRow) MutateColumns(columns []string) (string, []string, []interface{}) {
	if len(columns) == 0 {
		columns = r.ColumnNames()
	}
	values := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		switch column {
		case "shipper_id":
			values = append(values, r.ShipperId)
		case "shipment_id":
			values = append(values, r.ShipmentId)
		case "create_time":
			values = append(values, r.CreateTime)
		case "update_time":
			values = append(values, r.UpdateTime)
		case "delete_time":
			values = append(values, r.DeleteTime)
		default:
			panic(fmt.Errorf("table shipments does not have column %s", column))
		}
	}
	return "shipments", columns, values
}

func (r *ShipmentsRow) Key() ShipmentsKey {
	return ShipmentsKey{
		ShipperId:  r.ShipperId,
		ShipmentId: r.ShipmentId,
	}
}

type ShippersKey struct {
	ShipperId string
}

func (k ShippersKey) SpannerKey() spanner.Key {
	return spanner.Key{
		k.ShipperId,
	}
}

func (k ShippersKey) SpannerKeySet() spanner.KeySet {
	return k.SpannerKey()
}

func (k ShippersKey) Delete() *spanner.Mutation {
	return spanner.Delete("shippers", k.SpannerKey())
}

func (ShippersKey) Order() []spansql.Order {
	return []spansql.Order{
		{Expr: spansql.ID("shipper_id"), Desc: false},
	}
}

func (k ShippersKey) BoolExpr() spansql.BoolExpr {
	cmp0 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("shipper_id"),
		RHS: spansql.StringLiteral(k.ShipperId),
	})
	b := cmp0
	return spansql.Paren{Expr: b}
}

type ShipmentsKey struct {
	ShipperId  string
	ShipmentId string
}

func (k ShipmentsKey) SpannerKey() spanner.Key {
	return spanner.Key{
		k.ShipperId,
		k.ShipmentId,
	}
}

func (k ShipmentsKey) SpannerKeySet() spanner.KeySet {
	return k.SpannerKey()
}

func (k ShipmentsKey) Delete() *spanner.Mutation {
	return spanner.Delete("shipments", k.SpannerKey())
}

func (ShipmentsKey) Order() []spansql.Order {
	return []spansql.Order{
		{Expr: spansql.ID("shipper_id"), Desc: false},
		{Expr: spansql.ID("shipment_id"), Desc: false},
	}
}

func (k ShipmentsKey) BoolExpr() spansql.BoolExpr {
	cmp0 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("shipper_id"),
		RHS: spansql.StringLiteral(k.ShipperId),
	})
	cmp1 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("shipment_id"),
		RHS: spansql.StringLiteral(k.ShipmentId),
	})
	b := cmp0
	b = spansql.LogicalOp{
		Op:  spansql.And,
		LHS: b,
		RHS: cmp1,
	}
	return spansql.Paren{Expr: b}
}

type ShippersRowIterator interface {
	Next() (*ShippersRow, error)
	Do(f func(row *ShippersRow) error) error
	Stop()
}

type streamingShippersRowIterator struct {
	*spanner.RowIterator
}

func (i *streamingShippersRowIterator) Next() (*ShippersRow, error) {
	spannerRow, err := i.RowIterator.Next()
	if err != nil {
		return nil, err
	}
	var row ShippersRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

func (i *streamingShippersRowIterator) Do(f func(row *ShippersRow) error) error {
	return i.RowIterator.Do(func(spannerRow *spanner.Row) error {
		var row ShippersRow
		if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
			return err
		}
		return f(&row)
	})
}

type ShipmentsRowIterator interface {
	Next() (*ShipmentsRow, error)
	Do(f func(row *ShipmentsRow) error) error
	Stop()
}

type streamingShipmentsRowIterator struct {
	*spanner.RowIterator
}

func (i *streamingShipmentsRowIterator) Next() (*ShipmentsRow, error) {
	spannerRow, err := i.RowIterator.Next()
	if err != nil {
		return nil, err
	}
	var row ShipmentsRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

func (i *streamingShipmentsRowIterator) Do(f func(row *ShipmentsRow) error) error {
	return i.RowIterator.Do(func(spannerRow *spanner.Row) error {
		var row ShipmentsRow
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

func (t ReadTransaction) ReadShippersRows(
	ctx context.Context,
	keySet spanner.KeySet,
) ShippersRowIterator {
	return &streamingShippersRowIterator{
		RowIterator: t.Tx.Read(
			ctx,
			"shippers",
			keySet,
			((*ShippersRow)(nil)).ColumnNames(),
		),
	}
}

type GetShippersRowQuery struct {
	Key       ShippersKey
	Shipments bool
}

func (q *GetShippersRowQuery) hasInterleavedTables() bool {
	return q.Shipments
}

func (t ReadTransaction) GetShippersRow(
	ctx context.Context,
	query GetShippersRowQuery,
) (*ShippersRow, error) {
	if query.hasInterleavedTables() {
		return t.getShippersRowInterleaved(ctx, query)
	}
	spannerRow, err := t.Tx.ReadRow(
		ctx,
		"shippers",
		query.Key.SpannerKey(),
		((*ShippersRow)(nil)).ColumnNames(),
	)
	if err != nil {
		return nil, err
	}
	var row ShippersRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

type BatchGetShippersRowsQuery struct {
	Keys      []ShippersKey
	Shipments bool
}

func (q *BatchGetShippersRowsQuery) hasInterleavedTables() bool {
	return q.Shipments
}

func (t ReadTransaction) BatchGetShippersRows(
	ctx context.Context,
	query BatchGetShippersRowsQuery,
) (map[ShippersKey]*ShippersRow, error) {
	if query.hasInterleavedTables() {
		return t.batchGetShippersRowsInterleaved(ctx, query)
	}
	spannerKeys := make([]spanner.KeySet, 0, len(query.Keys))
	for _, key := range query.Keys {
		spannerKeys = append(spannerKeys, key.SpannerKey())
	}
	foundRows := make(map[ShippersKey]*ShippersRow, len(query.Keys))
	if err := t.ReadShippersRows(ctx, spanner.KeySets(spannerKeys...)).Do(func(row *ShippersRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

type ListShippersRowsQuery struct {
	Where       spansql.BoolExpr
	Order       []spansql.Order
	Limit       int32
	Offset      int64
	Params      map[string]interface{}
	ShowDeleted bool
	Shipments   bool
}

func (q *ListShippersRowsQuery) hasInterleavedTables() bool {
	return q.Shipments
}

func (t ReadTransaction) ListShippersRows(
	ctx context.Context,
	query ListShippersRowsQuery,
) ShippersRowIterator {
	if query.hasInterleavedTables() {
		return t.listShippersRowsInterleaved(ctx, query)
	}
	if len(query.Order) == 0 {
		query.Order = ShippersKey{}.Order()
	}
	params := make(map[string]interface{}, len(query.Params)+2)
	params["__limit"] = int64(query.Limit)
	params["__offset"] = int64(query.Offset)
	for param, value := range query.Params {
		if _, ok := params[param]; ok {
			panic(fmt.Errorf("invalid param: %s", param))
		}
		params[param] = value
	}
	if query.Where == nil {
		query.Where = spansql.True
	}
	if !query.ShowDeleted {
		query.Where = spansql.LogicalOp{
			Op:  spansql.And,
			LHS: spansql.Paren{Expr: query.Where},
			RHS: spansql.IsOp{
				LHS: spansql.ID("delete_time"),
				RHS: spansql.Null,
			},
		}
	}
	stmt := spanner.Statement{
		SQL: spansql.Query{
			Select: spansql.Select{
				List: ((*ShippersRow)(nil)).ColumnExprs(),
				From: []spansql.SelectFrom{
					spansql.SelectFromTable{Table: "shippers"},
				},
				Where: query.Where,
			},
			Order:  query.Order,
			Limit:  spansql.Param("__limit"),
			Offset: spansql.Param("__offset"),
		}.SQL(),
		Params: params,
	}
	return &streamingShippersRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

func (t ReadTransaction) listShippersRowsInterleaved(
	ctx context.Context,
	query ListShippersRowsQuery,
) ShippersRowIterator {
	if len(query.Order) == 0 {
		query.Order = ShippersKey{}.Order()
	}
	var q strings.Builder
	_, _ = q.WriteString(`
SELECT
    shipper_id,
    create_time,
    update_time,
    delete_time,
`)
	if query.Shipments {
		_, _ = q.WriteString(`
    ARRAY(
        SELECT AS STRUCT
            shipper_id,
            shipment_id,
            create_time,
            update_time,
            delete_time,
`)
		_, _ = q.WriteString(`
        FROM 
            shipments
        WHERE 
`)
		if !query.ShowDeleted {
			_, _ = q.WriteString(`
            delete_time IS NULL AND
`)
		}
		_, _ = q.WriteString(`
            shipments.shipper_id = shippers.shipper_id
        ORDER BY 
            shipper_id,
            shipment_id
    ) AS shipments,
`)
	}
	_, _ = q.WriteString(`
FROM
    shippers
`)
	if query.Where == nil {
		query.Where = spansql.True
	}
	if !query.ShowDeleted {
		query.Where = spansql.LogicalOp{
			Op:  spansql.And,
			LHS: spansql.Paren{Expr: query.Where},
			RHS: spansql.IsOp{
				LHS: spansql.ID("delete_time"),
				RHS: spansql.Null,
			},
		}
	}
	_, _ = q.WriteString("WHERE (")
	_, _ = q.WriteString(query.Where.SQL())
	_, _ = q.WriteString(") ")
	if len(query.Order) > 0 {
		_, _ = q.WriteString("ORDER BY ")
		for i, order := range query.Order {
			_, _ = q.WriteString(order.SQL())
			if i < len(query.Order)-1 {
				_, _ = q.WriteString(", ")
			} else {
				_, _ = q.WriteString(" ")
			}
		}
	}
	_, _ = q.WriteString("LIMIT @__limit ")
	_, _ = q.WriteString("OFFSET @__offset ")
	params := make(map[string]interface{}, len(query.Params)+2)
	params["__limit"] = int64(query.Limit)
	params["__offset"] = int64(query.Offset)
	for param, value := range query.Params {
		if _, ok := params[param]; ok {
			panic(fmt.Errorf("invalid param: %s", param))
		}
		params[param] = value
	}
	stmt := spanner.Statement{
		SQL:    q.String(),
		Params: params,
	}
	return &streamingShippersRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

func (t ReadTransaction) getShippersRowInterleaved(
	ctx context.Context,
	query GetShippersRowQuery,
) (*ShippersRow, error) {
	it := t.listShippersRowsInterleaved(ctx, ListShippersRowsQuery{
		Limit:       1,
		Where:       query.Key.BoolExpr(),
		ShowDeleted: true,
		Shipments:   query.Shipments,
	})
	defer it.Stop()
	row, err := it.Next()
	if err != nil {
		if err == iterator.Done {
			return nil, status.Errorf(codes.NotFound, "not found: %v", query.Key)
		}
		return nil, err
	}
	return row, nil
}

func (t ReadTransaction) batchGetShippersRowsInterleaved(
	ctx context.Context,
	query BatchGetShippersRowsQuery,
) (map[ShippersKey]*ShippersRow, error) {
	if len(query.Keys) == 0 {
		return nil, nil
	}
	where := query.Keys[0].BoolExpr()
	for _, key := range query.Keys[1:] {
		where = spansql.LogicalOp{
			Op:  spansql.Or,
			LHS: where,
			RHS: key.BoolExpr(),
		}
	}
	foundRows := make(map[ShippersKey]*ShippersRow, len(query.Keys))
	if err := t.ListShippersRows(ctx, ListShippersRowsQuery{
		Where:       spansql.Paren{Expr: where},
		Limit:       int32(len(query.Keys)),
		ShowDeleted: true,
		Shipments:   query.Shipments,
	}).Do(func(row *ShippersRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

func (t ReadTransaction) ReadShipmentsRows(
	ctx context.Context,
	keySet spanner.KeySet,
) ShipmentsRowIterator {
	return &streamingShipmentsRowIterator{
		RowIterator: t.Tx.Read(
			ctx,
			"shipments",
			keySet,
			((*ShipmentsRow)(nil)).ColumnNames(),
		),
	}
}

type GetShipmentsRowQuery struct {
	Key ShipmentsKey
}

func (t ReadTransaction) GetShipmentsRow(
	ctx context.Context,
	query GetShipmentsRowQuery,
) (*ShipmentsRow, error) {
	spannerRow, err := t.Tx.ReadRow(
		ctx,
		"shipments",
		query.Key.SpannerKey(),
		((*ShipmentsRow)(nil)).ColumnNames(),
	)
	if err != nil {
		return nil, err
	}
	var row ShipmentsRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

type BatchGetShipmentsRowsQuery struct {
	Keys []ShipmentsKey
}

func (t ReadTransaction) BatchGetShipmentsRows(
	ctx context.Context,
	query BatchGetShipmentsRowsQuery,
) (map[ShipmentsKey]*ShipmentsRow, error) {
	spannerKeys := make([]spanner.KeySet, 0, len(query.Keys))
	for _, key := range query.Keys {
		spannerKeys = append(spannerKeys, key.SpannerKey())
	}
	foundRows := make(map[ShipmentsKey]*ShipmentsRow, len(query.Keys))
	if err := t.ReadShipmentsRows(ctx, spanner.KeySets(spannerKeys...)).Do(func(row *ShipmentsRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

type ListShipmentsRowsQuery struct {
	Where       spansql.BoolExpr
	Order       []spansql.Order
	Limit       int32
	Offset      int64
	Params      map[string]interface{}
	ShowDeleted bool
}

func (t ReadTransaction) ListShipmentsRows(
	ctx context.Context,
	query ListShipmentsRowsQuery,
) ShipmentsRowIterator {
	if len(query.Order) == 0 {
		query.Order = ShipmentsKey{}.Order()
	}
	params := make(map[string]interface{}, len(query.Params)+2)
	params["__limit"] = int64(query.Limit)
	params["__offset"] = int64(query.Offset)
	for param, value := range query.Params {
		if _, ok := params[param]; ok {
			panic(fmt.Errorf("invalid param: %s", param))
		}
		params[param] = value
	}
	if query.Where == nil {
		query.Where = spansql.True
	}
	if !query.ShowDeleted {
		query.Where = spansql.LogicalOp{
			Op:  spansql.And,
			LHS: spansql.Paren{Expr: query.Where},
			RHS: spansql.IsOp{
				LHS: spansql.ID("delete_time"),
				RHS: spansql.Null,
			},
		}
	}
	stmt := spanner.Statement{
		SQL: spansql.Query{
			Select: spansql.Select{
				List: ((*ShipmentsRow)(nil)).ColumnExprs(),
				From: []spansql.SelectFrom{
					spansql.SelectFromTable{Table: "shipments"},
				},
				Where: query.Where,
			},
			Order:  query.Order,
			Limit:  spansql.Param("__limit"),
			Offset: spansql.Param("__offset"),
		}.SQL(),
		Params: params,
	}
	return &streamingShipmentsRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

type SpannerReadTransaction interface {
	Read(ctx context.Context, table string, keys spanner.KeySet, columns []string) *spanner.RowIterator
	ReadRow(ctx context.Context, table string, key spanner.Key, columns []string) (*spanner.Row, error)
	Query(ctx context.Context, statement spanner.Statement) *spanner.RowIterator
}
