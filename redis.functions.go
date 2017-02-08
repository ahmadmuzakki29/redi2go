package redi2go

func (r Redis) Do(command string, args ...interface{}) (interface{}, error) {
	if r.Mock != nil {
		return r.Mock.Do(command, args...)
	}
	c := r.Pool.Get()
	defer c.Close()

	return c.Do(command, args...)
}

func (r *Redis) cmd(command string, args ...interface{}) *Reply {
	reply := &Reply{}
	data, err := r.Do(command, args...)
	if err != nil {
		reply.err = err
		return reply
	}

	reply.value = data

	return reply
}

func (r *Redis) Get(key string, args ...interface{}) *Reply {
	args = append([]interface{}{key}, args...)
	return r.cmd("GET", args...)
}

func (r Redis) HGet(key string, args ...interface{}) *Reply {
	args = append([]interface{}{key}, args...)
	return r.cmd("HGET", args...)
}

func (r Redis) Set(key string, value interface{}, args ...interface{}) error {
	args = append([]interface{}{key, value}, args...)
	return r.cmd("SET", args...).err
}

func (r Redis) HSet(key, field string, value interface{}, args ...interface{}) error {
	args = append([]interface{}{key, field, value}, args...)
	return r.cmd("HSET", args...).err
}

func (r Redis) SMembers(key string, args ...interface{}) *Reply {
	args = append([]interface{}{key}, args...)
	return r.cmd("SMEMBERS", args...)
}

func (r Redis) Del(keys ...string) error {
	args := make([]interface{}, len(keys))
	for i, v := range keys {
		args[i] = v
	}
	return r.cmd("DEL", args...).err
}

func (r Redis) HKeys(key string, args ...interface{}) *Reply {
	args = append([]interface{}{key}, args...)
	return r.cmd("HKEYS", args...)
}

func (r Redis) Expire(key string, args ...interface{}) error {
	args = append([]interface{}{key}, args...)
	return r.cmd("EXPIRE", args...).err
}

func (r Redis) Publish(channel string, args ...interface{}) error {
	args = append([]interface{}{channel}, args...)
	return r.cmd("PUBLISH", args...).err
}
