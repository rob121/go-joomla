package joomla

import (
	"gopkg.in/ini.v1"
	"io/ioutil"
	"strings"
)

/*
 Loads the joomla config given and makes it available to go
 Path to config
*/

var Config *IniConfig

type IniConfig struct {
	Src *ini.File
}

func (c *IniConfig) Get(key string, def ...interface{}) interface{} {

	if !Config.Src.Section("").HasKey(key) {
		if len(def) > 0 {

			return def[0]
		}
	}

	return Config.Src.Section("").Key(key)

}

func (c *IniConfig) GetString(key string, def ...string) string {

	if !Config.Src.Section("").HasKey(key) {

		if len(def) > 0 {

			return def[0]
		}

	}

	return Config.Src.Section("").Key(key).String()

}

func (c *IniConfig) GetInt(key string, def ...int) (int, error) {

	if !Config.Src.Section("").HasKey(key) {
		if len(def) > 0 {
			return def[0], nil
		}
	}

	return Config.Src.Section("").Key(key).Int()

}

func LoadConfig(filename string) error {

	var out []string

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	for _, v := range lines {
		if strings.Contains(v, "=") {

			tmp := strings.Replace(v, `';`, `'`, -1)
			segs := strings.Split(tmp, "$")

			if len(segs) > 1 {
				out = append(out, segs[1])
			}
		}

	}

	ini_string := []byte(strings.Join(out, "\n"))

	cfg, err := ini.Load(ini_string)

	if err != nil {

		return err

	}

	Config = &IniConfig{Src: cfg}

	return nil

}
