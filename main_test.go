package main

import (
	"encoding/xml"
	"net/http/httptest"
	"strings"
	"testing"
)

type doc struct {
	Region struct {
		ID  string `xml:"id,attr"`
		Pup struct {
			Version string `xml:"version,attr"`
			Image   string `xml:"update_data>image"`
		} `xml:"system_pup"`
		Rec string `xml:"recovery_pup>image"`
	} `xml:"region"`
}

func get(t *testing.T, path string) (int, string) {
	rec := httptest.NewRecorder()
	handler().ServeHTTP(rec, httptest.NewRequest("GET", path, nil))
	return rec.Code, rec.Body.String()
}

func TestList(t *testing.T) {
	t.Setenv("PS4_UPDATE_VERSION", "09.000.000")
	code, body := get(t, "/update/ps4/list/us/ps4-updatelist.xml")
	var d doc
	if err := xml.Unmarshal([]byte(body), &d); code != 200 || err != nil {
		t.Fatalf("code %d err %v", code, err)
	}
	if d.Region.ID != "us" || d.Region.Pup.Version != "09.000.000" || d.Region.Pup.Image != "" || d.Region.Rec != "" {
		t.Fatalf("unexpected %+v", d)
	}
	t.Setenv("PS4_UPDATE_IMAGE_URL", "http://x/PS4UPDATE.PUP")
	_, body = get(t, "/update/ps4/list/us/ps4-updatelist.xml")
	d = doc{}
	if err := xml.Unmarshal([]byte(body), &d); err != nil || !strings.HasSuffix(d.Region.Pup.Image, "?dest=us") || !strings.HasSuffix(d.Region.Rec, "?dest=us") {
		t.Fatalf("image missing: %v %+v", err, d)
	}
	if c, _ := get(t, "/nope"); c != 404 {
		t.Fatalf("want 404, got %d", c)
	}
	if c, b := get(t, "/healthz"); c != 200 || b != "ok" {
		t.Fatalf("healthz %d %q", c, b)
	}
}
