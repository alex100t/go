package main

import (
	"testing"
)

func TestCvtNumber2Time(t *testing.T) {
	type args struct {
		inp string
	}
	tests := []struct {
		name     string
		args     args
		wantResp string
	}{
		// TODO: Add test cases.
		{
			name: "Test for normal time",
			args: args{
				inp: "10800",
			},
			wantResp: "03:00:00",
		},
		{
			name: "Test for negative number",
			args: args{
				inp: "-10800",
			},
			wantResp: "the number of seconds cannot be less than zero",
		},
		{
			name: "Test for too big number",
			args: args{
				inp: "1050010800",
			},
			wantResp: "the number of seconds is greater than the number of seconds in a day (86399)",
		},
		{
			name: "Test for zero",
			args: args{
				inp: "0",
			},
			wantResp: "00:00:00",
		},
		{
			name: "Test for Maxval",
			args: args{
				inp: "86399",
			},
			wantResp: "23:59:59",
		},
		{
			name: "Test for 24H",
			args: args{
				inp: "86400",
			},
			wantResp: "the number of seconds is greater than the number of seconds in a day (86399)",
		},
		{
			name: "Test for trash",
			args: args{
				inp: "zzzzz",
			},
			wantResp: "invalid Number format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResp := CvtNumber2Time(tt.args.inp); gotResp != tt.wantResp {
				t.Errorf("CvtNumber2Time() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestCvtTime2Number(t *testing.T) {
	type args struct {
		inp string
	}
	tests := []struct {
		name     string
		args     args
		wantResp string
	}{
		// TODO: Add test cases.
		{
			name: "Test for 17:32:54",
			args: args{
				inp: "17:32:54",
			},
			wantResp: "63174",
		},
		{
			name: "Test for 24:00:00",
			args: args{
				inp: "24:00:00",
			},
			wantResp: "invalid Time format",
		},
		{
			name: "Test for MaxVal",
			args: args{
				inp: "23:59:59",
			},
			wantResp: "86399",
		},
		{
			name: "Test for MinVal",
			args: args{
				inp: "00:00:00",
			},
			wantResp: "0",
		},
		{
			name: "Test for trash",
			args: args{
				inp: "zzzzz",
			},
			wantResp: "invalid Time format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResp := CvtTime2Number(tt.args.inp); gotResp != tt.wantResp {
				t.Errorf("CvtTime2Number() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_cvtJD(t *testing.T) {
	type args struct {
		jdDate int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test for JD = 1",
			args: args{
				jdDate: 1,
			},
			wantRes: "1 01-01-1841 Friday",
			wantErr: false,
		},
		{
			name: "Test for JD = 65000",
			args: args{
				jdDate: 65000,
			},
			wantRes: "65000 12-18-2018 Tuesday",
			wantErr: false,
		},
		{
			name: "Test for JD = 66000",
			args: args{
				jdDate: 66000,
			},
			wantRes: "66000 09-13-2021 Monday",
			wantErr: false,
		},
		{
			name: "Test for JD = MaxJD",
			args: args{
				jdDate: MaxJD,
			},
			wantRes: "106751 12-14-3209 Monday",
			wantErr: false,
		},
		{
			name: "Test for JD > MaxJD",
			args: args{
				jdDate: MaxJD + 1,
			},
			wantRes: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := cvtJD(tt.args.jdDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("cvtJD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("cvtJD() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_cvtDATE(t *testing.T) {
	type args struct {
		strDate string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test for Date = 12-18-2018",
			args: args{
				strDate: "12-18-2018",
			},
			wantRes: "65000 12-18-2018 Tuesday",
			wantErr: false,
		},
		{
			name: "Test for Date = 09-13-2021",
			args: args{
				strDate: "09-13-2021",
			},
			wantRes: "66000 09-13-2021 Monday",
			wantErr: false,
		},
		{
			name: "Test for invalid Date = -10",
			args: args{
				strDate: "-10",
			},
			wantRes: "",
			wantErr: true,
		},
		{
			name: "Test for inalid month in Date = 10-Bru-2021",
			args: args{
				strDate: "10-Bru-2021",
			},
			wantRes: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := cvtDATE(tt.args.strDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("cvtDATE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("cvtDATE() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_cvtUNIX(t *testing.T) {
	type args struct {
		uDate int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test for begin Unix",
			args: args{
				uDate: 0,
			},
			wantRes: "47117 01-01-1970 Thursday",
			wantErr: false,
		},
		{
			name: "Test for current Unix 26-Sep-2021",
			args: args{
				uDate: 1632678820,
			},
			wantRes: "66013 09-26-2021 Sunday",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := cvtUNIX(tt.args.uDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("cvtUNIX() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("cvtUNIX() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
