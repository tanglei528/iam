package utils

import "strconv"

func IntToBool(p int) (bool, error) {
    s := strconv.Itoa(p)
    b, err := strconv.ParseBool(s)
    return b, err
}

func ConvErrorToSlice(err error, errorsInfo []string) []string {
    errorsInfo = append(errorsInfo, err.Error())
    return errorsInfo
}