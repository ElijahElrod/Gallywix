package trend

import (
	"reflect"
	"testing"

	"github.com/elijahelrod/vespene/internal/algo/signal"
	"github.com/elijahelrod/vespene/pkg/model"
)

func TestNewDonchian(t *testing.T) {

	type args struct {
		tick model.Tick
	}

	tests := []struct {
		name    string
		args    args
		want    *Donchian
		wantErr bool
	}{
		{
			name: "Test New Donchian",
			args: args{},
			want: &Donchian{upper: 0,
				lower:      0,
				mid:        0,
				lowPeriod:  40,
				highPeriod: 50,
				channel:    signal.DonchianChannel,
				lows:       make([]float64, 0, 40),
				highs:      make([]float64, 0, 50)},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDonchian(50, 40)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDonchian() got = %v, want %v", got, tt.want)
			}
		})
	}
}
