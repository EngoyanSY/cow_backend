package models

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Cow struct {
	ID        uint      `gorm:"primaryKey" example:"1"` // ID коровы
	CreatedAt time.Time `example:"2007-01-01"`          // Время создания коровы в базе данных

	Farm   *Farm `json:"-"`
	FarmID *uint `example:"1"` // ID фермы, которой корова принадлежит

	FarmGroup   Farm `json:"-"`
	FarmGroupId uint `example:"1"` // ID хозяйства, которому корова принадлежит

	Holding   *Farm
	HoldingID *uint

	Breed   Breed `json:"-"`
	BreedId uint  `example:"1"` // ID породы коровы

	Sex   Sex  `json:"-"`
	SexId uint `example:"1"` // ID пола коровы

	Events []Event `json:"-"`

	GradeRegion   *Grade `json:"-"`
	GradeRegionId *uint  `example:"1"` // оценка по региону

	GradeHoz   *Grade `json:"-"`
	GradeHozId *uint  `example:"1"` // оценка по хозяйству

	FatherSelecs *uint64 // ID коровы отца коровы

	MotherSelecs *uint64 // ID коровы матери коровы

	// CreatedBy   *User `json:"-"` // пользователь, создавший корову
	// CreatedByID *uint `example:"1"`

	Genetic   *Genetic
	Exterior  *Exterior
	Lactation []Lactation `json:"-"`

	IdentificationNumber *string // он все-таки есть! это какой-то не российский номер коровы
	InventoryNumber      *string `example:"1213321"`    // Инвентарный номер коровы
	SelecsNumber         *uint64 `example:"98989"`      // Селекс номер коровы
	RSHNNumber           *string `example:"1323323232"` // РСХН номер коровы
	Name                 string  `example:"Дима"`       // Кличка коровы

	// Exterior                float64  `example:"3.14"` // Оценка экстерьера коровы, будет переделано в ID экстерьера коровы
	InbrindingCoeffByFamily *float64 `example:"3.14"` // Коэф. инбриндинга по роду

	Approved    int       `example:"1"` // Целое число, что-то для админов, чтобы подтверждать коров
	BirthDate   DateOnly  // День рождения
	DepartDate  *DateOnly // День отбытия из коровника
	DeathDate   *DateOnly // Дата смерти
	BirkingDate *DateOnly // Дата перебирковки

	// Новые поля
	PreviousHoz   *Farm   `json:"-"`
	PreviousHozId *uint   // ID предыдущего хозяйства, когда корову продают, она переходит к новому владельцу и становится "новой коровой"
	BirthHoz      *Farm   `json:"-"`
	BirthHozId    *uint   // ID хозяйства рождения
	BirthMethod   *string // способ зачатия: клон, эмбрион, искусственное осеменени, естественное осеменение

	PreviousInventoryNumber *string `json:"-"` // Одна и та же реальная корова имеет разные инвент. номера, это указатель на эту же корову в другом хоз-ве с другим инв. номером

	Documents []Document `json:"-"` // документы коровы
}

type Document struct {
	ID    uint
	CowID uint
	Path  string
}

func (c *Cow) BeforeCreate(tx *gorm.DB) error {
	if c.RSHNNumber == nil {
		c.RSHNNumber = new(string)
		if c.SelecsNumber == nil {
			return errors.New("нет ни селекса ни РСХН")
		}
		*c.RSHNNumber = "!" + strconv.FormatUint(uint64(*c.SelecsNumber), 10)
	}
	return nil
}
