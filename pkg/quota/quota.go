// Package quota contains usage data about the Battle.net API, which comes back as part of the response headers after
// you make a call with one of the gobattlenet clients. The following is a breakdown of the headers:
//
// X-Plan-Qps-Allotted: The maximum QPS (queries per second) capacity set on your API key. This is the maximum
// number of calls that can be made in any given second.
// X-Plan-Qps-Current: The current count of calls being applied against the above limit; in this case, it
// is a representation of how many calls your key is making at that particular second in time.
// X-Plan-Quota-Allotted: The maximum number of calls that can be made on daily basis.
// X-Plan-Quota-Current: The current count of calls made in the current limit period, i.e. the current day.
// X-Plan-Quota-Reset: The time when the quota count will reset to 0.
package quota

import (
	"net/http"
	"time"
)

const timeFormat = "Monday, Jan 02, 2006 15:04:05 AM GMT"

// Quota contains data about the Battle.net API usage.
type Quota struct {
	QPSAllotted  int
	QPSCurrent   int
	QuotaAlloted int
	QuotaCurrent int
	QuotaReset   time.Time
}

// Set sets the quota according to values returned by the given http.Response.
func (q *Quota) Set(r *http.Response) error {
	var err error

	q.QPSAllotted = r.Header["X-Plan-Qps-Allotted"]
	q.QPSCurrent = r.Header["X-Plan-Qps-Current"]
	q.QuotaAlloted = r.Header["X-Plan-Quota-Allotted"]
	q.QuotaCurrent = r.Header["X-Plan-Quota-Current"]
	q.QuotaReset, err = time.Parse(timeFormat, r.Header["X-Plan-Quota-Reset"])

	return err
}
