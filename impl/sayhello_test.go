package impl

import "testing"

func TestSayHello_SayHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		s       SayHello
		args    args
		want    string
		wantErr bool
	}{
		{name: "should return Hello Mistadobalina : ", s: SayHello{}, args: args{name: "Mistadobalina"}, want: "Hello Mistadobalina", wantErr: false},
		{name: "should return Error : ", s: SayHello{}, args: args{name: "wrong"}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SayHello{}
			got, err := s.SayHello(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SayHello.SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SayHello.SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
