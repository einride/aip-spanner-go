package tablecodegen

import (
	"fmt"
	"reflect"
	"strconv"

	"cloud.google.com/go/spanner/spansql"
	"github.com/stoewer/go-strcase"
	"go.einride.tech/aip-spanner/internal/codegen"
	"go.einride.tech/aip-spanner/internal/codegen/typescodegen"
	"go.einride.tech/aip-spanner/spanddl"
)

type PrimaryKeyCodeGenerator struct {
	Table *spanddl.Table
}

func (g PrimaryKeyCodeGenerator) Type() string {
	return strcase.UpperCamelCase(string(g.Table.Name)) + "PrimaryKey"
}

func (g PrimaryKeyCodeGenerator) FieldName(keyPart spansql.KeyPart) string {
	return strcase.UpperCamelCase(string(keyPart.Column))
}

func (g PrimaryKeyCodeGenerator) GenerateCode(f *codegen.File) {
	spannerPkg := f.Import("cloud.google.com/go/spanner")
	f.P()
	f.P("type ", g.Type(), " struct {")
	for _, keyPart := range g.Table.PrimaryKey {
		f.P(g.FieldName(keyPart), " ", g.columnType(f, keyPart))
	}
	f.P("}")
	f.P()
	f.P("func (k ", g.Type(), ") SpannerKey() ", spannerPkg, ".Key {")
	f.P("return ", spannerPkg, ".Key{")
	for _, keyPart := range g.Table.PrimaryKey {
		f.P("k.", g.FieldName(keyPart), ",")
	}
	f.P("}")
	f.P("}")
	f.P()
	f.P("func (k ", g.Type(), ") SpannerKeySet() ", spannerPkg, ".KeySet {")
	f.P("return k.SpannerKey()")
	f.P("}")
	g.generateDeleteMethod(f)
	g.generateBoolExprMethod(f)
	g.generateQualifiedBoolExprMethod(f)
}

func (g PrimaryKeyCodeGenerator) generateDeleteMethod(f *codegen.File) {
	spannerPkg := f.Import("cloud.google.com/go/spanner")
	f.P()
	f.P("func (k ", g.Type(), ") Delete() *", spannerPkg, ".Mutation {")
	f.P("return ", spannerPkg, ".Delete(", strconv.Quote(string(g.Table.Name)), ", k.SpannerKey())")
	f.P("}")
}

func (g PrimaryKeyCodeGenerator) generateBoolExprMethod(f *codegen.File) {
	spansqlPkg := f.Import("cloud.google.com/go/spanner/spansql")
	f.P()
	k0 := g.Table.PrimaryKey[0]
	f.P("func (k ", g.Type(), ") BoolExpr() ", spansqlPkg, ".BoolExpr {")
	f.P("b := ", spansqlPkg, ".BoolExpr(", spansqlPkg, ".ComparisonOp{")
	f.P("Op: ", spansqlPkg, ".Eq,")
	f.P("LHS: ", spansqlPkg, ".ID(", strconv.Quote(string(k0.Column)), "),")
	f.P(
		"RHS: ", g.columnSpanSQLType(f, k0),
		"(k.", g.FieldName(k0), typescodegen.ValueAccessor(g.keyColumn(k0)), "),",
	)
	f.P("})")
	for _, keyPart := range g.Table.PrimaryKey[1:] {
		f.P("b = ", spansqlPkg, ".LogicalOp{")
		f.P("Op: ", spansqlPkg, ".And,")
		f.P("LHS: b,")
		f.P("RHS: ", spansqlPkg, ".ComparisonOp{")
		f.P("Op: ", spansqlPkg, ".Eq,")
		f.P("LHS: ", spansqlPkg, ".ID(", strconv.Quote(string(keyPart.Column)), "),")
		f.P(
			"RHS: ", g.columnSpanSQLType(f, keyPart),
			"(k.", g.FieldName(keyPart), typescodegen.ValueAccessor(g.keyColumn(keyPart)), "),",
		)
		f.P("},")
		f.P("}")
	}
	f.P("return ", spansqlPkg, ".Paren{Expr: b}")
	f.P("}")
}

func (g PrimaryKeyCodeGenerator) generateQualifiedBoolExprMethod(f *codegen.File) {
	spansqlPkg := f.Import("cloud.google.com/go/spanner/spansql")
	f.P()
	k0 := g.Table.PrimaryKey[0]
	f.P("func (k ", g.Type(), ") QualifiedBoolExpr(prefix ", spansqlPkg, ".PathExp) ", spansqlPkg, ".BoolExpr {")
	f.P("b := ", spansqlPkg, ".BoolExpr(", spansqlPkg, ".ComparisonOp{")
	f.P("Op: ", spansqlPkg, ".Eq,")
	f.P("LHS: append(prefix, ", spansqlPkg, ".ID(", strconv.Quote(string(k0.Column)), ")),")
	f.P(
		"RHS: ", g.columnSpanSQLType(f, k0),
		"(k.", g.FieldName(k0), typescodegen.ValueAccessor(g.keyColumn(k0)), "),",
	)
	f.P("})")
	for _, keyPart := range g.Table.PrimaryKey[1:] {
		f.P("b = ", spansqlPkg, ".LogicalOp{")
		f.P("Op: ", spansqlPkg, ".And,")
		f.P("LHS: b,")
		f.P("RHS: ", spansqlPkg, ".ComparisonOp{")
		f.P("Op: ", spansqlPkg, ".Eq,")
		f.P("LHS: append(prefix, ", spansqlPkg, ".ID(", strconv.Quote(string(keyPart.Column)), ")),")
		f.P(
			"RHS: ", g.columnSpanSQLType(f, keyPart),
			"(k.", g.FieldName(keyPart), typescodegen.ValueAccessor(g.keyColumn(keyPart)), "),",
		)
		f.P("},")
		f.P("}")
	}
	f.P("return ", spansqlPkg, ".Paren{Expr: b}")
	f.P("}")
}

func (g PrimaryKeyCodeGenerator) keyColumn(keyPart spansql.KeyPart) *spanddl.Column {
	column, ok := g.Table.Column(keyPart.Column)
	if !ok {
		panic(fmt.Errorf("table %s has no column %s", g.Table.Name, keyPart.Column))
	}
	return column
}

func (g PrimaryKeyCodeGenerator) columnType(f *codegen.File, keyPart spansql.KeyPart) reflect.Type {
	t := typescodegen.GoType(g.keyColumn(keyPart))
	if t.PkgPath() != "" {
		_ = f.Import(t.PkgPath())
	}
	return t
}

func (g PrimaryKeyCodeGenerator) columnSpanSQLType(f *codegen.File, keyPart spansql.KeyPart) reflect.Type {
	t := typescodegen.SpanSQLType(g.keyColumn(keyPart))
	if t.PkgPath() != "" {
		_ = f.Import(t.PkgPath())
	}
	return t
}
