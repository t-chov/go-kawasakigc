package db

type Garbage struct {
	GarbageType string
	DetailType  string
	Description string
	Url         string
}

type GarbageDB map[string]Garbage

func (db GarbageDB) Find(name string) (*Garbage, bool) {
	garbage, ok := db[name]
	if !ok {
		return nil, false
	}
	return &garbage, true
}

func InitDb() GarbageDB {
	db := map[string]Garbage{
		"アームスタンド": {
			GarbageType: "小物金属",
			DetailType:  "粗大ごみ",
			Description: "最長辺30cm以上のものは粗大ごみとして出してください。",
			Url:         "",
		},
		"モニタースタンド": {
			GarbageType: "小物金属",
			DetailType:  "粗大ごみ",
			Description: "最長辺30cm以上のものは粗大ごみとして出してください。",
			Url:         "",
		},
		"アームバンド": {
			GarbageType: "普通ごみ",
			DetailType:  "",
			Description: "",
			Url:         "",
		},
		"IH調理器": {
			GarbageType: "小物金属",
			DetailType:  "粗大ごみ",
			Description: "最長辺30cm以上のものは粗大ごみとして出してください。/長辺30cm未満で、30cm×15cmの回収ボックスの投入口に入るものは小型家電としても出すことができます。",
			Url:         "http://www.city.kawasaki.jp/kurashi/category/24-1-23-1-1-6-9-0-0-0.html",
		},
		"電磁調理器": {
			GarbageType: "小物金属",
			DetailType:  "粗大ごみ",
			Description: "最長辺30cm以上のものは粗大ごみとして出してください。/長辺30cm未満で、30cm×15cmの回収ボックスの投入口に入るものは小型家電としても出すことができます。",
			Url:         "http://www.city.kawasaki.jp/kurashi/category/24-1-23-1-1-6-9-0-0-0.html",
		},
		"IHクッキングヒーター": {
			GarbageType: "小物金属",
			DetailType:  "粗大ごみ",
			Description: "最長辺30cm以上のものは粗大ごみとして出してください。/長辺30cm未満で、30cm×15cmの回収ボックスの投入口に入るものは小型家電としても出すことができます。",
			Url:         "http://www.city.kawasaki.jp/kurashi/category/24-1-23-1-1-6-9-0-0-0.html",
		},
		"電気調理器": {
			GarbageType: "小物金属",
			DetailType:  "粗大ごみ",
			Description: "最長辺30cm以上のものは粗大ごみとして出してください。/長辺30cm未満で、30cm×15cmの回収ボックスの投入口に入るものは小型家電としても出すことができます。",
			Url:         "http://www.city.kawasaki.jp/kurashi/category/24-1-23-1-1-6-9-0-0-0.html",
		},
	}
	return db
}
