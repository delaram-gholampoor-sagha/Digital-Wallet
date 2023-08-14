package protocol

import (
	"context"
)

type HTTP interface {
	Start() error
	Shutdown(ctx context.Context) error
	Close() error
}

type Success struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Error struct {
	Message string `json:"message"`
}
