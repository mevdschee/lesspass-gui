package main

import "github.com/andlabs/ui"
import "github.com/mevdschee/lesspass.go/lesspass"

func main() {
	err := ui.Main(func() {
		site := ui.NewEntry()
		login := ui.NewEntry()
		masterPassword := ui.NewEntry()
		button := ui.NewButton("Generate")
		password := ui.NewEntry()
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Site"), false)
		box.Append(site, false)
		box.Append(ui.NewLabel("Login"), false)
		box.Append(login, false)
		box.Append(ui.NewLabel("Master password"), false)
		box.Append(masterPassword, false)
		box.Append(button, false)
		box.Append(password, false)
		window := ui.NewWindow("LessPass GUI", 200, 100, false)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			passwordProfile := lesspass.PasswordProfile{}
			password.SetText(lesspass.GeneratePassword(site.Text(), login.Text(), masterPassword.Text(), passwordProfile))
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
