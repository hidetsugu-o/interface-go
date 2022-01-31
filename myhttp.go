package main

type httpClient struct {
	SecretPassword string // SecretPasswordを用意
}

// Interfaceを返却するコンストラクタを定義
func New() HttpClient {
	client := httpClient{SecretPassword: "pass"}
	return &client
}

type HttpClient interface {
	Get()
	Post()
}

func (c *httpClient) Get() {}

func (c *httpClient) Post() {}
