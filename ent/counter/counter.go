// Code generated by ent, DO NOT EDIT.

package counter

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the counter type in the database.
	Label = "counter"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// Table holds the table name of the counter in the database.
	Table = "counters"
)

// Columns holds all SQL columns for counter fields.
var Columns = []string{
	FieldID,
	FieldCount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// CountValidator is a validator for the "count" field. It is called by the builders before save.
	CountValidator func(int) error
)

// OrderOption defines the ordering options for the Counter queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCount orders the results by the count field.
func ByCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCount, opts...).ToFunc()
}
