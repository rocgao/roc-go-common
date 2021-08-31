package strext

import "testing"

func TestNewAesKey(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试1"},
		{name: "测试2"},
		{name: "测试3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAesKey(); len(got) != 24 {
				t.Errorf("NewAesKey() = %v, want 24", got)
			}
		})
	}
}

func TestNewSalt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试1"},
		{name: "测试2"},
		{name: "测试3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSalt(); len(got) != 20 {
				t.Errorf("NewSalt() = %v, want 20", got)
			}
		})
	}
}

func TestNewTraceID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试1"},
		{name: "测试2"},
		{name: "测试3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTraceID(); len(got) != 12 {
				t.Errorf("NewTraceID() = %v, want 12", got)
			}
		})
	}
}
