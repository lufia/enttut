package schema

import "entgo.io/ent"

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return nil
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return nil
}
