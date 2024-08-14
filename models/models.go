package models

import "gorm.io/gorm"

type RoleType string

type User struct {
	gorm.Model
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Image        string `json:"image"`
	Boards       []Board
	BoardMembers []BoardMember
	CardMembers  []CardMember
	Comments     []Comment
	Attachments  []Attachment
}

type Board struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	CoverImg    string `json:"coverImg"`
	Visibility  bool   `json:"visibility"`
	UserID      uint   `json:"userId"`
	User        User   `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
	Members     []BoardMember
	Lists       []List
	Labels      []Label
}

type Card struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	CoverImg    string `json:"coverImg"`
	ListID      uint   `json:"listId"`
	List        List   `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
	Attachments []Attachment
	Comments    []Comment
	Labels      []CardLabel
	Members     []CardMember
}

type List struct {
	gorm.Model
	Title   string `json:"title"`
	BoardID uint   `json:"boardId"`
	Board   Board  `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
	Cards   []Card
}

type Label struct {
	gorm.Model
	Name    string `json:"name"`
	BgColor string `json:"bgColor"`
	BoardID uint   `json:"boardId"`
	Board   Board  `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
	Cards   []CardLabel
}

type BoardMember struct {
	gorm.Model
	UserID  uint     `json:"userId"`
	User    User     `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
	BoardID uint     `json:"boardId"`
	Board   Board    `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
	Role    RoleType `json:"role" gorm:"type:enum('Author', 'Member')"`
}

type CardLabel struct {
	gorm.Model
	LabelID uint  `json:"labelId"`
	Label   Label `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
	CardID  uint  `json:"cardId"`
	Card    Card  `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
}

type CardMember struct {
	gorm.Model
	UserID uint `json:"userId"`
	User   User `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
	CardID uint `json:"cardId"`
	Card   Card `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
}

type Attachment struct {
	gorm.Model
	Name   string `json:"name"`
	Url    string `json:"url"`
	UserID uint   `json:"userId"`
	User   User   `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
	CardID uint   `json:"cardId"`
	Card   Card   `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
}

type Comment struct {
	gorm.Model
	Body   string `json:"body"`
	UserID uint   `json:"userId"`
	User   User   `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
	CardID uint   `json:"cardId"`
	Card   Card   `gorm:"onUpdate:CASCADE,onDelete:CASCADE"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Board{}, &List{}, &Card{}, &Label{}, &BoardMember{}, &CardLabel{}, &CardMember{}, &Attachment{}, &Comment{})
}
