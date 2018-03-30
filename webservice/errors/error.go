package errors

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/acsellers/inflections"

	"gopkg.in/go-playground/validator.v9"
)

type ErrInfo interface {
	GetCode() int
	GetError() error
	GetStackTrace() string
}

type errDesc struct {
	Desc string `json:"desc"`
}

type errDescs []errDesc

type errInfo struct {
	Code       int    `json:"code"`
	Err        error  `json:"err"`
	StackTrace string `json:"stack_track"`
}

func ErrorBadParams(err error) error {
	if err == nil {
		return nil
	}
	ret := errInfo{
		Code: http.StatusBadRequest,
		Err: &errDesc{
			Desc: err.Error(),
		},
		StackTrace: stackTrace(2),
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		ret.Err = err
	} else if structErr, ok := err.(validator.ValidationErrors); ok {
		ret.Err = describeStructErrors(structErr)
	} else {
		switch err.Error() {
		case "EOF":
			ret.Err = errDescs{
				errDesc{
					Desc: "body is empty",
				},
			}
		default:
			ret.Err = errDescs{
				errDesc{
					Desc: err.Error(),
				},
			}
		}
	}
	return &ret
}

func ErrorAccessDenied(err error) error {
	if err == nil {
		return nil
	}
	return &errInfo{
		Code: http.StatusForbidden,
		Err: &errDesc{
			Desc: err.Error(),
		},
		StackTrace: stackTrace(2),
	}
}

func ErrorUnauthorized(err error) error {
	if err == nil {
		return nil
	}
	return &errInfo{
		Code: http.StatusUnauthorized,
		Err: &errDesc{
			Desc: err.Error(),
		},
		StackTrace: stackTrace(2),
	}
}

func ErrorNotFound(err error) error {
	if err == nil {
		return nil
	}
	return &errInfo{
		Code: http.StatusNotFound,
		Err: &errDesc{
			Desc: err.Error(),
		},
		StackTrace: stackTrace(2),
	}
}

func ErrorInternalServer(err error) error {
	if err == nil {
		return nil
	}
	return &errInfo{
		Code: http.StatusInternalServerError,
		Err: &errDesc{
			Desc: err.Error(),
		},
		StackTrace: stackTrace(2),
	}
}

// ===================================

var _ ErrInfo = errInfo{}

func (e *errDesc) Error() string {
	if e == nil {
		return "<nil>"
	}
	bytes, _ := json.Marshal(e)
	return string(bytes)
}

func (errors errDescs) Error() string {
	if len(errors) == 0 {
		return "<nil>"
	}
	bytes, _ := json.Marshal(errors)
	return string(bytes)
}

func (err *errInfo) Error() string {
	if err == nil {
		return "<nil>"
	}
	bytes, _ := json.Marshal(err)
	return string(bytes)
}

func (err errInfo) GetCode() int {
	return err.Code
}

func (err errInfo) GetError() error {
	return err.Err
}

func (err errInfo) GetStackTrace() string {
	return err.StackTrace
}

// ===================================

func describeStructErrors(errs validator.ValidationErrors) errDescs {
	rets := errDescs{}
	for _, v := range errs {
		switch v.Kind() {
		case reflect.String:
			switch v.Tag() {
			case "numeric":
				rets = append(rets, errDesc{
					Desc: inflections.Underscore(v.Field()) + ": required numeric format",
				})
			case "required":
				rets = append(rets, errDesc{
					Desc: inflections.Underscore(v.Field()) + ": required",
				})
			default:
				rets = append(rets, errDesc{
					Desc: inflections.Underscore(v.Field()) + ": " + v.Tag() + " " + v.Param(),
				})
			}
		case reflect.Int:
			switch v.Tag() {
			case "required":
				rets = append(rets, errDesc{
					Desc: inflections.Underscore(v.Field()) + ": required",
				})
			default:
				rets = append(rets, errDesc{
					Desc: inflections.Underscore(v.Field()) + ": " + v.Tag() + " " + v.Param(),
				})
			}
		case reflect.Float64, reflect.Float32:
			switch v.Tag() {
			case "required":
				rets = append(rets, errDesc{
					Desc: inflections.Underscore(v.Field()) + ": required",
				})
			default:
				rets = append(rets, errDesc{
					Desc: inflections.Underscore(v.Field()) + ": " + v.Tag() + " " + v.Param(),
				})
			}
		default:
			rets = append(rets, errDesc{
				Desc: inflections.Underscore(v.Field()) + "(" + v.Kind().String() + "): " + v.Tag() + " " + v.Param(),
			})
		}
	}

	return rets
}
