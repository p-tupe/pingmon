package setup

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/EMPAT94/pingmon/internal/app/config"
)

var sample_config = config.Config{
	Sites:    []config.Site{{URL: "https://www.example.com"}},
	Interval: 30,
}

func Setup() {
	fmt.Println("Initializing pingmon...")

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error finding user config directory:", err)
		return
	}

	pingmonConfigDir := filepath.Join(userConfigDir, "pingmon")

	err = os.MkdirAll(pingmonConfigDir, os.ModeDir.Perm())
	if err != nil {
		fmt.Println("Error creating pingmon config directory:", err)
		return
	}

	configFilePath := filepath.Join(pingmonConfigDir, "/config.json")
	cf, err := os.Create(configFilePath)
	defer cf.Close()
	if err != nil {
		fmt.Println("Error creating pingmon config file:", err)
		return
	}

	readVal("Add a website's full url to monitor (default https://www.example.com): ", &sample_config.Sites[0].URL)
	readVal("Site check interval (default 30 min): ", &sample_config.Interval)
	readVal("Add your slack webhook (eg https://hooks.slack.com/your/custom/id): ", &sample_config.SlackWebhook)

	fmt.Print("Add email config now? y/n: ")
	var yes string
	_, err = fmt.Scan(&yes)
	if err != nil {
		fmt.Println("Error reading user input: ", err)
		return
	}
	if yes == "yes" || yes == "y" || yes == "Y" {
		readVal("Host (smtp.mail.host): ", &sample_config.Mailer.Host)
		readVal("Port (587): ", &sample_config.Mailer.Port)
		readVal("Username: ", &sample_config.Mailer.Username)
		readVal("Password: ", &sample_config.Mailer.Password)
		readVal("From Id: ", &sample_config.Mailer.From)
		readVal("To Id: ", &sample_config.EmailTo[0])
	}

	err = json.NewEncoder(cf).Encode(sample_config)
	if err != nil {
		fmt.Println("Error writing pingmon config file:", err)
		return
	}

	fmt.Println("Successfully created config at", configFilePath)

	fmt.Print(os.ExpandEnv("Open config file in $EDITOR? y/n: "))
	_, err = fmt.Scan(&yes)
	if err != nil {
		fmt.Println("Error reading user input: ", err)
	}

	if yes == "yes" || yes == "y" || yes == "Y" {
		editor, ok := os.LookupEnv("EDITOR")
		if !ok {
			fmt.Println("No editor found! Please open the file on above path and edit manually.")
		} else {
			cmd := exec.Command(editor, configFilePath)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				fmt.Println("Unexpected error while opening file for editing: ", err)
			}
		}
	}
}

func readVal(msg string, v any) {
	fmt.Print(msg)
	_, err := fmt.Scanln(v)
	if err != nil && err.Error() != "unexpected newline" {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
}
