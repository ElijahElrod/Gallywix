package config

import (
	"context"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
