package sqlx

import (
	"github.com/doug-martin/goqu/v9"
	// import the dialect
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

// PreparedDialect 包装DialectWrappper
// 这样在调用ToSQL() 方法的时候，返回的args 需要保留下来，作为参数使用
type PreparedDialect struct {
	rawDialect goqu.DialectWrapper
}

// PostgresBuilder postgres builder
func PostgresBuilder() PreparedDialect {
	dialect := goqu.Dialect("postgres")

	return PreparedDialect{
		rawDialect: dialect,
	}
}

// RawDialect 原始的Dialect
func (pd PreparedDialect) RawDialect() goqu.DialectWrapper {
	return pd.rawDialect
}

// From 转发给原始Dialect,包装Prepared(true)
func (pd PreparedDialect) From(table ...interface{}) *goqu.SelectDataset {
	return pd.RawDialect().From(table...).Prepared(true)
}

// Select 封装原始Select方法，包装Prepared(true)
func (pd PreparedDialect) Select(cols ...interface{}) *goqu.SelectDataset {
	return pd.RawDialect().Select(cols...).Prepared(true)
}

// Insert 包装原始Insert方法，增加Prepared(true)
func (pd PreparedDialect) Insert(table interface{}) *goqu.InsertDataset {
	return pd.RawDialect().Insert(table).Prepared(true)
}

// Update 包装原始Update方法，增加Prepared(true)
func (pd PreparedDialect) Update(table interface{}) *goqu.UpdateDataset {
	return pd.RawDialect().Update(table).Prepared(true)
}

// Delete 包装原始Delete方法，增加Prepared(true)
func (pd PreparedDialect) Delete(table interface{}) *goqu.DeleteDataset {
	return pd.RawDialect().Delete(table).Prepared(true)
}
