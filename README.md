# 📬 SMTP Validator

✅ A lightweight, interactive Go program for validating and testing SMTP (Simple Mail Transfer Protocol) account credentials. This tool is designed to help developers and sysadmins verify email server configurations, test authentication, and send test emails — all from the terminal.

---

## 🚀 Features

- 🔐 Validate SMTP credentials using TLS (port 465) or STARTTLS (port 587)
- 📤 Send test emails to verify full SMTP functionality
- 📄 Load multiple SMTP configurations from a JSON file
- 🧑‍💻 Interactive terminal interface with menu-driven options
- 🧪 Logs full connection and authentication process
- 🛡️ Secure: avoids committing sensitive credentials by ignoring `/json` directory

---

## 📁 Project Structure

```
smtp-validator/
├── go.mod
├── main.go
├── json/                         # Ignored by Git (contains sensitive configs)
│   └── smtp_config.json         # Your actual SMTP credentials (not committed)
├── smtp/
│   ├── json.go                  # Loads SMTP configs from JSON
│   ├── send.go                  # Sends test emails
│   ├── validate.go              # Validates SMTP credentials
│   └── types.go                 # Defines SMTPConfig struct
├── terminal/
│   └── home.go                  # Interactive terminal UI
├── README.md
└── LICENSE
```

---

## 🧾 JSON Configuration

Create a file at `json/smtp_config.json` with the following structure:

```json
[
  {
    "host": "smtp.gmail.com",
    "port": 587,
    "user": "your-email@gmail.com",
    "pass": "your-password"
  },
  {
    "host": "smtp.mail.yahoo.com",
    "port": 465,
    "user": "your-email@yahoo.com",
    "pass": "your-password"
  }
]
```

> ⚠️ This file is ignored by Git. Do not commit real credentials.

---

## 🖥️ How to Use

1. **Clone the repository** and navigate into it:

   ```bash
   git clone https://github.com/davieid/smtpTest.git
   cd smtpTest
   ```

2. **Initialize Go module (if not already):**

   ```bash
   go mod tidy
   ```

3. **Run the program:**

   ```bash
   go run main.go
   ```

4. **Follow the interactive prompts:**

   ```
   📁 Enter path to your SMTP JSON config file: json/smtp_config.json

   📬 Welcome to SMTP Validator
   Please choose an option:
   1. View Configuration
   2. Validate SMTP Credentials
   3. Validate and Send Test Email
   4. Exit
   ```

---

## 🛠️ Tech Stack

- Language: [Go](https://golang.org/)
- Standard libraries: `net/smtp`, `crypto/tls`, `encoding/json`, `bufio`, `os`, `fmt`, `log`

---

## 📦 Future Improvements

- Add support for OAuth2 authentication
- Export results to a log file
- Add CLI flags for headless operation
- Support for port 25 with STARTTLS fallback

---

## 🧑‍⚖️ License & Usage

This project is licensed under the MIT License. However:

> **Commercial use of this software or its derivatives requires prior written permission from the author. Attribution is required in all distributions.**

---

## 🙌 Contributions

Contributions are welcome! Feel free to fork the repo, submit pull requests, or open issues for bugs and feature requests.

---

## 👤 Author

**Bolaji**  
Built in 🆖 with ❤️ and Go.
```
