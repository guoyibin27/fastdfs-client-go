package common

import (
	"github.com/akkuman/parseConfig"
	"strings"
)

type ConfigReader struct {
	config parseConfig.Config
}

// constructor
func NewConfigReader() *ConfigReader {
	return new(ConfigReader)
}

// load config file
func (cr *ConfigReader) LoadConfigFile(filePath string) {
	config := parseConfig.New(filePath)
	cr.config = config
}

// return default value if value dose not exists
func (cr *ConfigReader) GetIntValue(name string, defaultValue int) int {
	value := cr.config.Get(name)
	if value != nil {
		return int(value.(float64))
	}
	return defaultValue
}

// return default value if value dose not exists
func (cr *ConfigReader) GetStringValue(name string, defaultValue string) string {
	value := cr.config.Get(name)
	if value != nil && len(strings.TrimSpace(value.(string))) > 0 {
		return value.(string)
	}
	return defaultValue
}

func (cr *ConfigReader) GetBoolValue(name string, defaultValue bool) bool {
	value := cr.config.Get(name)
	if value == nil {
		return defaultValue
	}
	return value.(bool)
}

// get value array
func (cr *ConfigReader) GetValues(name string) []string {
	value := cr.config.Get(name)
	if value == nil {
		return nil
	}

	temp := value.([]interface{})
	var values []string
	for _, v := range temp {
		values = append(values, v.(string))
	}
	return values
}
