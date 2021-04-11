package resp

type Status string

var (
	Success Status = "Success"
	Error          = "Error"
)

func (s Status) String() string {
	return string(s)
}
