package pkg


import


type Metadata struct {
	RequestUUID string
}

func NewMetadata(requestUUID string) *Metadata {
	return &Metadata{requestUUID}
}
