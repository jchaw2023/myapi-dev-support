package configbase

import (
	"github.com/magiconair/properties"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"os"
	"strings"
)

type ConfigBase struct {
	Mode string
}

func (c *ConfigBase) GetEtcdEndpoints(env string) []string {
	endpointsConfig := os.Getenv("etcd.endpoints")
	if endpointsConfig == "" {
		props, err := properties.LoadFile("./etc/etcd-"+env+".properties", properties.ISO_8859_1)
		if err != nil || props.Len() == 0 {
			if err != nil {
				log.Fatalf("./etc/etcd-"+env+".properties==>%#v", err)
			}
			props, err = properties.LoadFile("./etcd-"+env+".properties", properties.ISO_8859_1)
			if err != nil {
				log.Fatalf("./etcd-"+env+".properties==>%#v", err)
			}
		}
		if props.Len() > 0 {
			endpoints := props.GetString("etcd.endpoints", "")
			return strings.Split(endpoints, ";")
		}
		return nil
	} else {
		return strings.Split(endpointsConfig, ";")
	}
}
func (c *ConfigBase) Load(projectName, mode string, config interface{}) bool {
	c.Mode = mode
	if mode == "pod" || mode == "test" {
		endpoints := c.GetEtcdEndpoints(mode)
		if len(endpoints) == 0 {
			log.Fatalln("etcd-endpoint is nil")
			return false
		}
		fileKey := "/config/" + projectName + "/conf-" + mode + ".yaml"
		viper.AddRemoteProvider("etcd3", endpoints[0], fileKey)
		viper.SetConfigType("yaml")
		err := viper.ReadRemoteConfig()
		if err != nil {
			log.Fatalln(err)
			return false
		}
	} else {
		viper.SetConfigFile("./etc/conf-" + mode + ".yaml")
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalln(err)
			return false
		}
	}
	err := viper.Unmarshal(config, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "yaml"
	})
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}
