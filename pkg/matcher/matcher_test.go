package matcher

import "testing"

func TestMatch(t *testing.T) {
	type args struct {
		comp Comparable
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"it returns true if it matches",
			args{
				Comparable{
					Response{"foo"},
					Solution{Type{EXACT}, "foo"},
				},
			},
			true,
			false,
		},
		{
			"it returns false if doesn't match",
			args{
				Comparable{
					Response{"bar"},
					Solution{Type{EXACT}, "foo"},
				},
			},
			false,
			false,
		},
		{
			"it returns false if doesn't match",
			args{
				Comparable{
					Response{"bar"},
					Solution{Type{"bad type"}, "foo"},
				},
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Match(tt.args.comp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Match() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
