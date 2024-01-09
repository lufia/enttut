package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Time("paid_date"),
		field.Int("amount"),
		field.Text("memo"),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("wallet", Wallet.Type).Ref("transactions").Unique(),
	}
}

// Indexes of the Transaction.
func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("paid_date"),
	}
}
