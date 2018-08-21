package main

import (
		_ "github.com/go-sql-driver/mysql"
	"github.com/mifly/v2exs"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/mifly/v2exs/config"
	"github.com/labstack/echo"
)

func main() {
	var Engine *xorm.Engine

	var err error
	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8", config.Get("db_user"),
		config.Get("db_passwd"), config.Get("db_name"))
	Engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = Engine.Sync2(new(v2exs.Node))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	App := echo.New()
	App.GET("/v2ex/:t", v2exs.GetTopics)
	//sigs := make(chan os.Signal, 1)
	//// `signal.Notify` registers the given channel to
	//// receive notifications of the specified signals.
	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//<-sigs
	App.Logger.Fatal(App.Start(":1323"))
}

