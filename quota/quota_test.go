package quota

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestQuota(t *testing.T) {
	r := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString("body")),
		ContentLength: int64(4),
		Request:       nil,
		Header:        make(http.Header, 0),
	}

	r.Header.Add("X-Plan-Qps-Allotted", "100")
	r.Header.Add("X-Plan-Qps-Current", "1")
	r.Header.Add("X-Plan-Quota-Allotted", "36000")
	r.Header.Add("X-Plan-Quota-Current", "1")
	r.Header.Add("X-Plan-Quota-Reset", "Friday, July 27, 2018 5:00:00 AM GMT")

	q := &Quota{}

	q.Set(r)

	if q.qpsAllotted != 100 {
		t.Fatal("qpsAllotted wasn't properly set")
	}
	if q.qpsCurrent != 1 {
		t.Fatal("qpsCurrent wasn't properly set")
	}
	if q.quotaAlloted != 36000 {
		t.Fatal("quotaAlloted wasn't properly set")
	}
	if q.quotaCurrent != 1 {
		t.Fatal("quotaCurrent wasn't properly set")
	}
	if q.quotaReset.Unix() != 1532667600 {
		t.Fatal("quotaReset wasn't properly set")
	}
}
