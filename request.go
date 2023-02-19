package bingx

import "context"

// Request - used as common contract.
type Request interface {
	Process(ctx context.Context) (interface{}, error)
}
