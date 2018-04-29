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

	var rs []Reel
	for n := 0; n < 2; n++ {
		r := createDummyReel()
		r.UserID = u.ID
		DB.Create(&r)
		rs = append(rs, r)
	}

	// Test
	var dbrs []Reel
	DB.Model(&u).Related(&dbrs)
	assert.Equal(t, 2, len(dbrs))
}

func TestReelPictures(t *testing.T) {
	// Setup
	DB.AutoMigrate(User{}, Reel{}, Picture{})
	defer DB.DropTable(User{}, Reel{}, Picture{})

	u := createDummyUser()
	DB.Create(&u)

	var ps []Picture
	for idx := 0; idx < 2; idx++ {
		p := createDummyPicture()
		p.UserID = u.ID
		DB.Create(&p)
		ps = append(ps, p)
	}
	r := createDummyReel()
	r.UserID = u.ID
	r.Pictures = ps
	DB.Create(&r)

	// Tests
	var pics []Picture
	DB.Model(&r).Related(&pics, "Pictures")
	assert.Equal(t, 2, len(pics))
}

func TestUserPictures(t *testing.T) {
	// Setup
	DB.AutoMigrate(User{}, Picture{})
	defer DB.DropTable(User{}, Picture{})

	u := createDummyUser()
	DB.Create(&u)

	var ps []Picture
	for idx := 0; idx < 2; idx++ {
		p := createDummyPicture()
		p.UserID = u.ID
		DB.Create(&p)
		ps = append(ps, p)
	}

	// Tests
	var pics []Picture
	DB.Model(&u).Related(&pics)
	assert.Equal(t, 2, len(pics))
}
