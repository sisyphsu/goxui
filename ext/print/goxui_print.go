package print

/*
#cgo darwin LDFLAGS: -L./darwin
#cgo darwin LDFLAGS: -F/usr/local/opt/qt/lib
#cgo darwin LDFLAGS: -framework Carbon
#cgo darwin LDFLAGS: -framework Cocoa
#cgo darwin LDFLAGS: -lc++
#cgo darwin LDFLAGS: -framework QtCore
#cgo darwin LDFLAGS: -framework QtWidgets
#cgo darwin LDFLAGS: -framework QtQuick
#cgo darwin LDFLAGS: -framework QtGui
#cgo darwin LDFLAGS: -framework QtQml
#cgo darwin LDFLAGS: -framework QtNetwork
#cgo darwin LDFLAGS: -framework QtConcurrent
#cgo darwin LDFLAGS: -framework QtPrintSupport

#cgo LDFLAGS: -lgoxui-print

#include "goxui_print.h"
*/
import "C"
import "github.com/go-eden/goxui/core"

func init() {
	core.AddInitCallback(doInit)
}

func doInit() {
	C.ui_init_print()
}
