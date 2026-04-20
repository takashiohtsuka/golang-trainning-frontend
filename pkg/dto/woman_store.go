package dto

import fvo "golang-trainning-frontend/pkg/dto/valueobject"

// WomanStore は Woman コンテキストで使用する Store の軽量 DTO
type WomanStore struct {
	ID           uint               `json:"id"`
	Name         string             `json:"name"`
	BusinessType fvo.BusinessType   `json:"business_type"`
}
