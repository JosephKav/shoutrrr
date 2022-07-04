// Code generated by "shoutrrr-gen --lang go ../../../spec/telegram.yml"; DO NOT EDIT.
package telegram

import (
	"fmt"
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"
	"github.com/containrrr/shoutrrr/pkg/format"
)

type Config struct {
	Chats        []string
	Notification bool
	ParseMode    int64
	Preview      bool
	Title        string
	Token        string
}

type configProp int
const (
	propChats        configProp = 0
	propNotification configProp = 1
	propParseMode    configProp = 2
	propPreview      configProp = 3
	propTitle        configProp = 4
	propToken        configProp = 5
	propCount = 6
)
var propNames = []string{
	"Chats",
	"Notification",
	"ParseMode",
	"Preview",
	"Title",
	"Token",
}

// Note that propKeys may not align with propNames, as a property can have no or multiple keys
var propKeys = []string{
	"channels",
	"chats",
	"notification",
	"parsemode",
	"preview",
	"title",
}

var defaultValues = []string{
	"",
	"Yes",
	"None",
	"Yes",
	"",
	"",
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL {
	return &url.URL{
		User: url.User(config.Token),
	}
}

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error {
	updates := make(map[string]string, propCount)
	updates["Token"] = url.User.Username()

	if qv := url.Query()["channels"]; len(qv) == 1 {
        updates["Chats"] = qv[0]
	}
	if qv := url.Query()["chats"]; len(qv) == 1 {
        updates["Chats"] = qv[0]
	}
	if qv := url.Query()["notification"]; len(qv) == 1 {
        updates["Notification"] = qv[0]
	}
	if qv := url.Query()["parsemode"]; len(qv) == 1 {
        updates["ParseMode"] = qv[0]
	}
	if qv := url.Query()["preview"]; len(qv) == 1 {
        updates["Preview"] = qv[0]
	}
	if qv := url.Query()["title"]; len(qv) == 1 {
        updates["Title"] = qv[0]
	}

	return nil
}

func (config *Config) Enums() map[string]types.EnumFormatter {
	return map[string]types.EnumFormatter{
		"ParseMode": ParseModeFormatter,
	}
}

var (
	ParseModeFormatter = format.CreateEnumFormatter([]string{
		"None",
		"Markdown",
		"HTML",
		"MarkdownV2",
	})
)
// Update updates the Config from a map of it's properties
func (config *Config) Update(updates map[string]string) error {
	var last_err error
	for key, value := range updates {
		switch key {
		case "Chats":
			if val, err := format.ParseListValue(value); err != nil {
				last_err = err
			} else {
				config.Chats = val
			}
		case "Notification":
			if val, err := format.ParseToggleValue(value); err != nil {
				last_err = err
			} else {
				config.Notification = val
			}
		case "ParseMode":
			if val, err := format.ParseNumberValue(value, 10); err != nil {
				last_err = err
			} else {
				config.ParseMode = val
			}
		case "Preview":
			if val, err := format.ParseToggleValue(value); err != nil {
				last_err = err
			} else {
				config.Preview = val
			}
		case "Title":
			if val, err := format.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Title = val
			}
		case "Token":
			if val, err := format.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Token = val
			}
		default:
			last_err = fmt.Errorf("invalid key")
		}
		if last_err != nil {
			return fmt.Errorf("failed to set value for %q: %v", key, last_err)
		}
	}
	return nil
}

// Init sets all the Config properties to their default values
func (config *Config) Init() {
	updates := make(map[string]string, propCount)
	for i, name := range propNames {
		updates[name] = defaultValues[i]
	}
	config.Update(updates)
}
