# stopwatch [![GoDoc](https://godoc.org/github.com/1800alex/go-utilities-stopwatch?status.svg)](https://godoc.org/github.com/1800alex/go-utilities-stopwatch) [![Build Status](https://travis-ci.com/1800alex/go-utilities-stopwatch.png?branch=master)](https://travis-ci.com/1800alex/go-utilities-stopwatch)
Package stopwatch is a package that makes it simple to measure elapsed times.

Download:
```shell
go get github.com/1800alex/go-utilities-stopwatch
```

* * *
Package stopwatch is a package that makes it simple to measure elapsed times.





# Examples

Stopwatch
Code:

```
{
	sw := NewStopwatch()
	sw.Start()
	time.Sleep(250 * time.Millisecond)
	fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())
	sw.Stop()
	time.Sleep(250 * time.Millisecond)
	fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())
	sw.Start()
	time.Sleep(250 * time.Millisecond)
	fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())
	sw.Reset()
	sw.Start()
	time.Sleep(250 * time.Millisecond)
	fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())
}
```



