package envload

import (
	"os"
	"testing"
)

func TestLoadWithNormal(t *testing.T) {
	type Config struct {
		SampleId string `env:"SAMPLE_ID"`
	}
	expectedSampleId := "100"
	os.Setenv("SAMPLE_ID", expectedSampleId)
	cfg := Config{}
	if err := Load(&cfg); err != nil {
		t.Error(err)
	}
	if cfg.SampleId != expectedSampleId {
		t.Errorf("SampleId: Actual=%s, Expected=%s", cfg.SampleId, expectedSampleId)
	}
}

func TestLoadWithNoTag(t *testing.T) {
	type Config struct {
		SampleId string
	}
	cfg := Config{}
	if err := Load(&cfg); err != nil {
		t.Error(err)
	}
	if cfg.SampleId != "" {
		t.Errorf("SampleId: Actual=%s, Expected=\"\"", cfg.SampleId)
	}

}
func TestLoadWithDefaultValue(t *testing.T) {
	type Config struct {
		DefaultValue string `env:"DEFAULT_VALUE,default=100"`
	}
	cfg := Config{}
	if err := Load(&cfg); err != nil {
		t.Error(err)
	}

	if cfg.DefaultValue != "100" {
		t.Errorf("DefaultValue: Actual=%v, Expected=100", cfg.DefaultValue)
	}
}
func TestLoadWithRequired(t *testing.T) {
	type Config struct {
		RequiredField string `env:"REQUIRED,required"`
	}
	cfg := Config{}
	if err := Load(&cfg); err == nil {
		t.Errorf("RequiredField: Actual=%v, Expected=nil", err)
	}
}

func TestLoadWithDefaultAndRequired(t *testing.T) {
	type Config struct {
		DefaultAndRequired string `env:"DEFAULT_AND_REQUIRED,required,default=100"`
	}
	cfg := Config{}
	if err := Load(&cfg); err != nil {
		t.Errorf("RequiredField: Actual=%v, Expected=nil", err)
	}
	if cfg.DefaultAndRequired != "100" {
		t.Errorf("DefaultAndRequired: Actual=%v, Expected=100", cfg.DefaultAndRequired)
	}

}
