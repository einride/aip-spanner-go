// +build testdata.4.sql.database

package testdata

// Code generated by TestDatabaseCodeGenerator_GenerateCode/database/testdata/4.sql. DO NOT EDIT.

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

type SingersRow struct {
	SingerId   int64              `spanner:"SingerId"`
	FirstName  spanner.NullString `spanner:"FirstName"`
	LastName   spanner.NullString `spanner:"LastName"`
	SingerInfo []uint8            `spanner:"SingerInfo"`
	Albums     []*AlbumsRow       `spanner:"Albums"`
	Singles    []*SinglesRow      `spanner:"Singles"`
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
		case "Albums":
			if err := row.Column(i, &r.Albums); err != nil {
				return fmt.Errorf("unmarshal Singers interleaved row: Albums column: %w", err)
			}
		case "Singles":
			if err := row.Column(i, &r.Singles); err != nil {
				return fmt.Errorf("unmarshal Singers interleaved row: Singles column: %w", err)
			}
		default:
			return fmt.Errorf("unmarshal Singers row: unhandled column: %s", row.ColumnName(i))
		}
	}
	return nil
}

func (r *SingersRow) Mutate() (string, []string, []interface{}) {
	return "Singers", r.ColumnNames(), []interface{}{
		r.SingerId,
		r.FirstName,
		r.LastName,
		r.SingerInfo,
	}
}

func (r *SingersRow) MutateColumns(columns []string) (string, []string, []interface{}) {
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

type AlbumsRow struct {
	SingerId   int64              `spanner:"SingerId"`
	AlbumId    int64              `spanner:"AlbumId"`
	AlbumTitle spanner.NullString `spanner:"AlbumTitle"`
	Songs      []*SongsRow        `spanner:"Songs"`
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
		case "Songs":
			if err := row.Column(i, &r.Songs); err != nil {
				return fmt.Errorf("unmarshal Albums interleaved row: Songs column: %w", err)
			}
		default:
			return fmt.Errorf("unmarshal Albums row: unhandled column: %s", row.ColumnName(i))
		}
	}
	return nil
}

func (r *AlbumsRow) Mutate() (string, []string, []interface{}) {
	return "Albums", r.ColumnNames(), []interface{}{
		r.SingerId,
		r.AlbumId,
		r.AlbumTitle,
	}
}

func (r *AlbumsRow) MutateColumns(columns []string) (string, []string, []interface{}) {
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

type SongsRow struct {
	SingerId int64              `spanner:"SingerId"`
	AlbumId  int64              `spanner:"AlbumId"`
	TrackId  int64              `spanner:"TrackId"`
	SongName spanner.NullString `spanner:"SongName"`
}

func (*SongsRow) ColumnNames() []string {
	return []string{
		"SingerId",
		"AlbumId",
		"TrackId",
		"SongName",
	}
}

func (*SongsRow) ColumnIDs() []spansql.ID {
	return []spansql.ID{
		"SingerId",
		"AlbumId",
		"TrackId",
		"SongName",
	}
}

func (*SongsRow) ColumnExprs() []spansql.Expr {
	return []spansql.Expr{
		spansql.ID("SingerId"),
		spansql.ID("AlbumId"),
		spansql.ID("TrackId"),
		spansql.ID("SongName"),
	}
}

func (r *SongsRow) Validate() error {
	return nil
}

func (r *SongsRow) UnmarshalSpannerRow(row *spanner.Row) error {
	for i := 0; i < row.Size(); i++ {
		switch row.ColumnName(i) {
		case "SingerId":
			if err := row.Column(i, &r.SingerId); err != nil {
				return fmt.Errorf("unmarshal Songs row: SingerId column: %w", err)
			}
		case "AlbumId":
			if err := row.Column(i, &r.AlbumId); err != nil {
				return fmt.Errorf("unmarshal Songs row: AlbumId column: %w", err)
			}
		case "TrackId":
			if err := row.Column(i, &r.TrackId); err != nil {
				return fmt.Errorf("unmarshal Songs row: TrackId column: %w", err)
			}
		case "SongName":
			if err := row.Column(i, &r.SongName); err != nil {
				return fmt.Errorf("unmarshal Songs row: SongName column: %w", err)
			}
		default:
			return fmt.Errorf("unmarshal Songs row: unhandled column: %s", row.ColumnName(i))
		}
	}
	return nil
}

func (r *SongsRow) Mutate() (string, []string, []interface{}) {
	return "Songs", r.ColumnNames(), []interface{}{
		r.SingerId,
		r.AlbumId,
		r.TrackId,
		r.SongName,
	}
}

func (r *SongsRow) MutateColumns(columns []string) (string, []string, []interface{}) {
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
		case "TrackId":
			values = append(values, r.TrackId)
		case "SongName":
			values = append(values, r.SongName)
		default:
			panic(fmt.Errorf("table Songs does not have column %s", column))
		}
	}
	return "Songs", columns, values
}

