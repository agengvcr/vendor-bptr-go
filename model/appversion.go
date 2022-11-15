package model

type AppVersion struct {
	Version   string `gorm:"column:latestappversion"`
	ChangeLog string `gorm:"column:latestappchangelog"`
	ApkUrl    string `gorm:"column:latestappapkurl"`
}
