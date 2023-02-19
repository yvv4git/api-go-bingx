package bingx

// PayloadURL - used as common contract for all url payloads.
type PayloadURL interface {
	Create() string
}
