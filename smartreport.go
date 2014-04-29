
package smartreport

/*
#include <smartmontools/smart_report.h>
#cgo linux LDFLAGS: -L/usr/local/lib -lsmartreport -lstdc++
#cgo darwin LDFLAGS: -L/usr/local/lib -lsmartreport -framework CoreFoundation -framework IOKit -lstdc++
*/
import "C"

func init() {
    C.smart_init();
}

/**
 * Wrap smart_report_create from libsmartreport
 */
func Report(path string) string {
    report := C.smart_report_create(C.CString(path))
    var result string;
    if report != nil {
        result = C.GoString(report.body)
    } else {
        result = ""
    }
    C.smart_report_free(report)
    return result
}

