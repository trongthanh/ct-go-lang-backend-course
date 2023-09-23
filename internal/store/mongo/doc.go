package mongostore

import (
	"gosocial/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doc struct {
	Id        primitive.ObjectID `bson:"_id"`
	Version   int64              `bson:"version"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func NewDoc() Doc {
	docId := primitive.NewObjectID()
	return Doc{
		Id:        docId,
		Version:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type UserDoc struct {
	Doc         `bson:",inline"`
	entity.User `bson:",inline"`
}

func NewUserDoc(user entity.User) *UserDoc {
	return &UserDoc{
		Doc:  NewDoc(),
		User: user,
	}
}

type ProfileDoc struct {
	Doc            `bson:",inline"`
	entity.Profile `bson:",inline"`
}

func NewProfileDoc(profile entity.Profile) *ProfileDoc {
	return &ProfileDoc{
		Doc:     NewDoc(),
		Profile: profile,
	}
}

func (pd *ProfileDoc) ToProfile() entity.Profile {
	return pd.Profile
}
