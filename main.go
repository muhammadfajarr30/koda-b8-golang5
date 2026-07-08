package main

import (
	"fmt"
	"os"
)

type User struct {
	first, last, email, password string
}

var dataUser []User

func (u User) fullName() string {
	return u.first + " " + u.last
}

func register() {
	fmt.Print("\033[H\033[2J")
	var user User

	fmt.Println("====== Register ======")

	fmt.Print("First Name: ")
	fmt.Scan(&user.first)
	fmt.Print("Last Name: ")
	fmt.Scan(&user.last)
	fmt.Print("Email: ")
	fmt.Scan(&user.email)
	fmt.Print("Password: ")
	fmt.Scan(&user.password)
	fmt.Println("Register successed!")
	dataUser = append(dataUser, user)

}

func login() {
	for {

		var email, password string
		fmt.Print("enter your email: ")
		fmt.Scan(&email)
		fmt.Print("enter your password: ")
		fmt.Scan(&password)
		for _, user := range dataUser {
			if user.email == email && user.password == password {
				fmt.Println("login success")
				loginSuccess(user)
				return
			}
		}
		fmt.Println("wrong email or password!")
	}

}

func listAllUsers() {
	fmt.Println("=== List of Users ===")

	if len(dataUser) == 0 {
		fmt.Println("Belum ada user yang terdaftar.")
		return
	}

	for i, user := range dataUser {
		fmt.Printf("%d. %s\n", i+1, user.fullName())
		fmt.Printf("   Email: %s\n", user.email)
	}
}

func loginSuccess(user User) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()

			var option int

			fmt.Println("=== Welcome To System ===")
			fmt.Printf("Welcome, %s\n", user.fullName())

			fmt.Println("1. List All Users")
			fmt.Println("0. Logout")
			fmt.Print("Select menu: ")
			fmt.Scan(&option)

			switch option {
			case 1:
				listAllUsers()
			case 0:
				logout()
				return
			default:
				panic("menu tidak tersedia")
			}
		}()
	}
}

func logout() {
	fmt.Println("Logout berhasil!")
}

func main() {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Error: %v\n\n", r)
				}
			}()

			var choose int

			fmt.Println("=== Welcome To System ===")
			fmt.Println("1. Register")
			fmt.Println("2. Login")
			fmt.Println("0. Exit")
			fmt.Print("Select your menu: ")

			fmt.Scan(&choose)

			switch choose {
			case 1:
				register()
			case 2:
				login()
			case 0:
				fmt.Println("Thank you!")
				os.Exit(0)
			default:
				panic(fmt.Sprintf("menu %d tidak tersedia", choose))
			}
		}()
	}
}
