// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/lufia/enttut/ent/schema"
	"github.com/lufia/enttut/ent/wallet"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	walletFields := schema.Wallet{}.Fields()
	_ = walletFields
	// walletDescName is the schema descriptor for name field.
	walletDescName := walletFields[0].Descriptor()
	// wallet.NameValidator is a validator for the "name" field. It is called by the builders before save.
	wallet.NameValidator = func() func(string) error {
		validators := walletDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
