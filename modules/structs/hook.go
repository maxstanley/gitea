// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package structs

import (
	"errors"
	"time"
)

// ErrInvalidReceiveHook FIXME
var ErrInvalidReceiveHook = errors.New("Invalid JSON payload received over webhook")

// Hook a hook is a web hook when one repository changed
type Hook struct {
	ID                  int64             `json:"id"`
	Type                string            `json:"type"`
	BranchFilter        string            `json:"branch_filter"`
	URL                 string            `json:"-"`
	Config              map[string]string `json:"config"`
	Events              []string          `json:"events"`
	AuthorizationHeader string            `json:"authorization_header"`
	Active              bool              `json:"active"`
	// swagger:strfmt date-time
	Updated time.Time `json:"updated_at"`
	// swagger:strfmt date-time
	Created time.Time `json:"created_at"`
}

// HookList represents a list of API hook.
type HookList []*Hook

// CreateHookOptionConfig has all config options in it
// required are "content_type" and "url" Required
type CreateHookOptionConfig map[string]string

// CreateHookOption options when create a hook
type CreateHookOption struct {
	// required: true
	// enum: dingtalk,discord,gitea,gogs,msteams,slack,telegram,feishu,wechatwork,packagist
	Type string `json:"type" binding:"Required"`
	// required: true
	Config              CreateHookOptionConfig `json:"config" binding:"Required"`
	Events              []string               `json:"events"`
	BranchFilter        string                 `json:"branch_filter" binding:"GlobPattern"`
	AuthorizationHeader string                 `json:"authorization_header"`
	// default: false
	Active bool `json:"active"`
}

// EditHookOption options when modify one hook
type EditHookOption struct {
	Config              map[string]string `json:"config"`
	Events              []string          `json:"events"`
	BranchFilter        string            `json:"branch_filter" binding:"GlobPattern"`
	AuthorizationHeader string            `json:"authorization_header"`
	Active              *bool             `json:"active"`
}

// Payloader payload is some part of one hook
type Payloader interface {
	JSONPayload() ([]byte, error)
}
