package main

import (
	"fmt"
	"os"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-tools/go-steputils/stepconf"
	"github.com/magiconair/properties"
)

// Config - variables should be defined in bitrise secrets.
type Config struct {
	PropertiesFilePath  string `env:"properties_file_path"`
	PropertiesFileKey   string `env:"properties_file_key"`
	PropertiesFileValue string `env:"properties_file_value"`
}

func main() {

	var conf Config
	// Parse the environment variables to a config instance.
	if err := stepconf.Parse(&conf); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	stepconf.Print(conf)

	// Load and parse the properties file
	p := properties.MustLoadFile(conf.PropertiesFilePath, properties.UTF8)

	// Set/Update the new value based on the given
	p.Set(conf.PropertiesFileKey, conf.PropertiesFileValue)

	// Create a writer.
	f, err := os.Create(conf.PropertiesFilePath)

	if err != nil {
		log.Errorf("Error: %s\n", err)
		os.Exit(1)
	}

	_, err2 := p.WriteComment(f, "# ", properties.UTF8)

	if err2 != nil {
		log.Errorf("Error: %s\n", err2)
		os.Exit(1)
	}

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	defer f.Close()
	os.Exit(0)
}
