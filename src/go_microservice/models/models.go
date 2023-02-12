package models

type Item struct {
	Id          int     `json:"id"`
	Description string  `json:"Description"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type Order struct {
	Id       string  `json:"id,omitempty"`
	Status   string  `json:"status,omitempty"`
	Items    []Item  `json:"items,omitempty"`
	Total    float32 `json:"total,omitempty"`
	Currency string  `json:"currencyUnit,omitempty"`
}

type OrderQuery struct {
	Id       string   `json:"id,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Items    *[]Item  `json:"items,omitempty"`
	Total    *float32 `json:"total,omitempty"`
	Currency *string  `json:"currencyUnit,omitempty"`
}

type SearchOptions struct {
	OrderASC  *string `json:"orderAsc,omitempty"`
	OrderDESC *string `json:"orderDesc,omitempty"` // ASC or DESC
	Limit     *int    `json:"limit,omitempty"`
}

type SearchValues struct {
	Status   *string `json:"status,omitempty"`
	Currency *string `json:"currencyUnit,omitempty"`
}
