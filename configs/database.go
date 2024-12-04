package configs

inport (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func initDB() {
	dsn := "root:anthrachite@tcp(Localhost:3306)/evermos_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected")
}

func Migrate() {
	DB.AutoMigrate(&domain.User{})
}

func Migrate() {
	DB.AutoMigrate(&domain.User{}, &domain.Store{})
}

func ConnectDB() {
	configs.DB.AutoMigrate(&domain.User{})
}