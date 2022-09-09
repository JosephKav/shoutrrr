package bark

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/pkr"
	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config for use within the telegram plugin
type LegacyConfig struct {
	standard.EnumlessConfig
	Title     string `key:"title"    default:""      desc:"Notification title, optionally set by the sender"`
	Host      string `url:"host"                     desc:"Server hostname and port"`
	Path      string `url:"path"     default:"/"     desc:"Server path"`
	DeviceKey string `url:"password"                 desc:"The key for each device"`
	Scheme    string `key:"scheme"   default:"https" desc:"Server protocol, http or https"`
	Sound     string `key:"sound"    default:""      desc:"Value from https://github.com/Finb/Bark/tree/master/Sounds"`
	Badge     int64  `key:"badge"    default:"0"     desc:"The number displayed next to App icon"`
	Icon      string `key:"icon"     default:""      desc:"An url to the icon, available only on iOS 15 or later"`
	Group     string `key:"group"    default:""      desc:"The group of the notification"`
	URL       string `key:"url"      default:""      desc:"Url that will jump when click notification"`
	Category  string `key:"category" default:""      desc:"Reserved field, no use yet"`
	Copy      string `key:"copy"     default:""      desc:"The value to be copied"`
}

// GetURL returns a URL representation of it's current field values
func (config *LegacyConfig) GetURL() *url.URL {
	resolver := pkr.NewPropKeyResolver(config)
	return config.getURL(&resolver)
}

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *LegacyConfig) SetURL(url *url.URL) error {
	resolver := pkr.NewPropKeyResolver(config)
	return config.setURL(&resolver, url)
}

func (config *LegacyConfig) getURL(resolver types.ConfigQueryResolver) *url.URL {
	return &url.URL{
		User:       url.UserPassword("", config.DeviceKey),
		Host:       config.Host,
		Scheme:     Scheme,
		ForceQuery: true,
		Path:       config.Path,
		RawQuery:   pkr.BuildQuery(resolver),
	}

}

func (config *LegacyConfig) setURL(resolver types.ConfigQueryResolver, url *url.URL) error {

	password, _ := url.User.Password()
	config.DeviceKey = password
	config.Host = url.Host
	config.Path = url.Path

	for key, vals := range url.Query() {
		if err := resolver.Set(key, vals[0]); err != nil {
			return err
		}
	}

	return nil
}

func (service *Service) GetLegacyConfig() types.ServiceConfig {
	return &LegacyConfig{}
}