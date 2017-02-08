package redi2go

import redigo "github.com/garyburd/redigo/redis"

func (r *Reply) String() (string, error) {
	if r.err != nil {
		return "", r.err
	}

	return redigo.String(r.value, nil)
}

func (r *Reply) Bytes() ([]byte, error) {
	if r.err != nil {
		return nil, r.err
	}

	return redigo.Bytes(r.value, nil)
}

func (r *Reply) Float64() (float64, error) {
	if r.err != nil {
		return 0, r.err
	}

	return redigo.Float64(r.value, nil)
}

func (r *Reply) ByteSlice() ([][]byte, error) {
	if r.err != nil {
		return [][]byte(nil), r.err
	}

	return redigo.ByteSlices(r.value, nil)
}

func (r *Reply) Strings() ([]string, error) {
	if r.err != nil {
		return []string{}, r.err
	}

	return redigo.Strings(r.value, nil)
}

func (r *Reply) Int() (int, error) {
	if r.err != nil {
		return 0, r.err
	}

	return redigo.Int(r.value, nil)
}
