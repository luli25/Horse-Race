package race

import (
	horse "horse-race/internal/model"
)

type Update struct {
	Horses []horse.Horse `json:"horses"`
	Winner string        `json:"winner,omitempty"`
}
