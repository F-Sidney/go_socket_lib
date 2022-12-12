package NSockets

import "testing"

func TestConnectUDP(t *testing.T) {
	type args struct {
		ipAndPort string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{"test1", args{"192.168.1.4:44402"}, "connect succeed"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ConnectUDP(tt.args.ipAndPort); gotResult != tt.wantResult {
				t.Errorf("ConnectUDP() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
