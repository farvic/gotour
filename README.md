# Go Tour

All the exercises from [Go Tour](https://go.dev/tour/list).
In order to run the a code, just put the .go file on a separate and rename the whateverIsThereMain function to just main, since Golang only allows you to have one package per folder and only one function with the same name for each folder, so the main functions in each file should belong to different folders.

Example:

```
package main

import (
	"fmt"
	"time"
)

func mainhello() {

	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
}
```

The main function here is mainhello. You should rename it to main.

The files are on the same folder just so that you can navigate through them without having to open different folders.
