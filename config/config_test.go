package config

import (
	"reflect"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantNil bool
	}{
		{
			name:    "test_01",
			wantNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultConfig(); !reflect.DeepEqual(got == nil, tt.wantNil) {
				t.Errorf("DefaultConfig() = %v, want %v", got, tt.wantNil)
			}
		})
	}
}

func TestDefaultConsumerConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantNil bool
	}{
		{
			name:    "test_01",
			wantNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultConsumerConfig(); !reflect.DeepEqual(got == nil, tt.wantNil) {
				t.Errorf("DefaultConsumerConfig() = %v, want %v", got, tt.wantNil)
			}
		})
	}
}

func TestDefaultProducerConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantNil bool
	}{
		{
			name:    "test_01",
			wantNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultProducerConfig(); !reflect.DeepEqual(got == nil, tt.wantNil) {
				t.Errorf("DefaultProducerConfig() = %v, want %v", got, tt.wantNil)
			}
		})
	}
}
