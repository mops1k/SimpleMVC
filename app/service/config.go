package service

import (
    "io/ioutil"
    "path/filepath"
    "time"

    "github.com/arthurkushman/pgo"
    "github.com/spf13/viper"
)

type Config struct {
    reader *viper.Viper
}

var configuration *Config

func initConfig() *Config {
        configuration = &Config{reader: viper.New()}
        configuration.reader.AutomaticEnv()
        configuration.reader.SetConfigType("yaml")

        dirName := "./config/"

        files, err := ioutil.ReadDir(dirName)
        if err != nil {
            Container.GetLogger().App.Panic(err)
        }

        for _, file := range files {
            if file.IsDir() {
                continue
            }

            if !pgo.InArray(filepath.Ext(file.Name()), []string{".yaml", ".yml"}) {
                continue
            }

            configuration.reader.SetConfigFile(dirName + file.Name())
            err = configuration.reader.MergeInConfig()
            if err != nil {
                Container.GetLogger().App.Fatal(err)
            }
        }

    return configuration
}

// Get config parameter value
func (c *Config) Get(key string) interface{} {
    return c.reader.Get(key)
}

func (c *Config) GetString(key string) string {
    return c.reader.GetString(key)
}

func (c *Config) GetInt(key string) int {
    return c.reader.GetInt(key)
}

func (c *Config) GetBool(key string) bool {
    return c.reader.GetBool(key)
}

func (c *Config) GetStringSlice(key string) []string {
    return c.reader.GetStringSlice(key)
}

func (c *Config) GetDuration(key string) time.Duration {
    return c.reader.GetDuration(key)
}

// Add file to config
func (c *Config) AddFile(path string) {
    c.reader.SetConfigFile(path)
    err := c.reader.MergeInConfig()
    if err != nil {
        Container.GetLogger().App.Fatal(err)
    }

    err = c.reader.MergeInConfig()
    if err != nil {
        Container.GetLogger().App.Fatal(err)
    }
}
