package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ShareInfo struct {
	Link                 string `json:"link"`
	ShareText            string `json:"share_text"`
	OriginalContentOwner string `json:"original_content_owner"`
}
type Message struct {
	SenderName     string    `json:"sender_name"`
	TimeStamp      int64     `json:"timestamp_ms"`
	Content        string    `json:"content"`
	Share          ShareInfo `json:"share"`
	GeoBlocked     bool      `json:"is_geoblocked_for_viewer"`
	UnsentByParent bool      `json:"is_unsent_image_by_messenger_kid_parent"`
}

type Participant struct {
	Name string `json:"name"`
}
type MessagesList struct {
	Participants     []Participant `json:"participants"`
	Messages         []Message     `json:"messages"`
	Title            string        `json:"title"`
	StillParticipant bool          `json:"is_still_participant"`
	ThreadPath       string        `json:"thread_path"`
	MagicWords       []string      `json:"magic_words"`
}

type Person struct {
	Name          string
	MessagesCount int64
}

func (p *Person) AddMessage() {
	p.MessagesCount++
}

func SearchPerson(persons []*Person, name string) (*Person, error) {
	for _, v := range persons {
		if v.Name == name {
			return v, nil
		}
	}
	return &Person{}, fmt.Errorf("Couldn't find a person: ", name)
}

func main() {
	args := os.Args[1:]
	path := args[0]

	var result MessagesList

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		panic(err)
	}
	persons := make([]*Person, 0)
	for _, v := range result.Participants {
		persons = append(persons, &Person{
			Name:          v.Name,
			MessagesCount: 0,
		})
	}

	for _, m := range result.Messages {
		p, err := SearchPerson(persons, m.SenderName)
		if err != nil {
			panic(err)
		}
		p.AddMessage()
	}

	for _, p := range persons {
		fmt.Println("User: ", p.Name, ", Messages count: ", p.MessagesCount)
	}
}
