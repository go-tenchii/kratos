package httpclient

import (
	"context"
	"github.com/go-tenchii/kratos/pkg/net/metadata"
	"github.com/go-tenchii/kratos/pkg/net/trace"
	"net/http"
	"strconv"
)

const _defaultComponentName = "net/http"

func Do(ctx context.Context, dc *http.Client, req *http.Request) (*http.Response, error) {
	t, err := trace.Extract(trace.HTTPFormat, req.Header)
	if err != nil {
		var opts []trace.Option
		if ok, _ := strconv.ParseBool(trace.KratosTraceDebug); ok {
			opts = append(opts, trace.EnableDebug())
		}
		t = trace.New("external/"+req.URL.Path, opts...)
	}
	t.SetTitle(req.URL.Path)
	t.SetTag(trace.String(trace.TagComponent, _defaultComponentName))
	t.SetTag(trace.String(trace.TagHTTPMethod, req.Method))
	t.SetTag(trace.String(trace.TagHTTPURL, req.URL.String()))
	t.SetTag(trace.String(trace.TagSpanKind, "server"))
	// business tag
	t.SetTag(trace.String("caller", metadata.String(ctx, metadata.Caller)))
	// export trace id to user.
	ctx = trace.NewContext(ctx, t)
	resp, err := dc.Do(req)
	t.Finish(&err)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
