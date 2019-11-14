package utils

import (
    "fmt"
    "github.com/astaxie/beego/validation"
    "iam/pkg/logging"
    "reflect"
)

func MarkErrors(form interface{}, errors []*validation.Error) []string {
    errorsMsg := []string{}
    st := reflect.TypeOf(form).Elem()
    for _, err := range errors {
        //logging.Error(err.Key, err.Message)
        field, _ := st.FieldByName(err.Field)
        tag := field.Tag.Get("json")
        logging.Error(tag, err.Message)
        errorsMsg = append(errorsMsg, fmt.Sprintf("%s %s", tag, err.Message))
    }
    return errorsMsg
}
