package usecase

// UpdateItemInput はアイテムの部分更新用の入力構造体
// nil の場合はそのフィールドを更新しないことを意味する
type UpdateItemInput struct {
	Name          *string `json:"name,omitempty"`
	Brand         *string `json:"brand,omitempty"`
	PurchasePrice *int    `json:"purchase_price,omitempty"`
}

// HasUpdates は更新するフィールドが1つでも存在するかチェック
func (u *UpdateItemInput) HasUpdates() bool {
	return u.Name != nil ||
		u.Brand != nil ||
		u.PurchasePrice != nil
}

