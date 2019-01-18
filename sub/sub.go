package sub

import "golang.org/x/sync/singleflight"

var grp singleflight.Group

func FlightTest(key string) (string, error) {
	out, err, _ := grp.Do(key, func() (interface{}, error) {
		return key + ":RANDOM", nil
	})
	if err != nil {
		return "", err
	}
	return out.(string), nil
}
