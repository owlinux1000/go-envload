package envload

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

const TAG_KEYWORD = "env"

type ParsedTag struct {
	EnvName      string
	Required     bool
	DefaultValue string
}

func NewParsedTag(tag string) *ParsedTag {
	tag_splited := strings.Split(tag, ",")
	env_name := tag_splited[0]
	isRequired := false
	defaultValue := ""
	for _, t := range tag_splited {
		if strings.HasPrefix(t, "default=") {
			defaultValue = strings.Split(t, "=")[1]
		}
		if t == "required" {
			isRequired = true
		}
	}
	return &ParsedTag{
		EnvName:      env_name,
		Required:     isRequired,
		DefaultValue: defaultValue,
	}
}

func Load(config any) error {
	t := reflect.TypeOf(config).Elem()
	v := reflect.ValueOf(config).Elem()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(TAG_KEYWORD)

		parsedTag := NewParsedTag(tag)
		if !f.IsExported() {
			return fmt.Errorf("The field must be exported: %s\n", parsedTag.EnvName)
		}

		env, ok := os.LookupEnv(parsedTag.EnvName)

		if !ok && parsedTag.DefaultValue != "" {
			env = parsedTag.DefaultValue
		}

		if !ok && parsedTag.Required && env == "" {
			return fmt.Errorf("The field must be required: %s\n", parsedTag.EnvName)
		}
		v.FieldByName(f.Name).Set(reflect.ValueOf(env))
	}

	return nil
}
