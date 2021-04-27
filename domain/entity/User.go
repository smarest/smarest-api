package entity

import "github.com/smarest/smarest-common/domain/value"

type User struct {
	UserName   string          `db:"user_name" json:"userName"`
	Role       string          `db:"role" json:"role"`
	Password   string          `db:"password" json:"-"`
	Name       string          `db:"name" json:"name"`
	Available  bool            `db:"available" json:"available"`
	SalaryType string          `db:"salary_type" json:"salaryType"`
	JoinedDate value.DateTime  `db:"joined_date" json:"joinedDate"`
	LeftDate   *value.DateTime `db:"left_date" json:"leftDate"`
}
