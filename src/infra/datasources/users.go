package datasource

type User struct {
	ID       int
	Email    string
	Password string
	Role     string
}

func MakeUsersDataSource() []User {
	return []User{
		{
			ID:       1,
			Email:    "user1@example.com",
			Password: "$argon2i$v=19$m=16,t=2,p=1$Z3ZJaFh4MnZsVmRKWkkwQg$AvCX7g6TUgGshY0RGexGSw",
			Role:     "admin",
		},
		{
			ID:       2,
			Email:    "user2@example.com",
			Password: "$argon2i$v=19$m=16,t=2,p=1$Z3ZJaFh4MnZsVmRKWkkwQg$AvCX7g6TUgGshY0RGexGSw",
			Role:     "user",
		},
	}
}
