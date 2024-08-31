package model

import (
	"time"
)

type Account struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    string `gorm:"unique;not null"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"default:null"`
	Password  string `gorm:"not null"`
	IsBot     bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Events    []Event `gorm:"many2many:account_events"`
	Pays      []Pay   `gorm:"many2many:account_pays"`
}

type Event struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Pays        []Pay     `gorm:"foreignKey:EventID"`
	Accounts    []Account `gorm:"many2many:accounts_events"`
}

type Pay struct {
	ID         uint `gorm:"primaryKey"`
	PaidUserID uint `gorm:"not null"`
	EventID    uint `gorm:"not null"`
	Amount     uint `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PaidUser   Account   `gorm:"foreignKey:PaidUserID"`
	Event      Event     `gorm:"foreignKey:EventID"`
	Accounts   []Account `gorm:"many2many:account_pays"`
}

type AccountEvent struct {
	ID          uint `gorm:"primaryKey"`
	AccountID   uint `gorm:"not null"`
	EventID     uint `gorm:"not null"`
	AuthorityID uint `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Account     Account   `gorm:"foreignKey:AccountID"`
	Event       Event     `gorm:"foreignKey:EventID"`
	Authority   Authority `gorm:"foreignKey:AuthorityID"`
}

type Authority struct {
	ID         uint `gorm:"primaryKey"`
	AddPays    bool `gorm:"not null;default:true"`
	EditPays   bool `gorm:"not null;default:true"`
	DeletePays bool `gorm:"not null;default:true"`
	AddUser    bool `gorm:"not null;default:true"`
	EditEvent  bool `gorm:"not null;default:true"`
	DeleteUser bool `gorm:"not null;default:true"`
}

type Friend struct {
	ID                uint `gorm:"primaryKey"`
	SendAccountID     uint `gorm:"not null"`
	ReceivedAccountID uint `gorm:"not null"`
	SendAt            time.Time
	AcceptedAt        time.Time
	SendAccount       Account `gorm:"foreignKey:SendAccountID"`
	ReceivedAccount   Account `gorm:"foreignKey:ReceivedAccountID"`
}

type AccountPay struct {
	ID        uint `gorm:"primaryKey"`
	AccountID uint `gorm:"not null"`
	PayID     uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Account   Account `gorm:"foreignKey:AccountID"`
	Pay       Pay     `gorm:"foreignKey:PayID"`
}

func (ae *AccountEvent) TableName() string {
	return "accounts_events"
}
func (ap *AccountPay) TableName() string {
	return "accounts_pays"
}
