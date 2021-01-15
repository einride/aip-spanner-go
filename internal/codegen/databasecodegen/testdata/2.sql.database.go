// +build testdata.2.sql.database

package testdata

// Code generated by TestDatabaseCodeGenerator_GenerateCode/database/testdata/2.sql. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
	"google.golang.org/api/iterator"
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
	key SingersKey,
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
	keys []SingersKey,
) (map[SingersKey]*SingersRow, error) {
	spannerKeys := make([]spanner.KeySet, 0, len(keys))
	for _, key := range keys {
		spannerKeys = append(spannerKeys, key.SpannerKey())
	}
	foundRows := make(map[SingersKey]*SingersRow, len(keys))
	if err := t.Read(ctx, spanner.KeySets(spannerKeys...)).Do(func(row *SingersRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

func (t SingersReadTransaction) List(
	ctx context.Context,
	query ListQuery,
) *SingersRowIterator {
	if len(query.Order) == 0 {
		query.Order = SingersKey{}.Order()
	}
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

func (SingersKey) Order() []spansql.Order {
	return []spansql.Order{
		{Expr: spansql.ID("SingerId"), Desc: false},
	}
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

func (r *SingersRow) Key() SingersKey {
	return SingersKey{
		SingerId: r.SingerId,
	}
}

type SingersAndAlbumsReadTransaction struct {
	Tx SpannerReadTransaction
}

func SingersAndAlbums(tx SpannerReadTransaction) SingersAndAlbumsReadTransaction {
	return SingersAndAlbumsReadTransaction{Tx: tx}
}

func (t SingersAndAlbumsReadTransaction) List(
	ctx context.Context,
	query ListQuery,
) *SingersAndAlbumsRowIterator {
	if len(query.Order) == 0 {
		query.Order = SingersKey{}.Order()
	}
	var q strings.Builder
	_, _ = q.WriteString("SELECT ")
	_, _ = q.WriteString("SingerId, ")
	_, _ = q.WriteString("FirstName, ")
	_, _ = q.WriteString("LastName, ")
	_, _ = q.WriteString("SingerInfo, ")
	_, _ = q.WriteString("ARRAY( ")
	_, _ = q.WriteString("SELECT AS STRUCT ")
	_, _ = q.WriteString("SingerId, ")
	_, _ = q.WriteString("AlbumId, ")
	_, _ = q.WriteString("AlbumTitle, ")
	_, _ = q.WriteString("FROM Albums ")
	_, _ = q.WriteString("WHERE ")
	_, _ = q.WriteString("SingerId = Singers.SingerId ")
	_, _ = q.WriteString("ORDER BY ")
	_, _ = q.WriteString("SingerId")
	_, _ = q.WriteString(", ")
	_, _ = q.WriteString("AlbumId")
	_, _ = q.WriteString(" ")
	_, _ = q.WriteString(") AS Albums, ")
	_, _ = q.WriteString("FROM Singers ")
	if query.Where != nil {
		_, _ = q.WriteString("WHERE (")
		_, _ = q.WriteString(query.Where.SQL())
		_, _ = q.WriteString(") ")
	}
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
	_, _ = q.WriteString("LIMIT @limit ")
	_, _ = q.WriteString("OFFSET @offset ")
	stmt := spanner.Statement{
		SQL: q.String(),
		Params: map[string]interface{}{
			"limit":  query.Limit,
			"offset": query.Offset,
		},
	}
	return &SingersAndAlbumsRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

func (t SingersAndAlbumsReadTransaction) Get(
	ctx context.Context,
	key SingersKey,
) (*SingersAndAlbumsRow, error) {
	it := t.List(ctx, ListQuery{
		Where: key.BoolExpr(),
		Limit: 1,
	})
	defer it.Stop()
	row, err := it.Next()
	if err != nil {
		if err == iterator.Done {
			return nil, status.Errorf(codes.NotFound, "not found: %v", key)
		}
		return nil, err
	}
	return row, nil
}

func (t SingersAndAlbumsReadTransaction) BatchGet(
	ctx context.Context,
	keys []SingersKey,
) (map[SingersKey]*SingersAndAlbumsRow, error) {
	if len(keys) == 0 {
		return nil, nil
	}
	where := keys[0].BoolExpr()
	for _, key := range keys[1:] {
		where = spansql.LogicalOp{
			Op:  spansql.Or,
			LHS: where,
			RHS: key.BoolExpr(),
		}
	}
	foundRows := make(map[SingersKey]*SingersAndAlbumsRow, len(keys))
	if err := t.List(ctx, ListQuery{
		Where: spansql.Paren{Expr: where},
		Limit: int64(len(keys)),
	}).Do(func(row *SingersAndAlbumsRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

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

func (r *SingersAndAlbumsRow) Key() SingersKey {
	return SingersKey{
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

func (r SingersAndAlbumsRow) SingersRow() *SingersRow {
	return &SingersRow{
		SingerId:   r.SingerId,
		FirstName:  r.FirstName,
		LastName:   r.LastName,
		SingerInfo: r.SingerInfo,
	}
}

func (r SingersAndAlbumsRow) Insert() []*spanner.Mutation {
	n := 1
	n += len(r.Albums)
	mutations := make([]*spanner.Mutation, 0, n)
	mutations = append(mutations, r.SingersRow().Insert())
	for _, interleavedRow := range r.Albums {
		mutations = append(mutations, interleavedRow.Insert())
	}
	return mutations
}

func (r SingersAndAlbumsRow) Update() []*spanner.Mutation {
	n := 2 // one delete mutation per interleaved table
	n += len(r.Albums)
	mutations := make([]*spanner.Mutation, 0, n)
	mutations = append(mutations, r.SingersRow().Update())
	mutations = append(mutations, spanner.Delete("Albums", r.Key().SpannerKey().AsPrefix()))
	for _, interleavedRow := range r.Albums {
		mutations = append(mutations, interleavedRow.Insert())
	}
	return mutations
}

type AlbumsReadTransaction struct {
	Tx SpannerReadTransaction
}

func Albums(tx SpannerReadTransaction) AlbumsReadTransaction {
	return AlbumsReadTransaction{Tx: tx}
}

func (t AlbumsReadTransaction) Read(
	ctx context.Context,
	keySet spanner.KeySet,
) *AlbumsRowIterator {
	return &AlbumsRowIterator{
		RowIterator: t.Tx.Read(
			ctx,
			"Albums",
			keySet,
			((*AlbumsRow)(nil)).ColumnNames(),
		),
	}
}

func (t AlbumsReadTransaction) Get(
	ctx context.Context,
	key AlbumsKey,
) (*AlbumsRow, error) {
	spannerRow, err := t.Tx.ReadRow(
		ctx,
		"Albums",
		key.SpannerKey(),
		((*AlbumsRow)(nil)).ColumnNames(),
	)
	if err != nil {
		return nil, err
	}
	var row AlbumsRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

func (t AlbumsReadTransaction) BatchGet(
	ctx context.Context,
	keys []AlbumsKey,
) (map[AlbumsKey]*AlbumsRow, error) {
	spannerKeys := make([]spanner.KeySet, 0, len(keys))
	for _, key := range keys {
		spannerKeys = append(spannerKeys, key.SpannerKey())
	}
	foundRows := make(map[AlbumsKey]*AlbumsRow, len(keys))
	if err := t.Read(ctx, spanner.KeySets(spannerKeys...)).Do(func(row *AlbumsRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

func (t AlbumsReadTransaction) List(
	ctx context.Context,
	query ListQuery,
) *AlbumsRowIterator {
	if len(query.Order) == 0 {
		query.Order = AlbumsKey{}.Order()
	}
	stmt := spanner.Statement{
		SQL: spansql.Query{
			Select: spansql.Select{
				List: ((*AlbumsRow)(nil)).ColumnExprs(),
				From: []spansql.SelectFrom{
					spansql.SelectFromTable{Table: "Albums"},
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
	return &AlbumsRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

type AlbumsRowIterator struct {
	*spanner.RowIterator
}

func (i *AlbumsRowIterator) Next() (*AlbumsRow, error) {
	spannerRow, err := i.RowIterator.Next()
	if err != nil {
		return nil, err
	}
	var row AlbumsRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

func (i *AlbumsRowIterator) Do(f func(row *AlbumsRow) error) error {
	return i.RowIterator.Do(func(spannerRow *spanner.Row) error {
		var row AlbumsRow
		if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
			return err
		}
		return f(&row)
	})
}

type AlbumsKey struct {
	SingerId int64
	AlbumId  int64
}

func (k AlbumsKey) SpannerKey() spanner.Key {
	return spanner.Key{
		k.SingerId,
		k.AlbumId,
	}
}

func (k AlbumsKey) SpannerKeySet() spanner.KeySet {
	return k.SpannerKey()
}

func (k AlbumsKey) Delete() *spanner.Mutation {
	return spanner.Delete("Albums", k.SpannerKey())
}

func (AlbumsKey) Order() []spansql.Order {
	return []spansql.Order{
		{Expr: spansql.ID("SingerId"), Desc: false},
		{Expr: spansql.ID("AlbumId"), Desc: false},
	}
}

func (k AlbumsKey) BoolExpr() spansql.BoolExpr {
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

func (k AlbumsKey) QualifiedBoolExpr(prefix spansql.PathExp) spansql.BoolExpr {
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

func (r *AlbumsRow) Mutation() (string, []string, []interface{}) {
	return "Albums", r.ColumnNames(), []interface{}{
		r.SingerId,
		r.AlbumId,
		r.AlbumTitle,
	}
}

func (r *AlbumsRow) MutationForColumns(columns []string) (string, []string, []interface{}) {
	if len(columns) == 0 {
		columns = r.ColumnNames()
	}
	values := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		switch column {
		case "SingerId":
			values = append(values, r.SingerId)
		case "AlbumId":
			values = append(values, r.AlbumId)
		case "AlbumTitle":
			values = append(values, r.AlbumTitle)
		default:
			panic(fmt.Errorf("table Albums does not have column %s", column))
		}
	}
	return "Albums", columns, values
}

func (r *AlbumsRow) Key() AlbumsKey {
	return AlbumsKey{
		SingerId: r.SingerId,
		AlbumId:  r.AlbumId,
	}
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