func (r *SongsRow) Key() SongsKey {
	return SongsKey{
		SingerId: r.SingerId,
		AlbumId:  r.AlbumId,
		TrackId:  r.TrackId,
	}
}

type SinglesRow struct {
	SingerId int64              `spanner:"SingerId"`
	AlbumId  int64              `spanner:"AlbumId"`
	SingleId int64              `spanner:"SingleId"`
	SongName spanner.NullString `spanner:"SongName"`
}

func (*SinglesRow) ColumnNames() []string {
	return []string{
		"SingerId",
		"AlbumId",
		"SingleId",
		"SongName",
	}
}

func (*SinglesRow) ColumnIDs() []spansql.ID {
	return []spansql.ID{
		"SingerId",
		"AlbumId",
		"SingleId",
		"SongName",
	}
}

func (*SinglesRow) ColumnExprs() []spansql.Expr {
	return []spansql.Expr{
		spansql.ID("SingerId"),
		spansql.ID("AlbumId"),
		spansql.ID("SingleId"),
		spansql.ID("SongName"),
	}
}

func (r *SinglesRow) Validate() error {
	return nil
}

func (r *SinglesRow) UnmarshalSpannerRow(row *spanner.Row) error {
	for i := 0; i < row.Size(); i++ {
		switch row.ColumnName(i) {
		case "SingerId":
			if err := row.Column(i, &r.SingerId); err != nil {
				return fmt.Errorf("unmarshal Singles row: SingerId column: %w", err)
			}
		case "AlbumId":
			if err := row.Column(i, &r.AlbumId); err != nil {
				return fmt.Errorf("unmarshal Singles row: AlbumId column: %w", err)
			}
		case "SingleId":
			if err := row.Column(i, &r.SingleId); err != nil {
				return fmt.Errorf("unmarshal Singles row: SingleId column: %w", err)
			}
		case "SongName":
			if err := row.Column(i, &r.SongName); err != nil {
				return fmt.Errorf("unmarshal Singles row: SongName column: %w", err)
			}
		default:
			return fmt.Errorf("unmarshal Singles row: unhandled column: %s", row.ColumnName(i))
		}
	}
	return nil
}

func (r *SinglesRow) Mutate() (string, []string, []interface{}) {
	return "Singles", r.ColumnNames(), []interface{}{
		r.SingerId,
		r.AlbumId,
		r.SingleId,
		r.SongName,
	}
}

func (r *SinglesRow) MutateColumns(columns []string) (string, []string, []interface{}) {
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
		case "SingleId":
			values = append(values, r.SingleId)
		case "SongName":
			values = append(values, r.SongName)
		default:
			panic(fmt.Errorf("table Singles does not have column %s", column))
		}
	}
	return "Singles", columns, values
}

func (r *SinglesRow) Key() SinglesKey {
	return SinglesKey{
		SingerId: r.SingerId,
		AlbumId:  r.AlbumId,
		SingleId: r.SingleId,
	}
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
	cmp0 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("SingerId"),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	b := cmp0
	return spansql.Paren{Expr: b}
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
	cmp0 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("SingerId"),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	cmp1 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("AlbumId"),
		RHS: spansql.IntegerLiteral(k.AlbumId),
	})
	b := cmp0
	b = spansql.LogicalOp{
		Op:  spansql.And,
		LHS: b,
		RHS: cmp1,
	}
	return spansql.Paren{Expr: b}
}

