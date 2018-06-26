package keycloak

import (
	"strconv"
	"strings"
	"time"
)

// AttributeMap represents a map of attributes
type AttributeMap map[string]interface{}

// UnixTime is an alias for a date time from Keycloak
type UnixTime time.Time

// MarshalJSON lets UnixTime implement the json.Marshaler interface
func (t UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON lets UnixTime implement the json.Unmarshaler interface
func (t *UnixTime) UnmarshalJSON(s []byte) error {
	r := strings.Replace(string(s), `"`, ``, -1)

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q/1000, 0)
	return nil
}

func (t UnixTime) String() string {
	return time.Time(t).String()
}