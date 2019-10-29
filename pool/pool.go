package pool

type sig struct{}

type f func() error

type Pool struct {
	capacity   int32
	running    int32
	freeSignal chan sig
}
