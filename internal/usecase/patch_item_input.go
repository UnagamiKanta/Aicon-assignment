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
func (p *PatchItemInput) HasUpdates() bool {
	return p.Name != nil ||
		p.Brand != nil ||
		p.PurchasePrice != nil
}

func (p *PatchItemInput) ApplyTo(item *entity.Item) {
	if p.Name != nil {
		item.Name = *p.Name
	}
	if p.Brand != nil {
		item.Brand = *p.Brand
	}
	if p.PurchasePrice != nil {
		item.PurchasePrice = *p.PurchasePrice
	}
}
