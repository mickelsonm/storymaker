package stories

type Collages []Collage
type Collage struct {
	Title   string `bson:"title" json:"title" xml:"title"`
	Content string `bson:"content" json:"content" xml:"content"`
}
