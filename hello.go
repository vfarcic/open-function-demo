package hello

import (
	"fmt"
	"net/http"

	ofctx "github.com/OpenFunction/functions-framework-go/context"

	"github.com/OpenFunction/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloWorld", HelloWorld,
		functions.WithFunctionPath("/{greeting}"))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	vars := ofctx.VarsFromCtx(r.Context())
	fmt.Fprintf(w, "Hello, %s!\n", vars["greeting"])
}
