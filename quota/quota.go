package quota

import (
	"net/http"
	"strconv"
	"time"

	"github.com/munsy/gobattlenet/errors"
)

const timeFormat = "Monday, January 2, 2006 3:04:05 PM GMT"

// Quota contains data about the Battle.net API usage. The following is a breakdown of the headers:
//
// X-Plan-Qps-Allotted: The maximum QPS (queries per second) capacity set on your API key. This is the maximum
// number of calls that can be made in any given second.
//
// X-Plan-Qps-Current: The current count of calls being applied against the above limit; in this case, it
// is a representation of how many calls your key is making at that particular second in time.
//
// X-Plan-Quota-Allotted: The maximum number of calls that can be made on daily basis.
//
// X-Plan-Quota-Current: The current count of calls made in the current limit period, i.e. the current day.
//
// X-Plan-Quota-Reset: The time when the quota count will reset to 0.
type Quota struct {
	QPSAllotted  int
	QPSCurrent   int
	QuotaAlloted int
	QuotaCurrent int
	QuotaReset   time.Time
}

// Set sets the quota according to values returned by the given http.Response.
func (q *Quota) Set(r *http.Response) error {
	if r.Header.Get("X-Mashery-Error-Code") != "" {
		return errors.ErrNoTokenSupplied
	}

	qpsAllotted := r.Header.Get("X-Plan-Qps-Allotted")
	qpsCurrent := r.Header.Get("X-Plan-Qps-Current")
	quotaAlloted := r.Header.Get("X-Plan-Quota-Allotted")
	quotaCurrent := r.Header.Get("X-Plan-Quota-Current")
	quotaReset := r.Header.Get("X-Plan-Quota-Reset")

	var err error

	q.QPSAllotted, err = strconv.Atoi(qpsAllotted)
	if nil != err {
		return err
	}

	q.QPSCurrent, err = strconv.Atoi(qpsCurrent)
	if nil != err {
		return err
	}

	q.QuotaAlloted, err = strconv.Atoi(quotaAlloted)

	if nil != err {
		return err
	}

	q.QuotaCurrent, err = strconv.Atoi(quotaCurrent)

	if nil != err {
		return err
	}

	q.QuotaReset, err = time.Parse(timeFormat, quotaReset)

	if nil != err {
		return err
	}

	return nil
}
