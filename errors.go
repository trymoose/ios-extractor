package main

func Defer(fn func() error) {
	Must(fn())
}

func Get[T any](t T, err error) T {
	Must(err)
	return t
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func OnFail(fn func()) (success func(), defr func()) {
	fail := true
	return func() { fail = false }, func() {
		if fail {
			fn()
		}
	}
}
