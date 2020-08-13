package main

import (
	"gotrading/app/controllers"
	"gotrading/config"
	"gotrading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile) // ログ設定
	controllers.StreamIngestionData()            // データ取得
	log.Println(controllers.StartWebServer())    // ウェブサーバー起動
}
