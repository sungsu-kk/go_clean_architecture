package repository

import (
	"arch_sample/domain"
	"arch_sample/member/repository/entity"
	"gorm.io/gorm"
	"log"
)

type MemberRepository struct {
	DB *gorm.DB
}

func CreateMemberRepository(DB *gorm.DB) domain.MemberRepository {
	if err := DB.AutoMigrate(entity.MemberEntity{}); err != nil {
		panic(err)
	}
	return &MemberRepository{
		DB: DB,
	}
}

func (m *MemberRepository) Save(entity entity.MemberEntity) (entity.MemberEntity, error) {
	//TODO implement me
	log.Print("MemberRepository Save method called")
	result := m.DB.Create(&entity)
	if result.Error != nil {
		return entity, result.Error
	}
	return entity, nil
}

func (m *MemberRepository) FindOneById(id int64) (entity.MemberEntity, error) {
	//TODO implement me
	log.Print("MemberRepository FindById method called")
	return entity.MemberEntity{}, nil
}

func (m *MemberRepository) Update(id int64, entity entity.MemberEntity) error {
	//TODO implement me
	log.Print("MemberRepository Update method called")
	return nil
}

func (m *MemberRepository) Delete(id int64) error {
	//TODO implement me
	log.Print("MemberRepository Delete method called")
	return nil
}
