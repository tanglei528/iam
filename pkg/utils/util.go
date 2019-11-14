package utils

import "strconv"

func IntToBool(p int) (bool, error) {
    s := strconv.Itoa(p)
    b, err := strconv.ParseBool(s)
    return b, err
}
