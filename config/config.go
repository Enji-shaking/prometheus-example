package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Conf struct {
	ArticleList []Article
}

var ArticleList []Article
var _config Conf

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.Unmarshal(&_config); err != nil {
		panic(fmt.Errorf("Unable to Unmarshal config file: %+v \n", err))
	}
	ArticleList = _config.ArticleList
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&_config); err != nil {
			panic(fmt.Errorf("Unable to Unmarshal config file: %+v \n", err))
		}
		ArticleList = _config.ArticleList
		fmt.Println("Config file changed:", e.Name)
		fmt.Printf("Article list %+v\n", ArticleList)
	})
}
