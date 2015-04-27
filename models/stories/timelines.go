package stories

import (
	"errors"

	"github.com/mickelsonm/storymaker/helpers/database"
	"gopkg.in/mgo.v2/bson"
)

type Timelines []Timeline
type Timeline struct {
	ID      bson.ObjectId `bson:"_id" json:"id" xml:"id,attr"`
	Author  string        `bson:"author" json:"author" xml:"author"`
	Title   string        `bson:"title" json:"title" xml:"title"`
	Stories Stories       `bson:"stories" json:"stories,omitempty" xml:"stories"`
}

func GetAllTimelines() (timelines Timelines, err error) {
	session := database.MongoSession.Copy()
	defer session.Close()

	timelines = make(Timelines, 0)
	err = session.DB(database.MongoDatabase).C(database.TimelinesCollectionName).Find(bson.M{}).All(&timelines)

	return
}

func (t *Timeline) Get() error {
	if t.ID.Hex() == "" {
		return errors.New("Invalid timeline ID")
	}

	session := database.MongoSession.Copy()
	defer session.Close()

	err := session.DB(database.MongoDatabase).C(database.TimelinesCollectionName).Find(bson.M{
		"_id": t.ID,
	}).One(&t)

	return err
}

func init() {
	database.InitMongo()
}
