package ast

type SQL struct {
	Statements []SQLStatement
}

type SQLStatement struct {
	Explain     bool
	AlterTable  *AlterTableStatement
	Create      *CreateStatement
	Transaction *TransactionStatement
}

type AlterTableStatement struct {
	RenameTable  bool
	RenameColumn bool
	AddColumn    bool
	Source       *ColumnInfo
	Destination  *ColumnInfo
	Definition   *ColumnDefinition
}

type CreateStatement struct {
	Index       *IndexInfo
	Table       *TableInfo
	IfNotExists bool
}

type TransactionStatement struct {
	BeginDeferred  bool // Default
	BeginImmediate bool
	BeginExclusive bool
	Commit         bool
	End            bool
	Rollback       bool
	SavePoint      string
}

type IndexInfo struct {
	Name           string
	Unique         bool
	Schema         string
	Database       string
	Table          string
	IndexedColumns []IndexedColumn
	Condition      *Expression
}

type IndexedColumn struct {
	Name       string
	Expression *Expression
	Collation  string
	Asc        bool
	Desc       bool
}

type TableInfo struct {
	Temporary  bool
	Schema     string
	Database   string
	Table      string
	Columns    []ColumnInfo
	Constraint []TableConstraint
	RowID      bool
}

type ColumnInfo struct {
	Name       string
	Definition *ColumnDefinition
}

type ColumnDefinition struct {
	Type       ColumnType
	Constraint ColumnConstraint
}

type ColumnType struct {
	Integer         bool
	UnsignedInteger bool
	FloatingPoint   bool
	String          bool
	DateTime        bool
	Date            bool
	Time            bool
	Length          int
	Precision       int
	Scale           int
}

type TableConstraint struct {
	PrimaryKey    bool
	UniqueKey     bool
	IndexedColumn []IndexedColumn
}

type ColumnConstraint struct {
	Name       string
	PrimaryKey *PrimaryKey
	NotNull    bool
	Unique     bool
	Check      *Expression
	Default    *DefaultValue
	Collate    string
	ForeginKey *ForeignKeyClause
	Generated  *Generated
	Conflict   *ConflictClause
}

type PrimaryKey struct {
	Asc           bool
	Desc          bool
	AutoIncrement bool
}

type ConflictClause struct {
	RollBack      bool
	Abort         bool // Default
	Fail          bool
	Ignore        bool
	Replace       bool
	AutoIncrement bool
}

type Expression struct {
}

type DefaultValue struct {
	Integer          bool
	IntegerValue     int
	Float            bool
	FloatValue       float64
	String           bool
	StringValue      string
	Null             bool
	True             bool
	False            bool
	CurrentTime      bool
	CurrentDate      bool
	CurrentTimestamp bool
	Expression       *Expression
}

type ForeignKeyClause struct {
	ForeignColumn []ColumnInfo
	Delete        bool
	Update        bool
	Match         string
	SetNull       bool
	SetDefault    bool
	Cascade       bool
	Restrict      bool
	NoAction      bool
	Deferrable    bool
	NotDeferrable bool
	Deferred      bool
	Immediate     bool
}

type Generated struct {
	Expression *Expression
	Stored     bool
	Virtual    bool
}
