package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"myapi-dev-support/dev-support/models/api_model"
	"reflect"
	"strings"
)

type MyBaseController struct {
}

func (c *MyBaseController) HandleParam(ctx iris.Context, param any) bool {
	err := ctx.ReadJSON(param)
	if paramErrorResult := ParamFault(ctx, err, param); paramErrorResult != nil {
		ctx.JSON(paramErrorResult)
		return false
	}
	return true
}
func (c *MyBaseController) HandleParamError(ctx iris.Context, param any) *api_model.ApiResult {
	err := ctx.ReadJSON(param)
	if paramErrorResult := ParamFault(ctx, err, param); paramErrorResult != nil {
		return paramErrorResult
	}
	return nil
}

func (c *MyBaseController) PanicParamError(ctx iris.Context, param any) {
	err := ctx.ReadJSON(param)
	paramErr := WrapParamError(ctx, err, param)
	if paramErr != nil {
		panic(ApiParamErr{ParamError: paramErr})
	}
}

func FailParam(err *ParamError) *api_model.ApiResult {
	if err != nil {
		outMsg := "请求错误"
		r := &api_model.ApiResult{
			Success: false,
			Data:    err,
			Msg:     outMsg,
		}
		return r
	}
	return nil
}

type ParamError struct {
	Param string `json:"param,omitempty"`
	Error string `json:"error,omitempty"`
}
type ApiParamErr struct {
	ParamError *ParamError
}

func (e *ApiParamErr) Error() string {
	if e.ParamError.Param != "" {
		return strings.Join([]string{e.ParamError.Param, e.ParamError.Error}, ":")
	}
	return e.ParamError.Error
}

func ParamFault(ctx iris.Context, err error, param interface{}) *api_model.ApiResult {
	paramErr := WrapParamError(ctx, err, param)
	return FailParam(paramErr)
}
func WrapParamError(ctx iris.Context, err error, param interface{}) (paramError *ParamError) {
	_ = ctx
	if errors, ok := err.(validator.ValidationErrors); ok {
		for _, validationErr := range errors {
			_t := reflect.TypeOf(param)
			if _t.Kind() == reflect.Ptr {
				_t = _t.Elem()
			}
			field, _ := _t.FieldByName(validationErr.StructField())
			errJsonParam := field.Tag.Get("json")
			message := field.Tag.Get("message")
			if errJsonParam == "" {
				errJsonParam = validationErr.StructField()
			}
			paramError = &ParamError{
				Param: errJsonParam,
				Error: message,
			}
			return
		}
	}

	if err != nil {
		paramError = &ParamError{
			Error: err.Error(),
		}
	}
	return
}
