package models

type UserData struct {
	IdUser        string  `json:"idUser"`
	TankX         float32 `json:"tankX"`
	TankY         float32 `json:"tankY"`
	BulletX       float32 `json:"bulletX"`
	BulletY       float32 `json:"bulletY"`
	AngleTank     float32 `json:"angle"`
	AnimateStatus string  `json:"animateStatus"`
	IsAlive       bool    `json:"isAlive"`
}

type Coordinates struct {
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
	Id string  `json:"id"`
}
