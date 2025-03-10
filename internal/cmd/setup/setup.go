package setup

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

const sample_config = `;; global settings (required)

interval=30 ; minutes
slack_webhook=
email_to=

[Mailer]
server=
port=
username=
password=
from=


;; site specific settings (optional)

[https://www.example.com]
interval=10

[https://www.google.com]
email_to=`

func Init() {
	fmt.Println("Initializing pingmon...")

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error finding user config directory: ", err)
		return
	}

	pingmonConfigDir := path.Join(userConfigDir, "pingmon")

	err = os.MkdirAll(pingmonConfigDir, os.ModeDir.Perm())
	if err != nil {
		fmt.Println("Error creating pingmon config directory: ", err)
		return
	}

	configFilePath := path.Join(pingmonConfigDir, "/config.ini")
	cf, err := os.Create(configFilePath)
	defer cf.Close()
	if err != nil {
		fmt.Println("Error creating pingmon config file: ", err)
		return
	}

	cf.Write([]byte(sample_config))

	fmt.Println("Successfully created config at ", configFilePath)

	fmt.Print("Open config for editing? y/n: ")
	var yes string
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
