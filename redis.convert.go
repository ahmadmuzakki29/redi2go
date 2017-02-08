package redi2go

import redigo "github.com/garyburd/redigo/redis"

func (r *Reply) String() (string, *Error) {
	if r.err != nil {
		return "", r.err
	}

	val, err := redigo.String(r.value, nil)
	return val, &Error{err}
}

func (r *Reply) Bytes() ([]byte, *Error) {
	if r.err != nil {
		return nil, r.err
	}

	val, err := redigo.Bytes(r.value, nil)
	return val, &Error{err}
}

func (r *Reply) Float64() (float64, *Error) {
	if r.err != nil {
		return 0, r.err
	}

	val, err := redigo.Float64(r.value, nil)
	return val, &Error{err}
}

func (r *Reply) ByteSlice() ([][]byte, *Error) {
	if r.err != nil {
		return [][]byte(nil), r.err
	}

	val, err := redigo.ByteSlices(r.value, nil)
	return val, &Error{err}
}

func (r *Reply) Strings() ([]string, *Error) {
	if r.err != nil {
		return []string{}, r.err
	}

	val, err := redigo.Strings(r.value, nil)
	return val, &Error{err}
}

func (r *Reply) Int() (int, *Error) {
	if r.err != nil {
		return 0, r.err
	}

	val, err := redigo.Int(r.value, nil)
	return val, &Error{err}
}
