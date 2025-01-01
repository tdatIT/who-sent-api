package pagable

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
)

const (
	StringType = 1
	NumberType = 2
)

type QueryCondition struct {
	Condition string
	Value     interface{}
}

func (q *Query) ORMConditions(conds map[string]int) []*QueryCondition {
	var conditions []*QueryCondition

	for _, filter := range q.ExpressionFilters {
		filter.Field = CamelToSnake(filter.Field)
		// Check if the field is in the allowed fields
		if _, ok := conds[filter.Field]; !ok {
			continue
		}

		var value interface{}

		switch conds[filter.Field] {
		case StringType:
			value = fmt.Sprintf("%s", filter.Value)
		case NumberType:
			value, _ = strconv.Atoi(fmt.Sprintf("%v", filter.Value))
		}

		if filter.Operation == Search {
			value = "%" + fmt.Sprint(value) + "%"
		}

		condition := q.ParseCondition(filter)
		queryCdt := &QueryCondition{
			Condition: condition,
			Value:     value,
		}
		conditions = append(conditions, queryCdt)
	}

	return conditions
}

func (q *Query) ParseCondition(ft Filter) string {
	condition := ""

	switch ft.Operation {
	case Equal:
		condition = ft.Field + " = ?"
	case NotEqual:
		condition = ft.Field + " <> ?"
	case LT:
		condition = ft.Field + " < ? "
	case LTE:
		condition = ft.Field + " <= ?"
	case GT:
		condition = ft.Field + " > ?"
	case GTE:
		condition = ft.Field + " >=  ?"
	case In:
		condition = ft.Field + " IN ?"
	case NotIn:
		condition = ft.Field + " NOT IN ?"
	case NotContains:
		condition = ft.Field + " NOT LIKE ?"
	case IsNotNull:
		condition = ft.Field + " IS NOT NULL ?"
	case Search:
		condition = ft.Field + " LIKE ?"
	}
	return condition
}

func (q *Query) ParseQueryParams() (map[string]string, error) {
	conditions := map[string]string{}
	for _, filter := range q.ExpressionFilters {
		if filter.Operation != Equal {
			return conditions, errors.New("only $eq filtering is supported")
		}
		conditions[filter.Field] = fmt.Sprint(filter.Value)
	}
	return conditions, nil
}

func ApplyWhereCondition(db *gorm.DB, conditions []*QueryCondition) {
	for _, i := range conditions {
		if conds := db.Statement.BuildCondition(i.Condition, i.Value); len(conds) > 0 {
			db.Statement.AddClause(clause.Where{Exprs: conds})
		}
	}
}
