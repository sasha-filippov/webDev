package telegramClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
	aboutMsg          = `Hello dear user,
My name is Sasha and I'm currently doing my best to learn GO!
Please use my bot to check out my links`
	helpMsg = `Currently I support such commands as 
/about
/help
/links
/start`
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: "bot" + token,
		client:   http.Client{},
	}
}

func (c *Client) SendMessage(chatID int, text string) error {
	
	switch text {
	case "/about":
		text = aboutMsg
	case "/start":
		text = "start"
	case "/help":
		text = helpMsg
	case "/links":
		text = "https://github.com/sasha-filippov"
	}

	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)
	_, err := c.makeRequest(sendMessageMethod, q)
	if err != nil {
		return fmt.Errorf("coulnd't send message: %w", err)
	}
	return nil
}

func (c *Client) Updates(offset int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))

	data, err := c.makeRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}
	var res UpdatesResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res.Result, nil

}

func (c *Client) makeRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't make a request: %w", err)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("couldn't make a request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read response: %w", err)
	}
	return body, nil

}
