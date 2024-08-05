package cmd

import (
	"log"
	"main/cmd/domain"
)

// 사용자 추가
func insertMember(member *domain.TMember) {
	result := DB.Create(member)
	log.Printf("row 추가 결과 %s", result.Error)
}

// 사용자 수정
func updateMember(member *domain.TMember) {
	result := DB.Model(&domain.TMember{}).Where("USER_ID = ?", member.UserId).Updates(member)
	log.Printf("row 수정 결과 %s", result.Error)
}

// 전체 읽기
func selectMember() []domain.TMember {
	var members []domain.TMember
	DB.Find(&members)

	log.Printf("Number of members: %d\n", len(members)) // 가져온 row 갯수
	// row 전체 출력
	for _, data := range members {
		// log.Printf("UserId: %s, FirstName: %s, LastName: %s\n", data.UserId, data.FirstName, data.LastName)
		log.Printf("UserId: %s, FirstName: %s, LastName: %s, Email: %s, CustomDate: %s\n", data.UserId, data.FirstName, data.LastName, data.Email, data.CustomDate)
	}
	return members
}

// ////////////////////////////////////////////////////////////////////////////////////
// 단일 읽기
// ////////////////////////////////////////////////////////////////////////////////////
func selectMemberOne(userId string) domain.TMember {
	var member domain.TMember
	DB.First(&member, "USER_ID = ?", userId)

	// log.Printf("UserId: %s, FirstName: %s, LastName: %s, Email: %s, CreateAt: %s\n",
	// 	member.UserId, member.FirstName, member.LastName, member.Email, member.CreateAt)
	return member
}

// ////////////////////////////////////////////////////////////////////////////////////
// 단건 삭제
// ////////////////////////////////////////////////////////////////////////////////////
func deleteMemberOne(userId string) bool {
	result := DB.Delete(&domain.TMember{}, "user_id = ?", userId)

	// 삭제 결과 확인
	if result.Error != nil {
		// fmt.Println("Error deleting record:", result.Error)
		return false
	} else {
		if result.RowsAffected == 0 {
			return false
		} else {
			return true
		}
	}
}
