package singleton_test

import (
	singleton "exercise/design/01_signleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	if !assert.Equal(t, singleton.GetInstance(), singleton.GetInstance()) {
		t.Errorf("instance not equal")
	}
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Errorf("instance not equal")
			}
		}
	})
}

func TestLazyGetInstance(t *testing.T) {
	if !assert.Equal(t, singleton.GetLazyInstance(), singleton.GetLazyInstance()) {
		t.Errorf("instance not equal")
	}
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetLazyInstance() != singleton.GetLazyInstance() {
				b.Errorf("instance not equal")
			}
		}
	})
}
