// Code generated by "shoutrrr-gen "; DO NOT EDIT.
package ifttt

import (
	"fmt"
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/conf"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// (‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾)
//  )  Props                          (
// (___________________________________)

type Config struct {
	Events            []string `key:"events" `
	Title             string   `key:"title" `
	UseMessageAsValue int64    `key:"messagevalue" `
	UseTitleAsValue   int64    `key:"titlevalue" `
	Value1            string   `key:"value1" `
	Value2            string   `key:"value2" `
	Value3            string   `key:"value3" `
	WebHookID         string   `url:"host" `
}

type configProp int

const (
	propEvents            configProp = 0
	propTitle             configProp = 1
	propUseMessageAsValue configProp = 2
	propUseTitleAsValue   configProp = 3
	propValue1            configProp = 4
	propValue2            configProp = 5
	propValue3            configProp = 6
	propWebHookID         configProp = 7
	propCount                        = 8
)

var propInfo = types.ConfigPropInfo{
	PropNames: []string{
		"Events",
		"Title",
		"UseMessageAsValue",
		"UseTitleAsValue",
		"Value1",
		"Value2",
		"Value3",
		"WebHookID",
	},

	// Note that propKeys may not align with propNames, as a property can have no or multiple keys
	Keys: []string{
		"events",
		"messagevalue",
		"title",
		"titlevalue",
		"value1",
		"value2",
		"value3",
	},

	DefaultValues: []string{
		"",
		"",
		"2",
		"0",
		"",
		"",
		"",
		"",
	},

	PrimaryKeys: []int{
		0,
		2,
		1,
		3,
		4,
		5,
		6,
		-1,
	},

	KeyPropIndexes: map[string]int{
		"events":       0,
		"messagevalue": 2,
		"title":        1,
		"titlevalue":   3,
		"value1":       4,
		"value2":       5,
		"value3":       6,
	},
}

func (_ *Config) PropInfo() *types.ConfigPropInfo {
	return &propInfo
}

// (‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾)
//  )  GetURL                         (
// (___________________________________)

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL {
	return &url.URL{
		// Userinfo fields are not used for configuration
		Host:     config.WebHookID,
		Path:     "",
		RawQuery: conf.QueryValues(config).Encode(),
		Scheme:   Scheme,
	}
}

// (‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾)
//  )  SetURL                         (
// (___________________________________)

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(configURL *url.URL) error {
	if lc, ok := (interface{})(config).(types.ConfigWithLegacyURLSupport); ok {
		configURL = lc.UpdateLegacyURL(configURL)
	}
	updates := make(map[int]string, propCount)
	updates[int(propWebHookID)] = configURL.Hostname()
	if configURL.Path != "" && configURL.Path != "/" {
		return fmt.Errorf("unexpected path in config URL: %v", configURL.Path)
	}

	for key, value := range configURL.Query() {

		if propIndex, found := propInfo.PropIndexFor(key); found {
			updates[propIndex] = value[0]
		} else if key != "title" {
			return fmt.Errorf("invalid key %q", key)
		}
	}

	err := config.Update(updates)
	if err != nil {
		return err
	}

	if len(config.Events) == 0 {
		return fmt.Errorf("events missing from config URL")
	}

	if config.UseMessageAsValue < 1 || config.UseMessageAsValue > 3 {
		return fmt.Errorf("value %v for useMessageAsValue is not in the range 1-3", config.UseMessageAsValue)
	}

	if config.UseTitleAsValue < 0 || config.UseTitleAsValue > 3 {
		return fmt.Errorf("value %v for useTitleAsValue is not in the range 0-3", config.UseTitleAsValue)
	}

	if config.UseTitleAsValue == config.UseMessageAsValue {
		return fmt.Errorf("value %v for useTitleAsValue is already used for UseMessageAsValue", config.UseTitleAsValue)
	}

	if config.WebHookID == "" {
		return fmt.Errorf("webHookID missing from config URL")
	}

	return nil
}

// (‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾)
//  )  Enums / Options                (
// (___________________________________)

func (config *Config) Enums() map[string]types.EnumFormatter {
	return map[string]types.EnumFormatter{}
}

// Update updates the Config from a map of it's properties
func (config *Config) Update(updates map[int]string) error {
	var last_err error
	for index, value := range updates {
		switch configProp(index) {
		case propEvents:
			if val, err := conf.ParseListValue(value, ","); err != nil {
				last_err = err
			} else {
				config.Events = val
			}
		case propTitle:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Title = val
			}
		case propUseMessageAsValue:
			if val, err := conf.ParseNumberValue(value, 10); err != nil {
				last_err = err
			} else {
				config.UseMessageAsValue = val
			}
		case propUseTitleAsValue:
			if val, err := conf.ParseNumberValue(value, 10); err != nil {
				last_err = err
			} else {
				config.UseTitleAsValue = val
			}
		case propValue1:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Value1 = val
			}
		case propValue2:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Value2 = val
			}
		case propValue3:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Value3 = val
			}
		case propWebHookID:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.WebHookID = val
			}
		default:
			return fmt.Errorf("invalid key")
		}
		if last_err != nil {
			return fmt.Errorf("failed to set value for %v: %v", propInfo.PropNames[index], last_err)
		}
	}
	return nil
}

// Update updates the Config from a map of it's properties
func (config *Config) PropValue(prop int) string {
	switch configProp(prop) {
	case propEvents:
		return conf.FormatListValue(config.Events, ",")
	case propTitle:
		return conf.FormatTextValue(config.Title)
	case propUseMessageAsValue:
		return conf.FormatNumberValue(config.UseMessageAsValue, 10)
	case propUseTitleAsValue:
		return conf.FormatNumberValue(config.UseTitleAsValue, 10)
	case propValue1:
		return conf.FormatTextValue(config.Value1)
	case propValue2:
		return conf.FormatTextValue(config.Value2)
	case propValue3:
		return conf.FormatTextValue(config.Value3)
	case propWebHookID:
		return conf.FormatTextValue(config.WebHookID)
	default:
		return ""
	}
}