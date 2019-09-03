package honey

import (
	"encoding/json"
)

type Solution struct {
	ID     uint64          `json:"id"`
	TaskID uint64          `json:"task_id" db:"task_id"`
	JobID  uint64          `json:"job_id" db:"job_id"`
	Data   json.RawMessage `json:"data"`
}
