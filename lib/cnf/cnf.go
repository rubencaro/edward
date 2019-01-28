package cnf

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

// Read parses all input from the outside world into a Config struct
func Read() (*Config, error) {
	// actual read of config file
	c, err := readFile()
	if err != nil {
		return nil, err
	}

	// fill in with default values
	config := applyDefaults(c)

	return config, nil
}

func readFile() (*Config, error) {
	var err error

	configFile := "edward.toml"

	// create 'edward.toml' if it does not exist
	err = ensureItExists(configFile)
	if err != nil {
		return nil, err
	}

	// read values
	config, err := readTOML(configFile)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func ensureItExists(file string) error {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return createSample(file)
	}
	return nil
}

func createSample(dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, strings.NewReader(sample))
	return err
}

func readTOML(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	inputBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = toml.Unmarshal(inputBytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func applyDefaults(c *Config) *Config {
	return &Config{
		ImgURL: c.ImgURL, // this goes as is, no default
	}
}
