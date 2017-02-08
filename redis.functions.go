package redi2go

func (r Redis) Do(command string, args ...interface{}) (interface{}, error) {
	if r.Mock != nil {
		return r.Mock.Do(command, args...)
	}
	c := r.pool.Get()
	defer c.Close()

	return c.Do(command, args...)
}

func (r *Redis) Cmd(command string, args ...interface{}) *Reply {
	reply := &Reply{}
	data, err := r.Do(command, args...)
	if err != nil {
		reply.err = &Error{err}
		return reply
	}

	reply.value = data

	return reply
}

func (r *Redis) Get(key string, args ...interface{}) *Reply {
	args = append([]interface{}{key}, args...)
	return r.Cmd("GET", args...)
}

func (r Redis) HGet(key string, args ...interface{}) *Reply {
	args = append([]interface{}{key}, args...)
	return r.Cmd("HGET", args...)
}

func (r Redis) Set(key string, value interface{}, args ...interface{}) *Reply {
	args = append([]interface{}{key, value}, args...)
	return r.Cmd("SET", args...)
}

func (r Redis) HSet(key, field string, value interface{}, args ...interface{}) *Reply {
	args = append([]interface{}{key, field, value}, args...)
	return r.Cmd("HSET", args...)
}

func (r Redis) SMembers(key string, args ...interface{}) *Reply {
	args = append([]interface{}{key}, args...)
	return r.Cmd("SMEMBERS", args...)
}

func (r Redis) Del(keys ...string) *Reply {
	args := make([]interface{}, len(keys))
	for i, v := range keys {
		args[i] = v
	}
	return r.Cmd("DEL", args...)
}

func (r Redis) HKeys(key string, args ...interface{}) *Reply {
	args = append([]interface{}{key}, args...)
	return r.Cmd("HKEYS", args...)
}

func (r Redis) Expire(key string, args ...interface{}) *Reply {
	args = append([]interface{}{key}, args...)
	return r.Cmd("EXPIRE", args...)
}

func (r Redis) Publish(channel string, args ...interface{}) *Reply {
	args = append([]interface{}{channel}, args...)
	return r.Cmd("PUBLISH", args...)
}
