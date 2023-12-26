package sendRequest

// Account model info
// @Description Request
type Request struct {
	ItemCount uint64 `json:"ItemCount" validate:"required,gt=0,lte=100000"`
}
