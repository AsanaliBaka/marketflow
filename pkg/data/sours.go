package data

type SourseConn interface {
	Sourse
	Close()
}

type SourseClient interface {
	SourseConn() SourseConn
	Close() error
}

type Sourse interface {
	Response() ([]string, error)
}
