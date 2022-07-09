// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joescharf/twitterprofile/v2/ent/predicate"
	"github.com/joescharf/twitterprofile/v2/ent/stripe"
	"github.com/joescharf/twitterprofile/v2/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetScreenName sets the "screen_name" field.
func (uu *UserUpdate) SetScreenName(s string) *UserUpdate {
	uu.mutation.SetScreenName(s)
	return uu
}

// SetTwitterUserID sets the "twitter_user_id" field.
func (uu *UserUpdate) SetTwitterUserID(i int64) *UserUpdate {
	uu.mutation.ResetTwitterUserID()
	uu.mutation.SetTwitterUserID(i)
	return uu
}

// AddTwitterUserID adds i to the "twitter_user_id" field.
func (uu *UserUpdate) AddTwitterUserID(i int64) *UserUpdate {
	uu.mutation.AddTwitterUserID(i)
	return uu
}

// SetDescription sets the "description" field.
func (uu *UserUpdate) SetDescription(s string) *UserUpdate {
	uu.mutation.SetDescription(s)
	return uu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDescription(s *string) *UserUpdate {
	if s != nil {
		uu.SetDescription(*s)
	}
	return uu
}

// ClearDescription clears the value of the "description" field.
func (uu *UserUpdate) ClearDescription() *UserUpdate {
	uu.mutation.ClearDescription()
	return uu
}

// SetToken sets the "token" field.
func (uu *UserUpdate) SetToken(s string) *UserUpdate {
	uu.mutation.SetToken(s)
	return uu
}

// SetTokenSecret sets the "token_secret" field.
func (uu *UserUpdate) SetTokenSecret(s string) *UserUpdate {
	uu.mutation.SetTokenSecret(s)
	return uu
}

// SetTwitterProfileImageURL sets the "twitter_profile_image_url" field.
func (uu *UserUpdate) SetTwitterProfileImageURL(s string) *UserUpdate {
	uu.mutation.SetTwitterProfileImageURL(s)
	return uu
}

// SetNillableTwitterProfileImageURL sets the "twitter_profile_image_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTwitterProfileImageURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetTwitterProfileImageURL(*s)
	}
	return uu
}

// ClearTwitterProfileImageURL clears the value of the "twitter_profile_image_url" field.
func (uu *UserUpdate) ClearTwitterProfileImageURL() *UserUpdate {
	uu.mutation.ClearTwitterProfileImageURL()
	return uu
}

// SetMin sets the "min" field.
func (uu *UserUpdate) SetMin(i int64) *UserUpdate {
	uu.mutation.ResetMin()
	uu.mutation.SetMin(i)
	return uu
}

// SetNillableMin sets the "min" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMin(i *int64) *UserUpdate {
	if i != nil {
		uu.SetMin(*i)
	}
	return uu
}

// AddMin adds i to the "min" field.
func (uu *UserUpdate) AddMin(i int64) *UserUpdate {
	uu.mutation.AddMin(i)
	return uu
}

// ClearMin clears the value of the "min" field.
func (uu *UserUpdate) ClearMin() *UserUpdate {
	uu.mutation.ClearMin()
	return uu
}

// SetMax sets the "max" field.
func (uu *UserUpdate) SetMax(i int64) *UserUpdate {
	uu.mutation.ResetMax()
	uu.mutation.SetMax(i)
	return uu
}

// SetNillableMax sets the "max" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMax(i *int64) *UserUpdate {
	if i != nil {
		uu.SetMax(*i)
	}
	return uu
}

// AddMax adds i to the "max" field.
func (uu *UserUpdate) AddMax(i int64) *UserUpdate {
	uu.mutation.AddMax(i)
	return uu
}

