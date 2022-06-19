package conversions

import "testing"

func TestPoundsToMetric(t *testing.T) {
	type args struct {
		lb float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "my weight", args: args{lb: 208}, want: 94.35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PoundsToMetric(tt.args.lb); got != tt.want {
				t.Errorf("PoundsToMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roundWeight(t *testing.T) {
	type args struct {
		weight float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "rounding up", args: args{weight: 94.3479887}, want: 94.35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundWeight(tt.args.weight); got != tt.want {
				t.Errorf("roundWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
