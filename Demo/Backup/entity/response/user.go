package response

import "e-commerce/entity"

// Role
type RoleBaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateRoleResponse struct {
	Role_ID   int64         `json:"role_id" gorm:"column:role_id;type:bigint;primaryKey;autoIncrement"`
	Role_Type string        `json:"role_type"`
	User      []entity.User `gorm:"constraint;foreignKey:Role_Ref;references:Role_ID"`
}

type GetRoleResponse struct {
	Role_ID   int64         `json:"role_id" gorm:"column:role_id;type:bigint;primaryKey;autoIncrement"`
	Role_Type string        `json:"role_type"`
	User      []entity.User `gorm:"constraint;foreignKey:Role_Ref;references:Role_ID"`
}

// User
type UserBaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateUserResponse struct {
	User_ID  int64  `json:"user_id" gorm:"column:user_id;type:bigint;primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	NIK      string `json:"NIK"`
	Role_Ref int64  `json:"role_ref"`
}

type UpdateUserResponse struct {
	User_ID  int64  `json:"user_id" gorm:"column:user_id;type:bigint;primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	NIK      string `json:"NIK"`
	Role_Ref int64  `json:"role_ref"`
}

type GetUserResponse struct {
	User_ID  int64  `json:"user_id" gorm:"column:user_id;type:bigint;primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	NIK      string `json:"NIK"`
	Role_Ref int64  `json:"role_ref"`
}

// Need to be reworked, delete by ID
type DeleteUserResponse struct {
	User_ID  int64  `json:"user_id" gorm:"column:user_id;type:bigint;primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	NIK      string `json:"NIK"`
	Role_Ref int64  `json:"role_ref"`
}
