package service

import (
	"log"
	"time"

	"gitlab.com/komentr/access/models"
)

//New User
func New(req models.UserRequest) (models.User, error) {
	s := models.User{}

	hashedPass, err := HashPassword(req.Password)
	if err != nil {
		log.Println("Registration failed, reason=", err)
		return s, err
	}

	s.Username = req.Username
	s.Password = hashedPass
	s.Name = req.Name
	s.Role = req.Role
	s.Image = req.Image
	s.Active = models.UserStatusActive
	if req.Active != 0 {
		s.Active = req.Active
	}
	s.RegBy = req.RegBy
	s.TID = req.TID
	s.Created = time.Now()

	return s, nil
}

// func Update(req *models.User) {
// 	s.ID = req.ID
// 	s.Username = req.Username
// 	s.Name = req.Name
// 	s.Role = req.Role
// 	s.Image = req.Image
// 	s.TID = req.TID
// 	s.Active = req.Active
// 	s.RegBy = req.RegBy
// 	s.Created = req.Created
// }
