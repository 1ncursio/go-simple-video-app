// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/1ncursio/go-simple-video-app/ent/video"
)

// Video is the model entity for the Video schema.
type Video struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Views holds the value of the "views" field.
	Views int `json:"views,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VideoQuery when eager-loading is set.
	Edges VideoEdges `json:"edges"`
}

// VideoEdges holds the relations/edges for other nodes in the graph.
type VideoEdges struct {
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e VideoEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Video) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case video.FieldID, video.FieldViews:
			values[i] = new(sql.NullInt64)
		case video.FieldURL, video.FieldTitle:
			values[i] = new(sql.NullString)
		case video.FieldCreatedAt, video.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Video", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Video fields.
func (v *Video) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case video.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case video.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				v.URL = value.String
			}
		case video.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				v.Title = value.String
			}
		case video.FieldViews:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field views", values[i])
			} else if value.Valid {
				v.Views = int(value.Int64)
			}
		case video.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case video.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				v.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTags queries the "tags" edge of the Video entity.
func (v *Video) QueryTags() *TagQuery {
	return (&VideoClient{config: v.config}).QueryTags(v)
}

// Update returns a builder for updating this Video.
// Note that you need to call Video.Unwrap() before calling this method if this Video
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Video) Update() *VideoUpdateOne {
	return (&VideoClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the Video entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Video) Unwrap() *Video {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Video is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Video) String() string {
	var builder strings.Builder
	builder.WriteString("Video(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", url=")
	builder.WriteString(v.URL)
	builder.WriteString(", title=")
	builder.WriteString(v.Title)
	builder.WriteString(", views=")
	builder.WriteString(fmt.Sprintf("%v", v.Views))
	builder.WriteString(", created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(v.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Videos is a parsable slice of Video.
type Videos []*Video

func (v Videos) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
