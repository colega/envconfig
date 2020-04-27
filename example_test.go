package envconfig_test

import (
	"fmt"
	"os"

	"github.com/colega/envconfig"
)

func ExampleUnused() {
	os.Clearenv()
	os.Setenv("APP_VERSION", "1.0.0")
	os.Setenv("APP_TAG", "this is not used")

	cfg := struct{ Version string }{}
	unused, _ := envconfig.Unused("app", &cfg)

	for _, v := range unused {
		fmt.Println(v)
	}
	// Output:
	// APP_TAG
}
