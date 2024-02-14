package models

type Department struct {
	DepartmentID   int    `gorm:"primaryKey" json:"id"`
	DepartmentName string `gorm:"type:varchar(50)" json:"deptName"`
}
