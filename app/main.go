package main

import (
	"arch_sample/app/config"
	"arch_sample/docs"
	"arch_sample/member/delivery/http"
	"arch_sample/member/repository"
	"arch_sample/member/usecase"
	"flag"
)

/***********************************************
* @author : sungs
* @date : 2023-07-24
* @name : main.go
* @Role : Application Main method
* @description :
	 - init() : Main method 호출 전
 		- role : config 파일 validation
	 - main() : Main method
*********************************************/

func init() {

}

/* main 메서드 */
func main() {

	configfile := flag.String("config", "./resources/application-dev.cfg", "config file path")
	flag.Parse()
	configuration := config.Configuration{}
	err := configuration.LoadConfiguration(*configfile)
	if err != nil {
		panic(err)
	}
	// Webserver
	webServer := configuration.WebServer
	app := webServer.Initialize()

	//Swagger
	swagger := configuration.Swagger
	docs.SwaggerInfo.BasePath = "/api/install/v1"
	docs.SwaggerInfo.Host = swagger.GetUsingIp() + ":" + webServer.Port

	// Database
	database := configuration.Database
	database.Connect()

	// Application Wired
	memberRepository := repository.CreateMemberRepository(config.DB)
	memberUseCase := usecase.CreateMemberUseCase(memberRepository)
	http.CreateMemberHandler(app, memberUseCase)

	// Application Run
	webServer.Run(app)

}
