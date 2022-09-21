package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/radovskyb/watcher"
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v3"
)

type smtpConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

type messageConfig struct {
	Subject      string
	From         string
	FromNickname string `yaml:"fromNickname"`
	To           string
	ToNickname   string `yaml:"toNickname"`
}

type config struct {
	Smtp      smtpConfig
	Message   messageConfig
	WatchPath string `yaml:"watchPath"`
}

type cliOptions struct {
	cfgPath string
}

func parseCliArgs() cliOptions {
	var opts cliOptions
	cfgPath := flag.String("config", "config.yaml", "Path to YAML config file")
	flag.Parse()
	opts.cfgPath = *cfgPath
	return opts
}

func loadConfig(filePath string) (config, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config{}, err
	}
	var cfg config
	err = yaml.Unmarshal(fileBytes, &cfg)
	if err != nil {
		return config{}, err
	}
	return cfg, nil
}

func readyMessage(cfg config, filePath string) *gomail.Message {
	m := gomail.NewMessage()
	c := cfg.Message

	if c.FromNickname != "" {
		m.SetHeader("From", m.FormatAddress(c.From, c.FromNickname))
	} else {
		m.SetHeader("From", c.From)
	}

	if c.ToNickname != "" {
		m.SetHeader("To", m.FormatAddress(c.To, c.ToNickname))
	} else {
		m.SetHeader("To", c.To)
	}

	m.SetHeader("Subject", c.Subject)
	m.Attach(filePath)

	return m
}

func readyDialer(cfg config) gomail.Dialer {
	d := gomail.Dialer{
		Host:      cfg.Smtp.Host,
		Port:      cfg.Smtp.Port,
		Username:  cfg.Smtp.Username,
		Password:  cfg.Smtp.Password,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return d
}

func main() {
	opts := parseCliArgs()

	cfg, err := loadConfig(opts.cfgPath)
	if err != nil {
		panic(err)
	}

	w := watcher.New()
	if err := w.Add(cfg.WatchPath); err != nil {
		panic(err)
	}
	w.FilterOps(watcher.Create)

	go func() {
		for {
			select {
			case event := <-w.Event:
				m := readyMessage(cfg, event.Path)
				d := readyDialer(cfg)
				if err := d.DialAndSend(m); err != nil {
					panic(err)
				}
				fmt.Printf("Sent %q\n", event.Path)
			case err := <-w.Error:
				panic(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
