package entity

type Satellite struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	NoardId  string `json:"noardId" gorm:"type:varchar(255)"`
	HexColor string `json:"hexColor"`
}

type Sensor struct {
	Id             int     `json:"id" gorm:"primary_key"`
	SatNoardId     string  `json:"satNoardId" gorm:"type:varchar(255)"`
	SatName        string  `json:"satName" gorm:"type:varchar(255)"`
	Name           string  `json:"name" gorm:"type:varchar(255)"`
	Resolution     float32 `json:"resolution"`
	Width          float32 `json:"width"`
	RightSideAngle float32 `json:"rightSideAngle"`
	LeftSideAngle  float32 `json:"leftSideAngle"`
	ObserveAngle   float32 `json:"observeAngle"`
	InitAngle      float32 `json:"initAngle"`
	HexColor       string  `json:"hexColor"`
}

type SysUser struct {
	Id       int    `json:"id" gorm:"primary_key"`
	UserName string `json:"userName" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	AdminId  int    `json:"adminId"`
	RoleId   int    `json:"roleId"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
}

type Tle struct {
	Id         int    `json:"id" gorm:"primary_key"`
	SatNoardId string `json:"satNoardId" gorm:"type:varchar(255)"`
	Time       int64  `json:"time"`
	Line1      string `json:"line1" gorm:"type:varchar(255)"`
	Line2      string `json:"line2" gorm:"type:varchar(255)"`
}

type TleSite struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Site        string `json:"site" gorm:"type:varchar(255)"`
	Url         string `json:"url" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
}

type Track struct {
	Id         int     `json:"id" gorm:"primary_key"`
	TimeOffset int64   `json:"timeOffset"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
	Z          float32 `json:"z"`
	Vx         float32 `json:"vx"`
	Vy         float32 `json:"vy"`
	Vz         float32 `json:"vz"`
	Lon        float32 `json:"lon"`
	Lat        float32 `json:"lat"`
	Alt        float32 `json:"alt"`
}

type TrackInfo struct {
	Id         int    `json:"id" gorm:"primary_key"`
	SatNoardId string `json:"satNoardId" gorm:"type:varchar(255)"`
	SatName    string `json:"satName" gorm:"type:varchar(255)"`
	StartTime  int64  `json:"startTime"`
}

type PathInfo struct {
	Id         int    `json:"id" gorm:"primary_key"`
	SatNoardId string `json:"satNoardId" gorm:"type:varchar(255)"`
	SatName    string `json:"satName" gorm:"type:varchar(255)"`
	SenName    string `json:"senName" gorm:"type:varchar(255)"`
	StartTime  int64  `json:"startTime"`
}

type SenPath struct {
	Id         int     `json:"id" gorm:"primary_key"`
	TimeOffset int64   `json:"timeOffset"`
	Lon1       float32 `json:"lon1"`
	Lat1       float32 `json:"lat1"`
	Lon2       float32 `json:"lon2"`
	Lat2       float32 `json:"lat2"`
}
