package domain

import (
	"arch_sample/member/repository/entity"
	"time"
)

/***********************************************
 * @author : sungs
 * @date : 2023-07-24
 * @name : member.go
 * @Role : Domain 관리, usercase, repository interface 관리
 * @Layer : ALL Layer
 * @description : 사용자는 Domain만 보고 해당 비즈니스의 흐름을 알수 있어야 함
 *********************************************/

/*Request 데이터*/
type RegisterMemberCmd struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

/*Return 데이터 */
type Member struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
}

/*
* Member Usecase Interface
* Return Type : Member (Domain)
 */
type MemberUsecase interface {
	Regist(cmd RegisterMemberCmd) (Member, error)
	GetById(id int64) (Member, error)
	Update(id int64, member Member) error
	Delete(id int64) error
}

/*
* Member Repository interfact
* Return Type : MemberEntity
 */

type MemberRepository interface {
	Save(entity.MemberEntity) (entity.MemberEntity, error)
	FindOneById(id int64) (entity.MemberEntity, error)
	Update(id int64, entity entity.MemberEntity) error
	Delete(id int64) error
}
