// +build testdata.2.sql.genericcolumn

package testdata

// Code generated by TestGenericColumnDescriptorCodeGenerator_GenerateCode/genericcolumn/testdata/2.sql. DO NOT EDIT.

import (
	"cloud.google.com/go/spanner/spansql"
)

type ColumnDescriptor interface {
	ColumnID() spansql.ID
	ColumnName() string
	ColumnType() spansql.Type
	NotNull() bool
	Options() spansql.ColumnOptions
}

type columnDescriptor struct {
	columnID   spansql.ID
	columnType spansql.Type
	notNull    bool
	options    spansql.ColumnOptions
}

func (d *columnDescriptor) ColumnName() string {
	return string(d.columnID)
}

func (d *columnDescriptor) ColumnID() spansql.ID {
	return d.columnID
}

func (d *columnDescriptor) ColumnType() spansql.Type {
	return d.columnType
}

func (d *columnDescriptor) ColumnExpr() spansql.Expr {
	return d.columnID
}

func (d *columnDescriptor) NotNull() bool {
	return d.notNull
}

func (d *columnDescriptor) Options() spansql.ColumnOptions {
	return d.options
}
