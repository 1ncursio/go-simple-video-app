// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/1ncursio/go-simple-video-app/ent/predicate"
	"github.com/1ncursio/go-simple-video-app/ent/tag"
	"github.com/1ncursio/go-simple-video-app/ent/video"
)

// VideoUpdate is the builder for updating Video entities.
type VideoUpdate struct {
	config
	hooks    []Hook
	mutation *VideoMutation
}

// Where appends a list predicates to the VideoUpdate builder.
func (vu *VideoUpdate) Where(ps ...predicate.Video) *VideoUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetURL sets the "url" field.
func (vu *VideoUpdate) SetURL(s string) *VideoUpdate {
	vu.mutation.SetURL(s)
	return vu
}

// SetTitle sets the "title" field.
func (vu *VideoUpdate) SetTitle(s string) *VideoUpdate {
	vu.mutation.SetTitle(s)
	return vu
}

// SetViews sets the "views" field.
func (vu *VideoUpdate) SetViews(i int) *VideoUpdate {
	vu.mutation.ResetViews()
	vu.mutation.SetViews(i)
	return vu
}

// SetNillableViews sets the "views" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableViews(i *int) *VideoUpdate {
	if i != nil {
		vu.SetViews(*i)
	}
	return vu
}

// AddViews adds i to the "views" field.
func (vu *VideoUpdate) AddViews(i int) *VideoUpdate {
	vu.mutation.AddViews(i)
	return vu
}

// SetCreatedAt sets the "created_at" field.
func (vu *VideoUpdate) SetCreatedAt(t time.Time) *VideoUpdate {
	vu.mutation.SetCreatedAt(t)
	return vu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableCreatedAt(t *time.Time) *VideoUpdate {
	if t != nil {
		vu.SetCreatedAt(*t)
	}
	return vu
}

// SetUpdatedAt sets the "updated_at" field.
func (vu *VideoUpdate) SetUpdatedAt(t time.Time) *VideoUpdate {
	vu.mutation.SetUpdatedAt(t)
	return vu
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (vu *VideoUpdate) AddTagIDs(ids ...int) *VideoUpdate {
	vu.mutation.AddTagIDs(ids...)
	return vu
}

// AddTags adds the "tags" edges to the Tag entity.
func (vu *VideoUpdate) AddTags(t ...*Tag) *VideoUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vu.AddTagIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vu *VideoUpdate) Mutation() *VideoMutation {
	return vu.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (vu *VideoUpdate) ClearTags() *VideoUpdate {
	vu.mutation.ClearTags()
	return vu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (vu *VideoUpdate) RemoveTagIDs(ids ...int) *VideoUpdate {
	vu.mutation.RemoveTagIDs(ids...)
	return vu
}

// RemoveTags removes "tags" edges to Tag entities.
func (vu *VideoUpdate) RemoveTags(t ...*Tag) *VideoUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vu.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VideoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	vu.defaults()
	if len(vu.hooks) == 0 {
		affected, err = vu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vu.mutation = mutation
			affected, err = vu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vu.hooks) - 1; i >= 0; i-- {
			if vu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VideoUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VideoUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VideoUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vu *VideoUpdate) defaults() {
	if _, ok := vu.mutation.UpdatedAt(); !ok {
		v := video.UpdateDefaultUpdatedAt()
		vu.mutation.SetUpdatedAt(v)
	}
}

func (vu *VideoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   video.Table,
			Columns: video.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: video.FieldID,
			},
		},
	}
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldURL,
		})
	}
	if value, ok := vu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldTitle,
		})
	}
	if value, ok := vu.mutation.Views(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: video.FieldViews,
		})
	}
	if value, ok := vu.mutation.AddedViews(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: video.FieldViews,
		})
	}
	if value, ok := vu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: video.FieldCreatedAt,
		})
	}
	if value, ok := vu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: video.FieldUpdatedAt,
		})
	}
	if vu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   video.TagsTable,
			Columns: video.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !vu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   video.TagsTable,
			Columns: video.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   video.TagsTable,
			Columns: video.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{video.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// VideoUpdateOne is the builder for updating a single Video entity.
type VideoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VideoMutation
}

