package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Annotations of the Transaction.
func (Transaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("transactionを扱うテーブル"),
	}
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New),
		field.UUID("wallet_id", uuid.UUID{}),
		field.Time("paid_date").Comment("お金を払った日付"),
		field.Int("amount").Comment("数量(金額)"),
		field.Text("memo").Optional().Comment("メモ"),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("wallet", Wallet.Type).Ref("transactions").Field("wallet_id").Unique().Required(),
	}
}

// Indexes of the Transaction.
func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("paid_date"),
	}
}
