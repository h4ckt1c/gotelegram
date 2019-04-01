package gotelegram

import (
    "encoding/json"
    "net/http"
    "fmt"
    "bytes"
)

const (
    APIURL = "https://api.telegram.org/"
)

type Telegram struct {
    APIToken string
}

type getMe struct {
    Ok bool
    Result getMeResult
}
type getMeResult struct {
    Id int
    Is_bot bool
    First_name string
    Username string
}

type Message struct {
    Text string `json:"text"`
    Chat_id string `json:"chat_id"`
}

func ValidateToken(t *Telegram) error {
    getMe := getMe {}
    getMeURL := fmt.Sprintf("%vbot%v/getMe", APIURL, t.APIToken)
    resp, err := http.Get(getMeURL)
    if err != nil {
        return fmt.Errorf("Get request failed")
    }
    defer resp.Body.Close()
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&getMe)
    if getMe.Ok != true {
        return fmt.Errorf("Invalid API Token")
    }
    return nil
}

func NewClient(token string) (*Telegram, error) {
    client := Telegram {}
    client.APIToken = token
    err := ValidateToken(&client)
    if err != nil {
        return nil, nil
    } else {
        return &client, nil
    }
}

//func SendMessage(t *Telegram, userid string, msg string) error {
func (t Telegram) SendMessage(userid string, msg string) error {
    SendMessageURL := fmt.Sprintf("%vbot%v/sendMessage", APIURL, t.APIToken)
    PlainMessage := Message{
        msg,
        userid,
    }
    jsonMessage, err := json.Marshal(PlainMessage)
    fmt.Println(bytes.NewBuffer(jsonMessage))
    if err != nil {
        return err
    }
    resp, err := http.Post(SendMessageURL, "application/json", bytes.NewBuffer(jsonMessage))
    if err != nil {
        return fmt.Errorf("Post request failed")
    }
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        return fmt.Errorf(resp.Status)
    }
    fmt.Println(resp.StatusCode)
    return nil
}
