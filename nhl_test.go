package gonhl

import "testing"

func Test_joinIntIDs(t *testing.T) {
	type args struct {
		ids []int
		sep string
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{ids: []int{}, sep: ","}, want: ""},
		{args: args{ids: []int{1}, sep: ","}, want: "1"},
		{args: args{ids: []int{1, 2}, sep: ","}, want: "1,2"},
		{args: args{ids: []int{22, 44, 66}, sep: ","}, want: "22,44,66"},
		{args: args{ids: []int{-1, -2, -3, 0}, sep: "Q"}, want: "-1Q-2Q-3Q0"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := joinIntIDs(tt.args.ids, tt.args.sep); got != tt.want {
				t.Errorf("joinIntIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
