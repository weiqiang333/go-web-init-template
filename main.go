// author: weiqiang; date: 2022-12
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"go-web-init-template/internal/api"
	"go-web-init-template/internal/metrics"
)

func init() {
	pflag.String("configFile", "configs/config.yaml", "go config file")
	pflag.String("listen_address", "0.0.0.0:8080", "server listen address.")
	pflag.ErrHelp.Error()
}

func main() {
	loadConfig()

	prometheus.MustRegister(metrics.NewExporter())

	listenAddress := viper.GetString("listen_address")
	router := engine()
	err := router.Run(listenAddress) // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(fmt.Errorf("Failed web server: %s ", err.Error()))
	}
}

// gin web run engine
func engine() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "./web/static")

	router.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})
	router.POST("/-/reload", reloadConfig)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/", api.Default)

	userAuth := loadAuthUsers()
	authorized := router.Group("/api/v1", gin.BasicAuth(userAuth))
	authorized.GET("/", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		c.String(200, "asd", user)
	})

	return router
}

// load config and flag config
func loadConfig() {
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Println(err.Error())
		panic(fmt.Errorf("Fatal error BindPFlags: %w \n", err))
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(viper.GetString("configFile"))
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

// reloadConfig 127.0.0.1:8080/-/reload
func reloadConfig(c *gin.Context) {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(fmt.Errorf("Fatal error config file: %w \n", err))
		c.String(http.StatusOK, fmt.Sprintf("Failed reload config file: %s, err: %s", viper.ConfigFileUsed(), err.Error()))
		return
	}
	fmt.Println(fmt.Sprintf("reload config file: %s", viper.ConfigFileUsed()))
	c.String(http.StatusOK, fmt.Sprintf("reload config file: %s", viper.ConfigFileUsed()))
}

func loadAuthUsers() map[string]string {
	return viper.GetStringMapString("auth.basic")
}
