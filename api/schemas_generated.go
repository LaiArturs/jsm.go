// auto generated 2020-04-27 14:52:37.047353 +0200 CEST m=+2.670812440

package api

import (
	"encoding/base64"
)

var schemas map[string][]byte

var schemaTypes = map[string]func() interface{}{
	"io.nats.server.advisory.v1.client_connect":              func() interface{} { return &ConnectEventMsg{} },
	"io.nats.server.advisory.v1.client_disconnect":           func() interface{} { return &DisconnectEventMsg{} },
	"io.nats.server.metric.v1.service_latency":               func() interface{} { return &ServiceLatency{} },
	"io.nats.jetstream.advisory.v1.api_audit":                func() interface{} { return &JetStreamAPIAudit{} },
	"io.nats.jetstream.advisory.v1.max_deliver":              func() interface{} { return &ConsumerDeliveryExceededAdvisory{} },
	"io.nats.jetstream.metric.v1.consumer_ack":               func() interface{} { return &ConsumerAckMetric{} },
	"io.nats.jetstream.api.v1.consumer_configuration":        func() interface{} { return &ConsumerConfig{} },
	"io.nats.jetstream.api.v1.stream_configuration":          func() interface{} { return &StreamConfig{} },
	"io.nats.jetstream.api.v1.stream_template_configuration": func() interface{} { return &StreamTemplateConfig{} },
	"io.nats.unknown_event":                                  func() interface{} { return &UnknownEvent{} },
}

