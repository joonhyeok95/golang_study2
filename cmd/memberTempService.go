package cmd

import (
	"log"
	"main/cmd/domain"
)

// 사용자 추가
func insertMemberTemp(memberTemp *domain.TMemberTemp) {
	result := DB.Save(memberTemp)
	log.Printf("T_MEMBER_TEMP Row 추가 결과 %s", result.Error)
}
func insertAllMemberTemp(members []domain.TMember) int {
	var cnt int
	for _, member := range members {
		tempMember := &domain.TMemberTemp{
			UserId:     member.UserId,
			FirstName:  member.FirstName,
			LastName:   member.LastName,
			Email:      member.Email,
			CustomDate: member.CustomDate,
		}
		DB.Save(&tempMember)
		cnt++
	}
	log.Printf("T_MEMBER_TEMP Row '%d' 건 변경됨", cnt)
	return cnt
}

// 전체 읽기
// func selectMemberTemp() []domain.TMemberTemp {
// 	var memberTemps []domain.TMemberTemp
// 	DB.Find(&memberTemps)

// 	log.Printf("Number of memberTemps: %d\n", len(memberTemps)) // 가져온 row 갯수
// 	// row 전체 출력
// 	for _, data := range memberTemps {
// 		// log.Printf("UserId: %s, FirstName: %s, LastName: %s\n", data.UserId, data.FirstName, data.LastName)
// 		log.Printf("UserId: %s, FirstName: %s, LastName: %s, Email: %s, CreateAt: %s\n", data.UserId, data.FirstName, data.LastName, data.Email, data.CreateAt)
// 	}
// 	return memberTemps
// }

// 단일 읽기
// func selectMemberTempOne(userId string) domain.TMemberTemp {
// 	var memberTemp domain.TMemberTemp
// 	DB.First(&memberTemp, "USER_ID = ?", userId)

// 	// log.Printf("UserId: %s, FirstName: %s, LastName: %s, Email: %s, CreateAt: %s\n",
// 	// 	memberTemp.UserId, memberTemp.FirstName, memberTemp.LastName, memberTemp.Email, memberTemp.CreateAt)
// 	return memberTemp
// }

// 단건 삭제
// func deleteMemberTempOne(userId string) bool {
// 	result := DB.Delete(&domain.TMemberTemp{}, "user_id = ?", userId)

// 	// 삭제 결과 확인
// 	if result.Error != nil {
// 		// fmt.Println("Error deleting record:", result.Error)
// 		return false
// 	} else {
// 		if result.RowsAffected == 0 {
// 			return false
// 		}
// 		return true
// 	}
// }