type SongsKey struct {
	SingerId int64
	AlbumId  int64
	TrackId  int64
}

func (k SongsKey) SpannerKey() spanner.Key {
	return spanner.Key{
		k.SingerId,
		k.AlbumId,
		k.TrackId,
	}
}

func (k SongsKey) SpannerKeySet() spanner.KeySet {
	return k.SpannerKey()
}

func (k SongsKey) Delete() *spanner.Mutation {
	return spanner.Delete("Songs", k.SpannerKey())
}

func (SongsKey) Order() []spansql.Order {
	return []spansql.Order{
		{Expr: spansql.ID("SingerId"), Desc: false},
		{Expr: spansql.ID("AlbumId"), Desc: false},
		{Expr: spansql.ID("TrackId"), Desc: false},
	}
}

func (k SongsKey) BoolExpr() spansql.BoolExpr {
	cmp0 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("SingerId"),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	cmp1 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("AlbumId"),
		RHS: spansql.IntegerLiteral(k.AlbumId),
	})
	cmp2 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("TrackId"),
		RHS: spansql.IntegerLiteral(k.TrackId),
	})
	b := cmp0
	b = spansql.LogicalOp{
		Op:  spansql.And,
		LHS: b,
		RHS: cmp1,
	}
	b = spansql.LogicalOp{
		Op:  spansql.And,
		LHS: b,
		RHS: cmp2,
	}
	return spansql.Paren{Expr: b}
}

type SinglesKey struct {
	SingerId int64
	AlbumId  int64
	SingleId int64
}

func (k SinglesKey) SpannerKey() spanner.Key {
	return spanner.Key{
		k.SingerId,
		k.AlbumId,
		k.SingleId,
	}
}

func (k SinglesKey) SpannerKeySet() spanner.KeySet {
	return k.SpannerKey()
}

func (k SinglesKey) Delete() *spanner.Mutation {
	return spanner.Delete("Singles", k.SpannerKey())
}

func (SinglesKey) Order() []spansql.Order {
	return []spansql.Order{
		{Expr: spansql.ID("SingerId"), Desc: false},
		{Expr: spansql.ID("AlbumId"), Desc: false},
		{Expr: spansql.ID("SingleId"), Desc: false},
	}
}

func (k SinglesKey) BoolExpr() spansql.BoolExpr {
	cmp0 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("SingerId"),
		RHS: spansql.IntegerLiteral(k.SingerId),
	})
	cmp1 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("AlbumId"),
		RHS: spansql.IntegerLiteral(k.AlbumId),
	})
	cmp2 := spansql.BoolExpr(spansql.ComparisonOp{
		Op:  spansql.Eq,
		LHS: spansql.ID("SingleId"),
		RHS: spansql.IntegerLiteral(k.SingleId),
	})
	b := cmp0
	b = spansql.LogicalOp{
		Op:  spansql.And,
		LHS: b,
		RHS: cmp1,
	}
	b = spansql.LogicalOp{
		Op:  spansql.And,
		LHS: b,
		RHS: cmp2,
	}
	return spansql.Paren{Expr: b}
}

type SingersRowIterator interface {
	Next() (*SingersRow, error)
	Do(f func(row *SingersRow) error) error
	Stop()
}

type streamingSingersRowIterator struct {
	*spanner.RowIterator
}

func (i *streamingSingersRowIterator) Next() (*SingersRow, error) {
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

func (i *streamingSingersRowIterator) Do(f func(row *SingersRow) error) error {
	return i.RowIterator.Do(func(spannerRow *spanner.Row) error {
		var row SingersRow
		if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
			return err
		}
		return f(&row)
	})
}

type bufferedSingersRowIterator struct {
	rows []*SingersRow
	err  error
}

func (i *bufferedSingersRowIterator) Next() (*SingersRow, error) {
	if i.err != nil {
		return nil, i.err
	}
	if len(i.rows) == 0 {
		return nil, iterator.Done
	}
	next := i.rows[0]
	i.rows = i.rows[1:]
	return next, nil
}

