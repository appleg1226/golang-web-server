package models

import (
	"awesomeProject/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

func dbConnection() {
	config.DB, _ = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfigTest())), &gorm.Config{})
	_ = config.DB.AutoMigrate(&User{})
}

func TestCreateUser(t *testing.T) {
	dbConnection()

	user := User{Username: "dmschd", Password: "dmschd"}
	if err := CreateUser(&user); err != nil {
		t.Error("not created")
	}

	var user2 User
	config.DB.First(&user2, 1)
	assert.Equal(t, user.Id, user2.Id)
	config.DB.Migrator().DropTable("user")
}

func TestGetAllUsers(t *testing.T) {
	dbConnection()

	users := []User{{Username: "dmschd1", Password: "dmschd"},
		{Username: "dmschd2", Password: "dmschd"},
		{Username: "dmschd3", Password: "dmschd"},
	}

	for _, i := range users {
		_ = CreateUser(&i)
	}

	var count int64
	config.DB.Table("user").Count(&count)

	assert.Equal(t, int64(3), count)

	config.DB.Migrator().DropTable("user")
}

func TestUpdateUser(t *testing.T) {
	dbConnection()

	user := User{Username: "dmschd", Password: "dmschd"}
	if err := CreateUser(&user); err != nil {
		t.Error("not created")
	}

	user.Password = "dmschd2"
	_ = UpdateUser(&user)
	assert.Equal(t, "dmschd2", user.Password)
	config.DB.Migrator().DropTable("user")
}

func TestDeleteUser(t *testing.T) {
	dbConnection()

	user := User{Username: "dmschd", Password: "dmschd"}
	if err := CreateUser(&user); err != nil {
		t.Error("not created")
	}

	_ = DeleteUser(&user, strconv.Itoa(user.Id))

	var count int64
	config.DB.Table("user").Count(&count)
	assert.Equal(t, int64(0), count)
	config.DB.Migrator().DropTable("user")
}