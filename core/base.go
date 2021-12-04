package core

import (
	"context"

	"github.com/canbefree/magazine/core/repo/repo_model"
)

type BaseInventoryIFace interface {
	TryTran(ctx context.Context, iid uint64)
	SubmitTran()
}

// 分包操作
type BaseInventory struct {
	Model *repo_model.MInventory `json:"model,omitempty"`
}

// TODO lock
