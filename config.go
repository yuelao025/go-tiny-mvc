package main

/*
	This contains all the functions
	related to loading and reloading
	configuration.
*/

// Imports

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	ERR_COULD_NOT_READ_CONFIG_FILE    = 30
	ERR_COULD_NOT_PROCESS_CONFIG_FILE = 31
)

// Structs
// Custom configuration structs should be
// set below and added to the Conf struct.

type Conf struct {
	Server  ServerConf  `yaml:"server"`
	Logging LoggingConf `yaml:"logging"`
}

type ServerConf struct {
	Port    int    `yaml:"port"`
	Address string `yaml:"address"`
}

type LoggingConf struct {
	Level string `yaml:"level"`
}

// Functions

func configure(path string) (ServerConf, error) {
	/*
		Extracts the configuration and passes
		it onto other functions that actually
		perform the configuration process.

		Everything except the actual HTTP 
		server should be configured in here.
		This allows for easy reloading by just
		calling the configure function without
		destroying the HTTP server.

		Custom configuration functions should
		be added in here.
	*/
	logrus.WithFields(logrus.Fields{
		"path": path,
	}).Info("Loading configuration")
	conf, err := readConfiguration(path)

	// Add in custom configuration functions here.
	configureLogging(conf.Logging)

	// End custom configuration functions.
	return conf.Server, err
}

func readConfiguration(path string) (Conf, error) {
	/*
		Reads the configuration from the file path.
	*/
	var conf Conf
	// Read the file
	confFile, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"errorcode": ERR_COULD_NOT_READ_CONFIG_FILE,
			"error":     err,
			"filename":  path,
		}).Error("Could not read configuration file")
		return conf, err
	}
	// Turn the YAML into the Conf struct.
	err = yaml.Unmarshal(confFile, &conf)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"errorcode": ERR_COULD_NOT_READ_CONFIG_FILE,
			"error":     err,
			"filename":  path,
		}).Error("Could not process configuration file")
		return conf, err
	}

	return conf, err

}

func configureLogging(loggingConf LoggingConf) {
	/*
		Sets the logging configuration such as
		the file to use and the level to log at.
		This can be done dynamically and can be
		reloaded during runtime.
	*/

	logrus.WithFields(logrus.Fields{
		"requested": loggingConf.Level,
	}).Debug("Setting logging level")

	// Set the logging level.
	switch loggingConf.Level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.WithFields(logrus.Fields{
			"setting": loggingConf.Level,
		}).Warn("No log level set defaulting to 'info'")
		logrus.SetLevel(logrus.InfoLevel)
	}
}
