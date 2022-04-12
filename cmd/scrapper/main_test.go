package main

import (
	"reflect"
	"testing"
)

func Test_getPerson(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Person
		wantErr bool
	}{
		{
			name:    "test",
			args:    args{id: "1"},
			want:    Person{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPerson(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPerson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
