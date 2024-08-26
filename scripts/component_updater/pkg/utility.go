package pkg

type SystemType int

const (
	Meshplay SystemType = iota
	Docs
	RemoteProvider
)

func (dt SystemType) String() string {
	switch dt {
	case Meshplay:
		return "meshplay"

	case Docs:
		return "docs"

	case RemoteProvider:
		return "remote-provider"
	}
	return ""
}
