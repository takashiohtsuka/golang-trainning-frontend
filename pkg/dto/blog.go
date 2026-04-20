package dto

import (
	"time"

	"golang-trainning-frontend/pkg/collection"
)

type Blog struct {
	ID          uint                      `json:"id"`
	WomanID     uint                      `json:"woman_id"`
	Title       string                    `json:"title"`
	Body        *string                   `json:"body"`
	IsPublished bool                      `json:"is_published"`
	Photos      collection.Collection[Photo] `json:"photos"`
	CreatedAt   *time.Time                `json:"created_at"`
	UpdatedAt   *time.Time                `json:"updated_at"`
	DeletedAt   *time.Time                `json:"deleted_at"`
}

func (b *Blog) IsNil() bool                                { return b.ID == 0 }
func (b *Blog) GetID() uint                                { return b.ID }
func (b *Blog) GetWomanID() uint                           { return b.WomanID }
func (b *Blog) GetTitle() string                           { return b.Title }
func (b *Blog) GetBody() *string                           { return b.Body }
func (b *Blog) GetIsPublished() bool                       { return b.IsPublished }
func (b *Blog) GetPhotos() collection.Collection[Photo] { return b.Photos }

type NilBlog struct{}

func (n *NilBlog) IsNil() bool                              { return true }
func (n *NilBlog) GetID() uint                              { return 0 }
func (n *NilBlog) GetWomanID() uint                         { return 0 }
func (n *NilBlog) GetTitle() string                         { return "" }
func (n *NilBlog) GetBody() *string                         { return nil }
func (n *NilBlog) GetIsPublished() bool                     { return false }
func (n *NilBlog) GetPhotos() collection.Collection[Photo] {
	return collection.NewCollection[Photo](nil)
}
