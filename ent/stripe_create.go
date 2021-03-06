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
	"github.com/joescharf/twitterprofile/v2/ent/stripe"
)

// StripeCreate is the builder for creating a Stripe entity.
type StripeCreate struct {
	config
	mutation *StripeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAPIKey sets the "api_key" field.
func (sc *StripeCreate) SetAPIKey(s string) *StripeCreate {
	sc.mutation.SetAPIKey(s)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *StripeCreate) SetCreatedAt(t time.Time) *StripeCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StripeCreate) SetNillableCreatedAt(t *time.Time) *StripeCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StripeCreate) SetUpdatedAt(t time.Time) *StripeCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StripeCreate) SetNillableUpdatedAt(t *time.Time) *StripeCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// Mutation returns the StripeMutation object of the builder.
func (sc *StripeCreate) Mutation() *StripeMutation {
	return sc.mutation
}

// Save creates the Stripe in the database.
func (sc *StripeCreate) Save(ctx context.Context) (*Stripe, error) {
	var (
		err  error
		node *Stripe
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StripeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StripeCreate) SaveX(ctx context.Context) *Stripe {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StripeCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StripeCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StripeCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := stripe.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := stripe.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StripeCreate) check() error {
	if _, ok := sc.mutation.APIKey(); !ok {
		return &ValidationError{Name: "api_key", err: errors.New(`ent: missing required field "Stripe.api_key"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Stripe.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Stripe.updated_at"`)}
	}
	return nil
}

func (sc *StripeCreate) sqlSave(ctx context.Context) (*Stripe, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *StripeCreate) createSpec() (*Stripe, *sqlgraph.CreateSpec) {
	var (
		_node = &Stripe{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: stripe.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stripe.FieldID,
			},
		}
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.APIKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stripe.FieldAPIKey,
		})
		_node.APIKey = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: stripe.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: stripe.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Stripe.Create().
//		SetAPIKey(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StripeUpsert) {
//			SetAPIKey(v+v).
//		}).
//		Exec(ctx)
//
func (sc *StripeCreate) OnConflict(opts ...sql.ConflictOption) *StripeUpsertOne {
	sc.conflict = opts
	return &StripeUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Stripe.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sc *StripeCreate) OnConflictColumns(columns ...string) *StripeUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &StripeUpsertOne{
		create: sc,
	}
}

type (
	// StripeUpsertOne is the builder for "upsert"-ing
	//  one Stripe node.
	StripeUpsertOne struct {
		create *StripeCreate
	}

	// StripeUpsert is the "OnConflict" setter.
	StripeUpsert struct {
		*sql.UpdateSet
	}
)

// SetAPIKey sets the "api_key" field.
func (u *StripeUpsert) SetAPIKey(v string) *StripeUpsert {
	u.Set(stripe.FieldAPIKey, v)
	return u
}

// UpdateAPIKey sets the "api_key" field to the value that was provided on create.
func (u *StripeUpsert) UpdateAPIKey() *StripeUpsert {
	u.SetExcluded(stripe.FieldAPIKey)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *StripeUpsert) SetCreatedAt(v time.Time) *StripeUpsert {
	u.Set(stripe.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StripeUpsert) UpdateCreatedAt() *StripeUpsert {
	u.SetExcluded(stripe.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StripeUpsert) SetUpdatedAt(v time.Time) *StripeUpsert {
	u.Set(stripe.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StripeUpsert) UpdateUpdatedAt() *StripeUpsert {
	u.SetExcluded(stripe.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Stripe.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *StripeUpsertOne) UpdateNewValues() *StripeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Stripe.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *StripeUpsertOne) Ignore() *StripeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StripeUpsertOne) DoNothing() *StripeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StripeCreate.OnConflict
// documentation for more info.
func (u *StripeUpsertOne) Update(set func(*StripeUpsert)) *StripeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StripeUpsert{UpdateSet: update})
	}))
	return u
}

// SetAPIKey sets the "api_key" field.
func (u *StripeUpsertOne) SetAPIKey(v string) *StripeUpsertOne {
	return u.Update(func(s *StripeUpsert) {
		s.SetAPIKey(v)
	})
}

// UpdateAPIKey sets the "api_key" field to the value that was provided on create.
func (u *StripeUpsertOne) UpdateAPIKey() *StripeUpsertOne {
	return u.Update(func(s *StripeUpsert) {
		s.UpdateAPIKey()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *StripeUpsertOne) SetCreatedAt(v time.Time) *StripeUpsertOne {
	return u.Update(func(s *StripeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StripeUpsertOne) UpdateCreatedAt() *StripeUpsertOne {
	return u.Update(func(s *StripeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StripeUpsertOne) SetUpdatedAt(v time.Time) *StripeUpsertOne {
	return u.Update(func(s *StripeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StripeUpsertOne) UpdateUpdatedAt() *StripeUpsertOne {
	return u.Update(func(s *StripeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *StripeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StripeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StripeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StripeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StripeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StripeCreateBulk is the builder for creating many Stripe entities in bulk.
type StripeCreateBulk struct {
	config
	builders []*StripeCreate
	conflict []sql.ConflictOption
}

// Save creates the Stripe entities in the database.
func (scb *StripeCreateBulk) Save(ctx context.Context) ([]*Stripe, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stripe, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StripeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StripeCreateBulk) SaveX(ctx context.Context) []*Stripe {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StripeCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StripeCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Stripe.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StripeUpsert) {
//			SetAPIKey(v+v).
//		}).
//		Exec(ctx)
//
func (scb *StripeCreateBulk) OnConflict(opts ...sql.ConflictOption) *StripeUpsertBulk {
	scb.conflict = opts
	return &StripeUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Stripe.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (scb *StripeCreateBulk) OnConflictColumns(columns ...string) *StripeUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &StripeUpsertBulk{
		create: scb,
	}
}

// StripeUpsertBulk is the builder for "upsert"-ing
// a bulk of Stripe nodes.
type StripeUpsertBulk struct {
	create *StripeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Stripe.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *StripeUpsertBulk) UpdateNewValues() *StripeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Stripe.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *StripeUpsertBulk) Ignore() *StripeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StripeUpsertBulk) DoNothing() *StripeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StripeCreateBulk.OnConflict
// documentation for more info.
func (u *StripeUpsertBulk) Update(set func(*StripeUpsert)) *StripeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StripeUpsert{UpdateSet: update})
	}))
	return u
}

// SetAPIKey sets the "api_key" field.
func (u *StripeUpsertBulk) SetAPIKey(v string) *StripeUpsertBulk {
	return u.Update(func(s *StripeUpsert) {
		s.SetAPIKey(v)
	})
}

// UpdateAPIKey sets the "api_key" field to the value that was provided on create.
func (u *StripeUpsertBulk) UpdateAPIKey() *StripeUpsertBulk {
	return u.Update(func(s *StripeUpsert) {
		s.UpdateAPIKey()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *StripeUpsertBulk) SetCreatedAt(v time.Time) *StripeUpsertBulk {
	return u.Update(func(s *StripeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StripeUpsertBulk) UpdateCreatedAt() *StripeUpsertBulk {
	return u.Update(func(s *StripeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StripeUpsertBulk) SetUpdatedAt(v time.Time) *StripeUpsertBulk {
	return u.Update(func(s *StripeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StripeUpsertBulk) UpdateUpdatedAt() *StripeUpsertBulk {
	return u.Update(func(s *StripeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *StripeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StripeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StripeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StripeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
