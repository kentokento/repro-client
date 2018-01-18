package repro

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	typeString   = "string"
	typeInt      = "int"
	typeDecimal  = "decimal"
	typeDatetime = "datetime"

	datetimeFormat = "2006-01-02T15:04:05-0700"
)

type UserProfiles struct {
	UserID       string  `json:"user_id"`
	UserProfiles []param `json:"user_profiles"`
}

type param struct {
	Key   string      `json:"key"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func NewUserProfiles(userID interface{}) UserProfiles {
	return UserProfiles{UserID: fmt.Sprintf("%v", userID)}
}

func (r *UserProfiles) add(key, typ string, val interface{}) {
	r.UserProfiles = append(r.UserProfiles, param{Key: key, Type: typ, Value: val})
}

func (r *UserProfiles) Add(key string, val interface{}) {
	switch t := val.(type) {
	case string:
		r.AddString(key, t)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		r.AddInt(key, t)
	case float32, float64:
		r.AddDecimal(key, t)
	case time.Time:
		r.AddDatetime(key, t)
	default:
		// invalid type
	}
}

func (r *UserProfiles) AddString(key, val string) {
	r.add(key, typeString, val)
}

func (r *UserProfiles) AddInt(key string, val interface{}) {
	r.add(key, typeInt, val)
}

func (r *UserProfiles) AddDecimal(key string, val interface{}) {
	r.add(key, typeDecimal, val)
}

func (r *UserProfiles) AddDatetime(key string, val time.Time) {
	r.add(key, typeDatetime, val.Format(datetimeFormat))
}

func (r *UserProfiles) Validate() error {
	if len(r.UserID) == 0 {
		return errors.New("invalid user_id")
	}
	if len(r.UserProfiles) == 0 {
		return errors.New("invalid user_profiles")
	}
	return nil
}

func (r *UserProfiles) Send() error {
	if err := r.Validate(); err != nil {
		return err
	}

	body, err := json.Marshal(r)
	if err != nil {
		return err
	}

	return SendUserProfile(body)
}
