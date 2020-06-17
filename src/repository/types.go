package repository

// HighWay struct
type HighWay struct {
	ID     string
	Name   string
	Rating float64
}

// HighWayType struct
type HighWayType struct {
	ID   string
	Name string
}

// Country struct
type Country struct {
	Name string
	Code string
}

// Feedback struct
type Feedback struct {
	Text     string
	Rating   float64
	Date     string
	RoadPart string
}

// IRepository interface
type IRepository interface {
	GetCountiesList() ([]Country, error)
	GetHighwaysList(countryCode string, highwayTypeID string) ([]HighWay, error)
	GetFeedbacksList(countryCode string, highwayTypeID string) ([]Feedback, error)
}