// ClearMax clears the value of the "max" field.
func (uu *UserUpdate) ClearMax() *UserUpdate {
	uu.mutation.ClearMax()
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// AddAccountIDs adds the "accounts" edge to the Stripe entity by IDs.
func (uu *UserUpdate) AddAccountIDs(ids ...int) *UserUpdate {
	uu.mutation.AddAccountIDs(ids...)
	return uu
}

// AddAccounts adds the "accounts" edges to the Stripe entity.
func (uu *UserUpdate) AddAccounts(s ...*Stripe) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddAccountIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearAccounts clears all "accounts" edges to the Stripe entity.
func (uu *UserUpdate) ClearAccounts() *UserUpdate {
	uu.mutation.ClearAccounts()
	return uu
}

// RemoveAccountIDs removes the "accounts" edge to Stripe entities by IDs.
func (uu *UserUpdate) RemoveAccountIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveAccountIDs(ids...)
	return uu
}

// RemoveAccounts removes "accounts" edges to Stripe entities.
func (uu *UserUpdate) RemoveAccounts(s ...*Stripe) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveAccountIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.ScreenName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldScreenName,
		})
	}
	if value, ok := uu.mutation.TwitterUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTwitterUserID,
		})
	}
	if value, ok := uu.mutation.AddedTwitterUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTwitterUserID,
		})
	}
	if value, ok := uu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDescription,
		})
	}
	if uu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldDescription,
		})
	}
	if value, ok := uu.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldToken,
		})
	}
	if value, ok := uu.mutation.TokenSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTokenSecret,
		})
	}
	if value, ok := uu.mutation.TwitterProfileImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTwitterProfileImageURL,
		})
	}
	if uu.mutation.TwitterProfileImageURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldTwitterProfileImageURL,
		})
	}
	if value, ok := uu.mutation.Min(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldMin,
		})
	}
	if value, ok := uu.mutation.AddedMin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldMin,
		})
	}
	if uu.mutation.MinCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: user.FieldMin,
		})
	}
	if value, ok := uu.mutation.Max(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldMax,
		})
	}
	if value, ok := uu.mutation.AddedMax(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldMax,
		})
	}
	if uu.mutation.MaxCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: user.FieldMax,
		})
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: stripe.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !uu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: stripe.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: stripe.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetScreenName sets the "screen_name" field.
func (uuo *UserUpdateOne) SetScreenName(s string) *UserUpdateOne {
	uuo.mutation.SetScreenName(s)
	return uuo
}

// SetTwitterUserID sets the "twitter_user_id" field.
func (uuo *UserUpdateOne) SetTwitterUserID(i int64) *UserUpdateOne {
	uuo.mutation.ResetTwitterUserID()
	uuo.mutation.SetTwitterUserID(i)
	return uuo
}

// AddTwitterUserID adds i to the "twitter_user_id" field.
func (uuo *UserUpdateOne) AddTwitterUserID(i int64) *UserUpdateOne {
	uuo.mutation.AddTwitterUserID(i)
	return uuo
}

// SetDescription sets the "description" field.
func (uuo *UserUpdateOne) SetDescription(s string) *UserUpdateOne {
	uuo.mutation.SetDescription(s)
	return uuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDescription(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDescription(*s)
	}
	return uuo
}

// ClearDescription clears the value of the "description" field.
func (uuo *UserUpdateOne) ClearDescription() *UserUpdateOne {
	uuo.mutation.ClearDescription()
	return uuo
}

// SetToken sets the "token" field.
func (uuo *UserUpdateOne) SetToken(s string) *UserUpdateOne {
	uuo.mutation.SetToken(s)
	return uuo
}

// SetTokenSecret sets the "token_secret" field.
func (uuo *UserUpdateOne) SetTokenSecret(s string) *UserUpdateOne {
	uuo.mutation.SetTokenSecret(s)
	return uuo
}

// SetTwitterProfileImageURL sets the "twitter_profile_image_url" field.
func (uuo *UserUpdateOne) SetTwitterProfileImageURL(s string) *UserUpdateOne {
	uuo.mutation.SetTwitterProfileImageURL(s)
	return uuo
}

