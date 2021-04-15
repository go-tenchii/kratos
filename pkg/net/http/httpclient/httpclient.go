package httpclient

import (
	"context"
	"github.com/go-tenchii/kratos/pkg/net/trace"
	"net/http"
	"strconv"
)

const _defaultComponentName = "net/http"

func Do(ctx context.Context, dc *http.Client, req *http.Request) (*http.Response, error) {
	var (
		t  trace.Trace
		ok bool
	)

	if t, ok = trace.FromContext(ctx); ok {
		req.Header.Set(trace.KratosTraceID, t.TraceID())
	}
	t, err := trace.Extract(trace.HTTPFormat, req.Header)
	if err != nil {
		var opts []trace.Option
		if ok, _ := strconv.ParseBool(trace.KratosTraceDebug); ok {
			opts = append(opts, trace.EnableDebug())
		}
		t = trace.New(req.URL.Path, opts...)
	}

	t.SetTitle("外部服务: " + req.URL.Path)
	t.SetTag(trace.String(trace.TagComponent, _defaultComponentName))
	t.SetTag(trace.String(trace.TagHTTPMethod, req.Method))
	t.SetTag(trace.String(trace.TagHTTPURL, req.URL.String()))
	ctx = trace.NewContext(ctx, t)
	resp, err := dc.Do(req)
	defer t.Finish(&err)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
