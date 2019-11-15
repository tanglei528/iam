package utils

import (
    "github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"
    e "iam/pkg/exception"
    "net/http"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int, []string) {
    err := c.Bind(form)
    if err != nil {
        return http.StatusBadRequest, e.InvalidParams, ConvErrorToSlice(err, []string{})
    }

    valid := validation.Validation{}
    check, err := valid.Valid(form)
    if err != nil {
        return http.StatusInternalServerError, e.Error, ConvErrorToSlice(err, []string{})
    }
    if !check {
        errorsMsg := MarkErrors(form, valid.Errors)
        return http.StatusBadRequest, e.InvalidParams, errorsMsg
    }

    return http.StatusOK, e.Success, nil
}
