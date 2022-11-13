// Code generated by go generate DO NOT EDIT.
// If you edit this file, write "// DO NOT OVERWRITE" on the first line.

package controllers

import usecases "github.com/rinoguchi/microblog/usecases/models"

func (getCommentsParams GetCommentsParams) ToUGetCommentsParams() usecases.UGetCommentsParams {
	return usecases.UGetCommentsParams{
		Query:     getCommentsParams.Query,
		Year:      getCommentsParams.Year,
		Yearmonth: getCommentsParams.Yearmonth,
	}
}

func FromUGetCommentsParams(uGetCommentsParams usecases.UGetCommentsParams) GetCommentsParams {
	return GetCommentsParams{
		Query:     uGetCommentsParams.Query,
		Year:      uGetCommentsParams.Year,
		Yearmonth: uGetCommentsParams.Yearmonth,
	}
}