package configuration

import (
	"github.com/magiconair/properties"
)

var AppLanguage *properties.Properties

func init() {
	AppLanguage = properties.MustLoadFile("./src/main/resources/message_en.properties", properties.UTF8)
	AppLanguage = properties.MustLoadFile("./src/main/resources/message_fr.properties", properties.UTF8)
}
