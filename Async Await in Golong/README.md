# Implementing Async/Await

To implement async/await in Golang, we will start with a package directory named async. The project structure looks like
```
├── async
│   └── async.go
├── main.go
└── README.md
```
In the async file, we write the simplest future interface that can handle async tasks.
```
package async

import "context"

// Future interface has the method signature for await
type Future interface {
	Await() interface{}
}

type future struct {
	await func(ctx context.Context) interface{}
}

func (f future) Await() interface{} {
	return f.await(context.Background())
}

// Exec executes the async function
func Exec(f func() interface{}) Future {
	var result interface{}
	c := make(chan struct{})
	go func() {
		defer close(c)
		result = f()
	}()
	return future{
		await: func(ctx context.Context) interface{} {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				return result
			}
		},
	}
}

```

Not a lot is happening here, we add a Future interface that has the Await method signature. Next, we add a future struct that holds one value, a function signature of the await function. Now futute struct implements Future interface's Await method by invoking its own await function.

Next in the Exec function, we execute the passed function asynchronously in goroutine. And we return the await function. It waits for the channel to close or context to read from. Based on whichever happens first, it either returns the error or the result which is an interface.

Now armed with this new async package, let's see how we can change our current go code
```
package main

import (
	"fmt"
	"time"

	"./async"
)

func DoneAsync() int {
	fmt.Println("Warming up...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done...")
	return 1
}

func main() {
	fmt.Println("Let's start...")
	future := async.Exec(func() interface{} {
		return DoneAsync()
	})
	fmt.Println("Done is running...")
	val := future.Await()
	fmt.Println(val)
}
```

At the first glance, it looks much cleaner, we are not explicitly working with goroutine or channels here. Our DoneAsync function has been changed to a completely synchronous nature. In the main function, we use the async package's Exect method to handle DoneAsync. Which starts execution of DoneAsync. The control flow is returned back to main function which can execute other pieces of codes until we call await on it. Finally, we make blocking call to Await and read back data.

Now the code looks much simpler and easier to read. We can modify our async package to incorporate a lot of other types of asynchronous tasks in Golang, but we would just stick to simple implementation for now in this tutorial.

## Conclusion
We have gone through what async/await it and implemented a simple version of that in Golang. I would encourage you to look into async/await a lot more and see how it can ease the readability of the codebase much better.