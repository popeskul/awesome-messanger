package querybuilder

import (
	"fmt"
	"strings"
)

type Builder interface {
	Build() (string, []interface{}, error)
}

type SQLBuilder struct {
	parts []string
	args  []interface{}
}

func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{
		parts: make([]string, 0),
		args:  make([]interface{}, 0),
	}
}

func (b *SQLBuilder) AddPart(part string, args ...interface{}) *SQLBuilder {
	b.parts = append(b.parts, part)
	b.args = append(b.args, args...)
	return b
}

func (b *SQLBuilder) Build() (string, []interface{}, error) {
	if len(b.parts) == 0 {
		return "", nil, fmt.Errorf("empty query")
	}
	return strings.Join(b.parts, " "), b.args, nil
}

type SelectBuilder struct {
	SQLBuilder
	table     string
	columns   []string
	where     string
	whereArgs []interface{}
}

func NewSelectBuilder(table string, columns ...string) *SelectBuilder {
	return &SelectBuilder{
		table:   table,
		columns: columns,
	}
}

func (sb *SelectBuilder) Where(condition string, args ...interface{}) *SelectBuilder {
	sb.where = condition
	sb.whereArgs = args
	return sb
}

func (sb *SelectBuilder) Build() (string, []interface{}, error) {
	sb.AddPart(fmt.Sprintf("SELECT %s FROM %s", strings.Join(sb.columns, ", "), sb.table))
	if sb.where != "" {
		sb.AddPart(fmt.Sprintf("WHERE %s", sb.where))
		sb.args = append(sb.args, sb.whereArgs...)
	}
	return sb.SQLBuilder.Build()
}
