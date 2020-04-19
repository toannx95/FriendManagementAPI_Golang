package repository

type UserRepository interface {
	GetAllUsers() []string
	CreateUser(email string) bool
	ExistsByEmail(email string) bool
	FindUserIdByEmail(email string) int64
	FindByIds(ids []int64) []string
}