package repository

import (
	"fmt"

	"github.com/1206yaya/go-echo-jwt-noteapp-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type INoteRepository interface {
	GetAllNotes(notes *[]model.Note, userId uint) error
	GetNoteById(note *model.Note, userId uint, noteId uint) error
	CreateNote(note *model.Note) error
	UpdateNote(note *model.Note, userId uint, noteId uint) error
	DeleteNote(userId uint, noteId uint) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) INoteRepository {
	return &noteRepository{db}
}

func (tr *noteRepository) GetAllNotes(notes *[]model.Note, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(notes).Error; err != nil {
		return err
	}
	return nil
}

func (tr *noteRepository) GetNoteById(note *model.Note, userId uint, noteId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(note, noteId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *noteRepository) CreateNote(note *model.Note) error {
	if err := tr.db.Create(note).Error; err != nil {
		return err
	}
	return nil
}

func (tr *noteRepository) UpdateNote(note *model.Note, userId uint, noteId uint) error {
	// result := tr.db.Model(note).Clauses(clause.Returning{}).Where("id=? AND user_id=?", noteId, userId).Update("title", note.Title)
	result := tr.db.Model(&model.Note{}).Clauses(clause.Returning{}).Where("id=? AND user_id=?", noteId, userId).
		Updates(map[string]interface{}{
			"title": note.Title,
			"body":  note.Body,
		})
		
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *noteRepository) DeleteNote(userId uint, noteId uint) error {
	result := tr.db.Where("id=? AND user_id=?", noteId, userId).Delete(&model.Note{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
