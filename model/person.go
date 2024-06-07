package model

type Person struct {
	ID   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (p *Person) TableName() string {
	return "people"
}
