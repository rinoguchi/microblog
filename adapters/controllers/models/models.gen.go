// Package controllers provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package controllers

import (
	"time"
)

// Comment defines model for Comment.
type Comment struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id        *int64     `json:"id,omitempty"`
	Text      string     `json:"text"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// CommonProps defines model for CommonProps.
type CommonProps struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id        *int64     `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// NewComment defines model for NewComment.
type NewComment struct {
	Text string `json:"text"`
}

// AddCommentJSONBody defines parameters for AddComment.
type AddCommentJSONBody = NewComment

// AddCommentJSONRequestBody defines body for AddComment for application/json ContentType.
type AddCommentJSONRequestBody = AddCommentJSONBody