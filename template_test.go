package jsm_test

import (
	"testing"

	"github.com/nats-io/jsm.go"
)

func TestNewStreamTemplate(t *testing.T) {
	srv, _, mgr := startJSServer(t)
	defer srv.Shutdown()

	_, err := mgr.NewStreamTemplate("orders_templ", 1, jsm.DefaultStream, jsm.FileStorage(), jsm.Subjects("ORDERS.*"))
	checkErr(t, err, "new stream template failed")

	templ, err := mgr.LoadStreamTemplate("orders_templ")
	checkErr(t, err, "load stream template failed")

	if templ.Name() != "orders_templ" {
		t.Fatalf("expected name==orders_templ got %q", templ.Name())
	}
}

func TestNewOrLoadStreamTemplate(t *testing.T) {
	srv, _, mgr := startJSServer(t)
	defer srv.Shutdown()

	first, err := mgr.NewStreamTemplate("orders_templ", 1, jsm.DefaultStream, jsm.FileStorage(), jsm.Subjects("ORDERS.*"))
	checkErr(t, err, "new stream template failed")

	second, err := mgr.LoadOrNewStreamTemplate("orders_templ", 1, jsm.DefaultStream, jsm.FileStorage(), jsm.Subjects("ORDERS.*"))
	checkErr(t, err, "load or new stream template failed")

	if first.Name() != second.Name() {
		t.Fatalf("got wrong template, expected %q got %q", first.Name(), second.Name())
	}
}

func TestStreamTemplate_Delete(t *testing.T) {
	srv, _, mgr := startJSServer(t)
	defer srv.Shutdown()

	templ, err := mgr.NewStreamTemplate("orders_templ", 1, jsm.DefaultStream, jsm.FileStorage(), jsm.Subjects("ORDERS.*"))
	checkErr(t, err, "new stream template failed")

	names, err := mgr.StreamTemplateNames()
	checkErr(t, err, "names failed")

	if len(names) != 1 || names[0] != "orders_templ" {
		t.Fatalf("expected [orders_templ] got %q", names)
	}

	err = templ.Delete()
	checkErr(t, err, "delete failed")

	names, err = mgr.StreamTemplateNames()
	checkErr(t, err, "names failed")
	if len(names) != 0 {
		t.Fatalf("expected [] got %q", names)
	}
}

func TestStreamTemplate_Reset(t *testing.T) {
	srv, nc, mgr := startJSServer(t)
	defer srv.Shutdown()

	templ, err := mgr.NewStreamTemplate("orders_templ", 1, jsm.DefaultStream, jsm.FileStorage(), jsm.Subjects("ORDERS.*"))
	checkErr(t, err, "new stream template failed")

	if len(templ.Streams()) != 0 {
		t.Fatalf("expected no streams got %q", templ.Streams())
	}

	err = nc.Publish("ORDERS.1", []byte("hello"))
	checkErr(t, err, "publish failed")

	err = templ.Reset()
	checkErr(t, err, "reset failed")

	if len(templ.Streams()) != 1 || templ.Streams()[0] != "ORDERS_1" {
		t.Fatalf("expected [ORDERS_1] got %q", templ.Streams())
	}
}
