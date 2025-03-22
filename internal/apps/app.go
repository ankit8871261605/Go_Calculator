package apps

type User struct {
	USERNAME	string
	PASSWORD	string
	EMAIL 		string
	ROLE		string
}

func GetUserDetails () *User {
	return &User{
		USERNAME: "Ankit",
		PASSWORD: "root",
		EMAIL: "ankit@go.dev",
		ROLE: "dev",
	}
}