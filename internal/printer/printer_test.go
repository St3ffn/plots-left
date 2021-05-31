package printer

import (
	"bytes"
	"github.com/St3ffn/plots-left/pkg/disk"
	"reflect"
	"testing"
)

func TestDefaultPrinter_Print(t *testing.T) {
	tests := []struct {
		name string
		info disk.PlotInfo
		want string
	}{
		{
			name: "none",
			info: disk.PlotInfo{
				Disk: &disk.Disk{
					Path:  "/tmp",
					Total: 0,
					Free:  0,
					Used:  0,
				},
			},
			want: "0\n",
		},
		{
			name: "one",
			info: disk.PlotInfo{
				Reserved: 2,
				Disk: &disk.Disk{
					Path:  "/tmp",
					Total: disk.SizeOfPlot * 5,
					Free:  disk.SizeOfPlot * 3,
					Used:  disk.SizeOfPlot,
				},
			},
			want: "1\n",
		},
		{
			name: "five",
			info: disk.PlotInfo{
				Reserved: 1,
				Disk: &disk.Disk{
					Path:  "/tmp",
					Total: disk.SizeOfPlot * 7,
					Free:  disk.SizeOfPlot * 6,
					Used:  disk.SizeOfPlot,
				},
			},
			want: "5\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			Stdout = b
			DefaultPrinter{}.Print(&tt.info)
			got := b.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Print() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerbosePrinter_Print(t *testing.T) {
	tests := []struct {
		name string
		info disk.PlotInfo
		want string
	}{
		{
			name: "none",
			info: disk.PlotInfo{
				Disk: &disk.Disk{
					Path:  "/tmp",
					Total: 0,
					Free:  0,
					Used:  0,
				},
			},
			want: "Path Total Stored Reserved Left\n/tmp 0     0      0        0\n",
		},
		{
			name: "one",
			info: disk.PlotInfo{
				Reserved: 2,
				Disk: &disk.Disk{
					Path:  "/tmp",
					Total: disk.SizeOfPlot * 5,
					Free:  disk.SizeOfPlot * 3,
					Used:  disk.SizeOfPlot,
				},
			},
			want: "Path Total Stored Reserved Left\n/tmp 5     1      2        1\n",
		},
		{
			name: "five",
			info: disk.PlotInfo{
				Reserved: 1,
				Disk: &disk.Disk{
					Path:  "/tmp",
					Total: disk.SizeOfPlot * 7,
					Free:  disk.SizeOfPlot * 6,
					Used:  disk.SizeOfPlot,
				},
			},
			want: "Path Total Stored Reserved Left\n/tmp 7     1      1        5\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			Stdout = b
			VerbosePrinter{}.Print(&tt.info)
			got := b.String()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Print() got = %v, want %v", got, tt.want)
			}
		})
	}
}
