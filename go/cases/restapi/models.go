package restapi

import "time"

type TODO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TODORequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
