package entity

import "time"

/* Persistence Layer Entity */
type MemberEntity struct {
	Id        int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement;NOT NULL;unique;comment:아이디"`
	Name      string    `json:"name" grom:"column:name;type:varchar(30);NOT NULL;comment:이름"`
	Age       int       `json:"age" grom:"column:age;type:Integer;NOT NULL;comment:이름"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;type:datetime;NOT NULL;default:CURRENT_TIMESTAMP;comment:생성 일자"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt;type:datetime;comment:수정 일자"`
	DeleteYn  string    `json:"deleteYn" gorm:"column:deleteYn;type:enum('Y','N');NOT NULL;default:'N';comment:삭제 여부"`
}
