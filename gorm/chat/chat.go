package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model        // a table with a name 'user' that have these fields and the default mode fields (which are id createdAt, updatedAt,deletedAt)
	Email      string `gorm:"unique"`
	Username   string `gorm:"unique"`
}
type Channel struct {
	gorm.Model
	Name        string
	Description string
}
type Message struct {
	gorm.Model
	Content   string
	UserId    uint // here the belong to relation is defined showing that every message belongs to a single user and is present in a single channel
	ChannelId uint
	User      User
	Channel   Channel
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Channel{}, &Message{})
	AddData(db)
}
func AddData(db *gorm.DB) {
	channels := []Channel{
		{Name: "Tech Channel", Description: "Everything related to IT"},
		{Name: "Team Channel", Description: "Everything related to the team"},
		{Name: "Support", Description: "For any kind of assistance"},
	}
	db.Create(&channels)
	users := []User{
		{Email: "test@test.com", Username: "tester404"},
		{Email: "supporter@rzp.com", Username: "supporter200"},
	}
	db.Create(&users)
	var tech, support Channel // the type of these variables will tell us what table we are querying from
	db.First(&tech, "Name=?", "Tech Channel")
	db.First(&support, "Name=?", "Support")

	var tester, supporter User
	db.First(&tester, "Username=?", "tester404")
	db.First(&supporter, "Username=?", "supporter200")

	messages := []Message{
		{Content: "Hello I am the IT Guy", Channel: tech, User: tester},
		{Content: "Hello I am from the Support Team", Channel: support, User: supporter},
	}
	db.Create(&messages)

}
func main() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("newDb1.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("There is some error in making a connection to the database")
	}
	sqldb, err1 := db.DB()
	if err1 != nil {
		panic(err1)
	}
	defer sqldb.Close()

	//setup(db)

	// to get all the users present
	var users []User
	db.Find(&users)
	for _, u := range users {
		fmt.Println("Email:", u.Email, "Username:", u.Username)
	}

	// if I want to find all the messages given by a user lets say (user[0]) we can do it becuase the belongs relation is defined
	var messages []Message
	// db.Model(&user[0]).Related(&messages)
	db.Where("user_id=?", users[0].ID).Find(&messages)
	fmt.Println(users[0].Email)
	fmt.Println(messages)
}
