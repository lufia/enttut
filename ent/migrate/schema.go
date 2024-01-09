// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "paid_date", Type: field.TypeTime},
		{Name: "amount", Type: field.TypeInt},
		{Name: "memo", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "wallet_id", Type: field.TypeInt},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_wallets_transactions",
				Columns:    []*schema.Column{TransactionsColumns[4]},
				RefColumns: []*schema.Column{WalletsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "transaction_paid_date",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[1]},
			},
		},
	}
	// WalletsColumns holds the columns for the "wallets" table.
	WalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 50, SchemaType: map[string]string{"postgres": "varchar(50)"}},
		{Name: "method", Type: field.TypeEnum, Enums: []string{"cash", "credit-card", "e-money"}},
	}
	// WalletsTable holds the schema information for the "wallets" table.
	WalletsTable = &schema.Table{
		Name:       "wallets",
		Columns:    WalletsColumns,
		PrimaryKey: []*schema.Column{WalletsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TransactionsTable,
		WalletsTable,
	}
)

func init() {
	TransactionsTable.ForeignKeys[0].RefTable = WalletsTable
}
