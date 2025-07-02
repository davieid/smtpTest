package smtp

import (
    "crypto/tls"
    "fmt"
    "log"
    "net"
    "net/smtp"
    "strings"
    "time"
)

// SendTestEmail sends a test email to the specified recipient using the given SMTP configuration.
// Logs the full process and returns true if successful, false otherwise.
func SendTestEmail(config SMTPConfig, recipient string) bool {
    address := net.JoinHostPort(config.Host, fmt.Sprintf("%d", config.Port))
    subject := "SMTP Test Email"
    body := "This is a test email sent by the SMTP validator."

    log.Println("🔌 Connecting to SMTP server:", address)

    var client *smtp.Client
    var err error

    switch config.Port {
    case 465:
        tlsConfig := &tls.Config{
            InsecureSkipVerify: true,
            ServerName:         config.Host,
        }
        conn, err := tls.Dial("tcp", address, tlsConfig)
        if err != nil {
            log.Println("❌ TLS connection failed:", err)
            return false
        }
        client, err = smtp.NewClient(conn, config.Host)
        if err != nil {
            log.Println("❌ Failed to create SMTP client over TLS:", err)
            return false
        }

    case 587:
        conn, err := net.DialTimeout("tcp", address, 5*time.Second)
        if err != nil {
            log.Println("❌ TCP connection failed:", err)
            return false
        }
        client, err = smtp.NewClient(conn, config.Host)
        if err != nil {
            log.Println("❌ Failed to create SMTP client:", err)
            return false
        }
        tlsConfig := &tls.Config{
            InsecureSkipVerify: true,
            ServerName:         config.Host,
        }
        if err = client.StartTLS(tlsConfig); err != nil {
            log.Println("❌ STARTTLS failed:", err)
            return false
        }

    default:
        log.Println("❌ Unsupported port:", config.Port)
        return false
    }

    defer client.Quit()

    log.Println("🔐 Authenticating...")
    auth := smtp.PlainAuth("", config.User, config.Pass, config.Host)
    if err := client.Auth(auth); err != nil {
        log.Println("❌ Authentication failed:", err)
        return false
    }

    log.Println("📤 Setting sender and recipient...")
    if err := client.Mail(config.User); err != nil {
        log.Println("❌ MAIL FROM failed:", err)
        return false
    }
    if err := client.Rcpt(recipient); err != nil {
        log.Println("❌ RCPT TO failed:", err)
        return false
    }

    log.Println("✍️ Writing message...")
    wc, err := client.Data()
    if err != nil {
        log.Println("❌ DATA command failed:", err)
        return false
    }

    message := strings.Join([]string{
        fmt.Sprintf("To: %s", recipient),
        fmt.Sprintf("From: %s", config.User),
        fmt.Sprintf("Subject: %s", subject),
        "",
        body,
    }, "\r\n")

    _, err = wc.Write([]byte(message))
    if err != nil {
        log.Println("❌ Writing message failed:", err)
        return false
    }

    err = wc.Close()
    if err != nil {
        log.Println("❌ Closing message writer failed:", err)
        return false
    }

    log.Println("✅ Test email sent successfully!")
    return true
}
