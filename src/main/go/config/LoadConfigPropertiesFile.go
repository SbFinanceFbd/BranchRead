package configuration

import (
	"github.com/magiconair/properties"
)

var (
	AppConfig *properties.Properties
	config    *properties.Properties
)

func init() {
	config = properties.MustLoadFile("./src/main/resources/config.properties", properties.UTF8)
	active := config.MustGetString("active")
	AppConfig = properties.MustLoadFile("./src/main/resources/config_"+active+".properties", properties.UTF8)

}
