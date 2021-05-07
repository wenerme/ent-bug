// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "uid", Type: field.TypeUUID, Unique: true, Default: "gen_random_uuid()"},
		{Name: "owner_id", Type: field.TypeString, Nullable: true},
		{Name: "owner_type", Type: field.TypeString, Nullable: true},
		{Name: "owner_uid", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "owning_user_id", Type: field.TypeString, Nullable: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pets_users_owningUser",
				Columns:    []*schema.Column{PetsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "pet_owner_type_owner_id",
				Unique:  false,
				Columns: []*schema.Column{PetsColumns[3], PetsColumns[2]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "uid", Type: field.TypeUUID, Unique: true, Default: "gen_random_uuid()"},
		{Name: "name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PetsTable,
		UsersTable,
	}
)

func init() {
	PetsTable.ForeignKeys[0].RefTable = UsersTable
}
