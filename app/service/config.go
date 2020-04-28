package service

import (
    "github.com/spf13/viper"
    "io/ioutil"
    "os"
    "time"
)

type Config struct {
    reader *viper.Viper
    files  []string
}

var Configuration *Config

func InitConfig() *Config {
    if Configuration == nil {
        var files []os.FileInfo
        Configuration = &Config{reader: viper.New()}
        Configuration.reader.AutomaticEnv()
        Configuration.reader.SetConfigType("yaml")

        dirName := "./config/"

        files, err := ioutil.ReadDir(dirName)
        if err != nil {
            Logger.App.Panic(err)
        }

        for _, file := range files {
            if file.IsDir() {
                continue
            }
            Configuration.reader.SetConfigFile(dirName + file.Name())
            err = Configuration.reader.MergeInConfig()
            if err != nil {
                Logger.App.Fatal(err)
            }
        }
    }

    return Configuration
}

// Get config parameter value
func (c *Config) Get(key string) interface{} {
    return c.reader.Get(key)
}

func (c *Config) GetInt(key string) int {
    return c.reader.GetInt(key)
}

func (c *Config) GetBool(key string) bool {
    return c.reader.GetBool(key)
}

func (c *Config) GetDuration(key string) time.Duration {
    return c.reader.GetDuration(key)
}

// Add file to config
func (c *Config) AddFile(path string) {
    c.reader.SetConfigFile(path)
    err := Configuration.reader.MergeInConfig()
    if err != nil {
        Logger.App.Fatal(err)
    }

    err = c.reader.MergeInConfig()
    if err != nil {
        Logger.App.Fatal(err)
    }
}
