package service

import (
	"golang-fiber-crud/data/request"
	"golang-fiber-crud/data/response"
	"golang-fiber-crud/helper"
	"golang-fiber-crud/model"
	"golang-fiber-crud/repository"

	"github.com/go-playground/validator/v10"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	validate       *validator.Validate
}

func NewNoteServiceImpl(noteRepository repository.NoteRepository, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{NoteRepository: noteRepository, validate: validate}
}

func (n *NoteServiceImpl) Create(note request.CreateNoteRequest) {
	err := n.validate.Struct(note)
	helper.ErrorPanic(err)
	noteModel := model.Note{
		Content: note.Content,
	}
	n.NoteRepository.Save(noteModel)
}

func (n *NoteServiceImpl) Update(note request.UpdateNoteRequest) {
	noteData, err := n.NoteRepository.FindById(note.Id)
	helper.ErrorPanic(err)
	noteData.Content = note.Content
	n.NoteRepository.Update(noteData)

}

func (n *NoteServiceImpl) Delete(noteId int) {
	n.NoteRepository.Delete(noteId)
}

func (n *NoteServiceImpl) FindById(noteId int) response.NoteResponse {
	noteData, err := n.NoteRepository.FindById(noteId)
	helper.ErrorPanic(err)
	noteResponse := response.NoteResponse{
		Id:      noteData.Id,
		Content: noteData.Content,
	}

	return noteResponse
}

func (n *NoteServiceImpl) FindAll() []response.NoteResponse {

	result := n.NoteRepository.FindAll()
	var notes []response.NoteResponse

	for _, value := range result {
		note := response.NoteResponse{
			Id:      value.Id,
			Content: value.Content,
		}
		notes = append(notes, note)
	}
	return notes
}
