package setup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/p-tupe/pingmon/internal/app/config"
)

var sample_config = config.Config{
	Sites:    []config.Site{{URL: "https://www.example.com"}},
	Interval: 30,
	Mailer: &config.Mailer{
		Port: 587,
	},
	EmailTo: make([]string, 1),
}

func Setup() {
	fmt.Println("Setting up pingmon...")

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
	//
	// _, err = os.Stat(configFilePath)
	// if err != nil && !errors.Is(err, fs.ErrNotExist) {
	// 	fmt.Println("Error accessing pingmon config file:", err)
	// 	return
	// } else {
	// 	fmt.Printf("File already exists at %s, overwrite? y/n: ", pingmonConfigDir)
	// 	yes := ""
	// 	_, err = fmt.Scan(&yes)
	// 	if err != nil {
	// 		fmt.Println("Error reading user input: ", err)
	// 		return
	// 	}
	// 	if yes != "yes" && yes != "y" && yes != "Y" {
	// 		return
	// 	}
	// }
	//
	// cf, err := os.Create(configFilePath)
	// if err != nil {
	// 	fmt.Println("Error creating pingmon config file:", err)
	// 	return
	// }
	// defer cf.Close()
	//
	// readVal("Add a website's full url to monitor (default https://www.example.com): ", &sample_config.Sites[0].URL)
	// readVal("Site check interval (default 30 min): ", &sample_config.Interval)
	// readVal("Add your slack webhook (eg https://hooks.slack.com/your/custom/id): ", &sample_config.SlackWebhook)
	//
	// fmt.Print("Add email config now? y/n: ")
	// yes := ""
	// _, err = fmt.Scan(&yes)
	// if err != nil {
	// 	fmt.Println("Error reading user input: ", err)
	// 	return
	// }
	// if yes == "yes" || yes == "y" || yes == "Y" {
	// 	readVal("Host (eg smtp.mail.host): ", &sample_config.Mailer.Host)
	// 	readVal("Port (default 587): ", &sample_config.Mailer.Port)
	// 	readVal("Username: ", &sample_config.Mailer.Username)
	// 	readVal("Password: ", &sample_config.Mailer.Password)
	// 	readVal("From Id: ", &sample_config.Mailer.From)
	// 	readVal("To Id: ", &sample_config.EmailTo[0])
	// }
	//
	// err = json.NewEncoder(cf).Encode(sample_config)
	// if err != nil {
	// 	fmt.Println("Error writing pingmon config file:", err)
	// 	return
	// }

	var yes string

	fmt.Println("Successfully created config at", configFilePath)

	fmt.Print("Configure systemd unit now? y/n: ")
	_, err = fmt.Scan(&yes)
	if err != nil {
		fmt.Println("Error reading user input: ", err)
		return
	}
	if yes == "yes" || yes == "y" || yes == "Y" {
		cp := exec.Command("sudo", "cp", "./pingmon.service", "/etc/systemd/system/")
		if err := cp.Run(); err != nil {
			fmt.Println("Error copying pingmon.service into /etc/systemd/system: ", err)
			return
		}

		err = os.Chmod("/etc/systemd/system/pingmon.service", 0640)
		if err != nil {
			fmt.Println("Error changing file permissions at /etc/systemd/system/pingmon.service: ", err)
			return
		}

		reload := exec.Command("sudo", "systemctl daemon-reload")
		if err := reload.Run(); err != nil {
			fmt.Println("Error opening directory /etc/systemd/system: ", err)
			return
		}

		enable := exec.Command("sudo", "systemctl enable pingmon")
		if err := enable.Run(); err != nil {
			fmt.Println("Error opening directory /etc/systemd/system: ", err)
			return
		}

		fmt.Println("Pingmon service configured successfully\nUse \"pingmon start\" to start monitoring")
	}

	fmt.Println("Pingmon setup successful!")
}

func readVal(msg string, v any) {
	fmt.Print(msg)
	_, err := fmt.Scanln(v)
	if err != nil && err.Error() != "unexpected newline" {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
}
