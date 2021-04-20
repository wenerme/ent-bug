// Code generated by entc, DO NOT EDIT.

package pet

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the pet type in the database.
	Label = "pet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOwnerID holds the string denoting the ownerid field in the database.
	FieldOwnerID = "owner_id"
	// FieldOwnerType holds the string denoting the ownertype field in the database.
	FieldOwnerType = "owner_type"
	// FieldOwningUserID holds the string denoting the owninguserid field in the database.
	FieldOwningUserID = "owning_user_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeOwningUser holds the string denoting the owninguser edge name in mutations.
	EdgeOwningUser = "owningUser"
	// Table holds the table name of the pet in the database.
	Table = "pets"
	// OwningUserTable is the table the holds the owningUser relation/edge.
	OwningUserTable = "pets"
	// OwningUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwningUserInverseTable = "users"
	// OwningUserColumn is the table column denoting the owningUser relation/edge.
	OwningUserColumn = "owning_user_id"
)

// Columns holds all SQL columns for pet fields.
var Columns = []string{
	FieldID,
	FieldOwnerID,
	FieldOwnerType,
	FieldOwningUserID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/wenerme/ent-demo/ent/runtime"
//
var (
	Hooks [1]ent.Hook
)