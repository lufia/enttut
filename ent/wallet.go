// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lufia/enttut/ent/wallet"
)

// Wallet is the model entity for the Wallet schema.
type Wallet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Method holds the value of the "method" field.
	Method string `json:"method,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WalletQuery when eager-loading is set.
	Edges        WalletEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WalletEdges holds the relations/edges for other nodes in the graph.
type WalletEdges struct {
	// Transactions holds the value of the transactions edge.
	Transactions []*Transaction `json:"transactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TransactionsOrErr returns the Transactions value or an error if the edge
// was not loaded in eager-loading.
func (e WalletEdges) TransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[0] {
		return e.Transactions, nil
	}
	return nil, &NotLoadedError{edge: "transactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Wallet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wallet.FieldID:
			values[i] = new(sql.NullInt64)
		case wallet.FieldName, wallet.FieldMethod:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Wallet fields.
func (w *Wallet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wallet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case wallet.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				w.Name = value.String
			}
		case wallet.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				w.Method = value.String
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Wallet.
// This includes values selected through modifiers, order, etc.
func (w *Wallet) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// QueryTransactions queries the "transactions" edge of the Wallet entity.
func (w *Wallet) QueryTransactions() *TransactionQuery {
	return NewWalletClient(w.config).QueryTransactions(w)
}

// Update returns a builder for updating this Wallet.
// Note that you need to call Wallet.Unwrap() before calling this method if this Wallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Wallet) Update() *WalletUpdateOne {
	return NewWalletClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Wallet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Wallet) Unwrap() *Wallet {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Wallet is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Wallet) String() string {
	var builder strings.Builder
	builder.WriteString("Wallet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("name=")
	builder.WriteString(w.Name)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(w.Method)
	builder.WriteByte(')')
	return builder.String()
}

// Wallets is a parsable slice of Wallet.
type Wallets []*Wallet
