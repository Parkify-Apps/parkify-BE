package data

type User struct {
	UserID   int `gorm:"primaryKey;"`
	Nama     string
	Email    string `gorm:"type:varchar(30);"`
	Password string
	// Role     bool
}
