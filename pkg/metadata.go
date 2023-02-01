package pkg

type Metadata struct {
	RequestUUID string
}

func NewMetadata(requestUUID string) *Metadata {
	return &Metadata{requestUUID}
}
