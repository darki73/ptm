package configuration

import (
	bi "github.com/darki73/ptm/pkg/configuration/base-image"
	ci "github.com/darki73/ptm/pkg/configuration/cloud-init"
	"github.com/darki73/ptm/pkg/configuration/downloader"
	"github.com/darki73/ptm/pkg/configuration/qemu"
	"github.com/darki73/ptm/pkg/configuration/repositories"
	uu "github.com/darki73/ptm/pkg/configuration/unattended-upgrades"
	"github.com/darki73/ptm/pkg/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

var (
	// configuration is the global variable that holds the configuration.
	configuration *Configuration
	// mutes is the global variable that holds the mutex for the configuration.
	mutex = &sync.RWMutex{}
	// ChangeChannel is the channel that is used to notify the application that the configuration has changed.
	ChangeChannel = make(chan bool)
)

// Configuration is the configuration for the application.
type Configuration struct {
	// BaseImage is a reference to the BaseImage configuration.
	BaseImage *bi.Configuration `json:"base_image" yaml:"base_image" toml:"base_image" mapstructure:"base_image"`
	// CloudInit is a reference to the CloudInit configuration.
	CloudInit *ci.Configuration `json:"cloud_init" yaml:"cloud_init" toml:"cloud_init" mapstructure:"cloud_init"`
	// Downloader is a reference to the Downloader configuration.
	Downloader *downloader.Configuration `json:"downloader" yaml:"downloader" toml:"downloader" mapstructure:"downloader"`
	// Qemu is a reference to the Qemu configuration.
	Qemu *qemu.Configuration `json:"qemu" yaml:"qemu" toml:"qemu" mapstructure:"qemu"`
	// Repositories is a list of repositories that will be added to the base image.
	Repositories []*repositories.Configuration `json:"repositories" yaml:"repositories" toml:"repositories" mapstructure:"repositories"`
	// UnattendedUpgrades is a reference to the UnattendedUpgrades configuration.
	UnattendedUpgrades *uu.Configuration `json:"unattended_upgrades" yaml:"unattended_upgrades" toml:"unattended_upgrades" mapstructure:"unattended_upgrades"`
	// BasePackages is a list of packages that will be installed in the base image.
	BasePackages []string `json:"base_packages" yaml:"base_packages" toml:"base_packages" mapstructure:"base_packages"`
	// ExtraPackages is a list of user specified packages that will be installed in the base image.
	ExtraPackages []string `json:"extra_packages" yaml:"extra_packages" toml:"extra_packages" mapstructure:"extra_packages"`
	// LogLevel is the log level for the application.
	LogLevel string `json:"log_level" yaml:"log_level" toml:"log_level" mapstructure:"log_level"`
}

// GetBaseImage returns the BaseImage configuration.
func (configuration *Configuration) GetBaseImage() *bi.Configuration {
	return configuration.BaseImage
}

// GetCloudInit returns the CloudInit configuration.
func (configuration *Configuration) GetCloudInit() *ci.Configuration {
	return configuration.CloudInit
}

// GetDownloader returns the Downloader configuration.
func (configuration *Configuration) GetDownloader() *downloader.Configuration {
	return configuration.Downloader
}

// GetQemu returns the Qemu configuration.
func (configuration *Configuration) GetQemu() *qemu.Configuration {
	return configuration.Qemu
}

// GetRepositories returns the list of repositories that will be added to the base image.
func (configuration *Configuration) GetRepositories() []*repositories.Configuration {
	return configuration.Repositories
}

// GetUnattendedUpgrades returns the UnattendedUpgrades configuration.
func (configuration *Configuration) GetUnattendedUpgrades() *uu.Configuration {
	return configuration.UnattendedUpgrades
}

// GetBasePackages returns the list of packages that will be installed in the base image.
func (configuration *Configuration) GetBasePackages() []string {
	return configuration.BasePackages
}

// GetExtraPackages returns the list of user specified packages that will be installed in the base image.
func (configuration *Configuration) GetExtraPackages() []string {
	return configuration.ExtraPackages
}

// GetLogLevel returns the log level for the application.
func (configuration *Configuration) GetLogLevel() string {
	return configuration.LogLevel
}

// IsConfigurationLoaded returns true if the configuration is loaded.
func IsConfigurationLoaded() bool {
	return configuration != nil
}

// LoadConfiguration loads the configuration from the given options.
func LoadConfiguration(options *Options) error {
	viper.SetConfigName(options.GetName())
	viper.SetConfigType(options.GetExtension())
	viper.AddConfigPath(options.GetPath())
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	configuration = &Configuration{
		BaseImage:          bi.InitializeWithDefaults(),
		CloudInit:          ci.InitializeWithDefaults(),
		Downloader:         downloader.InitializeWithDefaults(),
		Qemu:               qemu.InitializeWithDefaults(),
		Repositories:       []*repositories.Configuration{},
		UnattendedUpgrades: uu.InitializeWithDefaults(),
		BasePackages:       []string{},
		ExtraPackages:      []string{},
		LogLevel:           "i",
	}

	if err := viper.Unmarshal(configuration); err != nil {
		return err
	}

	setBasePackages()

	viper.WatchConfig()
	viper.OnConfigChange(configurationReloadHandler)

	return nil
}

// GetConfiguration returns the configuration for the application.
func GetConfiguration() *Configuration {
	mutex.RLock()
	defer mutex.RUnlock()

	return configuration
}

// setLogLevel sets the log level.
func setLogLevel() error {
	desiredLogLevel, err := log.ParseLevel(configuration.LogLevel)

	if err != nil {
		return err
	}

	log.SetLevel(desiredLogLevel)

	return nil
}

// configurationReloadHandler is a function that is called when the configuration is reloaded.
func configurationReloadHandler(event fsnotify.Event) {
	log.InfofWithFields(
		"configuration file changed: %s",
		log.FieldsMap{
			"source": "configuration",
		},
		event.Name,
	)

	mutex.Lock()
	defer mutex.Unlock()

	if err := viper.Unmarshal(configuration); err != nil {
		log.ErrorfWithFields(
			"error re-loading configuration: %s",
			log.FieldsMap{
				"source": "configuration",
			},
			err,
		)
	}

	setBasePackages()

	if err := setLogLevel(); err != nil {
		log.ErrorfWithFields(
			"error setting log level: %s",
			log.FieldsMap{
				"source": "configuration",
			},
			err,
		)
	}

	ChangeChannel <- true
}

// setBasePackages sets the base packages.
func setBasePackages() {
	if len(configuration.BasePackages) > 0 {
		return
	}
	configuration.BasePackages = []string{
		"apt-transport-https",
		"aptitude",
		"ca-certificates",
		"curl",
		"htop",
		"jq",
		"mc",
		"software-properties-common",
		"qemu-guest-agent",
		"wget",
	}
}
