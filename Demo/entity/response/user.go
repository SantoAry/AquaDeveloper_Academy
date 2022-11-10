package response

// Role
type RoleBaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateRoleResponse struct {
	Role_ID   int64  `json:"role_id" gorm:"column:role_id;type:bigint;primaryKey;autoIncrement"`
	Role_Type string `json:"role_type"`
}

type GetRoleResponse struct {
	Role_ID   int64  `json:"role_id" gorm:"column:role_id;type:bigint;primaryKey;autoIncrement"`
	Role_Type string `json:"role_type"`
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
	Cart_ID  int64
}

type UpdateUserResponse struct {
	User_ID  int64  `json:"user_id" gorm:"column:user_id;type:bigint;primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	NIK      string `json:"NIK"`
	Role_Ref int64  `json:"role_ref"`
	Cart_ID  int64
}

type GetUserResponse struct {
	User_ID  int64  `json:"user_id" gorm:"column:user_id;type:bigint;primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	NIK      string `json:"NIK"`
	Role_Ref int64  `json:"role_ref"`
	Cart_ID  int64
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
	Cart_ID  int64
}
