/*
The API in this package is not stable and may change without any notice.
*/
package orm

import (
	"context"
	"io"

	"github.com/go-pg/pg/v10/types"
)

// ColumnScanner is used to scan column values.
type ColumnScanner interface {
	// Scan assigns a column value from a row.
	//
	// An error should be returned if the value can not be stored
	// without loss of information.
	ScanColumn(col types.ColumnInfo, rd types.Reader, n int) error
}

type QueryAppender interface {
	AppendQuery(fmter QueryFormatter, b []byte) ([]byte, error)
}

type TemplateAppender interface {
	AppendTemplate(b []byte) ([]byte, error)
}

type QueryCommand interface {
	QueryAppender
	TemplateAppender
	String() string
	Operation() QueryOp
	Clone() QueryCommand
	Query() *Query
}

// DB is a common interface for pg.DB and pg.Tx types.
type DB interface {
	Model(model ...interface{}) *Query
	ModelContext(c context.Context, model ...interface{}) *Query

	Exec(query interface{}, params ...interface{}) (Result, error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (Result, error)
	ExecOne(query interface{}, params ...interface{}) (Result, error)
	ExecOneContext(c context.Context, query interface{}, params ...interface{}) (Result, error)
	Query(model, query interface{}, params ...interface{}) (Result, error)
	QueryContext(c context.Context, model, query interface{}, params ...interface{}) (Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (Result, error)
	QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (Result, error)

	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (Result, error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (Result, error)

	Context() context.Context
	Formatter() QueryFormatter
}

type QueryInterface interface {
	New() *Query
	Clone() *Query
	Context(c context.Context) *Query
	DB(db DB) *Query
	Model(model ...interface{}) *Query
	TableModel() TableModel
	Deleted() *Query
	AllWithDeleted() *Query
	With(name string, subq *Query) *Query
	WithInsert(name string, subq *Query) *Query
	WithUpdate(name string, subq *Query) *Query
	WithDelete(name string, subq *Query) *Query
	WrapWith(name string) *Query
	Table(tables ...string) *Query
	TableExpr(expr string, params ...interface{}) *Query
	Distinct() *Query
	DistinctOn(expr string, params ...interface{}) *Query
	Column(columns ...string) *Query
	ColumnExpr(expr string, params ...interface{}) *Query
	ExcludeColumn(columns ...string) *Query
	Relation(name string, apply ...func(*Query) (*Query, error)) *Query
	Set(set string, params ...interface{}) *Query
	Value(column string, value string, params ...interface{}) *Query
	Where(condition string, params ...interface{}) *Query
	WhereOr(condition string, params ...interface{}) *Query
	WhereGroup(fn func(*Query) (*Query, error)) *Query
	WhereNotGroup(fn func(*Query) (*Query, error)) *Query
	WhereOrGroup(fn func(*Query) (*Query, error)) *Query
	WhereOrNotGroup(fn func(*Query) (*Query, error)) *Query
	WhereIn(where string, slice interface{}) *Query
	WhereInMulti(where string, values ...interface{}) *Query
	WherePK() *Query
	Join(join string, params ...interface{}) *Query
	JoinOn(condition string, params ...interface{}) *Query
	JoinOnOr(condition string, params ...interface{}) *Query
	Group(columns ...string) *Query
	GroupExpr(group string, params ...interface{}) *Query
	Having(having string, params ...interface{}) *Query
	Union(other *Query) *Query
	UnionAll(other *Query) *Query
	Intersect(other *Query) *Query
	IntersectAll(other *Query) *Query
	Except(other *Query) *Query
	ExceptAll(other *Query) *Query
	Order(orders ...string) *Query
	OrderExpr(order string, params ...interface{}) *Query
	Limit(n int) *Query
	Offset(n int) *Query
	OnConflict(s string, params ...interface{}) *Query
	Returning(s string, params ...interface{}) *Query
	For(s string, params ...interface{}) *Query
	Apply(fn func(*Query) (*Query, error)) *Query
	Count() (int, error)
	First() error
	Last() error
	Select(values ...interface{}) error
	SelectAndCount(values ...interface{}) (count int, firstErr error)
	SelectAndCountEstimate(threshold int, values ...interface{}) (count int, firstErr error)
	ForEach(fn interface{}) error
	Insert(values ...interface{}) (Result, error)
	SelectOrInsert(values ...interface{}) (inserted bool, _ error)
	Update(scan ...interface{}) (Result, error)
	UpdateNotZero(scan ...interface{}) (Result, error)
	Delete(values ...interface{}) (Result, error)
	ForceDelete(values ...interface{}) (Result, error)
	CreateTable(opt *CreateTableOptions) error
	DropTable(opt *DropTableOptions) error
	CreateComposite(opt *CreateCompositeOptions) error
	DropComposite(opt *DropCompositeOptions) error
	Exec(query interface{}, params ...interface{}) (Result, error)
	ExecOne(query interface{}, params ...interface{}) (Result, error)
	Query(model, query interface{}, params ...interface{}) (Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (Result, error)
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (Result, error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (Result, error)
	AppendQuery(fmter QueryFormatter, b []byte) ([]byte, error)
	Exists() (bool, error)
}
