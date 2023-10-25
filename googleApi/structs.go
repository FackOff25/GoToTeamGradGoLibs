package googleApi

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type Photo struct {
	Height int64 `json:"height, omitempty"`
	Width int64 `json:"width, omitempty"`
	Reference string `json:"photo_reference, omitempty"`
}

type OpeningHours struct {
	WeekTimetable []string `json:"weekday_text, omitempty"`
}

type Summary struct {
	Language string `json:"editorial_summary, omitempty"`
	Overview string `json:"overview, omitempty"`
}

type Place struct {
	PlaceId string `json:"place_id, omitempty"`
	Name string `json:"name, omitempty"`
	Geometry Geometry `json:"geometry, omitempty"`
	OpeningHours OpeningHours `json:"opening_hours, omitempty"`
	Rating float64 `json:"rating, omitempty"`
	RatingCount int64 `json:"user_ratings_total, omitempty"`
	Types []string `json:"types, omitempty"`
	Address string `json:"formatted_address, omitempty"`
	Summary Summary `json:"editorial_summary, omitempty"`
	Photos []Photo `json:"photos, omitempty"`
}