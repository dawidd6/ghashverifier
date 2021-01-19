package main

//// This header is generated by glib-compile-resources in Makefile
// #cgo pkg-config: gio-2.0
// #include "resources.h"
import "C"

import (
	"os"
	"unsafe"

	"github.com/dawidd6/checksumo/src/controller"
	"github.com/dawidd6/checksumo/src/model"

	"github.com/dawidd6/checksumo/src/view"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// Those are set via -ldflags in Makefile
var (
	appID        string
	localeDomain string
	localeDir    string
	uiResource   string
)

func main() {
	// Initialize i18n
	glib.InitI18n(localeDomain, localeDir)

	// Create components
	m := model.New()
	v := view.New()
	c := controller.New(v, m)

	// Set program name to app ID,
	// so DEs could recognize the app and desktop file
	//
	// This should be upstreamed to gotk3
	appIDc := (*C.gchar)(C.CString(appID))
	C.g_set_prgname(appIDc)
	C.free(unsafe.Pointer(appIDc))

	// Create app
	app, _ := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	app.Connect("activate", v.Activate, uiResource)
	app.Connect("activate", c.Activate)

	// Run app
	os.Exit(app.Run(os.Args))
}
