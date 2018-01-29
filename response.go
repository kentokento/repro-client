package repro

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	errorCodeMissingPayload = 1001 // ペイロードの形式に誤りがあります
	errorCodeNotRegistered  = 1002 // ユーザーIDが登録されていません
	errorCodeInvalidValue   = 1003 // ユーザープロフィールの key, type, value が適切に設定されていません
)

type ReproError struct {
	Status string `json:"status"`
	Errors struct {
		Code     json.Number `json:"code"`
		Messages []string    `json:"messages"`
	} `json:"error"`
}

func (r *ReproError) Error() string {
	if len(r.Errors.Messages) == 0 {
		return ""
	}
	return r.Errors.Messages[0]
}

func (r *ReproError) Code() int {
	code, err := r.Errors.Code.Int64()
	if err != nil {
		return 0
	}
	return int(code)
}

type reproResponse struct {
	statusCode int

	// X-RateLimit-Limit     単位時間あたりのアクセス上限
	// X-RateLimit-Remaining アクセスできる残り回数
	// X-RateLimit-Reset     アクセス数がリセットされる時刻(unixtime)
	// Retry-After           再実行可能になるまでの秒数
	limit, remaining, reset, retryAfter string
}

func NewReproResponse(code int, header http.Header) *reproResponse {
	return &reproResponse{
		statusCode: code,
		limit:      header.Get("X-RateLimit-Limit"),
		remaining:  header.Get("X-RateLimit-Remaining"),
		reset:      header.Get("X-RateLimit-Reset"),
		retryAfter: header.Get("Retry-After"),
	}
}

func (r *reproResponse) IsOK() bool {
	return r.statusCode == http.StatusAccepted
}

func (r *reproResponse) IsBadRequest() bool {
	return r.statusCode == http.StatusBadRequest
}

func (r *reproResponse) IsTooManyRequests() bool {
	return r.statusCode == http.StatusTooManyRequests
}

func (r *reproResponse) StatusCode() int {
	return r.statusCode
}

func (r *reproResponse) Limit() int64 {
	v, err := strconv.ParseInt(r.limit, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func (r *reproResponse) Remaining() int64 {
	v, err := strconv.ParseInt(r.remaining, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func (r *reproResponse) Reset() int64 {
	v, err := strconv.ParseInt(r.reset, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func (r *reproResponse) RetryAfter() int64 {
	v, err := strconv.ParseInt(r.retryAfter, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

type ReproResponse interface {
	IsOK() bool
	IsBadRequest() bool
	IsTooManyRequests() bool
	StatusCode() int
	Limit() int64
	Remaining() int64
	Reset() int64
	RetryAfter() int64
}
