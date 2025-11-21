package usecase

import (
	"Aicon-assignment/internal/domain/entity"
)

// UpdateItemInput はアイテムの部分更新用の入力構造体
// nil の場合はそのフィールドを更新しないことを意味する
type PatchItemInput struct {
	Name          *string `json:"name,omitempty"`
	Brand         *string `json:"brand,omitempty"`
	PurchasePrice *int    `json:"purchase_price,omitempty"`
}

// HasUpdates は更新するフィールドが1つでも存在するかチェック
func (u *PatchItemInput) HasUpdates() bool {
	return u.Name != nil ||
		u.Brand != nil ||
		u.PurchasePrice != nil
}

func (u *PatchItemInput) ApplyTo(item *entity.Item) {
	if u.Name != nil {
		item.Name = *u.Name
	}
	if u.Brand != nil {
		item.Brand = *u.Brand
	}
	if u.PurchasePrice != nil {
		item.PurchasePrice = *u.PurchasePrice
	}
}
