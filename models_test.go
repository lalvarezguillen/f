package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randomIndex() string {
	return fmt.Sprint(rand.Int31n(1000))
}

func createDummyUser() User {
	idx := randomIndex()
	return User{
		FirstName: "Dummy" + idx,
		LastName:  "User",
		Email:     "dummy" + idx + "@user.com",
	}
}

func createDummyReel() Reel {
	idx := randomIndex()
	return Reel{
		Title:       "Dummy Reel " + idx,
		Description: "This is Dummy Reel for testing, number " + idx,
	}
}

func createDummyPicture() Picture {
	idx := randomIndex()
	return Picture{
		Caption: "This is a Dummy Picture for testing, number " + idx,
		URL:     "http://cdn.f.com/pictures/dummy" + idx + ".jpg",
	}
}

func TestUserReels(t *testing.T) {
	// Setup
	DB.AutoMigrate(User{}, Reel{})
	defer DB.DropTable(User{}, Reel{})
	u := createDummyUser()
	DB.Create(&u)
	rs := []Reel{createDummyReel(), createDummyReel()}
	for n := 0; n < len(rs); n++ {
		rs[n].UserID = u.ID
		DB.Create(&rs[n])
	}
	// Test
	var dbrs []Reel
	DB.Model(&u).Related(&dbrs)
	assert.Equal(t, 2, len(dbrs))
}
