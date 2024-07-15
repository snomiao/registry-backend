// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"registry-backend/ent/schema"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIsApproved holds the string denoting the is_approved field in the database.
	FieldIsApproved = "is_approved"
	// FieldIsAdmin holds the string denoting the is_admin field in the database.
	FieldIsAdmin = "is_admin"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgePublisherPermissions holds the string denoting the publisher_permissions edge name in mutations.
	EdgePublisherPermissions = "publisher_permissions"
	// EdgeReviews holds the string denoting the reviews edge name in mutations.
	EdgeReviews = "reviews"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PublisherPermissionsTable is the table that holds the publisher_permissions relation/edge.
	PublisherPermissionsTable = "publisher_permissions"
	// PublisherPermissionsInverseTable is the table name for the PublisherPermission entity.
	// It exists in this package in order to avoid circular dependency with the "publisherpermission" package.
	PublisherPermissionsInverseTable = "publisher_permissions"
	// PublisherPermissionsColumn is the table column denoting the publisher_permissions relation/edge.
	PublisherPermissionsColumn = "user_id"
	// ReviewsTable is the table that holds the reviews relation/edge.
	ReviewsTable = "node_reviews"
	// ReviewsInverseTable is the table name for the NodeReview entity.
	// It exists in this package in order to avoid circular dependency with the "nodereview" package.
	ReviewsInverseTable = "node_reviews"
	// ReviewsColumn is the table column denoting the reviews relation/edge.
	ReviewsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldEmail,
	FieldName,
	FieldIsApproved,
	FieldIsAdmin,
	FieldStatus,
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultIsApproved holds the default value on creation for the "is_approved" field.
	DefaultIsApproved bool
	// DefaultIsAdmin holds the default value on creation for the "is_admin" field.
	DefaultIsAdmin bool
)

const DefaultStatus schema.UserStatusType = "ACTIVE"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s schema.UserStatusType) error {
	switch s {
	case "ACTIVE", "BANNED":
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIsApproved orders the results by the is_approved field.
func ByIsApproved(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsApproved, opts...).ToFunc()
}

// ByIsAdmin orders the results by the is_admin field.
func ByIsAdmin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAdmin, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPublisherPermissionsCount orders the results by publisher_permissions count.
func ByPublisherPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPublisherPermissionsStep(), opts...)
	}
}

// ByPublisherPermissions orders the results by publisher_permissions terms.
func ByPublisherPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPublisherPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReviewsCount orders the results by reviews count.
func ByReviewsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReviewsStep(), opts...)
	}
}

// ByReviews orders the results by reviews terms.
func ByReviews(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReviewsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPublisherPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PublisherPermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PublisherPermissionsTable, PublisherPermissionsColumn),
	)
}
func newReviewsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReviewsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
	)
}
