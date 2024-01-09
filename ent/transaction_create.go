// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lufia/enttut/ent/transaction"
	"github.com/lufia/enttut/ent/wallet"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
}

// SetWalletID sets the "wallet_id" field.
func (tc *TransactionCreate) SetWalletID(i int) *TransactionCreate {
	tc.mutation.SetWalletID(i)
	return tc
}

// SetPaidDate sets the "paid_date" field.
func (tc *TransactionCreate) SetPaidDate(t time.Time) *TransactionCreate {
	tc.mutation.SetPaidDate(t)
	return tc
}

// SetAmount sets the "amount" field.
func (tc *TransactionCreate) SetAmount(i int) *TransactionCreate {
	tc.mutation.SetAmount(i)
	return tc
}

// SetMemo sets the "memo" field.
func (tc *TransactionCreate) SetMemo(s string) *TransactionCreate {
	tc.mutation.SetMemo(s)
	return tc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableMemo(s *string) *TransactionCreate {
	if s != nil {
		tc.SetMemo(*s)
	}
	return tc
}

// SetWallet sets the "wallet" edge to the Wallet entity.
func (tc *TransactionCreate) SetWallet(w *Wallet) *TransactionCreate {
	return tc.SetWalletID(w.ID)
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.WalletID(); !ok {
		return &ValidationError{Name: "wallet_id", err: errors.New(`ent: missing required field "Transaction.wallet_id"`)}
	}
	if _, ok := tc.mutation.PaidDate(); !ok {
		return &ValidationError{Name: "paid_date", err: errors.New(`ent: missing required field "Transaction.paid_date"`)}
	}
	if _, ok := tc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Transaction.amount"`)}
	}
	if _, ok := tc.mutation.WalletID(); !ok {
		return &ValidationError{Name: "wallet", err: errors.New(`ent: missing required edge "Transaction.wallet"`)}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(transaction.Table, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.PaidDate(); ok {
		_spec.SetField(transaction.FieldPaidDate, field.TypeTime, value)
		_node.PaidDate = value
	}
	if value, ok := tc.mutation.Amount(); ok {
		_spec.SetField(transaction.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if value, ok := tc.mutation.Memo(); ok {
		_spec.SetField(transaction.FieldMemo, field.TypeString, value)
		_node.Memo = value
	}
	if nodes := tc.mutation.WalletIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.WalletTable,
			Columns: []string{transaction.WalletColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.WalletID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	err      error
	builders []*TransactionCreate
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
