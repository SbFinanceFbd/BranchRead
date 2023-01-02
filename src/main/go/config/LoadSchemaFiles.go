package configuration

import (
	"github.com/magiconair/properties"
)

var AppSchema *properties.Properties

func init() {
	AppSchema = properties.MustLoadFile("./src/main/resources/schema_name.properties", properties.UTF8)
}
