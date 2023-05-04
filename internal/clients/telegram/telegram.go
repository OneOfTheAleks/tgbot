package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"tgbot/internal/clients"
)

const (
	errMsg          = "Request execution error"
	errParse        = "Parse error"
	getUpdateName   = "getUpdates"
	sendMessageName = "sendMessage"
)

type Client struct {
	host string
	path string
	cl   http.Client
}

func New(host, token string) *Client {

	return &Client{
		host: host,
		path: basePath(token),
		cl:   http.Client{},
	}
}

func (c *Client) Send(chatId int, msg string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", msg)
	_, err := c.Request(sendMessageName, q)
	if err != nil {
		return clients.Wrap("error sending message", err)
	}
	return nil
}

func (c *Client) Update(offset, limit int) ([]Message, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))
	data, err := c.Request(getUpdateName, q)
	if err != nil {
		return nil, clients.Wrap(errMsg, err)
	}

	var res Response
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, clients.Wrap(errParse, err)
	}

	return res.Result, nil
}

func (c *Client) Request(method string, q url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.path, method),
	}
	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, clients.Wrap(errMsg, err)
	}
	r.URL.RawQuery = q.Encode()
	resp, err := c.cl.Do(r)
	if err != nil {
		return nil, clients.Wrap(errMsg, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, clients.Wrap(errMsg, err)
	}
	return body, nil
}

func basePath(token string) string {
	return "bot" + token
}