// SetURL sets the "url" field.
func (vuo *VideoUpdateOne) SetURL(s string) *VideoUpdateOne {
	vuo.mutation.SetURL(s)
	return vuo
}

// SetTitle sets the "title" field.
func (vuo *VideoUpdateOne) SetTitle(s string) *VideoUpdateOne {
	vuo.mutation.SetTitle(s)
	return vuo
}

// SetViews sets the "views" field.
func (vuo *VideoUpdateOne) SetViews(i int) *VideoUpdateOne {
	vuo.mutation.ResetViews()
	vuo.mutation.SetViews(i)
	return vuo
}

// SetNillableViews sets the "views" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableViews(i *int) *VideoUpdateOne {
	if i != nil {
		vuo.SetViews(*i)
	}
	return vuo
}

// AddViews adds i to the "views" field.
func (vuo *VideoUpdateOne) AddViews(i int) *VideoUpdateOne {
	vuo.mutation.AddViews(i)
	return vuo
}

// SetCreatedAt sets the "created_at" field.
func (vuo *VideoUpdateOne) SetCreatedAt(t time.Time) *VideoUpdateOne {
	vuo.mutation.SetCreatedAt(t)
	return vuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableCreatedAt(t *time.Time) *VideoUpdateOne {
	if t != nil {
		vuo.SetCreatedAt(*t)
	}
	return vuo
}

// SetUpdatedAt sets the "updated_at" field.
func (vuo *VideoUpdateOne) SetUpdatedAt(t time.Time) *VideoUpdateOne {
	vuo.mutation.SetUpdatedAt(t)
	return vuo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (vuo *VideoUpdateOne) AddTagIDs(ids ...int) *VideoUpdateOne {
	vuo.mutation.AddTagIDs(ids...)
	return vuo
}

// AddTags adds the "tags" edges to the Tag entity.
func (vuo *VideoUpdateOne) AddTags(t ...*Tag) *VideoUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vuo.AddTagIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vuo *VideoUpdateOne) Mutation() *VideoMutation {
	return vuo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (vuo *VideoUpdateOne) ClearTags() *VideoUpdateOne {
	vuo.mutation.ClearTags()
	return vuo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (vuo *VideoUpdateOne) RemoveTagIDs(ids ...int) *VideoUpdateOne {
	vuo.mutation.RemoveTagIDs(ids...)
	return vuo
}

// RemoveTags removes "tags" edges to Tag entities.
func (vuo *VideoUpdateOne) RemoveTags(t ...*Tag) *VideoUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vuo.RemoveTagIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VideoUpdateOne) Select(field string, fields ...string) *VideoUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Video entity.
func (vuo *VideoUpdateOne) Save(ctx context.Context) (*Video, error) {
	var (
		err  error
		node *Video
	)
	vuo.defaults()
	if len(vuo.hooks) == 0 {
		node, err = vuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vuo.mutation = mutation
			node, err = vuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuo.hooks) - 1; i >= 0; i-- {
			if vuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VideoUpdateOne) SaveX(ctx context.Context) *Video {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VideoUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VideoUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vuo *VideoUpdateOne) defaults() {
	if _, ok := vuo.mutation.UpdatedAt(); !ok {
		v := video.UpdateDefaultUpdatedAt()
		vuo.mutation.SetUpdatedAt(v)
	}
}

func (vuo *VideoUpdateOne) sqlSave(ctx context.Context) (_node *Video, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   video.Table,
			Columns: video.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: video.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Video.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, video.FieldID)
		for _, f := range fields {
			if !video.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != video.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldURL,
		})
	}
	if value, ok := vuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldTitle,
		})
	}
	if value, ok := vuo.mutation.Views(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: video.FieldViews,
		})
	}
	if value, ok := vuo.mutation.AddedViews(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: video.FieldViews,
		})
	}
	if value, ok := vuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: video.FieldCreatedAt,
		})
	}
	if value, ok := vuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: video.FieldUpdatedAt,
		})
	}
	if vuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   video.TagsTable,
			Columns: video.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !vuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   video.TagsTable,
			Columns: video.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   video.TagsTable,
			Columns: video.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Video{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{video.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
