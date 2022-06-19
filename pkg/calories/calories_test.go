package calories

import (
	"testing"
)

func Test_BurnRate(t *testing.T) {
	type args struct {
		MET float64
		kg  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "mowing the lawn",
			args: args{MET: 4.5, kg: 94.25},
			want: 7.4221875,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BurnRate(tt.args.MET, tt.args.kg); got != tt.want {
				t.Errorf("BurnRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalBurn(t *testing.T) {
	type args struct {
		kg   float64
		MET  float64
		time float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "mowing the lawn",
			args: args{kg: 94.25, MET: 4.5, time: 30.00},
			want: 222.66562499999998,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalBurn(tt.args.kg, tt.args.MET, tt.args.time); got != tt.want {
				t.Errorf("TotalBurn() = %v, want %v", got, tt.want)
			}
		})
	}
}
