package output

import "context"

type ILogsRepository interface {
	WriteMessage(ctx context.Context, level string, sentence string, extra interface{}) (bool, error)
}
