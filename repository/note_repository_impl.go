package repository

import (
	"errors"
	"golang-fiber-crud/data/request"
	"golang-fiber-crud/helper"
	"golang-fiber-crud/model"

	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	DB *gorm.DB
}

func NewNoteRepositoryImpl(Db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{DB: Db}
}

func (n *NoteRepositoryImpl) Delete(noteId int) {
	var note model.Note
	result := n.DB.Where("id = ?", noteId).Delete(&note)
	helper.ErrorPanic(result.Error)
}

func (n *NoteRepositoryImpl) FindAll() []model.Note {
	var notes []model.Note
	result := n.DB.Find(&notes)
	helper.ErrorPanic(result.Error)
	return notes
}

func (n *NoteRepositoryImpl) FindById(noteId int) (model.Note, error) {
	var note model.Note
	result := n.DB.Find(&note, noteId)
	if result != nil {
		return note, nil
	} else {
		return note, errors.New("Note is not found")
	}
}

func (n *NoteRepositoryImpl) Save(note model.Note) {
	result := n.DB.Create(&note)
	helper.ErrorPanic(result.Error)
}

func (n *NoteRepositoryImpl) Update(note model.Note) {
	var updatedNote = request.UpdateNoteRequest{
		Id:      note.Id,
		Content: note.Content,
	}
	result := n.DB.Model(&note).Updates(updatedNote)
	helper.ErrorPanic(result.Error)
}
