package repository

import (
	"strings"

	"golang-trainning-frontend/pkg/usecase/query"
)

func buildWhereClause(conditions []query.Condition) (string, []any) {
	return buildWhereClauseWithPrefix(conditions, "")
}

func buildWhereClauseWithPrefix(conditions []query.Condition, tablePrefix string) (string, []any) {
	var clauses []string
	var args []any

	col := func(name string) string {
		if tablePrefix == "" {
			return name
		}
		return tablePrefix + "." + name
	}

	// KindWhereBetweenOr は同一カラムの BETWEEN 条件を OR でまとめるため別に収集する
	betweenOrClauses := make(map[string][]string)
	betweenOrArgs := make(map[string][]any)
	betweenOrOrder := make([]string, 0)

	for _, c := range conditions {
		switch c.Kind {
		case query.KindWhere:
			clauses = append(clauses, col(c.Column)+" = ?")
			args = append(args, c.Value)
		case query.KindWhereIn:
			clauses = append(clauses, col(c.Column)+" IN ?")
			args = append(args, c.Value)
		case query.KindWhereBetween:
			clauses = append(clauses, col(c.Column)+" BETWEEN ? AND ?")
			args = append(args, c.From, c.To)
		case query.KindWhereNotIn:
			clauses = append(clauses, col(c.Column)+" NOT IN ?")
			args = append(args, c.Value)
		case query.KindWhereBetweenOr:
			colName := col(c.Column)
			if _, exists := betweenOrClauses[colName]; !exists {
				betweenOrOrder = append(betweenOrOrder, colName)
			}
			betweenOrClauses[colName] = append(betweenOrClauses[colName], "("+colName+" BETWEEN ? AND ?)")
			betweenOrArgs[colName] = append(betweenOrArgs[colName], c.From, c.To)
		}
	}

	// OR グループを AND 条件として追加
	for _, colName := range betweenOrOrder {
		clauses = append(clauses, "("+strings.Join(betweenOrClauses[colName], " OR ")+")")
		args = append(args, betweenOrArgs[colName]...)
	}

	if len(clauses) == 0 {
		return "", nil
	}
	return " AND " + strings.Join(clauses, " AND "), args
}
