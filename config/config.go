package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"gopkg.in/yaml.v3"

	"redwood.dev/rpc"
)

type Config struct {
	BootstrapPeers []BootstrapPeer `yaml:"BootstrapPeers"`
	DataRoot       string          `yaml:"DataRoot"`
	DevMode        bool            `yaml:"DevMode"`

	KeyStore *KeyStoreConfig `yaml:"-"`

	P2PTransport  *P2PTransportConfig  `yaml:"P2PTransport"`
	HTTPTransport *HTTPTransportConfig `yaml:"HTTPTransport"`

	AuthProtocol *AuthProtocolConfig `yaml:"AuthProtocol"`
	BlobProtocol *BlobProtocolConfig `yaml:"BlobProtocol"`
	TreeProtocol *TreeProtocolConfig `yaml:"TreeProtocol"`

	HTTPRPC *rpc.HTTPConfig `yaml:"HTTPRPC"`

	configPath string       `yaml:"-"`
	mu         sync.RWMutex `yaml:"-"`
}

type BootstrapPeer struct {
	Transport     string   `yaml:"Transport"`
	DialAddresses []string `yaml:"DialAddresses"`
}

type KeyStoreConfig struct {
	Password             string `yaml:"Password"`
	InsecureScryptParams bool   `yaml:"InsecureScryptParams"`
}

type P2PTransportConfig struct {
	Enabled     bool   `yaml:"Enabled"`
	ListenAddr  string `yaml:"ListenAddr"`
	ListenPort  uint   `yaml:"ListenPort"`
	ReachableAt string `yaml:"ReachableAt"`
}

type HTTPTransportConfig struct {
	Enabled         bool   `yaml:"Enabled"`
	ListenHost      string `yaml:"ListenHost"`
	DefaultStateURI string `yaml:"DefaultStateURI"`
	ReachableAt     string `yaml:"ReachableAt"`
}

type AuthProtocolConfig struct {
	Enabled bool `yaml:"Enabled"`
}

type BlobProtocolConfig struct {
	Enabled bool `yaml:"Enabled"`
}

type TreeProtocolConfig struct {
	Enabled                 bool   `yaml:"Enabled"`
	MaxPeersPerSubscription uint64 `yaml:"MaxPeersPerSubscription"`
}

func DefaultConfig(appName string) Config {
	configRoot, err := DefaultConfigRoot(appName)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(configRoot, 0777|os.ModeDir)
	if err != nil {
		panic(err)
	}

	dataRoot, err := DefaultDataRoot(appName)
	if err != nil {
		panic(err)
	}

	return Config{
		BootstrapPeers: []BootstrapPeer{},
		DataRoot:       dataRoot,
		DevMode:        false,
		P2PTransport: &P2PTransportConfig{
			Enabled:    true,
			ListenAddr: "0.0.0.0",
			ListenPort: 21231,
		},
		HTTPTransport: &HTTPTransportConfig{
			Enabled:         true,
			ListenHost:      ":8080",
			DefaultStateURI: "",
		},
		AuthProtocol: &AuthProtocolConfig{
			Enabled: true,
		},
		BlobProtocol: &BlobProtocolConfig{
			Enabled: true,
		},
		TreeProtocol: &TreeProtocolConfig{
			Enabled:                 true,
			MaxPeersPerSubscription: 4,
		},
		HTTPRPC: &rpc.HTTPConfig{
			Enabled:    false,
			ListenHost: ":8081",
		},
	}
}

func DefaultConfigRoot(appName string) (root string, _ error) {
	configRoot, err := os.UserConfigDir()
	if err != nil {
		configRoot, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}
	configRoot = filepath.Join(configRoot, appName)
	return configRoot, nil
}

func DefaultConfigPath(appName string) (root string, _ error) {
	configRoot, err := DefaultConfigRoot(appName)
	if err != nil {
		return "", err
	}
	return filepath.Join(configRoot, ".redwoodrc"), nil
}

func DefaultDataRoot(appName string) (string, error) {
	switch runtime.GOOS {
	case "windows", "darwin", "plan9":
		configRoot, err := DefaultConfigRoot(appName)
		if err != nil {
			return "", err
		}
		return configRoot, nil

	default: // unix/linux
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		return filepath.Join(homeDir, ".local", "share", appName), nil
	}
}

func ReadConfigAtPath(appName, configPath string) (*Config, error) {
	if configPath == "" {
		var err error
		configPath, err = DefaultConfigRoot(appName)
		if err != nil {
			return nil, err
		}
		configPath = filepath.Join(configPath, ".redwoodrc")
	}

	// Copy the default config
	cfg := DefaultConfig(appName)

	bs, err := ioutil.ReadFile(configPath)
	// If the file can't be found, we ignore the error.  Otherwise, return it.
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	// Decode the config file on top of the defaults
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, err
	}

	// Save the file again in case it didn't exist or was missing fields
	cfg.configPath = configPath
	err = cfg.save()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) Read(fn func()) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	fn()
}

func (c *Config) Update(fn func() error) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := fn()
	if err != nil {
		return err
	}

	return c.save()
}

func (c *Config) save() error {
	f, err := os.OpenFile(c.configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	encoder.SetIndent(4)

	err = encoder.Encode(c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Path() string {
	return c.configPath
}

func (c *Config) BlobDataRoot() string {
	return filepath.Join(c.DataRoot, "blobs")
}

func (c *Config) TxDBRoot() string {
	return filepath.Join(c.DataRoot, "txs")
}

func (c *Config) StateDBRoot() string {
	return filepath.Join(c.DataRoot, "states")
}

func (c *Config) KeyStoreRoot() string {
	return filepath.Join(c.DataRoot, "keystore")
}

type Duration time.Duration

func (d Duration) MarshalText() ([]byte, error) {
	return []byte(time.Duration(d).String()), nil
}

func (d *Duration) UnmarshalText(text []byte) error {
	dur, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = Duration(dur)
	return nil
}
