package sql

import "github.com/go-kratos/kratos/pkg/net/trace"

func tracer(tt trace.Trace, operationName string) (t trace.Trace) {
	var internalTags []trace.Tag
	internalTags = append(internalTags, trace.TagString(trace.TagComponent, "databases/mysql"))
	internalTags = append(internalTags, trace.TagString(trace.TagPeerService, "mysql"))
	internalTags = append(internalTags, trace.TagString(trace.TagSpanKind, "client"))

	t = tt.Fork(_family, "Mysql:"+operationName)
	t.SetTag(internalTags...)
	return
}
