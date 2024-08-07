package cmd

import (
	"log"
	"main/cmd/domain"
	"os"

	apmmysql "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	DBNAME := os.Getenv("DBNAME")
	DBHOST := os.Getenv("DBHOST")
	DBPORT := os.Getenv("DBPORT")
	DBUSERNAME := os.Getenv("DBUSERNAME")
	DBPASSWORD := os.Getenv("DBPASSWORD")

	dsn := DBUSERNAME + ":" + DBPASSWORD + "@tcp(" + DBHOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := DBNAME + ":" + "Meta1915034@"@tcp(127.0.0.1:3306)/jh_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(apmmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}
	log.Printf("DB Connection Successed")
	db.AutoMigrate(&domain.TMember{}, &domain.TMemberTemp{}) // 테이블 자동 생성
	//db.Debug()                                               // 수행로그 debug 상태
	DB = db
}

// db.First(&tMember) // primary key기준으로 product 찾기
// result := db.First(&mem)
// log.Printf("발견된 record cnt : %d", result.RowsAffected)
// log.Printf("에러반환 %s", result.Error)
// log.Printf("데이터 %s", result.UserId)
// result.RowsAffected // returns count of records found
// result.Error        // returns error or nil

// for i, TMember := range mem {
// 	log.Println(i)
// }

// db.First(&product, "code = ?", "D42") // code가 D42인 product 찾기

// 수정 - product의 price를 200으로
// db.Model(&product).Update("Price", 200)
// 수정 - 여러개의 필드를 수정하기
// db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

// 삭제 - product 삭제하기
// db.Delete(&product, 1)
