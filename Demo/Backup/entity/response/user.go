package response

type UserBaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateUserResponse struct {
	User_ID int64  `json:"user_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	NIK     string `json:"NIK"`
	Role_ID int64  `json:"role_id"` //Foreign key to Role table
}

type UpdateUserResponse struct {
	User_ID int64  `json:"user_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	NIK     string `json:"NIK"`
	Role_ID int64  `json:"role_id"` //Foreign key to Role table
}

type GetUserResponse struct {
	User_ID int64  `json:"user_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	NIK     string `json:"NIK"`
	Role_ID int64  `json:"role_id"` //Foreign key to Role table
}

// Need to be reworked, delete by ID
type DeleteUserResponse struct {
	User_ID int64  `json:"user_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	NIK     string `json:"NIK"`
	Role_ID int64  `json:"role_id"` //Foreign key to Role table
}
