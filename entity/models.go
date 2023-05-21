package entity

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type PlanPara struct {
	CheckedSenIds *[]int  `json:"checkedSenIds"`
	Start         int64   `json:"start"`
	Stop          int64   `json:"stop"`
	Xmin          float32 `json:"xmin"`
	Xmax          float32 `json:"xmax"`
	Ymin          float32 `json:"ymin"`
	Ymax          float32 `json:"ymax"`
}

type NewSatDTO struct {
	Tle  string `json:"tle"`
	Name string `json:"name"`
}

type TleData struct {
	Line0 string `json:"line0"`
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
}

type SatDTO struct {
	SatName  string `json:"satName"`
	HexColor string `json:"hexColor"`
}

type NewSensorInDTO struct {
	SatId          string  `json:"satId"`
	Name           string  `json:"name"`
	Resolution     float32 `json:"resolution"`
	Width          float32 `json:"width"`
	RightSideAngle float32 `json:"rightSideAngle"`
	LeftSideAngle  float32 `json:"leftSideAngle"`
	ObserveAngle   float32 `json:"observeAngle"`
	InitAngle      float32 `json:"initAngle"`
	HexColor       string  `json:"hexColor"`
}

type SensorDTO struct {
	Name           string  `json:"name"`
	Resolution     float32 `json:"resolution"`
	Width          float32 `json:"width"`
	RightSideAngle float32 `json:"rightSideAngle"`
	LeftSideAngle  float32 `json:"leftSideAngle"`
	ObserveAngle   float32 `json:"observeAngle"`
	InitAngle      float32 `json:"initAngle"`
	HexColor       string  `json:"hexColor"`
}

type CurrentUserInfo struct {
	Id       int               `json:"id"`
	Name     string            `json:"name"`
	AdminId  int               `json:"adminId"`
	RoleId   int               `json:"roleId"`
	MenuList []PrivilegeMenuVO `json:"menuList"`
}

type PrivilegeMenuVO struct {
	Id  int    `json:"id"`
	PId int    `json:"pId"`
	Url string `json:"url"`
}

type PathUnit struct {
	SatId    string     `json:"satId"`
	SatName  string     `json:"satName"`
	SenName  string     `json:"senName"`
	HexColor string     `json:"hexColor"`
	Start    int64      `json:"start"`
	Stop     int64      `json:"stop"`
	PathGeo  *[]SenPath `json:"pathGeo"`
}

type SatSen struct {
	SatId    string    `json:"satId"`
	SenNames *[]string `json:"senNames"`
}

type SatItem struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	NoardId  string     `json:"noardId"`
	HexColor string     `json:"hexColor"`
	SenItems *[]SenItem `json:"senItems"`
}

type SenItem struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Resolution     float32 `json:"resolution"`
	Width          float32 `json:"width"`
	RightSideAngle float32 `json:"rightSideAngle"`
	LeftSideAngle  float32 `json:"leftSideAngle"`
	ObserveAngle   float32 `json:"observeAngle"`
	InitAngle      float32 `json:"initAngle"`
	HexColor       string  `json:"hexColor"`
}
