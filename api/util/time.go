package util

import (
	"strconv"
	"time"
)

func GetTimeFromEpoch(epoch string) (*time.Time, error) {
	i, err := strconv.ParseInt(epoch, 10, 64)
	if err != nil {
		return nil, err
	}
	tm := time.Unix(i, 0)

	return &tm, nil
}
