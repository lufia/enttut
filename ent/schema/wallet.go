package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Annotations of the Wallet.
func (Wallet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("walletを扱うテーブル"),
	}
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New),
		field.String("name").
			NotEmpty().
			MaxLen(50).
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(50)",
			}).
			Unique().
			Comment("名前"),
		field.Enum("payment_method").
			StorageKey("method").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(20)",
			}).
			Values("cash", "credit-card", "e-money").
			Comment("支払い方法"),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transactions", Transaction.Type),
	}
}
