package util

import (
	"arch_sample/domain"
	"arch_sample/member/repository/entity"
)

func toDomain(entity entity.MemberEntity) domain.Member {
	return domain.Member{
		ID:   entity.Id,
		Name: entity.Name,
	}
}
