package audits

import "time"

type Audit struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a Audit) AuditString() string {
	return "Created at: " + a.CreatedAt.Format("2006/01/02") + "\n" + "Updated at: " + a.UpdatedAt.Format("2006/01/02")
}