func (i *bufferedSingersRowIterator) Do(f func(row *SingersRow) error) error {
	for {
		row, err := i.Next()
		switch err {
		case iterator.Done:
			return nil
		case nil:
			if err = f(row); err != nil {
				return err
			}
		default:
			return err
		}
	}
}

func (i *bufferedSingersRowIterator) Stop() {}

type AlbumsRowIterator interface {
	Next() (*AlbumsRow, error)
	Do(f func(row *AlbumsRow) error) error
	Stop()
}

type streamingAlbumsRowIterator struct {
	*spanner.RowIterator
}

func (i *streamingAlbumsRowIterator) Next() (*AlbumsRow, error) {
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

func (i *streamingAlbumsRowIterator) Do(f func(row *AlbumsRow) error) error {
	return i.RowIterator.Do(func(spannerRow *spanner.Row) error {
		var row AlbumsRow
		if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
			return err
		}
		return f(&row)
	})
}

type bufferedAlbumsRowIterator struct {
	rows []*AlbumsRow
	err  error
}

func (i *bufferedAlbumsRowIterator) Next() (*AlbumsRow, error) {
	if i.err != nil {
		return nil, i.err
	}
	if len(i.rows) == 0 {
		return nil, iterator.Done
	}
	next := i.rows[0]
	i.rows = i.rows[1:]
	return next, nil
}

func (i *bufferedAlbumsRowIterator) Do(f func(row *AlbumsRow) error) error {
	for {
		row, err := i.Next()
		switch err {
		case iterator.Done:
			return nil
		case nil:
			if err = f(row); err != nil {
				return err
			}
		default:
			return err
		}
	}
}

func (i *bufferedAlbumsRowIterator) Stop() {}

type SongsRowIterator interface {
	Next() (*SongsRow, error)
	Do(f func(row *SongsRow) error) error
	Stop()
}

type streamingSongsRowIterator struct {
	*spanner.RowIterator
}

func (i *streamingSongsRowIterator) Next() (*SongsRow, error) {
	spannerRow, err := i.RowIterator.Next()
	if err != nil {
		return nil, err
	}
	var row SongsRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

func (i *streamingSongsRowIterator) Do(f func(row *SongsRow) error) error {
	return i.RowIterator.Do(func(spannerRow *spanner.Row) error {
		var row SongsRow
		if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
			return err
		}
		return f(&row)
	})
}

type bufferedSongsRowIterator struct {
	rows []*SongsRow
	err  error
}

func (i *bufferedSongsRowIterator) Next() (*SongsRow, error) {
	if i.err != nil {
		return nil, i.err
	}
	if len(i.rows) == 0 {
		return nil, iterator.Done
	}
	next := i.rows[0]
	i.rows = i.rows[1:]
	return next, nil
}

func (i *bufferedSongsRowIterator) Do(f func(row *SongsRow) error) error {
	for {
		row, err := i.Next()
		switch err {
		case iterator.Done:
			return nil
		case nil:
			if err = f(row); err != nil {
				return err
			}
		default:
			return err
		}
	}
}

func (i *bufferedSongsRowIterator) Stop() {}

type SinglesRowIterator interface {
	Next() (*SinglesRow, error)
	Do(f func(row *SinglesRow) error) error
	Stop()
}

type streamingSinglesRowIterator struct {
	*spanner.RowIterator
}

func (i *streamingSinglesRowIterator) Next() (*SinglesRow, error) {
	spannerRow, err := i.RowIterator.Next()
	if err != nil {
		return nil, err
	}
	var row SinglesRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

func (i *streamingSinglesRowIterator) Do(f func(row *SinglesRow) error) error {
	return i.RowIterator.Do(func(spannerRow *spanner.Row) error {
		var row SinglesRow
		if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
			return err
		}
		return f(&row)
	})
}

type bufferedSinglesRowIterator struct {
	rows []*SinglesRow
	err  error
}

func (i *bufferedSinglesRowIterator) Next() (*SinglesRow, error) {
	if i.err != nil {
		return nil, i.err
	}
	if len(i.rows) == 0 {
		return nil, iterator.Done
	}
	next := i.rows[0]
	i.rows = i.rows[1:]
	return next, nil
}

