// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lufia/enttut/ent/transaction"
	"github.com/lufia/enttut/ent/wallet"
)

// WalletCreate is the builder for creating a Wallet entity.
type WalletCreate struct {
	config
	mutation *WalletMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wc *WalletCreate) SetName(s string) *WalletCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetPaymentMethod sets the "payment_method" field.
func (wc *WalletCreate) SetPaymentMethod(wm wallet.PaymentMethod) *WalletCreate {
	wc.mutation.SetPaymentMethod(wm)
	return wc
}

// SetID sets the "id" field.
func (wc *WalletCreate) SetID(u uuid.UUID) *WalletCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WalletCreate) SetNillableID(u *uuid.UUID) *WalletCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (wc *WalletCreate) AddTransactionIDs(ids ...uuid.UUID) *WalletCreate {
	wc.mutation.AddTransactionIDs(ids...)
	return wc
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (wc *WalletCreate) AddTransactions(t ...*Transaction) *WalletCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddTransactionIDs(ids...)
}

// Mutation returns the WalletMutation object of the builder.
func (wc *WalletCreate) Mutation() *WalletMutation {
	return wc.mutation
}

// Save creates the Wallet in the database.
func (wc *WalletCreate) Save(ctx context.Context) (*Wallet, error) {
	wc.defaults()
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WalletCreate) SaveX(ctx context.Context) *Wallet {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WalletCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WalletCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WalletCreate) defaults() {
	if _, ok := wc.mutation.ID(); !ok {
		v := wallet.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WalletCreate) check() error {
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Wallet.name"`)}
	}
	if v, ok := wc.mutation.Name(); ok {
		if err := wallet.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Wallet.name": %w`, err)}
		}
	}
	if _, ok := wc.mutation.PaymentMethod(); !ok {
		return &ValidationError{Name: "payment_method", err: errors.New(`ent: missing required field "Wallet.payment_method"`)}
	}
	if v, ok := wc.mutation.PaymentMethod(); ok {
		if err := wallet.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "Wallet.payment_method": %w`, err)}
		}
	}
	return nil
}

func (wc *WalletCreate) sqlSave(ctx context.Context) (*Wallet, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WalletCreate) createSpec() (*Wallet, *sqlgraph.CreateSpec) {
	var (
		_node = &Wallet{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(wallet.Table, sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeUUID))
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wc.mutation.Name(); ok {
		_spec.SetField(wallet.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wc.mutation.PaymentMethod(); ok {
		_spec.SetField(wallet.FieldPaymentMethod, field.TypeEnum, value)
		_node.PaymentMethod = value
	}
	if nodes := wc.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.TransactionsTable,
			Columns: []string{wallet.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WalletCreateBulk is the builder for creating many Wallet entities in bulk.
type WalletCreateBulk struct {
	config
	err      error
	builders []*WalletCreate
}

// Save creates the Wallet entities in the database.
func (wcb *WalletCreateBulk) Save(ctx context.Context) ([]*Wallet, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Wallet, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WalletMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WalletCreateBulk) SaveX(ctx context.Context) []*Wallet {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WalletCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WalletCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
