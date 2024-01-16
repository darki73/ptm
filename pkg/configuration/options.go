package configuration

// Options is a struct that holds the configuration options for the application.
type Options struct {
	// Name is the name of the configuration file.
	Name string
	// Path is the path to the configuration file.
	Path string
	// Extension is the extension of the configuration file.
	Extension string
}

// NewOptions creates a new Options struct.
func NewOptions(name, path, extension string) Options {
	defaultOptions := NewDefaultOptions()
	if name == "" {
		name = defaultOptions.GetName()
	}
	if path == "" {
		path = defaultOptions.GetPath()
	}
	if extension == "" {
		extension = defaultOptions.GetExtension()
	}

	return Options{
		Name:      name,
		Path:      path,
		Extension: extension,
	}
}

// NewCustomOptions creates a new Options struct with custom values.
func NewCustomOptions(name, path, extension string) *Options {
	return &Options{
		Name:      name,
		Path:      path,
		Extension: extension,
	}
}

// NewDefaultOptions creates a new Options struct with default values.
func NewDefaultOptions() *Options {
	return &Options{
		Name:      "config",
		Path:      "/etc/ptm",
		Extension: "yaml",
	}
}

// GetName returns the name of the configuration file.
func (options *Options) GetName() string {
	return options.Name
}

// GetPath returns the path to the configuration file.
func (options *Options) GetPath() string {
	return options.Path
}

// GetExtension returns the extension of the configuration file.
func (options *Options) GetExtension() string {
	return options.Extension
}
