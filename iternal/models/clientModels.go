package models

//
//type MessageFromUser struct {
//	gorm.Model
//	Type   string `json:"type"`
//	IdUser string `json:"idUser"`
//
//	TankX         float32 `json:"tankX"`
//	TankY         float32 `json:"tankY"`
//	BulletX       float32 `json:"bulletX"`
//	BulletY       float32 `json:"bulletY"`
//	AngleTank     float64 `json:"angle"`
//	AnimateStatus string  `json:"animateStatus"`
//	IsAlive       bool    `json:"isAlive"`
//}

type UserData struct {
	//gorm.Model
	IdUser        string  `json:"idUser"`
	TankX         float32 `json:"tankX"`
	TankY         float32 `json:"tankY"`
	BulletX       float32 `json:"bulletX"`
	BulletY       float32 `json:"bulletY"`
	AngleTank     float64 `json:"angle"`
	AnimateStatus string  `json:"animateStatus"`
	IsAlive       bool    `json:"isAlive"`
}

type Coordinates struct {
	//gorm.Model
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
