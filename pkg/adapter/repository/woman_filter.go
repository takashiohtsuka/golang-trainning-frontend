package repository

import (
	"strings"
)

// buildWomanFilterCondition は blood_type・age_range の絞り込み条件文字列と args を返す。
// woman_district / woman_prefecture / woman_region など女性一覧リポジトリで共通利用する。
func buildWomanFilterCondition(bloodTypes []string, ageRanges []string) (string, []any) {
	var condition string
	var args []any

	if len(bloodTypes) > 0 {
		condition += " AND w.blood_type IN (?" + strings.Repeat(",?", len(bloodTypes)-1) + ")"
		for _, bt := range bloodTypes {
			args = append(args, bt)
		}
	}

	if len(ageRanges) > 0 {
		conditions := make([]string, 0, len(ageRanges))
		for _, ar := range ageRanges {
			parts := strings.Split(ar, "-")
			if len(parts) == 2 {
				conditions = append(conditions, "(w.age BETWEEN ? AND ?)")
				args = append(args, parts[0], parts[1])
			}
		}
		if len(conditions) > 0 {
			condition += " AND (" + strings.Join(conditions, " OR ") + ")"
		}
	}

	return condition, args
}
