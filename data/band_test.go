package data

import (
	"reflect"
	"testing"
)

func Test_band_characterizeSignal(t *testing.T) {
	type fields struct {
		start      int64
		end        int64
		resolution int64
	}
	type args struct {
		spike []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   signal
	}{
		{
			name:   "characterizes in band signal",
			fields: fields{start: int64(1e9), end: int64(2e9), resolution: 1000},
			args:   args{spike: []int{490, 510}},
			want:   signal{centerFrequency: int64(1.5e9), bandwidth: 20e6, power: -30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &band{
				start:      tt.fields.start,
				end:        tt.fields.end,
				resolution: tt.fields.resolution,
			}
			if got := b.characterizeSignal(tt.args.spike); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("band.characterizeSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}
