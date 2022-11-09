package entity

type Role struct {
	Role_ID   int64  `json:"role_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Role_Type string `json:"role_type"`
}

type User struct {
	User_ID int64  `json:"user_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	NIK     string `json:"NIK"`
	Role_ID int64  `json:"role_id"` //Foreign key to Role table
}
