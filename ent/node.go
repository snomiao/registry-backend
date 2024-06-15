// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"registry-backend/ent/node"
	"registry-backend/ent/publisher"
	"registry-backend/ent/schema"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Node is the model entity for the Node schema.
type Node struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// PublisherID holds the value of the "publisher_id" field.
	PublisherID string `json:"publisher_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
	// License holds the value of the "license" field.
	License string `json:"license,omitempty"`
	// RepositoryURL holds the value of the "repository_url" field.
	RepositoryURL string `json:"repository_url,omitempty"`
	// IconURL holds the value of the "icon_url" field.
	IconURL string `json:"icon_url,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// TotalInstall holds the value of the "total_install" field.
	TotalInstall int64 `json:"total_install,omitempty"`
	// TotalStar holds the value of the "total_star" field.
	TotalStar int64 `json:"total_star,omitempty"`
	// TotalReview holds the value of the "total_review" field.
	TotalReview int64 `json:"total_review,omitempty"`
	// Status holds the value of the "status" field.
	Status schema.NodeStatus `json:"status,omitempty"`
	// StatusDetail holds the value of the "status_detail" field.
	StatusDetail string `json:"status_detail,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NodeQuery when eager-loading is set.
	Edges        NodeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NodeEdges holds the relations/edges for other nodes in the graph.
type NodeEdges struct {
	// Publisher holds the value of the publisher edge.
	Publisher *Publisher `json:"publisher,omitempty"`
	// Versions holds the value of the versions edge.
	Versions []*NodeVersion `json:"versions,omitempty"`
	// Reviews holds the value of the reviews edge.
	Reviews []*NodeReview `json:"reviews,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PublisherOrErr returns the Publisher value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NodeEdges) PublisherOrErr() (*Publisher, error) {
	if e.Publisher != nil {
		return e.Publisher, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: publisher.Label}
	}
	return nil, &NotLoadedError{edge: "publisher"}
}

// VersionsOrErr returns the Versions value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) VersionsOrErr() ([]*NodeVersion, error) {
	if e.loadedTypes[1] {
		return e.Versions, nil
	}
	return nil, &NotLoadedError{edge: "versions"}
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) ReviewsOrErr() ([]*NodeReview, error) {
	if e.loadedTypes[2] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Node) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case node.FieldTags:
			values[i] = new([]byte)
		case node.FieldTotalInstall, node.FieldTotalStar, node.FieldTotalReview:
			values[i] = new(sql.NullInt64)
		case node.FieldID, node.FieldPublisherID, node.FieldName, node.FieldDescription, node.FieldCategory, node.FieldAuthor, node.FieldLicense, node.FieldRepositoryURL, node.FieldIconURL, node.FieldStatus, node.FieldStatusDetail:
			values[i] = new(sql.NullString)
		case node.FieldCreateTime, node.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Node fields.
func (n *Node) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case node.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				n.ID = value.String
			}
		case node.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				n.CreateTime = value.Time
			}
		case node.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				n.UpdateTime = value.Time
			}
		case node.FieldPublisherID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field publisher_id", values[i])
			} else if value.Valid {
				n.PublisherID = value.String
			}
		case node.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				n.Name = value.String
			}
		case node.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				n.Description = value.String
			}
		case node.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				n.Category = value.String
			}
		case node.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				n.Author = value.String
			}
		case node.FieldLicense:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field license", values[i])
			} else if value.Valid {
				n.License = value.String
			}
		case node.FieldRepositoryURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repository_url", values[i])
			} else if value.Valid {
				n.RepositoryURL = value.String
			}
		case node.FieldIconURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon_url", values[i])
			} else if value.Valid {
				n.IconURL = value.String
			}
		case node.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case node.FieldTotalInstall:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_install", values[i])
			} else if value.Valid {
				n.TotalInstall = value.Int64
			}
		case node.FieldTotalStar:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_star", values[i])
			} else if value.Valid {
				n.TotalStar = value.Int64
			}
		case node.FieldTotalReview:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_review", values[i])
			} else if value.Valid {
				n.TotalReview = value.Int64
			}
		case node.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				n.Status = schema.NodeStatus(value.String)
			}
		case node.FieldStatusDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status_detail", values[i])
			} else if value.Valid {
				n.StatusDetail = value.String
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Node.
// This includes values selected through modifiers, order, etc.
func (n *Node) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryPublisher queries the "publisher" edge of the Node entity.
func (n *Node) QueryPublisher() *PublisherQuery {
	return NewNodeClient(n.config).QueryPublisher(n)
}

// QueryVersions queries the "versions" edge of the Node entity.
func (n *Node) QueryVersions() *NodeVersionQuery {
	return NewNodeClient(n.config).QueryVersions(n)
}

// QueryReviews queries the "reviews" edge of the Node entity.
func (n *Node) QueryReviews() *NodeReviewQuery {
	return NewNodeClient(n.config).QueryReviews(n)
}

// Update returns a builder for updating this Node.
// Note that you need to call Node.Unwrap() before calling this method if this Node
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Node) Update() *NodeUpdateOne {
	return NewNodeClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Node entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Node) Unwrap() *Node {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Node is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Node) String() string {
	var builder strings.Builder
	builder.WriteString("Node(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("create_time=")
	builder.WriteString(n.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(n.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("publisher_id=")
	builder.WriteString(n.PublisherID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(n.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(n.Description)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(n.Category)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(n.Author)
	builder.WriteString(", ")
	builder.WriteString("license=")
	builder.WriteString(n.License)
	builder.WriteString(", ")
	builder.WriteString("repository_url=")
	builder.WriteString(n.RepositoryURL)
	builder.WriteString(", ")
	builder.WriteString("icon_url=")
	builder.WriteString(n.IconURL)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", n.Tags))
	builder.WriteString(", ")
	builder.WriteString("total_install=")
	builder.WriteString(fmt.Sprintf("%v", n.TotalInstall))
	builder.WriteString(", ")
	builder.WriteString("total_star=")
	builder.WriteString(fmt.Sprintf("%v", n.TotalStar))
	builder.WriteString(", ")
	builder.WriteString("total_review=")
	builder.WriteString(fmt.Sprintf("%v", n.TotalReview))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", n.Status))
	builder.WriteString(", ")
	builder.WriteString("status_detail=")
	builder.WriteString(n.StatusDetail)
	builder.WriteByte(')')
	return builder.String()
}

// Nodes is a parsable slice of Node.
type Nodes []*Node
