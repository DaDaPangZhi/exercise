package builder_test

import (
	"testing"

	builder "design/02_builder"
)

func TestBuilder(t *testing.T) {
	message := builder.Builder().GetMessage()
	message1 := builder.Builder().
		SetSrcAddr("192.168.0.1").
		SetSrcPort(1234).
		SetDestAddr("192.168.0.2").
		SetDestPort(8080).
		SetHeaderItems("contents", "application/json").
		SetBodyItems("record1").
		SetBodyItems("record2").
		GetMessage()
	t.Log(message, message1)
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			builder.Builder().
				SetSrcAddr("192.168.0.1").
				SetSrcPort(1234).
				SetDestAddr("192.168.0.2").
				SetDestPort(8080).
				SetHeaderItems("contents", "application/json").
				SetBodyItems("record1").
				SetBodyItems("record2").
				GetMessage()
		}
	})
}
