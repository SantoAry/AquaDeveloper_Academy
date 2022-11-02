package response

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateUserResponse struct {
	ID      int64  `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	NIK     string `json:"NIK"`
}

type UpdateUserResponse struct {
	ID      int64  `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	NIK     string `json:"NIK"`
}

type GetUserResponse struct {
	ID      int64  `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	NIK     string `json:"NIK"`
}

type DeleteUserResponse struct {
	ID      int64  `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	NIK     string `json:"NIK"`
}
