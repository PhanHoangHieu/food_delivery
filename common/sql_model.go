package common

import "time"

type SQLModel struct {
	Id int `json:"id"`
	//FakeId    *UID       `json:"id" gorm:"_"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty",gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty",gorm:"updated_at"`
}

//func (m *SQLModel) GenUID(dbType int) {
//	uid := NewUID(unit32(m.Id), dbType, 1)
//	m.FakeId = &uid
//}