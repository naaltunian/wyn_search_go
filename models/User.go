package UserModel

import (
	"time"

	"github.com/naaltunian/wyn-search-go/utils/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Email          string             `json:"email,omitempty" bson:"email,omitempty"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty"`
	LinkedIn       string             `json:"linkedIn,omitempty" bson:"linkedIn,omitempty"`
	GithubUsername string             `json:"githubUsername,omitempty" bson:"githubUsername,omitempty"`
	PersonalSite   string             `json:"personalSite,omitempty" bson:"personalSite,omitempty"`
	Bio            string             `json:"bio,omitempty" bson:"bio,omitempty"`
	DateCreated    time.Time          `json:"dateCreated,omitempty" bson:"dateCreated,omitempty"`
}

// UserModels is a slice of UserModel
type userModels []User

// Validate validates if the User model is valid
func (u *User) Validate() *errors.RestErr {
	// TODO: add further validations
	if u.Name == "" {
		return errors.NewBadRequestError("invalid name field")
	}
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email field")
	}
	if u.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	if u.GithubUsername == "" {
		return errors.NewBadRequestError("invalid github username field")
	}

	return nil
}
