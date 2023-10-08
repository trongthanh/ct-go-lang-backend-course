package mongostore

import (
	"gosocial/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doc struct {
	DocId     primitive.ObjectID `bson:"_id"`
	Version   int64              `bson:"version"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func NewDoc() Doc {
	docId := primitive.NewObjectID()
	return Doc{
		DocId:     docId,
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
	pd.Profile.Id = pd.DocId.Hex()
	return pd.Profile
}

type PostDoc struct {
	DocId       primitive.ObjectID `bson:"_id"`
	Version     int64              `bson:"version"`
	entity.Post `bson:",inline"`
}

func NewPostDoc(post entity.Post) *PostDoc {
	doc := NewDoc()
	post.CreatedAt = doc.CreatedAt
	post.UpdatedAt = doc.UpdatedAt
	return &PostDoc{
		DocId:   doc.DocId,
		Version: doc.Version,
		Post:    post,
	}
}

func (pd *PostDoc) ToPost() entity.Post {
	pd.Post.Id = pd.DocId.Hex()
	return pd.Post
}
