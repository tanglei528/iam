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
        return http.StatusBadRequest, e.InvalidParams, nil
    }

    valid := validation.Validation{}
    check, err := valid.Valid(form)
    if err != nil {
        return http.StatusInternalServerError, e.Error, nil
    }
    if !check {
        errorsMsg := MarkErrors(form, valid.Errors)
        return http.StatusBadRequest, e.InvalidParams, errorsMsg
    }

    return http.StatusOK, e.Success, nil
}
