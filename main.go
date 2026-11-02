// Command ps4-update serves the PlayStation 4 system-update list.
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var regionRe = regexp.MustCompile(`^[a-z]{2}$`)

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

// list renders the update list. Lines end in CRLF like Sony's, attributes in Sony's order.
func list(region string) string {
	ver := env("PS4_UPDATE_VERSION", "00.000.000")
	label := env("PS4_UPDATE_LABEL", "0.00")
	sdk := env("PS4_UPDATE_SDK_VERSION", ver)
	img := os.Getenv("PS4_UPDATE_IMAGE_URL")
	force := "00.000.000"
	pup := fmt.Sprintf(`system_pup ex_version="0" label="%s" sdk_version="%s" version="%s"`, label, sdk, ver)
	var b strings.Builder
	w := func(s string) { b.WriteString(s + "\r\n") }
	w(`<?xml version="1.0" ?>`)
	w(`<update_data_list>`)
	w(fmt.Sprintf("\t"+`<region id="%s">`, region))
	if img != "" {
		force = ver
	}
	w("\t\t<force_update>")
	w(fmt.Sprintf("\t\t\t"+`<system level0_system_ex_version="0" level0_system_version="%s" level1_system_ex_version="0" level1_system_version="%s"/>`, force, force))
	w("\t\t</force_update>")
	if img == "" {
		w("\t\t<" + pup + "/>")
	} else {
		// ponytail: one image for both PUPs, size="0"; consoles tolerate it, add a size env if one ever cares.
		url := fmt.Sprintf("%s?dest=%s", img, region)
		w("\t\t<" + pup + ">")
		w("\t\t\t" + `<update_data update_type="full">`)
		w("\t\t\t\t" + `<image size="0">` + url + `</image>`)
		w("\t\t\t" + `</update_data>`)
		w("\t\t</system_pup>")
		w("\t\t" + `<recovery_pup type="default">`)
		w("\t\t\t<" + pup + "/>")
		w("\t\t\t" + `<image size="0">` + url + `</image>`)
		w("\t\t</recovery_pup>")
	}
	w("\t</region>")
	w("</update_data_list>")
	return b.String()
}

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "ok") })
	mux.HandleFunc("GET /update/ps4/list/{region}/ps4-updatelist.xml", func(w http.ResponseWriter, r *http.Request) {
		region := r.PathValue("region")
		// ponytail: any two-letter code is accepted; Sony's region list is not enforced.
		if !regionRe.MatchString(region) {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprint(w, list(region))
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "path", r.URL.Path, "remote", r.RemoteAddr)
		mux.ServeHTTP(w, r)
	})
}

func main() {
	addr := env("PS4_UPDATE_ADDR", ":8540")
	slog.Info("listening", "addr", addr)
	if err := http.ListenAndServe(addr, handler()); err != nil {
		slog.Error("serve", "err", err)
		os.Exit(1)
	}
}
