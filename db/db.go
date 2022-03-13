package db

import "github.com/gocarina/gocsv"

type Garbage struct {
	Name        string `csv:"Name"`
	GarbageType string `csv:"GarbageType"`
	DetailType  string `csv:"DetailType"`
	Description string `csv:"Description"`
	Url         string `csv:"Url"`
}

type GarbageDB map[string]Garbage

func (db GarbageDB) Find(name string) (*Garbage, bool) {
	garbage, ok := db[name]
	if !ok {
		return nil, false
	}
	return &garbage, true
}

func InitDb(csvBytes []byte) (*GarbageDB, error) {
	garbages := []*Garbage{}
	if err := gocsv.UnmarshalBytes(csvBytes, &garbages); err != nil {
		return nil, err
	}

	db := make(GarbageDB)

	for _, garbage := range garbages {
		db[garbage.Name] = *garbage
	}

	return &db, nil
}
