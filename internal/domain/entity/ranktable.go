package entity

type RankTable struct {
	Id      string
	Name    string
	Public  bool
	Attrs   []Attribute
	Entries []Entry
}
