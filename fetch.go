package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"

	tea "github.com/charmbracelet/bubbletea"
)

type RawMailsDataMail struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Title string `json:"title"`
}

type RawMailsData struct {
	Mails []RawMailsDataMail `json:"mail"`
}

type fetchMsg struct {
	error error
	mails RawMailsData
}

type RawLetter struct {
	Letter RawLetterLetter
}

type RawLetterLetter struct {
	CreatedAt string `json:"created_at"`
	Status string `json:"status"`
	Tags []string `json:"tags"`
	Events []RawLetterLetterEvent `json:"events"`
}

type RawLetterLetterEvent struct {
	HappenedAt string `json:"happened_at"`
	Source string `json:"source"`
	Description string `json:"description"`
	Location string `json:"location"`
}

func fetchMails(api_key string) tea.Cmd {
	return func() tea.Msg {
		var serialized RawMailsData

		req, err := http.NewRequest("GET", "https://mail.hackclub.com/api/public/v1/mail", nil)
		if err != nil { return fetchMsg{err, RawMailsData{}} }

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api_key))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil { return fetchMsg{err, RawMailsData{}} }
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil { return fetchMsg{err, RawMailsData{}} }
		if err := json.Unmarshal(body, &serialized); err != nil { return fetchMsg{err, RawMailsData{}} }
		return fetchMsg{nil, serialized}
	}
}

func fetchLetter(letter_id string, api_key string) (RawLetter, error) {
	var serialized RawLetter

	req, err := http.NewRequest("GET", fmt.Sprintf("https://mail.hackclub.com/api/public/v1/letters/%s", letter_id), nil)
	if err != nil {
		return serialized, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api_key))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil { return serialized, err }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil { return serialized, err }
	if err := json.Unmarshal(body, &serialized); err != nil { return serialized, err }
	return serialized, nil
}
