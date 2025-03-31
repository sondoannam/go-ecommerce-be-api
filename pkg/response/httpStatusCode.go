package response

const (
	Success       = 201
	InvalidParams = 400
	NotFound      = 404
)

// messages
var msg = map[int]string{
	Success:       "Success",
	InvalidParams: "Invalid Parameters",
	NotFound:      "Not Found",
}
