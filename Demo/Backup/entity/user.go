package entity

import "e-commerce/entity/sub_entity"

type Role struct {
	Role_ID   int64  `json:"role_id" gorm:"column:role_id;type:bigint;primaryKey;autoIncrement"`
	Role_Type string `json:"role_type"`
	User      []User `gorm:"constraint;foreignKey:Role_Ref;references:Role_ID"`
}

type User struct {
	User_ID  int64           `json:"user_id" gorm:"column:user_id;type:bigint;primaryKey;autoIncrement"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Phone    string          `json:"phone"`
	Address  string          `json:"address"`
	NIK      string          `json:"NIK"`
	Role_Ref int64           `json:"role_ref"`
	Cart     sub_entity.Cart `gorm:"constraint;foreignKey:User_Ref;references:User_ID"`
}
