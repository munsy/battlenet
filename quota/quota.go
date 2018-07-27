package quota

import (
	"net/http"
	"strconv"
	"time"

	"github.com/munsy/battlenet/errors"
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
	qpsAllotted  int
	qpsCurrent   int
	quotaAlloted int
	quotaCurrent int
	quotaReset   time.Time
}

// QPSAllotted returns the allotted QPS.
func (q *Quota) QPSAllotted() int {
	return q.qpsAllotted
}

// QPSCurrent returns the current QPS.
func (q *Quota) QPSCurrent() int {
	return q.qpsCurrent
}

// QuotaAlloted returns the alloted quota.
func (q *Quota) QuotaAlloted() int {
	return q.quotaAlloted
}

// QuotaCurrent returns the current quota.
func (q *Quota) QuotaCurrent() int {
	return q.quotaCurrent
}

// QuotaReset returns the time the quota will reset.
func (q *Quota) QuotaReset() time.Time {
	return q.quotaReset
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

	q.qpsAllotted, err = strconv.Atoi(qpsAllotted)
	if nil != err {
		return err
	}

	q.qpsCurrent, err = strconv.Atoi(qpsCurrent)
	if nil != err {
		return err
	}

	q.quotaAlloted, err = strconv.Atoi(quotaAlloted)

	if nil != err {
		return err
	}

	q.quotaCurrent, err = strconv.Atoi(quotaCurrent)

	if nil != err {
		return err
	}

	q.quotaReset, err = time.Parse(timeFormat, quotaReset)

	if nil != err {
		return err
	}

	return nil
}