// SetNillableTwitterProfileImageURL sets the "twitter_profile_image_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTwitterProfileImageURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTwitterProfileImageURL(*s)
	}
	return uuo
}

// ClearTwitterProfileImageURL clears the value of the "twitter_profile_image_url" field.
func (uuo *UserUpdateOne) ClearTwitterProfileImageURL() *UserUpdateOne {
	uuo.mutation.ClearTwitterProfileImageURL()
	return uuo
}

// SetMin sets the "min" field.
func (uuo *UserUpdateOne) SetMin(i int64) *UserUpdateOne {
	uuo.mutation.ResetMin()
	uuo.mutation.SetMin(i)
	return uuo
}

// SetNillableMin sets the "min" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMin(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetMin(*i)
	}
	return uuo
}

// AddMin adds i to the "min" field.
func (uuo *UserUpdateOne) AddMin(i int64) *UserUpdateOne {
	uuo.mutation.AddMin(i)
	return uuo
}

// ClearMin clears the value of the "min" field.
func (uuo *UserUpdateOne) ClearMin() *UserUpdateOne {
	uuo.mutation.ClearMin()
	return uuo
}

// SetMax sets the "max" field.
func (uuo *UserUpdateOne) SetMax(i int64) *UserUpdateOne {
	uuo.mutation.ResetMax()
	uuo.mutation.SetMax(i)
	return uuo
}

// SetNillableMax sets the "max" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMax(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetMax(*i)
	}
	return uuo
}

// AddMax adds i to the "max" field.
func (uuo *UserUpdateOne) AddMax(i int64) *UserUpdateOne {
	uuo.mutation.AddMax(i)
	return uuo
}

// ClearMax clears the value of the "max" field.
func (uuo *UserUpdateOne) ClearMax() *UserUpdateOne {
	uuo.mutation.ClearMax()
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// AddAccountIDs adds the "accounts" edge to the Stripe entity by IDs.
func (uuo *UserUpdateOne) AddAccountIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddAccountIDs(ids...)
	return uuo
}

// AddAccounts adds the "accounts" edges to the Stripe entity.
func (uuo *UserUpdateOne) AddAccounts(s ...*Stripe) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddAccountIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearAccounts clears all "accounts" edges to the Stripe entity.
func (uuo *UserUpdateOne) ClearAccounts() *UserUpdateOne {
	uuo.mutation.ClearAccounts()
	return uuo
}

// RemoveAccountIDs removes the "accounts" edge to Stripe entities by IDs.
func (uuo *UserUpdateOne) RemoveAccountIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveAccountIDs(ids...)
	return uuo
}

// RemoveAccounts removes "accounts" edges to Stripe entities.
func (uuo *UserUpdateOne) RemoveAccounts(s ...*Stripe) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveAccountIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.ScreenName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldScreenName,
		})
	}
	if value, ok := uuo.mutation.TwitterUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTwitterUserID,
		})
	}
	if value, ok := uuo.mutation.AddedTwitterUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTwitterUserID,
		})
	}
	if value, ok := uuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDescription,
		})
	}
	if uuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldDescription,
		})
	}
	if value, ok := uuo.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldToken,
		})
	}
	if value, ok := uuo.mutation.TokenSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTokenSecret,
		})
	}
	if value, ok := uuo.mutation.TwitterProfileImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTwitterProfileImageURL,
		})
	}
	if uuo.mutation.TwitterProfileImageURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldTwitterProfileImageURL,
		})
	}
	if value, ok := uuo.mutation.Min(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldMin,
		})
	}
	if value, ok := uuo.mutation.AddedMin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldMin,
		})
	}
	if uuo.mutation.MinCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: user.FieldMin,
		})
	}
	if value, ok := uuo.mutation.Max(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldMax,
		})
	}
	if value, ok := uuo.mutation.AddedMax(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldMax,
		})
	}
	if uuo.mutation.MaxCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: user.FieldMax,
		})
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uuo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: stripe.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !uuo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: stripe.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: stripe.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
