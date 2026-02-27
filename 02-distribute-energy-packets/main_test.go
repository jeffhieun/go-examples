package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name    string
		packets []int64
		agents  int64
		want    int64
	}{
		{
			name:    "basic example",
			packets: []int64{7, 10, 4},
			agents:  4,
			want:    4,
		},
		{
			name:    "single packet",
			packets: []int64{100},
			agents:  5,
			want:    20,
		},
		{
			name:    "agents equals 1",
			packets: []int64{1, 2, 3},
			agents:  10,
			want:    0,
		},
		{
			name:    "empty packets",
			packets: []int64{},
			agents:  3,
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solution(tt.packets, tt.agents)
			if got != tt.want {
				t.Fatalf("Max Equal Energy() = %d, want %d", got, tt.want)
			}
		})
	}
}
