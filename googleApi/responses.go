package googleApi

const (
	STATUS_OK               = "OK"
	STATUS_ZERO_RESULTS     = "ZERO_RESULTS"
	STATUS_NOT_FOUND        = "NOT_FOUND"
	STATUS_INVALID_REQUEST  = "INVALID_REQUEST"
	STATUS_OVER_QUERY_LIMIT = "OVER_QUERY_LIMIT"
	STATUS_REQUEST_DENIED   = "REQUEST_DENIED"
	STATUS_UNKNOWN_ERROR    = "UNKNOWN_ERROR"
)

type PlaceAnswer struct {
	Status string `json:"status"`
	Result Place  `json:"result"`
}

type NearbyPlacesAnswer struct {
	Status string  `json:"status"`
	Result []Place `json:"results"`
}
