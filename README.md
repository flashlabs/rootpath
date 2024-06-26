# Rootpath - Testing Just Got Easier

The package in its current version is designed to make it easier to perform tests
that require loading external resources. It is not intended for production applications,
as its operation is based on the existence of the `go.mod` file in the main directory of the application.

## Usage

### Installation

```shell
go get github.com/flashlabs/rootpath
```

### Tests Bootstrap

Configure the test working directory to always be the project root directory. It doesn't matter whether you run tests
from the `IDE` or from the `CLI`, the context of their execution will always be the same.

Just import the package with a blank identifier (`_`) and the working directory of your tests 
will be the root directory of your project.

```go
package yours

import (
    _ "github.com/flashlabs/rootpath"
)
```

From now on you can load all files pointing to the rootpath from the root directory of your project, 
e.g. `"assets/file.json"`, no matter where the test will be called from.

See ["test/rootpath_test.go"](https://github.com/flashlabs/rootpath/blob/main/test/rootpath_test.go) for real life example.

### Manual Initialization

It is possible to manually chdir to `rootpath` with the `location.Chdir()` function.

All you need to do is to `NOT` `blank import` the `rootpath` package,
but instead call `location.Chdir()` when you want to:

```go
package yours

import (
	"fmt"

	"github.com/flashlabs/rootpath/location"
)

func main() {
	if err := location.Chdir(); err != nil {
		fmt.Errorf("error while executing location.Chdir: %w", err)
	}
}
```

See ["location/chdir_test.go"](https://github.com/flashlabs/rootpath/blob/main/location/chdir_test.go) for real life example.

## Issues and Ideas

If you have any questions or want to improve something, just submit 
an [Issue](https://github.com/flashlabs/rootpath/issues) with all the details.
