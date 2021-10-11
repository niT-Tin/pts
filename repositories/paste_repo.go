package repositories

import (
	"gorm.io/gorm"
	"pasteProject/models"
)

type IPasteRepo interface {
	CreateP(paste *models.Paste) error
	SelectPS() []models.Paste
	DeleteP(posterName string) error
	DeletePS() error
}

type PasteRepo struct {
	Db *gorm.DB
}

func NewPasteRepo(db *gorm.DB) IPasteRepo {
	return &PasteRepo{
		Db: db,
	}
}

func (p *PasteRepo) CreateP(paste *models.Paste) error {
	return NoRepeat("create paste failed", "paste_repo", p.Db.Create(paste))
}

func (p *PasteRepo) SelectPS() []models.Paste {
	var ps []models.Paste
	_ = NoRepeat("select whole paste thing failed", "paste_repo", p.Db.Find(&ps))
	return ps
}

func (p *PasteRepo) DeleteP(posterName string) error {
	return NoRepeat("delete paste failed", "paste_repo", p.Db.Where("poster = ?", posterName).Delete(&models.Paste{}))
}

func (p *PasteRepo) DeletePS() error {
	return NoRepeat("delete paste failed", "paste_repo", p.Db.Exec("DELETE FROM pastes"))
}
