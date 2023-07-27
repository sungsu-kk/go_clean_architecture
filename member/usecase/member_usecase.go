package usecase

import (
	"arch_sample/domain"
	"arch_sample/member/repository/entity"
	"log"
)

type memberUsecase struct {
	memberRepo domain.MemberRepository
}

func CreateMemberUseCase(memberRepository domain.MemberRepository) domain.MemberUsecase {
	return &memberUsecase{
		memberRepo: memberRepository,
	}
}

func (m memberUsecase) Regist(cmd domain.RegisterMemberCmd) (domain.Member, error) {
	//TODO implement me
	log.Println("MemberUseCase Save Called")
	var member domain.Member
	/*Command to Entity*/
	memberEntity := entity.MemberEntity{
		Name: cmd.Name,
		Age:  cmd.Age,
	}

	/*Repository call*/
	repoResult, err := m.memberRepo.Save(memberEntity)
	if err != nil {
		return member, nil
	}
	/*Entity to Domain*/
	member.ID = repoResult.Id
	member.Name = repoResult.Name
	member.Age = repoResult.Age
	member.CreatedAt = repoResult.CreatedAt

	/*return*/
	return member, nil
}

func (m memberUsecase) GetById(id int64) (domain.Member, error) {
	//TODO implement me
	panic("implement me")
}

func (m memberUsecase) Update(id int64, member domain.Member) error {
	//TODO implement me
	panic("implement me")
}

func (m memberUsecase) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
