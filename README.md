# futureGoLang
future interface implementation similar to Java Future.

This is a very basic implementation where Get, GetWithTimeout, isCompleted, isCancelled and cancel functions are implemented.

# how to use

```
future := newFuture( func() T{
    ...
  }
)

result := future.get()
```

# reference

https://docs.oracle.com/javase/8/docs/api/index.html?java/util/concurrent/Future.html
