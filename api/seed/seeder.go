package seed

import (
	"log"
	"time"

	"github.com/LFTrip/ProjectLFTrip/api/models"
	"github.com/jinzhu/gorm"
)

// Load : Validation and Join
func Load(db *gorm.DB) {

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	input := "2018-10-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	var users = []models.User{
		models.User{
			Firstname:        "Steve",
			Lastname:         "Victor",
			Email:            "steve@gmail.com",
			Accesslevel:      1,
			Password:         "password",
			Dateofbirth:      t,
			Sexe:             "Homme",
			City:             "Paris",
			PhoneNumber:      "0625321458",
			DepartureAirport: "Orly",
			Description:      "Developer Go",
		},
		models.User{
			Firstname:        "Kevin",
			Lastname:         "Feige",
			Email:            "feige@gmail.com",
			Accesslevel:      1,
			Password:         "password",
			Dateofbirth:      t,
			Sexe:             "Homme",
			City:             "Paris",
			PhoneNumber:      "0625321458",
			DepartureAirport: "Orly",
			Description:      "Student IT",
		},
		models.User{
			Firstname:        "Karim",
			Lastname:         "Benzema",
			Email:            "k@benzema.io",
			Accesslevel:      1,
			Password:         "Mostafa87",
			Dateofbirth:      t,
			Sexe:             "Homme",
			City:             "Paris",
			PhoneNumber:      "0625321458",
			DepartureAirport: "Orly",
			Description:      "Soccer player",
		},
	}

	var trips = []models.Trip{
		models.Trip{
			Country:     "France",
			Title:       "Voyage au sud de la France",
			Description: "Voyage en petit comités pour s'amuser, apprendre à se connaître et découvrir des choses",
			StartDate:   start,
			EndDate:     end,
			NbDays:      5,
			MiddleAge:   20,
			NbTraveler:  4,
			Program:     "On partirai de Nice et nous longerons toute la côte Est jusqu'à Menton.",
			Lodging:     "Principalement du airbnb",
			Budget:      350,
			AuthorID:    1,
		},
		models.Trip{
			Country:     "Espagne",
			Title:       "Road Trip détente dans le sud de l'Espagne",
			Description: "Flaner dans les plus belles plages de la côte sud d'Espagne, profiter du soleil et se reposer !",
			StartDate:   start,
			EndDate:     end,
			NbDays:      10,
			MiddleAge:   25,
			NbTraveler:  6,
			Program:     "Pas encore fait, mais si vous êtes intéressé, envoyer moi un MP et on y réfléchira ensemble.",
			Lodging:     "Principalement du airbnb ou auberge de jeunesse.",
			Budget:      500,
			AuthorID:    2,
		},
		models.Trip{
			Country:     "Italie",
			Title:       "Trek en Italie, ça vous dit ?",
			Description: "Voyage pour les fans de trekking et de nature !",
			StartDate:   start,
			EndDate:     end,
			NbDays:      7,
			MiddleAge:   28,
			NbTraveler:  3,
			Program:     "Pas encore décidé.",
			Lodging:     "Une tente !",
			Budget:      250,
			AuthorID:    1,
		},
	}

	err := db.Debug().DropTableIfExists(&models.Trip{}, &models.User{}, &models.Like{}, &models.Comment{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Trip{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Trip{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		trips[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Trip{}).Create(&trips[i]).Error
		if err != nil {
			log.Fatalf("cannot seed trips table: %v", err)
		}
	}
}
