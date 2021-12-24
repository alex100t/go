package main

import (
	"errors"
	"strconv"
	"strings"
)

// Update is an update response, from GetUpdates.
type Update struct {
	UpdateID int `json:"update_id"`
	// Message new incoming message of any kind — text, photo, sticker, etc.
	//
	// optional
	Message *Message `json:"message"`
	// EditedMessage
	//
	// optional
	EditedMessage *Message `json:"edited_message"`
}

// Message is returned by almost every request, and contains data about
// almost anything.
type Message struct {
	// MessageID is a unique message identifier inside this chat
	MessageID int `json:"message_id"`
	// From is a sender, empty for messages sent to channels;
	//
	// optional
	From *User `json:"from"`
	// Date of the message was sent in Unix time
	Date int `json:"date"`
	// Chat is the conversation the message belongs to
	Chat *Chat `json:"chat"`

	// Text is for text messages, the actual UTF-8 text of the message, 0-4096 characters;
	//
	// optional
	Text string `json:"text"`

	// Entities is for text messages, special entities like usernames,
	// URLs, bot commands, etc. that appear in the text;
	//
	// optional
	Entities *[]MessageEntity `json:"entities"`
}

// Chat contains information about the place a message was sent.
type Chat struct {
	// ID is a unique identifier for this chat
	ID int64 `json:"id"`
}

// MessageEntity contains information about data in a Message.
type MessageEntity struct {
	// Type of the entity.
	// Can be:
	//  “mention” (@username),
	//  “hashtag” (#hashtag),
	//  “cashtag” ($USD),
	//  “bot_command” (/start@jobs_bot),
	//  “url” (https://telegram.org),
	//  “email” (do-not-reply@telegram.org),
	//  “phone_number” (+1-212-555-0123),
	//  “bold” (bold text),
	//  “italic” (italic text),
	//  “underline” (underlined text),
	//  “strikethrough” (strikethrough text),
	//  “code” (monowidth string),
	//  “pre” (monowidth block),
	//  “text_link” (for clickable text URLs),
	//  “text_mention” (for users without usernames)
	Type string `json:"type"`
	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`
	// Length
	Length int `json:"length"`
	// URL for “text_link” only, url that will be opened after user taps on the text
	//
	// optional
	URL string `json:"url"`
	// User for “text_mention” only, the mentioned user
	//
	// optional
	User *User `json:"user"`
}

// IsCommand returns true if the type of the message entity is "bot_command".
func (e MessageEntity) IsCommand() bool {
	return e.Type == "bot_command"
}

// IsCommand returns true if message starts with a "bot_command" entity.
func (m *Message) IsCommand() bool {
	if m.Entities == nil || len(*m.Entities) == 0 {
		return false
	}

	entity := (*m.Entities)[0]
	return entity.Offset == 0 && entity.IsCommand()
}

// Command checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
//
// If the command contains the at name syntax, it is removed. Use
// CommandWithAt() if you do not want that.
func (m *Message) Command() string {
	command := m.CommandWithAt()

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}

// CommandWithAt checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
//
// If the command contains the at name syntax, it is not removed. Use Command()
// if you want that.
func (m *Message) CommandWithAt() string {
	if !m.IsCommand() {
		return ""
	}

	// IsCommand() checks that the message begins with a bot_command entity
	entity := (*m.Entities)[0]
	return m.Text[1:entity.Length]
}

// CommandArguments checks if the message was a command and if it was,
// returns all text after the command name. If the Message was not a
// command, it returns an empty string.
//
// Note: The first character after the command name is omitted:
// - "/foo bar baz" yields "bar baz", not " bar baz"
// - "/foo-bar baz" yields "bar baz", too
// Even though the latter is not a command conforming to the spec, the API
// marks "/foo" as command entity.
func (m *Message) CommandArguments() string {
	if !m.IsCommand() {
		return ""
	}

	// IsCommand() checks that the message begins with a bot_command entity
	entity := (*m.Entities)[0]
	if len(m.Text) == entity.Length {
		return "" // The command makes up the whole message
	}

	return m.Text[entity.Length+1:]
}

// User represents a Telegram user or bot.
type User struct {
	// ID is a unique identifier for this user or bot
	ID int `json:"id"`
	// FirstName user's or bot's first name
	FirstName string `json:"first_name"`
	// LastName user's or bot's last name
	//
	// optional
	LastName string `json:"last_name"`
	// UserName user's or bot's username
	//
	// optional
	UserName string `json:"username"`
	// LanguageCode IETF language tag of the user's language
	// more info: https://en.wikipedia.org/wiki/IETF_language_tag
	//
	// optional
	LanguageCode string `json:"language_code"`
	// IsBot true, if this user is a bot
	//
	// optional
	IsBot bool `json:"is_bot"`
}

func ProcMessage(m *Message) (string, error) {

	var resp string
	var err error

	if m.IsCommand() {
		resp, err = ProcCommand(m)
		if err != nil {
			return "", err
		}
	} else {
		resp, err = ProcDate(m)
		if err != nil {
			return "", err
		}
	}

	return resp, nil
}

func ProcDate(m *Message) (string, error) {

	var resp string
	var err error
	var jdDate int

	if strings.Contains(m.Text, "-") {

		resp, err = cvtDATE(m.Text)
		if err != nil {
			return "", err
		}

	} else if strings.ToUpper(m.Text) == "C" {

		resp, err = cvtUNIX(m.Date)
		if err != nil {
			return "", err
		}

	} else {

		jdDate, err = strconv.Atoi(m.Text)
		if err != nil {
			err = errors.New("invalid Number format")
			return "", err
		}

		resp, err = cvtJD(jdDate)
		if err != nil {
			return "", err
		}

	}

	return resp, nil
}

func ProcCommand(m *Message) (string, error) {

	var resp string
	var err error

	if m.Command() == "start" {

		resp = ` Hi!
 I 
 do ^JD 
 to convert dates between Julian and USA date formats 
 and something else...
 Let's start!`

	} else if m.Command() == "help" {

		resp = ` Bot emulates command 
GTM>do ^JD 
to convert dates between Julian and USA date formats. 
Send to bot 65000 and it will return 12-18-2018 (18 December 2018) 
or send 09-13-2021 (13 September 2021) and bot returns 66000 
or C to retrieve your current date`

	} else if m.Command() == "n2t" {

		resp = CvtNumber2Time(m.CommandArguments())

	} else if m.Command() == "t2n" {

		resp = CvtTime2Number(m.CommandArguments())

	} else {

		err = errors.New("invalid Command")
		return "", err

	}

	return resp, nil
}
