package http

import (
	"arch_sample/domain"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type MemberHandler struct {
	memberUsecase domain.MemberUsecase
}

func CreateMemberHandler(r *gin.Engine, usecase domain.MemberUsecase) {
	memberHandler := MemberHandler{usecase}
	member := r.Group("/member")
	{
		v1 := member.Group("/v1")
		v1.GET("/ping", memberHandler.Ping)
		v1.POST("/regist", memberHandler.RegisterMember)
	}
}

func (e *MemberHandler) Ping(g *gin.Context) {

	g.JSON(http.StatusOK,
		domain.CommonResponse{
			Status:  http.StatusOK,
			Message: "SUCCESS",
			Data:    "arch_sample return 'pong' current_time = " + time.Now().String(),
		})

	return
}
func (e *MemberHandler) RegisterMember(g *gin.Context) {

	log.Println("MemberHandler RegisterMember called ")
	/*Json to RequestData*/
	var request domain.RegisterMemberCmd
	if err := g.ShouldBindJSON(&request); err != nil {
		g.JSON(http.StatusBadRequest,
			domain.CommonResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Error:   domain.ErrInvalidRequest,
				Path:    "[" + g.Request.Method + "] " + g.Request.URL.Path,
			})
		return
	}
	/* Call Service 내부 memberUseCase 사용 바람 */
	member, err := e.memberUsecase.Regist(request)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			domain.CommonResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Error:   domain.ErrInternalServerError,
				Path:    "[" + g.Request.Method + "] " + g.Request.URL.Path,
			})
		return
	}
	/*결과 return*/
	g.JSON(http.StatusOK,
		domain.CommonResponse{
			Status:  http.StatusOK,
			Message: "SUCCESS",
			Data:    member,
		})

	return

}
