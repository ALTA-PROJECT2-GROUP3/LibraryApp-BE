package users

type Core struct {
	Id       uint
	Name     string `validate:"required"`
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Address  string `validate:"required"`
	Phone    string `validate:"required"`
	Password string `validate:"required,min=3"`
	Pictures string
}

type UserService interface {
	Register(newUser Core) error
	Login(username string, password string) (string, Core, error)
	Update(id int, updateUser Core) error
	UserByID(id int) (Core, error)
}

type UserData interface {
	Register(newUser Core) error
	Login(username string) (Core, error)
	Update(id int, updateUser Core) error
	UserByID(id int) (Core, error)
}
