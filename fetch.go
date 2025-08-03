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
