// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lufia/enttut/ent/predicate"
	"github.com/lufia/enttut/ent/transaction"
	"github.com/lufia/enttut/ent/wallet"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTransaction = "Transaction"
	TypeWallet      = "Wallet"
)

// TransactionMutation represents an operation that mutates the Transaction nodes in the graph.
type TransactionMutation struct {
	config
	op            Op
	typ           string
	id            *int
	paid_date     *time.Time
	amount        *int
	addamount     *int
	memo          *string
	clearedFields map[string]struct{}
	wallet        *int
	clearedwallet bool
	done          bool
	oldValue      func(context.Context) (*Transaction, error)
	predicates    []predicate.Transaction
}

var _ ent.Mutation = (*TransactionMutation)(nil)

// transactionOption allows management of the mutation configuration using functional options.
type transactionOption func(*TransactionMutation)

// newTransactionMutation creates new mutation for the Transaction entity.
func newTransactionMutation(c config, op Op, opts ...transactionOption) *TransactionMutation {
	m := &TransactionMutation{
		config:        c,
		op:            op,
		typ:           TypeTransaction,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTransactionID sets the ID field of the mutation.
func withTransactionID(id int) transactionOption {
	return func(m *TransactionMutation) {
		var (
			err   error
			once  sync.Once
			value *Transaction
		)
		m.oldValue = func(ctx context.Context) (*Transaction, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Transaction.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTransaction sets the old Transaction of the mutation.
func withTransaction(node *Transaction) transactionOption {
	return func(m *TransactionMutation) {
		m.oldValue = func(context.Context) (*Transaction, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TransactionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TransactionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TransactionMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TransactionMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Transaction.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetWalletID sets the "wallet_id" field.
func (m *TransactionMutation) SetWalletID(i int) {
	m.wallet = &i
}

// WalletID returns the value of the "wallet_id" field in the mutation.
func (m *TransactionMutation) WalletID() (r int, exists bool) {
	v := m.wallet
	if v == nil {
		return
	}
	return *v, true
}

// OldWalletID returns the old "wallet_id" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldWalletID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldWalletID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldWalletID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWalletID: %w", err)
	}
	return oldValue.WalletID, nil
}

// ResetWalletID resets all changes to the "wallet_id" field.
func (m *TransactionMutation) ResetWalletID() {
	m.wallet = nil
}

// SetPaidDate sets the "paid_date" field.
func (m *TransactionMutation) SetPaidDate(t time.Time) {
	m.paid_date = &t
}

// PaidDate returns the value of the "paid_date" field in the mutation.
func (m *TransactionMutation) PaidDate() (r time.Time, exists bool) {
	v := m.paid_date
	if v == nil {
		return
	}
	return *v, true
}

// OldPaidDate returns the old "paid_date" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldPaidDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPaidDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPaidDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPaidDate: %w", err)
	}
	return oldValue.PaidDate, nil
}

// ResetPaidDate resets all changes to the "paid_date" field.
func (m *TransactionMutation) ResetPaidDate() {
	m.paid_date = nil
}

// SetAmount sets the "amount" field.
func (m *TransactionMutation) SetAmount(i int) {
	m.amount = &i
	m.addamount = nil
}

// Amount returns the value of the "amount" field in the mutation.
func (m *TransactionMutation) Amount() (r int, exists bool) {
	v := m.amount
	if v == nil {
		return
	}
	return *v, true
}

// OldAmount returns the old "amount" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldAmount(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAmount: %w", err)
	}
	return oldValue.Amount, nil
}

// AddAmount adds i to the "amount" field.
func (m *TransactionMutation) AddAmount(i int) {
	if m.addamount != nil {
		*m.addamount += i
	} else {
		m.addamount = &i
	}
}

// AddedAmount returns the value that was added to the "amount" field in this mutation.
func (m *TransactionMutation) AddedAmount() (r int, exists bool) {
	v := m.addamount
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmount resets all changes to the "amount" field.
func (m *TransactionMutation) ResetAmount() {
	m.amount = nil
	m.addamount = nil
}

// SetMemo sets the "memo" field.
func (m *TransactionMutation) SetMemo(s string) {
	m.memo = &s
}

// Memo returns the value of the "memo" field in the mutation.
func (m *TransactionMutation) Memo() (r string, exists bool) {
	v := m.memo
	if v == nil {
		return
	}
	return *v, true
}

// OldMemo returns the old "memo" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldMemo(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMemo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMemo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMemo: %w", err)
	}
	return oldValue.Memo, nil
}

// ClearMemo clears the value of the "memo" field.
func (m *TransactionMutation) ClearMemo() {
	m.memo = nil
	m.clearedFields[transaction.FieldMemo] = struct{}{}
}

// MemoCleared returns if the "memo" field was cleared in this mutation.
func (m *TransactionMutation) MemoCleared() bool {
	_, ok := m.clearedFields[transaction.FieldMemo]
	return ok
}

// ResetMemo resets all changes to the "memo" field.
func (m *TransactionMutation) ResetMemo() {
	m.memo = nil
	delete(m.clearedFields, transaction.FieldMemo)
}

// ClearWallet clears the "wallet" edge to the Wallet entity.
func (m *TransactionMutation) ClearWallet() {
	m.clearedwallet = true
	m.clearedFields[transaction.FieldWalletID] = struct{}{}
}

// WalletCleared reports if the "wallet" edge to the Wallet entity was cleared.
func (m *TransactionMutation) WalletCleared() bool {
	return m.clearedwallet
}

// WalletIDs returns the "wallet" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// WalletID instead. It exists only for internal usage by the builders.
func (m *TransactionMutation) WalletIDs() (ids []int) {
	if id := m.wallet; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetWallet resets all changes to the "wallet" edge.
func (m *TransactionMutation) ResetWallet() {
	m.wallet = nil
	m.clearedwallet = false
}

// Where appends a list predicates to the TransactionMutation builder.
func (m *TransactionMutation) Where(ps ...predicate.Transaction) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TransactionMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TransactionMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Transaction, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TransactionMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TransactionMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Transaction).
func (m *TransactionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TransactionMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.wallet != nil {
		fields = append(fields, transaction.FieldWalletID)
	}
	if m.paid_date != nil {
		fields = append(fields, transaction.FieldPaidDate)
	}
	if m.amount != nil {
		fields = append(fields, transaction.FieldAmount)
	}
	if m.memo != nil {
		fields = append(fields, transaction.FieldMemo)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TransactionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case transaction.FieldWalletID:
		return m.WalletID()
	case transaction.FieldPaidDate:
		return m.PaidDate()
	case transaction.FieldAmount:
		return m.Amount()
	case transaction.FieldMemo:
		return m.Memo()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TransactionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case transaction.FieldWalletID:
		return m.OldWalletID(ctx)
	case transaction.FieldPaidDate:
		return m.OldPaidDate(ctx)
	case transaction.FieldAmount:
		return m.OldAmount(ctx)
	case transaction.FieldMemo:
		return m.OldMemo(ctx)
	}
	return nil, fmt.Errorf("unknown Transaction field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TransactionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case transaction.FieldWalletID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWalletID(v)
		return nil
	case transaction.FieldPaidDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPaidDate(v)
		return nil
	case transaction.FieldAmount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmount(v)
		return nil
	case transaction.FieldMemo:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMemo(v)
		return nil
	}
	return fmt.Errorf("unknown Transaction field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TransactionMutation) AddedFields() []string {
	var fields []string
	if m.addamount != nil {
		fields = append(fields, transaction.FieldAmount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TransactionMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case transaction.FieldAmount:
		return m.AddedAmount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TransactionMutation) AddField(name string, value ent.Value) error {
	switch name {
	case transaction.FieldAmount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmount(v)
		return nil
	}
	return fmt.Errorf("unknown Transaction numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TransactionMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(transaction.FieldMemo) {
		fields = append(fields, transaction.FieldMemo)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TransactionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TransactionMutation) ClearField(name string) error {
	switch name {
	case transaction.FieldMemo:
		m.ClearMemo()
		return nil
	}
	return fmt.Errorf("unknown Transaction nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TransactionMutation) ResetField(name string) error {
	switch name {
	case transaction.FieldWalletID:
		m.ResetWalletID()
		return nil
	case transaction.FieldPaidDate:
		m.ResetPaidDate()
		return nil
	case transaction.FieldAmount:
		m.ResetAmount()
		return nil
	case transaction.FieldMemo:
		m.ResetMemo()
		return nil
	}
	return fmt.Errorf("unknown Transaction field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TransactionMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.wallet != nil {
		edges = append(edges, transaction.EdgeWallet)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TransactionMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case transaction.EdgeWallet:
		if id := m.wallet; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TransactionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TransactionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TransactionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedwallet {
		edges = append(edges, transaction.EdgeWallet)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TransactionMutation) EdgeCleared(name string) bool {
	switch name {
	case transaction.EdgeWallet:
		return m.clearedwallet
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TransactionMutation) ClearEdge(name string) error {
	switch name {
	case transaction.EdgeWallet:
		m.ClearWallet()
		return nil
	}
	return fmt.Errorf("unknown Transaction unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TransactionMutation) ResetEdge(name string) error {
	switch name {
	case transaction.EdgeWallet:
		m.ResetWallet()
		return nil
	}
	return fmt.Errorf("unknown Transaction edge %s", name)
}

// WalletMutation represents an operation that mutates the Wallet nodes in the graph.
type WalletMutation struct {
	config
	op                  Op
	typ                 string
	id                  *int
	name                *string
	method              *string
	clearedFields       map[string]struct{}
	transactions        map[int]struct{}
	removedtransactions map[int]struct{}
	clearedtransactions bool
	done                bool
	oldValue            func(context.Context) (*Wallet, error)
	predicates          []predicate.Wallet
}

var _ ent.Mutation = (*WalletMutation)(nil)

// walletOption allows management of the mutation configuration using functional options.
type walletOption func(*WalletMutation)

// newWalletMutation creates new mutation for the Wallet entity.
func newWalletMutation(c config, op Op, opts ...walletOption) *WalletMutation {
	m := &WalletMutation{
		config:        c,
		op:            op,
		typ:           TypeWallet,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWalletID sets the ID field of the mutation.
func withWalletID(id int) walletOption {
	return func(m *WalletMutation) {
		var (
			err   error
			once  sync.Once
			value *Wallet
		)
		m.oldValue = func(ctx context.Context) (*Wallet, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Wallet.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWallet sets the old Wallet of the mutation.
func withWallet(node *Wallet) walletOption {
	return func(m *WalletMutation) {
		m.oldValue = func(context.Context) (*Wallet, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WalletMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WalletMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *WalletMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *WalletMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Wallet.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *WalletMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *WalletMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Wallet entity.
// If the Wallet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WalletMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *WalletMutation) ResetName() {
	m.name = nil
}

// SetMethod sets the "method" field.
func (m *WalletMutation) SetMethod(s string) {
	m.method = &s
}

// Method returns the value of the "method" field in the mutation.
func (m *WalletMutation) Method() (r string, exists bool) {
	v := m.method
	if v == nil {
		return
	}
	return *v, true
}

// OldMethod returns the old "method" field's value of the Wallet entity.
// If the Wallet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WalletMutation) OldMethod(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMethod is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMethod requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMethod: %w", err)
	}
	return oldValue.Method, nil
}

// ResetMethod resets all changes to the "method" field.
func (m *WalletMutation) ResetMethod() {
	m.method = nil
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by ids.
func (m *WalletMutation) AddTransactionIDs(ids ...int) {
	if m.transactions == nil {
		m.transactions = make(map[int]struct{})
	}
	for i := range ids {
		m.transactions[ids[i]] = struct{}{}
	}
}

// ClearTransactions clears the "transactions" edge to the Transaction entity.
func (m *WalletMutation) ClearTransactions() {
	m.clearedtransactions = true
}

// TransactionsCleared reports if the "transactions" edge to the Transaction entity was cleared.
func (m *WalletMutation) TransactionsCleared() bool {
	return m.clearedtransactions
}

// RemoveTransactionIDs removes the "transactions" edge to the Transaction entity by IDs.
func (m *WalletMutation) RemoveTransactionIDs(ids ...int) {
	if m.removedtransactions == nil {
		m.removedtransactions = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.transactions, ids[i])
		m.removedtransactions[ids[i]] = struct{}{}
	}
}

// RemovedTransactions returns the removed IDs of the "transactions" edge to the Transaction entity.
func (m *WalletMutation) RemovedTransactionsIDs() (ids []int) {
	for id := range m.removedtransactions {
		ids = append(ids, id)
	}
	return
}

// TransactionsIDs returns the "transactions" edge IDs in the mutation.
func (m *WalletMutation) TransactionsIDs() (ids []int) {
	for id := range m.transactions {
		ids = append(ids, id)
	}
	return
}

// ResetTransactions resets all changes to the "transactions" edge.
func (m *WalletMutation) ResetTransactions() {
	m.transactions = nil
	m.clearedtransactions = false
	m.removedtransactions = nil
}

// Where appends a list predicates to the WalletMutation builder.
func (m *WalletMutation) Where(ps ...predicate.Wallet) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the WalletMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *WalletMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Wallet, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *WalletMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *WalletMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Wallet).
func (m *WalletMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *WalletMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, wallet.FieldName)
	}
	if m.method != nil {
		fields = append(fields, wallet.FieldMethod)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *WalletMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case wallet.FieldName:
		return m.Name()
	case wallet.FieldMethod:
		return m.Method()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *WalletMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case wallet.FieldName:
		return m.OldName(ctx)
	case wallet.FieldMethod:
		return m.OldMethod(ctx)
	}
	return nil, fmt.Errorf("unknown Wallet field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WalletMutation) SetField(name string, value ent.Value) error {
	switch name {
	case wallet.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case wallet.FieldMethod:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMethod(v)
		return nil
	}
	return fmt.Errorf("unknown Wallet field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *WalletMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *WalletMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WalletMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Wallet numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *WalletMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *WalletMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *WalletMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Wallet nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *WalletMutation) ResetField(name string) error {
	switch name {
	case wallet.FieldName:
		m.ResetName()
		return nil
	case wallet.FieldMethod:
		m.ResetMethod()
		return nil
	}
	return fmt.Errorf("unknown Wallet field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *WalletMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.transactions != nil {
		edges = append(edges, wallet.EdgeTransactions)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *WalletMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case wallet.EdgeTransactions:
		ids := make([]ent.Value, 0, len(m.transactions))
		for id := range m.transactions {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *WalletMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtransactions != nil {
		edges = append(edges, wallet.EdgeTransactions)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *WalletMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case wallet.EdgeTransactions:
		ids := make([]ent.Value, 0, len(m.removedtransactions))
		for id := range m.removedtransactions {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *WalletMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtransactions {
		edges = append(edges, wallet.EdgeTransactions)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *WalletMutation) EdgeCleared(name string) bool {
	switch name {
	case wallet.EdgeTransactions:
		return m.clearedtransactions
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *WalletMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Wallet unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *WalletMutation) ResetEdge(name string) error {
	switch name {
	case wallet.EdgeTransactions:
		m.ResetTransactions()
		return nil
	}
	return fmt.Errorf("unknown Wallet edge %s", name)
}
