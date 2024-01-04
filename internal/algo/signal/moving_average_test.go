package signal

import (
	"reflect"
	"testing"

	"github.com/elijahelrod/vespene/pkg/model"
)

func TestNewLongMovingAverage(t *testing.T) {

	type args struct{}

	tests := []struct {
		name    string
		args    args
		want    *MovingAverage
		wantErr bool
	}{
		{
			name:    "Base Case",
			args:    args{},
			want:    &MovingAverage{Trend: Flat, WindowSize: Long, Ticks: make([]float64, 0, Long)},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLongMovingAverage()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLongMovingAverage() got = %v, want = %v", got, tt.want)
			}
		})
	}

}

func TestNewMediumMovingAverage(t *testing.T) {
	type args struct {
	}

	tests := []struct {
		name    string
		args    args
		want    *MovingAverage
		wantErr bool
	}{
		{
			name:    "Base Case",
			args:    args{},
			want:    &MovingAverage{Trend: Flat, WindowSize: Medium, Ticks: make([]float64, 0, Medium)},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMediumMovingAverage()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMediumMovingAverage() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestNewShortMovingAverage(t *testing.T) {
	type args struct{}

	tests := []struct {
		name    string
		args    args
		want    *MovingAverage
		wantErr bool
	}{
		{
			name:    "Base Case",
			args:    args{},
			want:    &MovingAverage{Trend: Flat, WindowSize: Short, Ticks: make([]float64, 0, Short)},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewShortMovingAverage()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShortMovingAverage() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestMovingAverage_SignalActive(t *testing.T) {
	type args struct {
		windowSize int
		ticks      []model.Tick
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Full Capacity",
			args: args{
				windowSize: 1,
				ticks: []model.Tick{{
					Timestamp: 0,
					Bid:       10.5,
					Ask:       20,
					Symbol:    "",
					Price:     23,
					DailyLow:  40,
					DailyHigh: 50,
					DailyVol:  100,
				}}},
			want:    true,
			wantErr: false,
		},
		{
			name: "Building Capacity",
			args: args{
				windowSize: 2,
				ticks: []model.Tick{{
					Timestamp: 0,
					Bid:       10.5,
					Ask:       20,
					Symbol:    "",
					Price:     23,
					DailyLow:  40,
					DailyHigh: 50,
					DailyVol:  100,
				}}},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			movingAverage := newMovingAverage(tt.args.windowSize)
			movingAverage.Update(tt.args.ticks[0])
			got := movingAverage.SignalActive()
			if got != tt.want {
				t.Errorf("SignalActive() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
