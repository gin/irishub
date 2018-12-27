package config

import (
	"bytes"
	"text/template"

	"github.com/spf13/viper"
	cmn "github.com/tendermint/tendermint/libs/common"
)

const defaultConfigTemplate = `# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# Validators reject any tx from the mempool with less than the minimum fee per gas.
minimum_fees = "{{ .BaseConfig.MinFees }}"


# Invariant failure handle level, valid values: 
# 	1. error: just print error message for invariant failure
# 	2. panic: panic and abort program for invariant failure 
# For any other values, invariant checking will be disabled
invariant_level = "{{ .BaseConfig.InvariantLevel }}"
`

var configTemplate *template.Template

func init() {
	var err error
	tmpl := template.New("irisConfigFileTemplate")
	if configTemplate, err = tmpl.Parse(defaultConfigTemplate); err != nil {
		panic(err)
	}
}

// ParseConfig retrieves the default environment configuration for iris.
func ParseConfig() (*Config, error) {
	conf := DefaultConfig()
	err := viper.Unmarshal(conf)
	return conf, err
}

// WriteConfigFile renders config using the template and writes it to configFilePath.
func WriteConfigFile(configFilePath string, config *Config) {
	var buffer bytes.Buffer

	if err := configTemplate.Execute(&buffer, config); err != nil {
		panic(err)
	}

	cmn.MustWriteFile(configFilePath, buffer.Bytes(), 0644)
}
