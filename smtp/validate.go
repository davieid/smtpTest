package smtp

import (
    "crypto/tls"
    "fmt"
    "net"
    "net/smtp"
    "time"
)

// Validate SMTPConfig connects to the SMTP server and authenticates using the provided credentials.
// Uses TLS for port 465 and STARTTLS for port 587.
func ValidateSMTPConfig(config SMTPConfig) error {
    address := net.JoinHostPort(config.Host, fmt.Sprintf("%d", config.Port))

    var client *smtp.Client

    switch config.Port {
    case 465:
        // Direct TLS connection
        tlsConfig := &tls.Config{
            InsecureSkipVerify: true, // You can set this to false for stricter security
            ServerName:         config.Host,
        }

        conn, err := tls.Dial("tcp", address, tlsConfig)
        if err != nil {
            return fmt.Errorf("TLS connection failed: %w", err)
        }

        client, err = smtp.NewClient(conn, config.Host)
        if err != nil {
            return fmt.Errorf("failed to create SMTP client over TLS: %w", err)
        }

    case 587:
        // Plain TCP connection, then upgrade to TLS using STARTTLS
        conn, err := net.DialTimeout("tcp", address, 5*time.Second)
        if err != nil {
            return fmt.Errorf("TCP connection failed: %w", err)
        }

        client, err = smtp.NewClient(conn, config.Host)
        if err != nil {
            return fmt.Errorf("failed to create SMTP client: %w", err)
        }

        tlsConfig := &tls.Config{
            InsecureSkipVerify: true,
            ServerName:         config.Host,
        }

        if err = client.StartTLS(tlsConfig); err != nil {
            return fmt.Errorf("STARTTLS failed: %w", err)
        }

    default:
        return fmt.Errorf("unsupported port: %d", config.Port)
    }

    defer client.Quit()

    // Authenticate
    auth := smtp.PlainAuth("", config.User, config.Pass, config.Host)
    if err := client.Auth(auth); err != nil {
        return fmt.Errorf("authentication failed: %w", err)
    }

    return nil // Success
}
