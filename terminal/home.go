package terminal

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/davieid/smtpTest/smtp"
)

func ShowHome() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("📁 Enter path to your SMTP JSON config file: ")
    jsonPath, _ := reader.ReadString('\n')
    jsonPath = strings.TrimSpace(jsonPath)

    for {
        fmt.Println("\n📬 Welcome to SMTP Validator")
        fmt.Println("Please choose an option:")
        fmt.Println("1. View Configuration")
        fmt.Println("2. Validate SMTP Credentials")
        fmt.Println("3. Validate and Send Test Email")
        fmt.Println("4. Exit")
        fmt.Print("Enter your choice (1-4): ")

        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        switch input {
        case "1":
            configs, err := smtp.LoadJSON(jsonPath)
            if err != nil {
                fmt.Println("❌ Failed to load configuration:", err)
                continue
            }

            fmt.Println("📄 Loaded SMTP Configurations:")
            for i, config := range configs {
                fmt.Printf("  %d. Host: %s | Port: %d | Username: %s\n", i+1, config.Host, config.Port, config.User)
            }

        case "2":
            configs, err := smtp.LoadJSON(jsonPath)
            if err != nil {
                fmt.Println("❌ Failed to load configuration:", err)
                continue
            }

            fmt.Println("🔐 Validating SMTP Credentials...")
            for i, config := range configs {
                fmt.Printf("  [%d] %s: ", i+1, config.User)
                err := smtp.ValidateSMTPConfig(config)
                if err != nil {
                    fmt.Println("❌", err)
                } else {
                    fmt.Println("✅ Valid")
                }
            }

        case "3":
            fmt.Print("Enter recipient email address: ")
            recipient, _ := reader.ReadString('\n')
            recipient = strings.TrimSpace(recipient)

            configs, err := smtp.LoadJSON(jsonPath)
            if err != nil {
                fmt.Println("❌ Failed to load configuration:", err)
                continue
            }

            fmt.Println("📤 Sending test emails...")
            for i, config := range configs {
                fmt.Printf("  [%d] %s: ", i+1, config.User)
                success := smtp.SendTestEmail(config, recipient)
                if success {
                    fmt.Println("✅ Sent")
                } else {
                    fmt.Println("❌ Failed")
                }
            }

        case "4":
            fmt.Println("👋 Exiting. Goodbye!")
            return

        default:
            fmt.Println("❗ Invalid choice. Please enter a number between 1 and 4.")
        }
    }
}
