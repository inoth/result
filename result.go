package result

type Result[T any] struct {
	val T
	err error
}

func (r Result[T]) Ok() bool {
	return r.err == nil
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.val
}

func (r Result[T]) Expect(errs ...error) T {
	if r.err != nil {
		if len(errs) > 0 {
			panic(errs[0])
		}
		panic(r.err)
	}
	return r.val
}

func (r Result[T]) Match(okFn func(val T), errFn func(err error)) {
	if r.Ok() {
		okFn(r.val)
	} else {
		errFn(r.err)
	}
}

func New[T any](val T, errs ...error) Result[T] {
	res := Result[T]{
		val: val,
	}
	if len(errs) > 0 {
		res.err = errs[0]
	}
	return res
}

func Must[T any](val T, err error) Result[T] {
	return Result[T]{
		val: val,
		err: err,
	}
}
