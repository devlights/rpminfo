package rpm_test

import (
	"fmt"

	"github.com/devlights/rpminfo/pkg/rpm"
)

func ExampleRpm() {
	const (
		pkgName = "openssl-1.1.1c-2.el8.x86_64"
	)

	r := rpm.Parse(pkgName)
	r.SetOutputPattern(rpm.RpmOutputNewLine)
	fmt.Println(r)

	// Output:
	// name   : openssl
	// version: 1.1.1c
	// rel    : 2.el8
	// arch   : x86_64
}
