package model

import "time"

type FileRecord struct {
	Name     string    `json:"name"`
	Modified time.Time `json:"modified"`
	Path     string    `json:"path"`
}