func (i *bufferedSinglesRowIterator) Do(f func(row *SinglesRow) error) error {
	for {
		row, err := i.Next()
		switch err {
		case iterator.Done:
			return nil
		case nil:
			if err = f(row); err != nil {
				return err
			}
		default:
			return err
		}
	}
}

func (i *bufferedSinglesRowIterator) Stop() {}

type ReadTransaction struct {
	Tx SpannerReadTransaction
}

func Query(tx SpannerReadTransaction) ReadTransaction {
	return ReadTransaction{Tx: tx}
}

func (t ReadTransaction) ReadSingersRows(
	ctx context.Context,
	keySet spanner.KeySet,
) SingersRowIterator {
	return &streamingSingersRowIterator{
		RowIterator: t.Tx.Read(
			ctx,
			"Singers",
			keySet,
			((*SingersRow)(nil)).ColumnNames(),
		),
	}
}

type GetSingersRowQuery struct {
	Key     SingersKey
	Albums  bool
	Songs   bool
	Singles bool
}

func (q *GetSingersRowQuery) hasInterleavedTables() bool {
	return q.Albums || q.Songs || q.Singles
}

func (t ReadTransaction) GetSingersRow(
	ctx context.Context,
	query GetSingersRowQuery,
) (*SingersRow, error) {
	if query.hasInterleavedTables() {
		return t.getSingersRowInterleaved(ctx, query)
	}
	spannerRow, err := t.Tx.ReadRow(
		ctx,
		"Singers",
		query.Key.SpannerKey(),
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

type BatchGetSingersRowsQuery struct {
	Keys    []SingersKey
	Albums  bool
	Songs   bool
	Singles bool
}

func (q *BatchGetSingersRowsQuery) hasInterleavedTables() bool {
	return q.Albums || q.Songs || q.Singles
}

func (t ReadTransaction) BatchGetSingersRows(
	ctx context.Context,
	query BatchGetSingersRowsQuery,
) (map[SingersKey]*SingersRow, error) {
	if query.hasInterleavedTables() {
		return t.batchGetSingersRowsInterleaved(ctx, query)
	}
	spannerKeys := make([]spanner.KeySet, 0, len(query.Keys))
	for _, key := range query.Keys {
		spannerKeys = append(spannerKeys, key.SpannerKey())
	}
	foundRows := make(map[SingersKey]*SingersRow, len(query.Keys))
	if err := t.ReadSingersRows(ctx, spanner.KeySets(spannerKeys...)).Do(func(row *SingersRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

type ListSingersRowsQuery struct {
	Where   spansql.BoolExpr
	Order   []spansql.Order
	Limit   int32
	Offset  int64
	Params  map[string]interface{}
	Albums  bool
	Songs   bool
	Singles bool
}

func (q *ListSingersRowsQuery) hasInterleavedTables() bool {
	return q.Albums || q.Songs || q.Singles
}

func (t ReadTransaction) ListSingersRows(
	ctx context.Context,
	query ListSingersRowsQuery,
) SingersRowIterator {
	if query.hasInterleavedTables() {
		return t.listSingersRowsInterleaved(ctx, query)
	}
	if len(query.Order) == 0 {
		query.Order = SingersKey{}.Order()
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
			Limit:  spansql.Param("__limit"),
			Offset: spansql.Param("__offset"),
		}.SQL(),
		Params: params,
	}
	return &streamingSingersRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

func (t ReadTransaction) listSingersRowsInterleaved(
	ctx context.Context,
	query ListSingersRowsQuery,
) SingersRowIterator {
	if len(query.Order) == 0 {
		query.Order = SingersKey{}.Order()
	}
	var q strings.Builder
	_, _ = q.WriteString(`
SELECT
    SingerId,
    FirstName,
    LastName,
    SingerInfo,
`)
	if query.Albums {
		_, _ = q.WriteString(`
    ARRAY(
        SELECT AS STRUCT
            SingerId,
            AlbumId,
            AlbumTitle,
`)
		if query.Songs {
			_, _ = q.WriteString(`
            ARRAY(
                SELECT AS STRUCT
                    SingerId,
                    AlbumId,
                    TrackId,
                    SongName,
`)
			_, _ = q.WriteString(`
                FROM 
                    Songs
                WHERE 
                    Songs.SingerId = Albums.SingerId AND
                    Songs.AlbumId = Albums.AlbumId
                ORDER BY 
                    SingerId,
                    AlbumId,
                    TrackId
            ) AS Songs,
`)
		}
		_, _ = q.WriteString(`
        FROM 
            Albums
        WHERE 
            Albums.SingerId = Singers.SingerId
        ORDER BY 
            SingerId,
            AlbumId
    ) AS Albums,
`)
	}
	if query.Singles {
		_, _ = q.WriteString(`
    ARRAY(
        SELECT AS STRUCT
            SingerId,
            AlbumId,
            SingleId,
            SongName,
`)
		_, _ = q.WriteString(`
        FROM 
            Singles
        WHERE 
            Singles.SingerId = Singers.SingerId
        ORDER BY 
            SingerId,
            AlbumId,
            SingleId
    ) AS Singles,
`)
	}
	_, _ = q.WriteString(`
FROM
    Singers
`)
	if query.Where == nil {
		query.Where = spansql.True
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
	return &streamingSingersRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

func (t ReadTransaction) getSingersRowInterleaved(
	ctx context.Context,
	query GetSingersRowQuery,
) (*SingersRow, error) {
	it := t.listSingersRowsInterleaved(ctx, ListSingersRowsQuery{
		Limit:   1,
		Where:   query.Key.BoolExpr(),
		Albums:  query.Albums,
		Songs:   query.Songs,
		Singles: query.Singles,
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

func (t ReadTransaction) batchGetSingersRowsInterleaved(
	ctx context.Context,
	query BatchGetSingersRowsQuery,
) (map[SingersKey]*SingersRow, error) {
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
	foundRows := make(map[SingersKey]*SingersRow, len(query.Keys))
	if err := t.ListSingersRows(ctx, ListSingersRowsQuery{
		Where:   spansql.Paren{Expr: where},
		Limit:   int32(len(query.Keys)),
		Albums:  query.Albums,
		Songs:   query.Songs,
		Singles: query.Singles,
	}).Do(func(row *SingersRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

func (t ReadTransaction) ReadAlbumsRows(
	ctx context.Context,
	keySet spanner.KeySet,
) AlbumsRowIterator {
	return &streamingAlbumsRowIterator{
		RowIterator: t.Tx.Read(
			ctx,
			"Albums",
			keySet,
			((*AlbumsRow)(nil)).ColumnNames(),
		),
	}
}

type GetAlbumsRowQuery struct {
	Key   AlbumsKey
	Songs bool
}

func (q *GetAlbumsRowQuery) hasInterleavedTables() bool {
	return q.Songs
}

func (t ReadTransaction) GetAlbumsRow(
	ctx context.Context,
	query GetAlbumsRowQuery,
) (*AlbumsRow, error) {
	if query.hasInterleavedTables() {
		return t.getAlbumsRowInterleaved(ctx, query)
	}
	spannerRow, err := t.Tx.ReadRow(
		ctx,
		"Albums",
		query.Key.SpannerKey(),
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

type BatchGetAlbumsRowsQuery struct {
	Keys  []AlbumsKey
	Songs bool
}

func (q *BatchGetAlbumsRowsQuery) hasInterleavedTables() bool {
	return q.Songs
}

func (t ReadTransaction) BatchGetAlbumsRows(
	ctx context.Context,
	query BatchGetAlbumsRowsQuery,
) (map[AlbumsKey]*AlbumsRow, error) {
	if query.hasInterleavedTables() {
		return t.batchGetAlbumsRowsInterleaved(ctx, query)
	}
	spannerKeys := make([]spanner.KeySet, 0, len(query.Keys))
	for _, key := range query.Keys {
		spannerKeys = append(spannerKeys, key.SpannerKey())
	}
	foundRows := make(map[AlbumsKey]*AlbumsRow, len(query.Keys))
	if err := t.ReadAlbumsRows(ctx, spanner.KeySets(spannerKeys...)).Do(func(row *AlbumsRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

type ListAlbumsRowsQuery struct {
	Where  spansql.BoolExpr
	Order  []spansql.Order
	Limit  int32
	Offset int64
	Params map[string]interface{}
	Songs  bool
}

func (q *ListAlbumsRowsQuery) hasInterleavedTables() bool {
	return q.Songs
}

func (t ReadTransaction) ListAlbumsRows(
	ctx context.Context,
	query ListAlbumsRowsQuery,
) AlbumsRowIterator {
	if query.hasInterleavedTables() {
		return t.listAlbumsRowsInterleaved(ctx, query)
	}
	if len(query.Order) == 0 {
		query.Order = AlbumsKey{}.Order()
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
			Limit:  spansql.Param("__limit"),
			Offset: spansql.Param("__offset"),
		}.SQL(),
		Params: params,
	}
	return &streamingAlbumsRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

func (t ReadTransaction) listAlbumsRowsInterleaved(
	ctx context.Context,
	query ListAlbumsRowsQuery,
) AlbumsRowIterator {
	if len(query.Order) == 0 {
		query.Order = AlbumsKey{}.Order()
	}
	var q strings.Builder
	_, _ = q.WriteString(`
SELECT
    SingerId,
    AlbumId,
    AlbumTitle,
`)
	if query.Songs {
		_, _ = q.WriteString(`
    ARRAY(
        SELECT AS STRUCT
            SingerId,
            AlbumId,
            TrackId,
            SongName,
`)
		_, _ = q.WriteString(`
        FROM 
            Songs
        WHERE 
            Songs.SingerId = Albums.SingerId AND
            Songs.AlbumId = Albums.AlbumId
        ORDER BY 
            SingerId,
            AlbumId,
            TrackId
    ) AS Songs,
`)
	}
	_, _ = q.WriteString(`
FROM
    Albums
`)
	if query.Where == nil {
		query.Where = spansql.True
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
	return &streamingAlbumsRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

func (t ReadTransaction) getAlbumsRowInterleaved(
	ctx context.Context,
	query GetAlbumsRowQuery,
) (*AlbumsRow, error) {
	it := t.listAlbumsRowsInterleaved(ctx, ListAlbumsRowsQuery{
		Limit: 1,
		Where: query.Key.BoolExpr(),
		Songs: query.Songs,
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

func (t ReadTransaction) batchGetAlbumsRowsInterleaved(
	ctx context.Context,
	query BatchGetAlbumsRowsQuery,
) (map[AlbumsKey]*AlbumsRow, error) {
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
	foundRows := make(map[AlbumsKey]*AlbumsRow, len(query.Keys))
	if err := t.ListAlbumsRows(ctx, ListAlbumsRowsQuery{
		Where: spansql.Paren{Expr: where},
		Limit: int32(len(query.Keys)),
		Songs: query.Songs,
	}).Do(func(row *AlbumsRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

func (t ReadTransaction) ReadSongsRows(
	ctx context.Context,
	keySet spanner.KeySet,
) SongsRowIterator {
	return &streamingSongsRowIterator{
		RowIterator: t.Tx.Read(
			ctx,
			"Songs",
			keySet,
			((*SongsRow)(nil)).ColumnNames(),
		),
	}
}

type GetSongsRowQuery struct {
	Key SongsKey
}

func (t ReadTransaction) GetSongsRow(
	ctx context.Context,
	query GetSongsRowQuery,
) (*SongsRow, error) {
	spannerRow, err := t.Tx.ReadRow(
		ctx,
		"Songs",
		query.Key.SpannerKey(),
		((*SongsRow)(nil)).ColumnNames(),
	)
	if err != nil {
		return nil, err
	}
	var row SongsRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

type BatchGetSongsRowsQuery struct {
	Keys []SongsKey
}

func (t ReadTransaction) BatchGetSongsRows(
	ctx context.Context,
	query BatchGetSongsRowsQuery,
) (map[SongsKey]*SongsRow, error) {
	spannerKeys := make([]spanner.KeySet, 0, len(query.Keys))
	for _, key := range query.Keys {
		spannerKeys = append(spannerKeys, key.SpannerKey())
	}
	foundRows := make(map[SongsKey]*SongsRow, len(query.Keys))
	if err := t.ReadSongsRows(ctx, spanner.KeySets(spannerKeys...)).Do(func(row *SongsRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

type ListSongsRowsQuery struct {
	Where  spansql.BoolExpr
	Order  []spansql.Order
	Limit  int32
	Offset int64
	Params map[string]interface{}
}

func (t ReadTransaction) ListSongsRows(
	ctx context.Context,
	query ListSongsRowsQuery,
) SongsRowIterator {
	if len(query.Order) == 0 {
		query.Order = SongsKey{}.Order()
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
	stmt := spanner.Statement{
		SQL: spansql.Query{
			Select: spansql.Select{
				List: ((*SongsRow)(nil)).ColumnExprs(),
				From: []spansql.SelectFrom{
					spansql.SelectFromTable{Table: "Songs"},
				},
				Where: query.Where,
			},
			Order:  query.Order,
			Limit:  spansql.Param("__limit"),
			Offset: spansql.Param("__offset"),
		}.SQL(),
		Params: params,
	}
	return &streamingSongsRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

func (t ReadTransaction) ReadSinglesRows(
	ctx context.Context,
	keySet spanner.KeySet,
) SinglesRowIterator {
	return &streamingSinglesRowIterator{
		RowIterator: t.Tx.Read(
			ctx,
			"Singles",
			keySet,
			((*SinglesRow)(nil)).ColumnNames(),
		),
	}
}

type GetSinglesRowQuery struct {
	Key SinglesKey
}

func (t ReadTransaction) GetSinglesRow(
	ctx context.Context,
	query GetSinglesRowQuery,
) (*SinglesRow, error) {
	spannerRow, err := t.Tx.ReadRow(
		ctx,
		"Singles",
		query.Key.SpannerKey(),
		((*SinglesRow)(nil)).ColumnNames(),
	)
	if err != nil {
		return nil, err
	}
	var row SinglesRow
	if err := row.UnmarshalSpannerRow(spannerRow); err != nil {
		return nil, err
	}
	return &row, nil
}

type BatchGetSinglesRowsQuery struct {
	Keys []SinglesKey
}

func (t ReadTransaction) BatchGetSinglesRows(
	ctx context.Context,
	query BatchGetSinglesRowsQuery,
) (map[SinglesKey]*SinglesRow, error) {
	spannerKeys := make([]spanner.KeySet, 0, len(query.Keys))
	for _, key := range query.Keys {
		spannerKeys = append(spannerKeys, key.SpannerKey())
	}
	foundRows := make(map[SinglesKey]*SinglesRow, len(query.Keys))
	if err := t.ReadSinglesRows(ctx, spanner.KeySets(spannerKeys...)).Do(func(row *SinglesRow) error {
		foundRows[row.Key()] = row
		return nil
	}); err != nil {
		return nil, err
	}
	return foundRows, nil
}

type ListSinglesRowsQuery struct {
	Where  spansql.BoolExpr
	Order  []spansql.Order
	Limit  int32
	Offset int64
	Params map[string]interface{}
}

func (t ReadTransaction) ListSinglesRows(
	ctx context.Context,
	query ListSinglesRowsQuery,
) SinglesRowIterator {
	if len(query.Order) == 0 {
		query.Order = SinglesKey{}.Order()
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
	stmt := spanner.Statement{
		SQL: spansql.Query{
			Select: spansql.Select{
				List: ((*SinglesRow)(nil)).ColumnExprs(),
				From: []spansql.SelectFrom{
					spansql.SelectFromTable{Table: "Singles"},
				},
				Where: query.Where,
			},
			Order:  query.Order,
			Limit:  spansql.Param("__limit"),
			Offset: spansql.Param("__offset"),
		}.SQL(),
		Params: params,
	}
	return &streamingSinglesRowIterator{
		RowIterator: t.Tx.Query(ctx, stmt),
	}
}

type SpannerReadTransaction interface {
	Read(ctx context.Context, table string, keys spanner.KeySet, columns []string) *spanner.RowIterator
	ReadRow(ctx context.Context, table string, key spanner.Key, columns []string) (*spanner.Row, error)
	Query(ctx context.Context, statement spanner.Statement) *spanner.RowIterator
}
