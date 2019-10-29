package pool

import (
	"fmt"
	"testing"
)

type people struct {
	Name string `json:"name"`
}

type Option func(sets *Settings)

type Settings struct {
	Size int    `json:"size"`
	Name string `json:"name"`
}

func WithSize(size int) Option {
	return func(sets *Settings) {
		sets.Size = size
	}
}

func WithName(name string) Option {
	return func(sets *Settings) {
		sets.Name = name
	}
}

func TestPool(t *testing.T) {
	arr := []*people{{Name: "xch"}}
	c := arr[0]
	arr[0] = nil
	fmt.Println(c)
}

func TestSetting(t *testing.T) {

}
