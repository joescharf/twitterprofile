// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StripesColumns holds the columns for the "stripes" table.
	StripesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "api_key", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_accounts", Type: field.TypeInt, Nullable: true},
	}
	// StripesTable holds the schema information for the "stripes" table.
	StripesTable = &schema.Table{
		Name:       "stripes",
		Columns:    StripesColumns,
		PrimaryKey: []*schema.Column{StripesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "stripes_users_accounts",
				Columns:    []*schema.Column{StripesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "screen_name", Type: field.TypeString, Unique: true},
		{Name: "twitter_user_id", Type: field.TypeInt64, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "token", Type: field.TypeString},
		{Name: "token_secret", Type: field.TypeString},
		{Name: "twitter_profile_image_url", Type: field.TypeString, Nullable: true},
		{Name: "min", Type: field.TypeInt64, Nullable: true},
		{Name: "max", Type: field.TypeInt64, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StripesTable,
		UsersTable,
	}
)

func init() {
	StripesTable.ForeignKeys[0].RefTable = UsersTable
}
