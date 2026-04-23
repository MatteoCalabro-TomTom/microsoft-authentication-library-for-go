package uuid

import "github.com/gofrs/uuid/v5"


func New() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}

func Parse(s string) (uuid.UUID, error) {
	return uuid.FromString(s)
}
