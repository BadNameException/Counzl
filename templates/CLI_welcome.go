package app

import (
	"./modules/counzl_chat"
	"./modules/users"
	"gopkg.in/abiosoft/ishell.v2"
	"utilities/cmd"
)

// Denne filen skal ha muligheten til å launche enhver modul.
func Welcome(online bool) {
	shell := ishell.New()
	shell.Print("\n" + cmd.ChangeColor("Velkommen til vårt Open Source prosjekt, Counzl!", "brightBlue") + "                                                  ")
	shell.Print("\n" + "Hvis dette er første gang du bruker Counzl; skriv '" + cmd.ChangeColor("hjelp", "green") + "'!\nDette vil gi deg en liste over kommandoer :)")
	shell.Println()
	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "loading",
		Help: "Test graphics for 'loading'",
		Func: func(c *ishell.Context) {
			cmd.Loading()
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "chat",
		Help: "Gå til chatte suiten",
		Func: func(c *ishell.Context) {
			switch online {
			case true:
				counzl_chat.Chatroom()
			case false:
				shell.Print("\n" + "Du kan ikke bruke 'chatroom' når brukeren din ikke er registrert på server")
			}

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "printusers",
		Help: "Print ut alle lokale brukere",
		Func: func(c *ishell.Context) {
			users.PrintLocalUsers()
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "Logg ut",
		Func: func(c *ishell.Context) {
			AccessCounzl()
		},
	})

	shell.Run()
}