func init() {
	schemas = make(map[string][]byte)
	schemas["io.nats.server.advisory.v1.client_connect"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9zZXJ2ZXIvYWR2aXNvcnkvdjEvY2xpZW50X2Nvbm5lY3QuanNvbiIsCiAgImRlc2NyaXB0aW9uIjogIkFkdmlzb3J5IHB1Ymxpc2hlZCBhIGNsaWVudCBjb25uZWN0cyB0byB0aGUgTkFUUyBTZXJ2ZXIiLAogICJ0aXRsZSI6ICJpby5uYXRzLnNlcnZlci5hZHZpc29yeS52MS5jbGllbnRfY29ubmVjdCIsCiAgInR5cGUiOiJvYmplY3QiLAogICJyZXF1aXJlZCI6WwogICAgInR5cGUiLAogICAgImlkIiwKICAgICJ0aW1lc3RhbXAiLAogICAgInNlcnZlciIsCiAgICAiY2xpZW50IgogIF0sCiAgImFkZGl0aW9uYWxJdGVtcyI6IGZhbHNlLAogICJwcm9wZXJ0aWVzIjogewogICAgInR5cGUiOiB7CiAgICAgICJ0eXBlIjoic3RyaW5nIiwKICAgICAgImVudW0iOiBbCiAgICAgICAgImlvLm5hdHMuc2VydmVyLmFkdmlzb3J5LnYxLmNsaWVudF9jb25uZWN0IgogICAgICBdCiAgICB9LAogICAgImlkIjogewogICAgICAidHlwZSI6InN0cmluZyIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJVbmlxdWUgY29ycmVsYXRpb24gSUQgZm9yIHRoaXMgZXZlbnQiCiAgICB9LAogICAgInRpbWVzdGFtcCI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImZvcm1hdCI6ICJkYXRlLXRpbWUiLAogICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHRpbWUgdGhpcyBldmVudCB3YXMgY3JlYXRlZCBpbiBSRkMzMzM5IGZvcm1hdCIKICAgIH0sCiAgICAic2VydmVyIjogewogICAgICAidHlwZSI6ICJvYmplY3QiLAogICAgICAiYWRkaXRpb25hbEl0ZW1zIjogZmFsc2UsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJEZXRhaWxzIGFib3V0IHRoZSBzZXJ2ZXIgdGhlIGNsaWVudCBjb25uZWN0ZWQgdG8iLAogICAgICAicmVxdWlyZWQiOiBbIm5hbWUiLCJob3N0IiwiaWQiLCJ2ZXIiLCJzZXEiLCJqZXRzdHJlYW0iLCJ0aW1lIl0sCiAgICAgICJwcm9wZXJ0aWVzIjogewogICAgICAgICJuYW1lIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgY29uZmlndXJlZCBuYW1lIGZvciB0aGUgc2VydmVyLCBtYXRjaGVzIElEIHdoZW4gdW5jb25maWd1cmVkIiwKICAgICAgICAgICJtaW5MZW5ndGgiOiAxCiAgICAgICAgfSwKICAgICAgICAiaG9zdCI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIGhvc3QgdGhpcyBzZXJ2ZXIgcnVucyBvbiwgdHlwaWNhbGx5IGEgSVAgYWRkcmVzcyIKICAgICAgICB9LAogICAgICAgICJpZCI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHVuaXF1ZSBzZXJ2ZXIgSUQgZm9yIHRoaXMgbm9kZSIKICAgICAgICB9LAogICAgICAgICJjbHVzdGVyIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgY2x1c3RlciB0aGUgc2VydmVyIGJlbG9uZ3MgdG8iCiAgICAgICAgfSwKICAgICAgICAidmVyIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgdmVyc2lvbiBOQVRTIHJ1bm5pbmcgb24gdGhlIHNlcnZlciIKICAgICAgICB9LAogICAgICAgICJzZXEiOiB7CiAgICAgICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJJbnRlcm5hbCBzZXJ2ZXIgc2VxdWVuY2UgSUQiCiAgICAgICAgfSwKICAgICAgICAiamV0c3RyZWFtIjogewogICAgICAgICAgInR5cGUiOiAiYm9vbGVhbiIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiSW5kaWNhdGVzIGlmIHRoaXMgc2VydmVyIGhhcyBKZXRTdHJlYW0gZW5hYmxlZCIKICAgICAgICB9LAogICAgICAgICJ0aW1lIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJmb3JtYXQiOiAiZGF0ZS10aW1lIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgbG9jYWwgdGltZSBvZiB0aGUgc2VydmVyIgogICAgICAgIH0KICAgICAgfQogICAgfSwKICAgICJjbGllbnQiOiB7CiAgICAgICJ0eXBlIjogIm9iamVjdCIsCiAgICAgICJhZGRpdGlvbmFsSXRlbXMiOiBmYWxzZSwKICAgICAgImRlc2NyaXB0aW9uIjogIkRldGFpbHMgYWJvdXQgdGhlIGNsaWVudCB0aGF0IGNvbm5lY3RlZCB0byB0aGUgc2VydmVyIiwKICAgICAgInJlcXVpcmVkIjogWwogICAgICAgICJpZCIsCiAgICAgICAgImFjYyIKICAgICAgXSwKICAgICAgInByb3BlcnRpZXMiOiB7CiAgICAgICAgInN0YXJ0IjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJmb3JtYXQiOiAiZGF0ZS10aW1lIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaW1lc3RhbXAgd2hlbiB0aGUgY2xpZW50IGNvbm5lY3RlZCIKICAgICAgICB9LAogICAgICAgICJzdG9wIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJmb3JtYXQiOiAiZGF0ZS10aW1lIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaW1lc3RhbXAgd2hlbiB0aGUgY2xpZW50IGRpc2Nvbm5lY3RlZCIKICAgICAgICB9LAogICAgICAgICJob3N0IjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgcmVtb3RlIGhvc3QgdGhlIGNsaWVudCBpcyBjb25uZWN0ZWQgZnJvbSIKICAgICAgICB9LAogICAgICAgICJ1c2VyIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgY2xpZW50cyB1c2VybmFtZSIKICAgICAgICB9LAogICAgICAgICJuYW1lIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgbmFtZSBwcmVzZW50ZWQgYnkgdGhlIGNsaWVudCBkdXJpbmcgY29ubmVjdGlvbiIKICAgICAgICB9LAogICAgICAgICJsYW5nIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgcHJvZ3JhbW1pbmcgbGFuZ3VhZ2UgbGlicmFyeSBpbiB1c2UgYnkgdGhlIGNsaWVudCIKICAgICAgICB9LAogICAgICAgICJ2ZXIiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSB2ZXJzaW9uIG9mIHRoZSBjbGllbnQgbGlicmFyeSBpbiB1c2UiCiAgICAgICAgfSwKICAgICAgICAicnR0IjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgbGFzdCBrbm93biBsYXRlbmN5IGJldHdlZW4gdGhlIE5BVFMgU2VydmVyIGFuZCB0aGUgQ2xpZW50IgogICAgICAgIH0KICAgICAgfQogICAgfQogIH0KfQo=")
	schemas["io.nats.server.advisory.v1.client_disconnect"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9zZXJ2ZXIvYWR2aXNvcnkvdjEvY2xpZW50X2Rpc2Nvbm5lY3QuanNvbiIsCiAgImRlc2NyaXB0aW9uIjogIkFkdmlzb3J5IHB1Ymxpc2hlZCBhIGNsaWVudCBkaXNjb25uZWN0cyB0byB0aGUgTkFUUyBTZXJ2ZXIiLAogICJ0aXRsZSI6ICJpby5uYXRzLnNlcnZlci5hZHZpc29yeS52MS5jbGllbnRfZGlzY29ubmVjdCIsCiAgInR5cGUiOiJvYmplY3QiLAogICJkZWZpbml0aW9ucyI6IHsKICAgICJkYXRhc3RhdHMiOiB7CiAgICAgICJ0eXBlIjogIm9iamVjdCIsCiAgICAgICJhZGRpdGlvbmFsSXRlbXMiOiBmYWxzZSwKICAgICAgInByb3BlcnRpZXMiOiB7CiAgICAgICAgIm1zZ3MiOiB7CiAgICAgICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgbnVtYmVyIG9mIG1lc3NhZ2VzIGhhbmRsZWQgYnkgdGhlIGNsaWVudCIKICAgICAgICB9LAogICAgICAgICJieXRlcyI6IHsKICAgICAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBudW1iZXIgb2YgYnl0ZXMgaGFuZGxlZCBieSB0aGUgY2xpZW50IgogICAgICAgIH0KICAgICAgfQogICAgfQogIH0sCiAgInJlcXVpcmVkIjpbCiAgICAidHlwZSIsCiAgICAiaWQiLAogICAgInRpbWVzdGFtcCIsCiAgICAic2VydmVyIiwKICAgICJjbGllbnQiLAogICAgInNlbnQiLAogICAgInJlY2VpdmVkIiwKICAgICJyZWFzb24iCiAgXSwKICAiYWRkaXRpb25hbEl0ZW1zIjogZmFsc2UsCiAgInByb3BlcnRpZXMiOiB7CiAgICAidHlwZSI6IHsKICAgICAgInR5cGUiOiJzdHJpbmciLAogICAgICAiZW51bSI6IFsKICAgICAgICAiaW8ubmF0cy5zZXJ2ZXIuYWR2aXNvcnkudjEuY2xpZW50X2Rpc2Nvbm5lY3QiCiAgICAgIF0KICAgIH0sCiAgICAiaWQiOiB7CiAgICAgICJ0eXBlIjoic3RyaW5nIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlVuaXF1ZSBjb3JyZWxhdGlvbiBJRCBmb3IgdGhpcyBldmVudCIKICAgIH0sCiAgICAidGltZXN0YW1wIjogewogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAiZm9ybWF0IjogImRhdGUtdGltZSIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgdGltZSB0aGlzIGV2ZW50IHdhcyBjcmVhdGVkIGluIFJGQzMzMzkgZm9ybWF0IgogICAgfSwKICAgICJzZXJ2ZXIiOiB7CiAgICAgICJ0eXBlIjogIm9iamVjdCIsCiAgICAgICJhZGRpdGlvbmFsSXRlbXMiOiBmYWxzZSwKICAgICAgImRlc2NyaXB0aW9uIjogIkRldGFpbHMgYWJvdXQgdGhlIHNlcnZlciB0aGUgY2xpZW50IGNvbm5lY3RlZCB0byIsCiAgICAgICJyZXF1aXJlZCI6IFsibmFtZSIsImhvc3QiLCJpZCIsInZlciIsInNlcSIsImpldHN0cmVhbSIsInRpbWUiXSwKICAgICAgInByb3BlcnRpZXMiOiB7CiAgICAgICAgIm5hbWUiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBjb25maWd1cmVkIG5hbWUgZm9yIHRoZSBzZXJ2ZXIsIG1hdGNoZXMgSUQgd2hlbiB1bmNvbmZpZ3VyZWQiLAogICAgICAgICAgIm1pbkxlbmd0aCI6IDEKICAgICAgICB9LAogICAgICAgICJob3N0IjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgaG9zdCB0aGlzIHNlcnZlciBydW5zIG9uLCB0eXBpY2FsbHkgYSBJUCBhZGRyZXNzIgogICAgICAgIH0sCiAgICAgICAgImlkIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgdW5pcXVlIHNlcnZlciBJRCBmb3IgdGhpcyBub2RlIgogICAgICAgIH0sCiAgICAgICAgImNsdXN0ZXIiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBjbHVzdGVyIHRoZSBzZXJ2ZXIgYmVsb25ncyB0byIKICAgICAgICB9LAogICAgICAgICJ2ZXIiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSB2ZXJzaW9uIE5BVFMgcnVubmluZyBvbiB0aGUgc2VydmVyIgogICAgICAgIH0sCiAgICAgICAgInNlcSI6IHsKICAgICAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkludGVybmFsIHNlcnZlciBzZXF1ZW5jZSBJRCIKICAgICAgICB9LAogICAgICAgICJqZXRzdHJlYW0iOiB7CiAgICAgICAgICAidHlwZSI6ICJib29sZWFuIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJJbmRpY2F0ZXMgaWYgdGhpcyBzZXJ2ZXIgaGFzIEpldFN0cmVhbSBlbmFibGVkIgogICAgICAgIH0sCiAgICAgICAgInRpbWUiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImZvcm1hdCI6ICJkYXRlLXRpbWUiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBsb2NhbCB0aW1lIG9mIHRoZSBzZXJ2ZXIiCiAgICAgICAgfQogICAgICB9CiAgICB9LAogICAgImNsaWVudCI6IHsKICAgICAgInR5cGUiOiAib2JqZWN0IiwKICAgICAgImFkZGl0aW9uYWxJdGVtcyI6IGZhbHNlLAogICAgICAiZGVzY3JpcHRpb24iOiAiRGV0YWlscyBhYm91dCB0aGUgY2xpZW50IHRoYXQgY29ubmVjdGVkIHRvIHRoZSBzZXJ2ZXIiLAogICAgICAicmVxdWlyZWQiOiBbImlkIiwgImFjYyJdLAogICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAic3RhcnQiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImZvcm1hdCI6ICJkYXRlLXRpbWUiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRpbWVzdGFtcCB3aGVuIHRoZSBjbGllbnQgY29ubmVjdGVkIgogICAgICAgIH0sCiAgICAgICAgInN0b3AiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImZvcm1hdCI6ICJkYXRlLXRpbWUiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRpbWVzdGFtcCB3aGVuIHRoZSBjbGllbnQgZGlzY29ubmVjdGVkIgogICAgICAgIH0sCiAgICAgICAgImhvc3QiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSByZW1vdGUgaG9zdCB0aGUgY2xpZW50IGlzIGNvbm5lY3RlZCBmcm9tIgogICAgICAgIH0sCiAgICAgICAgInVzZXIiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBjbGllbnRzIHVzZXJuYW1lIgogICAgICAgIH0sCiAgICAgICAgIm5hbWUiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBuYW1lIHByZXNlbnRlZCBieSB0aGUgY2xpZW50IGR1cmluZyBjb25uZWN0aW9uIgogICAgICAgIH0sCiAgICAgICAgImxhbmciOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBwcm9ncmFtbWluZyBsYW5ndWFnZSBsaWJyYXJ5IGluIHVzZSBieSB0aGUgY2xpZW50IgogICAgICAgIH0sCiAgICAgICAgInZlciI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHZlcnNpb24gb2YgdGhlIGNsaWVudCBsaWJyYXJ5IGluIHVzZSIKICAgICAgICB9LAogICAgICAgICJydHQiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBsYXN0IGtub3duIGxhdGVuY3kgYmV0d2VlbiB0aGUgTkFUUyBTZXJ2ZXIgYW5kIHRoZSBDbGllbnQiCiAgICAgICAgfQogICAgICB9CiAgICB9LAogICAgInNlbnQiOiB7CiAgICAgICJkZXNjcmlwdGlvbiI6ICJEYXRhIHNlbnQgdG8gdGhlIGNsaWVudCIsCiAgICAgICIkcmVmIjogIiMvZGVmaW5pdGlvbnMvZGF0YXN0YXRzIgogICAgfSwKICAgICJyZWNlaXZlZCI6IHsKICAgICAgImRlc2NyaXB0aW9uIjogIkRhdGEgcmVjZWl2ZWQgZnJvbSB0aGUgY2xpZW50IiwKICAgICAgIiRyZWYiOiAiIy9kZWZpbml0aW9ucy9kYXRhc3RhdHMiCiAgICB9LAogICAgInJlYXNvbiI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSByZWFzb24gdGhlIGNsaWVudCBkaXNjb25uZWN0ZWQiCiAgICB9CiAgfQp9Cg==")
	schemas["io.nats.server.metric.v1.service_latency"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9zZXJ2ZXIvbWV0cmljL3YxL3NlcnZpY2VfbGF0ZW5jeS5qc29uIiwKICAiZGVzY3JpcHRpb24iOiAiTWV0cmljIHB1Ymxpc2hlZCBhYm91dCBzYW1wbGVkIHNlcnZpY2UgcmVxdWVzdHMgc2hvd2luZyByZXF1ZXN0IHN0YXR1cyBhbmQgbGF0ZW5jaWVzIiwKICAidGl0bGUiOiAiaW8ubmF0cy5zZXJ2ZXIubWV0cmljLnYxLnNlcnZpY2VfbGF0ZW5jeSIsCiAgInR5cGUiOiJvYmplY3QiLAogICJyZXF1aXJlZCI6WwogICAgInR5cGUiLAogICAgImlkIiwKICAgICJ0aW1lc3RhbXAiLAogICAgInN0YXR1cyIsCiAgICAic3RhcnQiLAogICAgInN2YyIsCiAgICAibmF0cyIsCiAgICAidG90YWwiCiAgXSwKICAiYWRkaXRpb25hbEl0ZW1zIjogZmFsc2UsCiAgInByb3BlcnRpZXMiOiB7CiAgICAidHlwZSI6IHsKICAgICAgInR5cGUiOiJzdHJpbmciLAogICAgICAiY29uc3QiOiJpby5uYXRzLnNlcnZlci5tZXRyaWMudjEuc2VydmljZV9sYXRlbmN5IgogICAgfSwKICAgICJpZCI6IHsKICAgICAgInR5cGUiOiJzdHJpbmciLAogICAgICAiZGVzY3JpcHRpb24iOiAiVW5pcXVlIGNvcnJlbGF0aW9uIElEIGZvciB0aGlzIGV2ZW50IgogICAgfSwKICAgICJ0aW1lc3RhbXAiOiB7CiAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICJmb3JtYXQiOiAiZGF0ZS10aW1lIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSB0aW1lIHRoaXMgZXZlbnQgd2FzIGNyZWF0ZWQgaW4gUkZDMzMzOSBmb3JtYXQiCiAgICB9LAogICAgInN0YXR1cyI6IHsKICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgc3RhdHVzIG9mIHRoZSByZXF1ZXN0LiAyMDAgT0ssIDQwMCBCYWQgUmVxdWVzdCwgbm8gcmVwbHkgc3ViamVjdC4gNDA4IFJlcXVlc3QgVGltZW91dCwgcmVxdWVzdGVyIGxvc3QgaW50ZXJlc3QgYmVmb3JlIHJlcXVlc3QgY29tcGxldGVkLiA1MDMgU2VydmljZSBVbmF2YWlsYWJsZS4gNTA0IFNlcnZpY2UgVGltZW91dC4iLAogICAgICAiZW51bSI6IFsyMDAsIDQwMCwgNDA4LCA1MDMsIDUwNF0KICAgIH0sCiAgICAiZXJyb3IiOiB7CiAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJBIGRlc2NyaXB0aW9uIG9mIHRoZSBzdGF0dXMgY29kZSB3aGVuIG5vdCAyMDAiCiAgICB9LAogICAgImFwcCI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBuYW1lIHJlZ2lzdGVyZWQgYnkgdGhlIHNlcnZpY2Ugd2hlbiBjb25uZWN0aW5nIHRvIHRoZSBOQVRTIG5ldHdvcmsiCiAgICB9LAogICAgInN0YXJ0IjogewogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAiZm9ybWF0IjogImRhdGUtdGltZSIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgdGltZSB0aGUgcmVxdWVzdCBzdGFydGVkIgogICAgfSwKICAgICJzdmMiOiB7CiAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHRpbWUgdGFrZW4gYnkgdGhlIHNlcnZpY2UgdG8gcGVyZm9ybSB0aGUgcmVxdWVzdCBpbiBuYW5vc2Vjb25kcyIKICAgIH0sCiAgICAibmF0cyI6IHsKICAgICAgInR5cGUiOiAib2JqZWN0IiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSB0aW1lIHRha2VuIGJ5IHRoZSBOQVRTIG5ldHdvcmtpbmcgbGF5ZXIgdG8gZGVsaXZlciB0aGUgbWVzc2FnZSBpbiBuYW5vc2Vjb25kcyIsCiAgICAgICJhZGRpdGlvbmFsSXRlbXMiOiBmYWxzZSwKICAgICAgInByb3BlcnRpZXMiOiB7CiAgICAgICAgInJlcSI6IHsKICAgICAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkxhdGVuY3kgYmV0d2VlbiB0aGUgcmVxdWVzdGVyIGFuZCB0aGUgTkFUUyBzZXJ2ZXIgaXQgaXMgY29ubmVjdGVkIHRvIGluIG5hbm9zZWNvbmRzIgogICAgICAgIH0sCiAgICAgICAgInJlc3AiOiB7CiAgICAgICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJMYXRlbmN5IGJldHdlZW4gdGhlIHNlcnZpY2UgYW5kIHRoZSBOQVRTIHNlcnZlciBpdCBpcyBjb25uZWN0ZWQgdG8gaW4gbmFub3NlY29uZHMiCiAgICAgICAgfSwKICAgICAgICAic3lzIjogewogICAgICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGltZSBzcGVuZCB0cmF2ZXJzaW5nIHRoZSBOQVRTIG5ldHdvcmsgaW4gbmFub3NlY29uZHMiCiAgICAgICAgfQogICAgICB9CiAgICB9LAogICAgInRvdGFsIjogewogICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBvdmVyYWxsIHJlcXVlc3QgZHVyYXRpb24gaW4gbmFub3NlY29uZHMiCiAgICB9CiAgfQp9Cg==")
	schemas["io.nats.jetstream.advisory.v1.api_audit"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9qZXRzdHJlYW0vYWR2aXNvcnkvdjEvYXBpX2F1ZGl0Lmpzb24iLAogICJkZXNjcmlwdGlvbiI6ICJBZHZpc29yeSBwdWJsaXNoZWQgd2hlbiB0aGUgSmV0U3RyZWFtIEFQSSBpcyBhY2Nlc3NlZCBhY3Jvc3MgdGhlIG5ldHdvcmsiLAogICJ0aXRsZSI6ICJpby5uYXRzLmpldHN0cmVhbS5hZHZpc29yeS52MS5hcGlfYXVkaXQiLAogICJ0eXBlIjoib2JqZWN0IiwKICAicmVxdWlyZWQiOlsKICAgICJ0eXBlIiwKICAgICJpZCIsCiAgICAidGltZXN0YW1wIiwKICAgICJzZXJ2ZXIiLAogICAgImNsaWVudCIsCiAgICAic3ViamVjdCIsCiAgICAicmVzcG9uc2UiCiAgXSwKICAiYWRkaXRpb25hbEl0ZW1zIjogZmFsc2UsCiAgInByb3BlcnRpZXMiOiB7CiAgICAidHlwZSI6IHsKICAgICAgInR5cGUiOiJzdHJpbmciLAogICAgICAiZW51bSI6IFsKICAgICAgICAiaW8ubmF0cy5qZXRzdHJlYW0uYWR2aXNvcnkudjEuYXBpX2F1ZGl0IgogICAgICBdCiAgICB9LAogICAgImlkIjogewogICAgICAidHlwZSI6InN0cmluZyIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJVbmlxdWUgY29ycmVsYXRpb24gSUQgZm9yIHRoaXMgZXZlbnQiCiAgICB9LAogICAgInRpbWVzdGFtcCI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImZvcm1hdCI6ICJkYXRlLXRpbWUiLAogICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHRpbWUgdGhpcyBldmVudCB3YXMgY3JlYXRlZCBpbiBSRkMzMzM5IGZvcm1hdCIKICAgIH0sCiAgICAic2VydmVyIjogewogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHNlcnZlciB0aGlzIGV2ZW50IG9yaWdpbmF0ZXMgZnJvbSwgZWl0aGVyIGEgZ2VuZXJhdGVkIElEIG9yIHRoZSBjb25maWd1cmVkIG5hbWUiLAogICAgICAibWluTGVuZ3RoIjogMQogICAgfSwKICAgICJzdWJqZWN0IjogewogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHN1YmplY3QgdGhlIGFkbWluIEFQSSByZXF1ZXN0IHdhcyByZWNlaXZlZCBvbiIsCiAgICAgICJtaW5MZW5ndGgiOiAxCiAgICB9LAogICAgInJlc3BvbnNlIjp7CiAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgZnVsbCB1bnBhcnNlZCBib2R5IG9mIHRoZSByZXNwb25zZSBzZW50IHRvIHRoZSBjYWxsZXIiCiAgICB9LAogICAgInJlcXVlc3QiOiB7CiAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgZnVsbCB1bnBhcnNlZCBib2R5IG9mIHRoZSByZXF1ZXN0IHJlY2VpdmVkIGZyb20gdGhlIGNsaWVudCIKICAgIH0sCiAgICAiY2xpZW50IjogewogICAgICAidHlwZSI6ICJvYmplY3QiLAogICAgICAicmVxdWlyZWQiOiBbCiAgICAgICAgImhvc3QiLAogICAgICAgICJwb3J0IiwKICAgICAgICAiY2lkIiwKICAgICAgICAiYWNjb3VudCIKICAgICAgXSwKICAgICAgImFkZGl0aW9uYWxJdGVtcyI6IGZhbHNlLAogICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAiaG9zdCI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIElQIGFkZHJlc3Mgd2hlcmUgdGhlIGNsaWVudCBjb25uZWN0cyBmcm9tIgogICAgICAgIH0sCiAgICAgICAgInBvcnQiOiB7CiAgICAgICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgcG9ydCBudW1iZXIgd2hlcmUgdGhlIGNsaWVudCBjb25uZWN0cyBmcm9tIgogICAgICAgIH0sCiAgICAgICAgImNpZCI6IHsKICAgICAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSB1bmlxdWUgY2xpZW50IElEIHRoZSBzZXJ2ZXIgYXNzaWduZWQgdG8gdGhlIGNvbm5lY3Rpb24iCiAgICAgICAgfSwKICAgICAgICAiYWNjb3VudCI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIGFjY291bnQgdGhlIHVzZXIgYmVsb25ncyB0byIKICAgICAgICB9LAogICAgICAgICJ1c2VyIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgdXNlcm5hbWUgdGhhdCB3YXMgdXNlZCBkdXJpbmcgYXV0aGVudGljYXRpb24sIGlmIGFueSIKICAgICAgICB9LAogICAgICAgICJuYW1lIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgbmFtZSB0aGUgY2xpZW50IGFzc2lnbmVkIHRvIHRoZSBjb25uZWN0aW9uIGR1cmluZyBjb25uZWN0aW9uIG5lZ290aWF0aW9uIgogICAgICAgIH0sCiAgICAgICAgImxhbmd1YWdlIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgY2xpZW50IGxpYnJhcnkgbGFuZ3VhZ2UgdXNlZCB0byBjcmVhdGUgdGhlIGNvbm5lY3Rpb24iCiAgICAgICAgfSwKICAgICAgICAidmVyc2lvbiI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHZlcnNpb24gY2xpZW50IGxpYnJhcnkgdXNlZCB0byBjcmVhdGUgdGhlIGNvbm5lY3Rpb24iCiAgICAgICAgfQogICAgICB9CiAgICB9CiAgfQp9Cg==")
	schemas["io.nats.jetstream.advisory.v1.max_deliver"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9qZXRzdHJlYW0vYWR2aXNvcnkvdjEvbWF4X2RlbGl2ZXIuanNvbiIsCiAgImRlc2NyaXB0aW9uIjogIkFkdmlzb3J5IHB1Ymxpc2hlZCB3aGVuIGEgbWVzc2FnZSBoYXZlIHJlYWNoZWQgaXRzIG1heGltdW0gZGVsaXZlcnkgYXR0ZW1wdHMiLAogICJ0aXRsZSI6ICJpby5uYXRzLmpldHN0cmVhbS5hZHZpc29yeS52MS5tYXhfZGVsaXZlciIsCiAgInR5cGUiOiJvYmplY3QiLAogICJyZXF1aXJlZCI6WwogICAgInR5cGUiLAogICAgImlkIiwKICAgICJ0aW1lc3RhbXAiLAogICAgInN0cmVhbSIsCiAgICAiY29uc3VtZXIiLAogICAgInN0cmVhbV9zZXEiLAogICAgImRlbGl2ZXJpZXMiCiAgXSwKICAiYWRkaXRpb25hbEl0ZW1zIjogZmFsc2UsCiAgInByb3BlcnRpZXMiOiB7CiAgICAidHlwZSI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImVudW0iOiBbCiAgICAgICAgImlvLm5hdHMuamV0c3RyZWFtLmFkdmlzb3J5LnYxLm1heF9kZWxpdmVyIgogICAgICBdCiAgICB9LAogICAgImlkIjogewogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAiZGVzY3JpcHRpb24iOiAiVW5pcXVlIGNvcnJlbGF0aW9uIElEIGZvciB0aGlzIGV2ZW50IgogICAgfSwKICAgICJ0aW1lc3RhbXAiOiB7CiAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICJmb3JtYXQiOiAiZGF0ZS10aW1lIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSB0aW1lIHRoaXMgZXZlbnQgd2FzIGNyZWF0ZWQgaW4gUkZDMzMzOSBmb3JtYXQiCiAgICB9LAogICAgInN0cmVhbSI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBuYW1lIG9mIHRoZSBzdHJlYW0gd2hlcmUgdGhlIG1lc3NhZ2UgaXMgc3RvcmVkIgogICAgfSwKICAgICJjb25zdW1lciI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBuYW1lIG9mIHRoZSBjb25zdW1lciB3aGVyZSB0aGUgbWVzc2FnZSByZWFjaGVkIGl0cyBsaW1pdCIKICAgIH0sCiAgICAic3RyZWFtX3NlcSI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgIm1pbmltdW0iOiAxLAogICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHNlcXVlbmNlIG9mIHRoZSBtZXNzYWdlIGluIHRoZSBzdHJlYW0gdGhhdCBmYWlsZWQiCiAgICB9LAogICAgImRlbGl2ZXJpZXMiOiB7CiAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAibWluaW11bSI6IDEsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgbnVtYmVyIG9mIGRlbGl2ZXJpZXMgdGhhdCB3ZXJlIGF0dGVtcHRlZCIKICAgIH0KICB9Cn0K")
	schemas["io.nats.jetstream.metric.v1.consumer_ack"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9qZXRzdHJlYW0vbWV0cmljL3YxL2NvbnN1bWVyX2Fjay5qc29uIiwKICAiZGVzY3JpcHRpb24iOiAiTWV0cmljIHB1Ymxpc2hlZCB3aGVuIGEgbWVzc2FnZSB3YXMgYWNrbm93bGVkZ2VkIHRvIGEgY29uc3VtZXIgd2l0aCBBY2sgU2FtcGxpbmcgZW5hYmxlZCIsCiAgInRpdGxlIjogImlvLm5hdHMuamV0c3RyZWFtLm1ldHJpYy52MS5jb25zdW1lcl9hY2siLAogICJ0eXBlIjoib2JqZWN0IiwKICAicmVxdWlyZWQiOlsKICAgICJ0eXBlIiwKICAgICJpZCIsCiAgICAidGltZXN0YW1wIiwKICAgICJzdHJlYW0iLAogICAgImNvbnN1bWVyIiwKICAgICJzdHJlYW1fc2VxIiwKICAgICJjb25zdW1lcl9zZXEiLAogICAgImFja190aW1lIiwKICAgICJkZWxpdmVyaWVzIgogIF0sCiAgImFkZGl0aW9uYWxJdGVtcyI6IGZhbHNlLAogICJwcm9wZXJ0aWVzIjogewogICAgInR5cGUiOiB7CiAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICJlbnVtIjogWwogICAgICAgICJpby5uYXRzLmpldHN0cmVhbS5tZXRyaWMudjEuY29uc3VtZXJfYWNrIgogICAgICBdCiAgICB9LAogICAgImlkIjogewogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAiZGVzY3JpcHRpb24iOiAiVW5pcXVlIGNvcnJlbGF0aW9uIElEIGZvciB0aGlzIGV2ZW50IgogICAgfSwKICAgICJ0aW1lc3RhbXAiOiB7CiAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICJmb3JtYXQiOiAiZGF0ZS10aW1lIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSB0aW1lIHRoaXMgZXZlbnQgd2FzIGNyZWF0ZWQgaW4gUkZDMzMzOSBmb3JtYXQiCiAgICB9LAogICAgInN0cmVhbSI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBuYW1lIG9mIHRoZSBzdHJlYW0gd2hlcmUgdGhlIG1lc3NhZ2UgaXMgc3RvcmVkIgogICAgfSwKICAgICJjb25zdW1lciI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBuYW1lIG9mIHRoZSBjb25zdW1lciB3aGVyZSB0aGUgbWVzc2FnZSBpcyBoZWxkIgogICAgfSwKICAgICJzdHJlYW1fc2VxIjogewogICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgIm1pbmltdW0iOiAxLAogICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHNlcXVlbmNlIG9mIHRoZSBtZXNzYWdlIGluIHRoZSBzdHJlYW0gdGhhdCB3ZXJlIGFja25vd2xlZGdlZCIKICAgIH0sCiAgICAiY29uc3VtZXJfc2VxIjogewogICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgIm1pbmltdW0iOiAxLAogICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHNlcXVlbmNlIG9mIHRoZSBtZXNzYWdlIGluIHRoZSBjb25zdW1lciB0aGF0IHdlcmUgYWNrbm93bGVkZ2VkIgogICAgfSwKICAgICJhY2tfdGltZSI6IHsKICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICJtaW5pbXVtIjogMSwKICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSB0aW1lIGl0IHRvb2sgb24gdGhlIGZpbmFsIGRlbGl2ZXJ5IGZvciB0aGUgbWVzc2FnZSB0byBiZSBhY2tub3dsZWRnZWQgaW4gbmFub3NlY29uZHMiCiAgICB9LAogICAgImRlbGl2ZXJpZXMiOiB7CiAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAibWluaW11bSI6IDEsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgbnVtYmVyIG9mIGRlbGl2ZXJpZXMgdGhhdCB3ZXJlIGF0dGVtcHRlZCBiZWZvcmUgYmVpbmcgYWNrbm93bGVkZ2VkIgogICAgfQogIH0KfQo=")
	schemas["io.nats.jetstream.api.v1.consumer_configuration"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9qZXRzdHJlYW0vYXBpL3YxL2NvbnN1bWVyX2NvbmZpZ3VyYXRpb24uanNvbiIsCiAgImRlc2NyaXB0aW9uIjogIlRoZSBkYXRhIHN0cnVjdHVyZSB0aGF0IGRlc2NyaWJlIHRoZSBjb25maWd1cmF0aW9uIG9mIGEgTkFUUyBKZXRTdHJlYW0gQ29uc3VtZXIiLAogICJ0aXRsZSI6ICJpby5uYXRzLmpldHN0cmVhbS5hcGkudjEuY29uc3VtZXJfY29uZmlndXJhdGlvbiIsCiAgInR5cGUiOiJvYmplY3QiLAogICIkcmVmIjogImRlZmluaXRpb25zLmpzb24jL2RlZmluaXRpb25zL2NvbnN1bWVyX2NvbmZpZ3VyYXRpb24iCn0K")
	schemas["io.nats.jetstream.api.v1.stream_configuration"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9qZXRzdHJlYW0vYXBpL3YxL3N0cmVhbV9jb25maWd1cmF0aW9uLmpzb24iLAogICJkZXNjcmlwdGlvbiI6ICJUaGUgZGF0YSBzdHJ1Y3R1cmUgdGhhdCBkZXNjcmliZSB0aGUgY29uZmlndXJhdGlvbiBvZiBhIE5BVFMgSmV0U3RyZWFtIFN0cmVhbSIsCiAgInRpdGxlIjogImlvLm5hdHMuamV0c3RyZWFtLmFwaS52MS5zdHJlYW1fY29uZmlndXJhdGlvbiIsCiAgInR5cGUiOiJvYmplY3QiLAogICIkcmVmIjogImRlZmluaXRpb25zLmpzb24jL2RlZmluaXRpb25zL3N0cmVhbV9jb25maWd1cmF0aW9uIgp9Cg==")
	schemas["io.nats.jetstream.api.v1.stream_template_configuration"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9qZXRzdHJlYW0vYXBpL3YxL3N0cmVhbV90ZW1wbGF0ZV9jb25maWd1cmF0aW9uLmpzb24iLAogICJkZXNjcmlwdGlvbiI6ICJUaGUgZGF0YSBzdHJ1Y3R1cmUgdGhhdCBkZXNjcmliZSB0aGUgY29uZmlndXJhdGlvbiBvZiBhIE5BVFMgSmV0U3RyZWFtIFN0cmVhbSBUZW1wbGF0ZSIsCiAgInRpdGxlIjogImlvLm5hdHMuamV0c3RyZWFtLmFwaS52MS5zdHJlYW1fdGVtcGxhdGVfY29uZmlndXJhdGlvbiIsCiAgInR5cGUiOiJvYmplY3QiLAogICJyZXF1aXJlZCI6WwogICAgIm5hbWUiLAogICAgImNvbmZpZyIsCiAgICAibWF4X3N0cmVhbXMiCiAgXSwKICAiYWRkaXRpb25hbEl0ZW1zIjogZmFsc2UsCiAgInByb3BlcnRpZXMiOiB7CiAgICAibmFtZSI6IHsKICAgICAgImRlc2NyaXB0aW9uIjogIkEgdW5pcXVlIG5hbWUgZm9yIHRoZSBTdHJlYW0gVGVtcGxhdGUuIiwKICAgICAgIiRyZWYiOiAiZGVmaW5pdGlvbnMuanNvbiMvZGVmaW5pdGlvbnMvYmFzaWNfbmFtZSIKICAgIH0sCiAgICAibWF4X3N0cmVhbXMiOiB7CiAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgbWF4aW11bSBudW1iZXIgb2YgU3RyZWFtcyB0aGlzIFRlbXBsYXRlIGNhbiBjcmVhdGUsIC0xIGZvciB1bmxpbWl0ZWQuIiwKICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICJtaW5pbXVtIjogLTEsCiAgICAgICJkZWZhdWx0IjogLTEKICAgIH0sCiAgICAiY29uZmlnIjogewogICAgICAiJHJlZiI6ICJkZWZpbml0aW9ucy5qc29uIy9kZWZpbml0aW9ucy9zdHJlYW1fY29uZmlndXJhdGlvbiIKICAgIH0KICB9Cn0K")
	schemas["io.nats.jetstream.api.v1.definitions"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9qZXRzdHJlYW0vYXBpL3YxL2RlZmluaXRpb25zLmpzb24iLAogICJ0aXRsZSI6ICJpby5uYXRzLmpldHN0cmVhbS5hcGkudjEuZGVmaW5pdGlvbnMiLAogICJkZXNjcmlwdGlvbiI6ICJTaGFyZWQgZGVmaW5pdGlvbnMgZm9yIHRoZSBKZXRTdHJlYW0gQVBJIiwKICAidHlwZSI6ICJvYmplY3QiLAogICJkZWZpbml0aW9ucyI6IHsKICAgICJiYXNpY19uYW1lIjogewogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAicGF0dGVybiI6ICJeW14uKj5dKyQiLAogICAgICAibWluTGVuZ3RoIjogMQogICAgfSwKICAgICJkZWxpdmVyX3BvbGljeSI6IHsKICAgICAgIm9uZU9mIjogWwogICAgICAgIHsiJHJlZiI6ICIjL2RlZmluaXRpb25zL2FsbF9kZWxpdmVyX3BvbGljeSJ9LAogICAgICAgIHsiJHJlZiI6ICIjL2RlZmluaXRpb25zL2xhc3RfZGVsaXZlcl9wb2xpY3kifSwKICAgICAgICB7IiRyZWYiOiAiIy9kZWZpbml0aW9ucy9uZXdfZGVsaXZlcl9wb2xpY3kifSwKICAgICAgICB7IiRyZWYiOiAiIy9kZWZpbml0aW9ucy9zdGFydF9zZXF1ZW5jZV9kZWxpdmVyX3BvbGljeSJ9LAogICAgICAgIHsiJHJlZiI6ICIjL2RlZmluaXRpb25zL3N0YXJ0X3RpbWVfZGVsaXZlcl9wb2xpY3kifQogICAgICBdCiAgICB9LAogICAgImFsbF9kZWxpdmVyX3BvbGljeSI6IHsKICAgICAgInJlcXVpcmVkIjogWyJkZWxpdmVyX3BvbGljeSJdLAogICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAiZGVsaXZlcl9wb2xpY3kiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImVudW0iOiBbImFsbCJdCiAgICAgICAgfQogICAgICB9CiAgICB9LAogICAgImxhc3RfZGVsaXZlcl9wb2xpY3kiOiB7CiAgICAgICJyZXF1aXJlZCI6IFsiZGVsaXZlcl9wb2xpY3kiXSwKICAgICAgInByb3BlcnRpZXMiOiB7CiAgICAgICAgImRlbGl2ZXJfcG9saWN5IjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJlbnVtIjogWyJsYXN0Il0KICAgICAgICB9CiAgICAgIH0KICAgIH0sCiAgICAibmV3X2RlbGl2ZXJfcG9saWN5IjogewogICAgICAicmVxdWlyZWQiOiBbImRlbGl2ZXJfcG9saWN5Il0sCiAgICAgICJwcm9wZXJ0aWVzIjogewogICAgICAgICJkZWxpdmVyX3BvbGljeSI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZW51bSI6IFsibmV3Il0KICAgICAgICB9CiAgICAgIH0KICAgIH0sCiAgICAic3RhcnRfc2VxdWVuY2VfZGVsaXZlcl9wb2xpY3kiOiB7CiAgICAgICJyZXF1aXJlZCI6IFsiZGVsaXZlcl9wb2xpY3kiLCAib3B0X3N0YXJ0X3NlcSJdLAogICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAiZGVsaXZlcl9wb2xpY3kiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImVudW0iOiBbImJ5X3N0YXJ0X3NlcXVlbmNlIl0KICAgICAgICB9LAogICAgICAgICJvcHRfc3RhcnRfc2VxIjogewogICAgICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICAgICAibWluaW11bSI6IDAKICAgICAgICB9CiAgICAgIH0KICAgIH0sCiAgICAic3RhcnRfdGltZV9kZWxpdmVyX3BvbGljeSI6IHsKICAgICAgInJlcXVpcmVkIjogWyJkZWxpdmVyX3BvbGljeSIsICJvcHRfc3RhcnRfdGltZSJdLAogICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAiZGVsaXZlcl9wb2xpY3kiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImVudW0iOiBbImJ5X3N0YXJ0X3RpbWUiXQogICAgICAgIH0sCiAgICAgICAgIm9wdF9zdGFydF90aW1lIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJmb3JtYXQiOiAiZGF0ZS10aW1lIgogICAgICAgIH0KICAgICAgfQogICAgfSwKICAgICJjb25zdW1lcl9jb25maWd1cmF0aW9uIjogewogICAgICAicmVxdWlyZWQiOlsKICAgICAgICAiZGVsaXZlcl9wb2xpY3kiLAogICAgICAgICJhY2tfcG9saWN5IiwKICAgICAgICAicmVwbGF5X3BvbGljeSIKICAgICAgXSwKICAgICAgImFsbE9mIjogW3siJHJlZiI6ICIjL2RlZmluaXRpb25zL2RlbGl2ZXJfcG9saWN5In1dLAogICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAiZHVyYWJsZV9uYW1lIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkEgdW5pcXVlIG5hbWUgZm9yIGEgZHVyYWJsZSBjb25zdW1lciIsCiAgICAgICAgICAiJHJlZiI6ICIjL2RlZmluaXRpb25zL2Jhc2ljX25hbWUiCiAgICAgICAgfSwKICAgICAgICAiZGVsaXZlcl9zdWJqZWN0IjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJtaW5MZW5ndGgiOiAxCiAgICAgICAgfSwKICAgICAgICAiYWNrX3BvbGljeSI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZW51bSI6IFsibm9uZSIsICJhbGwiLCAiZXhwbGljaXQiXQogICAgICAgIH0sCiAgICAgICAgImFja193YWl0IjogewogICAgICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICAgICAibWluaW11bSI6IDEKICAgICAgICB9LAogICAgICAgICJtYXhfZGVsaXZlciI6IHsKICAgICAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAgICAgIm1pbmltdW0iOiAtMQogICAgICAgIH0sCiAgICAgICAgImZpbHRlcl9zdWJqZWN0IjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIgogICAgICAgIH0sCiAgICAgICAgInJlcGxheV9wb2xpY3kiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImVudW0iOiBbImluc3RhbnQiLCAib3JpZ2luYWwiXQogICAgICAgIH0sCiAgICAgICAgInNhbXBsZV9mcmVxIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIgogICAgICAgIH0KICAgICAgfQogICAgfSwKICAgICJzdHJlYW1fY29uZmlndXJhdGlvbiI6IHsKICAgICAgInR5cGUiOiAib2JqZWN0IiwKICAgICAgInJlcXVpcmVkIjpbCiAgICAgICAgInJldGVudGlvbiIsCiAgICAgICAgIm1heF9jb25zdW1lcnMiLAogICAgICAgICJtYXhfbXNncyIsCiAgICAgICAgIm1heF9ieXRlcyIsCiAgICAgICAgIm1heF9hZ2UiLAogICAgICAgICJzdG9yYWdlIiwKICAgICAgICAibnVtX3JlcGxpY2FzIgogICAgICBdLAogICAgICAiYWRkaXRpb25hbEl0ZW1zIjogZmFsc2UsCiAgICAgICJwcm9wZXJ0aWVzIjogewogICAgICAgICJuYW1lIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkEgdW5pcXVlIG5hbWUgZm9yIHRoZSBTdHJlYW0sIGVtcHR5IGZvciBTdHJlYW0gVGVtcGxhdGVzLiIsCiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgInBhdHRlcm4iOiAiXlteLio+XSokIiwKICAgICAgICAgICJtaW5MZW5ndGgiOiAwCiAgICAgICAgfSwKICAgICAgICAic3ViamVjdHMiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiQSBsaXN0IG9mIHN1YmplY3RzIHRvIGNvbnN1bWUsIHN1cHBvcnRzIHdpbGRjYXJkcy4iLAogICAgICAgICAgInR5cGUiOiAiYXJyYXkiLAogICAgICAgICAgIm1pbkxlbmd0aCI6IDEsCiAgICAgICAgICAiaXRlbXMiOiB7CiAgICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAgICJtaW5MZW5ndGgiOiAxCiAgICAgICAgICB9CiAgICAgICAgfSwKICAgICAgICAicmV0ZW50aW9uIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkhvdyBtZXNzYWdlcyBhcmUgcmV0YWluZWQgaW4gdGhlIFN0cmVhbSwgb25jZSB0aGlzIGlzIGV4Y2VlZGVkIG9sZCBtZXNzYWdlcyBhcmUgcmVtb3ZlZC4iLAogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJlbnVtIjogWyJsaW1pdHMiLCAiaW50ZXJlc3QiLCAid29ya3F1ZXVlIl0sCiAgICAgICAgICAiZGVmYXVsdCI6ICJsaW1pdHMiCiAgICAgICAgfSwKICAgICAgICAibWF4X2NvbnN1bWVycyI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJIb3cgbWFueSBDb25zdW1lcnMgY2FuIGJlIGRlZmluZWQgZm9yIGEgZ2l2ZW4gU3RyZWFtLiAtMSBmb3IgdW5saW1pdGVkLiIsCiAgICAgICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgICAgICJtaW5pbXVtIjogLTEsCiAgICAgICAgICAiZGVmYXVsdCI6IC0xCiAgICAgICAgfSwKICAgICAgICAibWF4X21zZ3MiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiSG93IG1hbnkgbWVzc2FnZXMgbWF5IGJlIGluIGEgU3RyZWFtLCBvbGRlc3QgbWVzc2FnZXMgd2lsbCBiZSByZW1vdmVkIGlmIHRoZSBTdHJlYW0gZXhjZWVkcyB0aGlzIHNpemUuIC0xIGZvciB1bmxpbWl0ZWQuIiwKICAgICAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAgICAgIm1pbmltdW0iOiAtMSwKICAgICAgICAgICJkZWZhdWx0IjogLTEKICAgICAgICB9LAogICAgICAgICJtYXhfYnl0ZXMiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiSG93IGJpZyB0aGUgU3RyZWFtIG1heSBiZSwgd2hlbiB0aGUgY29tYmluZWQgc3RyZWFtIHNpemUgZXhjZWVkcyB0aGlzIG9sZCBtZXNzYWdlcyBhcmUgcmVtb3ZlZC4gLTEgZm9yIHVubGltaXRlZC4iLAogICAgICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICAgICAibWluaW11bSI6IC0xLAogICAgICAgICAgImRlZmF1bHQiOiAtMQogICAgICAgIH0sCiAgICAgICAgIm1heF9hZ2UiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTWF4aW11bSBhZ2Ugb2YgYW55IG1lc3NhZ2UgaW4gdGhlIHN0cmVhbSwgZXhwcmVzc2VkIGluIG1pY3Jvc2Vjb25kcy4gLTEgZm9yIHVubGltaXRlZC4iLAogICAgICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICAgICAibWluaW11bSI6IDAsCiAgICAgICAgICAiZGVmYXVsdCI6IDAKICAgICAgICB9LAogICAgICAgICJtYXhfbXNnX3NpemUiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIGxhcmdlc3QgbWVzc2FnZSB0aGF0IHdpbGwgYmUgYWNjZXB0ZWQgYnkgdGhlIFN0cmVhbS4gLTEgZm9yIHVubGltaXRlZC4iLAogICAgICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICAgICAibWluaW11bSI6IC0xLAogICAgICAgICAgImRlZmF1bHQiOiAtMQogICAgICAgIH0sCiAgICAgICAgInN0b3JhZ2UiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHN0b3JhZ2UgYmFja2VuZCB0byB1c2UgZm9yIHRoZSBTdHJlYW0uIiwKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZW51bSI6IFsiZmlsZSIsICJtZW1vcnkiXQogICAgICAgIH0sCiAgICAgICAgIm51bV9yZXBsaWNhcyI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJIb3cgbWFueSByZXBsaWNhcyB0byBrZWVwIGZvciBlYWNoIG1lc3NhZ2UuIiwKICAgICAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAgICAgIm1pbmltdW0iOiAxLAogICAgICAgICAgImRlZmF1bHQiOiAxCiAgICAgICAgfSwKICAgICAgICAibm9fYWNrIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkRpc2FibGVzIGFja25vd2xlZGdpbmcgbWVzc2FnZXMgdGhhdCBhcmUgcmVjZWl2ZWQgYnkgdGhlIFN0cmVhbS4iLAogICAgICAgICAgInR5cGUiOiAiYm9vbGVhbiIsCiAgICAgICAgICAiZGVmYXVsdCI6IGZhbHNlCiAgICAgICAgfSwKICAgICAgICAidGVtcGxhdGVfb3duZXIiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiV2hlbiB0aGUgU3RyZWFtIGlzIG1hbmFnZWQgYnkgYSBTdHJlYW0gVGVtcGxhdGUgdGhpcyBpZGVudGlmaWVzIHRoZSB0ZW1wbGF0ZSB0aGF0IG1hbmFnZXMgdGhlIFN0cmVhbS4iLAogICAgICAgICAgInR5cGUiOiAic3RyaW5nIgogICAgICAgIH0KICAgICAgfQogICAgfQogIH0KfQo=")
}
