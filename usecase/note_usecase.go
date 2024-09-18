package usecase

import (
	"github.com/1206yaya/go-echo-jwt-noteapp-api/model"
	"github.com/1206yaya/go-echo-jwt-noteapp-api/repository"
)

type INoteUsecase interface {
	GetAllNotes(userId uint) ([]model.NoteResponse, error)
	GetNoteById(userId uint, noteId uint) (model.NoteResponse, error)
	CreateNote(note model.Note) (model.NoteResponse, error)
	UpdateNote(note model.Note, userId uint, noteId uint) (model.NoteResponse, error)
	DeleteNote(userId uint, noteId uint) error
}

type noteUsecase struct {
	repository repository.INoteRepository
}

func NewNoteUsecase(repository repository.INoteRepository) INoteUsecase {
	return &noteUsecase{repository}
}

func (usecase *noteUsecase) GetAllNotes(userId uint) ([]model.NoteResponse, error) {
	notes := []model.Note{}
	if err := usecase.repository.GetAllNotes(&notes, userId); err != nil {
		return nil, err
	}
	resNotes := []model.NoteResponse{}
	for _, v := range notes {
		t := model.NoteResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resNotes = append(resNotes, t)
	}
	return resNotes, nil
}

func (usecase *noteUsecase) GetNoteById(userId uint, noteId uint) (model.NoteResponse, error) {
	note := model.Note{}
	if err := usecase.repository.GetNoteById(&note, userId, noteId); err != nil {
		return model.NoteResponse{}, err
	}
	resNote := model.NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
	return resNote, nil
}

func (usecase *noteUsecase) CreateNote(note model.Note) (model.NoteResponse, error) {
	if err := usecase.repository.CreateNote(&note); err != nil {
		return model.NoteResponse{}, err
	}
	resNote := model.NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
	return resNote, nil
}

func (usecase *noteUsecase) UpdateNote(note model.Note, userId uint, noteId uint) (model.NoteResponse, error) {

	if err := usecase.repository.UpdateNote(&note, userId, noteId); err != nil {
		return model.NoteResponse{}, err
	}
	resNote := model.NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
	return resNote, nil
}

func (usecase *noteUsecase) DeleteNote(userId uint, noteId uint) error {
	if err := usecase.repository.DeleteNote(userId, noteId); err != nil {
		return err
	}
	return nil
}
