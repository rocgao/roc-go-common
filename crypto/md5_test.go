package crypto

import "testing"

func TestCalcMD5(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "空数据1", args: args{text: ""}, want: ""},
		{name: "数据1", args: args{text: "12345LxiDHIMGFpXzpLGIehps"}, want: "88931f1c9d79a2de68e8a53496af1b29"},
		{name: "数据2", args: args{text: "1ArejUM9RRblbIO4Lkc2X4SpRuxP2NpbRLOBfyfXz6lgax@nBxUJF9c"}, want: "d6c137c2763ac7647ef756030b78e729"},
		{name: "数据3", args: args{text: "gzQnvGIUPL7lHyB$UU2tNNp0MfSVDGUry3AE7jl8Hpze2ZC2LttNot5DQfphy1yznhoSc7Lay7gp2Jqr$etNNEMEoCex$iyD6gy11Wug08UND1m87yJ8BmwufQTKG"}, want: "c9e09de9dd8f8b4a3f94ae24764faa42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcMD5(tt.args.text); got != tt.want {
				t.Errorf("CalcMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}
