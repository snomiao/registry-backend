// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CiWorkflowResultsColumns holds the columns for the "ci_workflow_results" table.
	CiWorkflowResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "operating_system", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "gpu_type", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "pytorch_version", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "workflow_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "run_id", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "status", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "start_time", Type: field.TypeInt64, Nullable: true},
		{Name: "end_time", Type: field.TypeInt64, Nullable: true},
		{Name: "ci_workflow_result_storage_file", Type: field.TypeUUID, Nullable: true},
		{Name: "git_commit_results", Type: field.TypeUUID, Nullable: true},
	}
	// CiWorkflowResultsTable holds the schema information for the "ci_workflow_results" table.
	CiWorkflowResultsTable = &schema.Table{
		Name:       "ci_workflow_results",
		Columns:    CiWorkflowResultsColumns,
		PrimaryKey: []*schema.Column{CiWorkflowResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ci_workflow_results_storage_files_storage_file",
				Columns:    []*schema.Column{CiWorkflowResultsColumns[11]},
				RefColumns: []*schema.Column{StorageFilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "ci_workflow_results_git_commits_results",
				Columns:    []*schema.Column{CiWorkflowResultsColumns[12]},
				RefColumns: []*schema.Column{GitCommitsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GitCommitsColumns holds the columns for the "git_commits" table.
	GitCommitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "commit_hash", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "branch_name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "repo_name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "commit_message", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "commit_timestamp", Type: field.TypeTime},
		{Name: "author", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "timestamp", Type: field.TypeTime, Nullable: true},
	}
	// GitCommitsTable holds the schema information for the "git_commits" table.
	GitCommitsTable = &schema.Table{
		Name:       "git_commits",
		Columns:    GitCommitsColumns,
		PrimaryKey: []*schema.Column{GitCommitsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "gitcommit_repo_name_commit_hash",
				Unique:  true,
				Columns: []*schema.Column{GitCommitsColumns[5], GitCommitsColumns[3]},
			},
		},
	}
	// NodesColumns holds the columns for the "nodes" table.
	NodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "description", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "category", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "author", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "license", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "repository_url", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "icon_url", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "tags", Type: field.TypeJSON, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "total_install", Type: field.TypeInt64, Default: 0},
		{Name: "total_star", Type: field.TypeInt64, Default: 0},
		{Name: "total_review", Type: field.TypeInt64, Default: 0},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "banned", "deleted"}, Default: "active"},
		{Name: "status_detail", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "publisher_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
	}
	// NodesTable holds the schema information for the "nodes" table.
	NodesTable = &schema.Table{
		Name:       "nodes",
		Columns:    NodesColumns,
		PrimaryKey: []*schema.Column{NodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "nodes_publishers_nodes",
				Columns:    []*schema.Column{NodesColumns[16]},
				RefColumns: []*schema.Column{PublishersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// NodeReviewsColumns holds the columns for the "node_reviews" table.
	NodeReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "star", Type: field.TypeInt, Default: 0},
		{Name: "node_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "user_id", Type: field.TypeString},
	}
	// NodeReviewsTable holds the schema information for the "node_reviews" table.
	NodeReviewsTable = &schema.Table{
		Name:       "node_reviews",
		Columns:    NodeReviewsColumns,
		PrimaryKey: []*schema.Column{NodeReviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_reviews_nodes_reviews",
				Columns:    []*schema.Column{NodeReviewsColumns[2]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "node_reviews_users_reviews",
				Columns:    []*schema.Column{NodeReviewsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "nodereview_node_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{NodeReviewsColumns[2], NodeReviewsColumns[3]},
			},
		},
	}
	// NodeVersionsColumns holds the columns for the "node_versions" table.
	NodeVersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "version", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "changelog", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "pip_dependencies", Type: field.TypeJSON, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "deprecated", Type: field.TypeBool, Default: false},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "banned", "deleted", "pending", "flagged"}, Default: "pending"},
		{Name: "status_reason", Type: field.TypeString, Default: "", SchemaType: map[string]string{"postgres": "text"}},
		{Name: "node_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "node_version_storage_file", Type: field.TypeUUID, Nullable: true},
	}
	// NodeVersionsTable holds the schema information for the "node_versions" table.
	NodeVersionsTable = &schema.Table{
		Name:       "node_versions",
		Columns:    NodeVersionsColumns,
		PrimaryKey: []*schema.Column{NodeVersionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_versions_nodes_versions",
				Columns:    []*schema.Column{NodeVersionsColumns[9]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "node_versions_storage_files_storage_file",
				Columns:    []*schema.Column{NodeVersionsColumns[10]},
				RefColumns: []*schema.Column{StorageFilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "nodeversion_node_id_version",
				Unique:  true,
				Columns: []*schema.Column{NodeVersionsColumns[9], NodeVersionsColumns[3]},
			},
		},
	}
	// PersonalAccessTokensColumns holds the columns for the "personal_access_tokens" table.
	PersonalAccessTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "description", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "token", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "publisher_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
	}
	// PersonalAccessTokensTable holds the schema information for the "personal_access_tokens" table.
	PersonalAccessTokensTable = &schema.Table{
		Name:       "personal_access_tokens",
		Columns:    PersonalAccessTokensColumns,
		PrimaryKey: []*schema.Column{PersonalAccessTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "personal_access_tokens_publishers_personal_access_tokens",
				Columns:    []*schema.Column{PersonalAccessTokensColumns[6]},
				RefColumns: []*schema.Column{PublishersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "personalaccesstoken_token",
				Unique:  true,
				Columns: []*schema.Column{PersonalAccessTokensColumns[5]},
			},
		},
	}
	// PublishersColumns holds the columns for the "publishers" table.
	PublishersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "description", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "website", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "support_email", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "source_code_repo", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "logo_url", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "BANNED"}, Default: "ACTIVE"},
	}
	// PublishersTable holds the schema information for the "publishers" table.
	PublishersTable = &schema.Table{
		Name:       "publishers",
		Columns:    PublishersColumns,
		PrimaryKey: []*schema.Column{PublishersColumns[0]},
	}
	// PublisherPermissionsColumns holds the columns for the "publisher_permissions" table.
	PublisherPermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "permission", Type: field.TypeEnum, Enums: []string{"owner", "member"}},
		{Name: "publisher_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "user_id", Type: field.TypeString},
	}
	// PublisherPermissionsTable holds the schema information for the "publisher_permissions" table.
	PublisherPermissionsTable = &schema.Table{
		Name:       "publisher_permissions",
		Columns:    PublisherPermissionsColumns,
		PrimaryKey: []*schema.Column{PublisherPermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "publisher_permissions_publishers_publisher_permissions",
				Columns:    []*schema.Column{PublisherPermissionsColumns[2]},
				RefColumns: []*schema.Column{PublishersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "publisher_permissions_users_publisher_permissions",
				Columns:    []*schema.Column{PublisherPermissionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// StorageFilesColumns holds the columns for the "storage_files" table.
	StorageFilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "bucket_name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "object_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "file_path", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "file_type", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "file_url", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
	}
	// StorageFilesTable holds the schema information for the "storage_files" table.
	StorageFilesTable = &schema.Table{
		Name:       "storage_files",
		Columns:    StorageFilesColumns,
		PrimaryKey: []*schema.Column{StorageFilesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "is_approved", Type: field.TypeBool, Default: false},
		{Name: "is_admin", Type: field.TypeBool, Default: false},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "BANNED"}, Default: "ACTIVE"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CiWorkflowResultsTable,
		GitCommitsTable,
		NodesTable,
		NodeReviewsTable,
		NodeVersionsTable,
		PersonalAccessTokensTable,
		PublishersTable,
		PublisherPermissionsTable,
		StorageFilesTable,
		UsersTable,
	}
)

func init() {
	CiWorkflowResultsTable.ForeignKeys[0].RefTable = StorageFilesTable
	CiWorkflowResultsTable.ForeignKeys[1].RefTable = GitCommitsTable
	NodesTable.ForeignKeys[0].RefTable = PublishersTable
	NodeReviewsTable.ForeignKeys[0].RefTable = NodesTable
	NodeReviewsTable.ForeignKeys[1].RefTable = UsersTable
	NodeVersionsTable.ForeignKeys[0].RefTable = NodesTable
	NodeVersionsTable.ForeignKeys[1].RefTable = StorageFilesTable
	PersonalAccessTokensTable.ForeignKeys[0].RefTable = PublishersTable
	PublisherPermissionsTable.ForeignKeys[0].RefTable = PublishersTable
	PublisherPermissionsTable.ForeignKeys[1].RefTable = UsersTable
}
