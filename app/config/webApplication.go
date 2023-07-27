package config

import (
	"github.com/gin-gonic/gin"
)

// 어플리케이션 설정 정보 구조체
type WebServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Mode string `json:"mode"`
}

// 어플리케이션 설정정보 출력
func (ws *WebServer) String() string {
	return "Host: " + ws.Host + ", Port: " + ws.Port
}

// 어플리케이션 모드 설정
func (ws *WebServer) SetMode(mode string) {
	gin.SetMode(mode)
}

// 어플리케이션 초기화
func (ws *WebServer) Initialize() *gin.Engine {
	gin.SetMode(ws.Mode)
	return gin.Default()
}

// 어플리케이션 실행
func (ws *WebServer) Run(g *gin.Engine) {
	g.Run(ws.Host + ":" + ws.Port)
}
