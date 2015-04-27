package stories

import (
	"errors"

	"github.com/mickelsonm/storymaker/helpers/database"
	"gopkg.in/mgo.v2/bson"
)

type Stories []Story
type Story struct {
	ID       bson.ObjectId `bson:"_id" json:"id" xml:"id,attr"`
	Title    string        `bson:"title" json:"title" xml:"title"`
	Content  string        `bson:"content" json:"content" xml:"content"`
	Collages Collages      `bson:"collages" json:"collages,omitempty" xml:"collages"`
}

func GetAllStories() (stories Stories, err error) {
	session := database.MongoSession.Copy()
	defer session.Close()

	stories = make(Stories, 0)
	err = session.DB(database.MongoDatabase).C(database.StoriesCollectionName).Find(bson.M{}).All(&stories)

	return
}

func (s *Story) Get() error {
	if s.ID.Hex() == "" {
		return errors.New("Invalid story ID")
	}

	session := database.MongoSession.Copy()
	defer session.Close()

	err := session.DB(database.MongoDatabase).C(database.StoriesCollectionName).Find(bson.M{
		"_id": s.ID,
	}).One(&s)

	return err
}

func init() {
	database.InitMongo()
}
