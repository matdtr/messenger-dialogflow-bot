package config

import (
	"flag"
	"os"
)

type Config struct {
	accessToken     string
	verifyToken     string
	pageID     		string
	projectID       string
	jsonFilePath	string
	host 			string
	apiPort			string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.accessToken, "accessToken", os.Getenv("ACCESS_TOKEN"), "FB app access token")
	flag.StringVar(&conf.verifyToken, "verifyToken", os.Getenv("VERIFY_TOKEN"), "FB app verify token")
	flag.StringVar(&conf.pageID, "pageID", os.Getenv("PAGE_ID"), "FB page ID")
	flag.StringVar(&conf.projectID, "projectID", os.Getenv("PROJECT_ID"), "DialogFlow project ID")
	flag.StringVar(&conf.jsonFilePath, "jsonFilePath", os.Getenv("JSON_KEY"), "DialogFlow service account JSON Key")
	flag.StringVar(&conf.host, "host", os.Getenv("HOST_ADDR"), "Host Address")
	flag.StringVar(&conf.apiPort, "apiPort", os.Getenv("API_PORT"), "API Port")

	flag.Parse()

	return conf
}

func (c *Config) GetAccessToken() string {
	return c.accessToken
}

func (c *Config) GetVerifyToken() string {
	return c.verifyToken

}

func (c *Config) GetPageID() string {
	return c.pageID
}

func (c *Config) GetProjectID() string {
	return c.projectID
}

func (c *Config) GetAuthJSONFilePath() string {
	return c.jsonFilePath
}

func (c *Config) GetHost() string {
	return c.host
}

func (c *Config) GetAPIPort() string {
	return ":" + c.apiPort
}