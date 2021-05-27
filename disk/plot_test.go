package disk

import "testing"

func TestPlotInfo_PlotsTotal(t *testing.T) {
	tests := []struct {
		name  string
		input PlotInfo
		want  uint64
	}{
		{
			name: "1000 plots none reserved",
			input: PlotInfo{
				Disk: Disk{
					Total: SizeOfPlot*1000 + (SizeOfPlot - 1),
				},
				reserved: 0,
			},
			want: 1000,
		},
		{
			name: "1000 plots none reserved point",
			input: PlotInfo{
				Disk: Disk{
					Total: SizeOfPlot * 1000,
				},
				reserved: 0,
			},
			want: 1000,
		},
		{
			name: "1000 plots 200 reserved",
			input: PlotInfo{
				Disk: Disk{
					Total: SizeOfPlot*1000 + (SizeOfPlot - 1),
				},
				reserved: 200,
			},
			want: 1000 - 200,
		},
		{
			name: "2 plots 2 reserved",
			input: PlotInfo{
				Disk: Disk{
					Total: SizeOfPlot*2 + (SizeOfPlot - 1),
				},
				reserved: 2,
			},
			want: 0,
		},
		{
			name: "1 plot 2 reserved",
			input: PlotInfo{
				Disk: Disk{
					Total: SizeOfPlot + (SizeOfPlot - 1),
				},
				reserved: 2,
			},
			want: 0,
		},
		{
			name: "1 plot 50 reserved",
			input: PlotInfo{
				Disk: Disk{
					Total: SizeOfPlot + (SizeOfPlot - 1),
				},
				reserved: 50,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.PlotsTotal(); got != tt.want {
				t.Errorf("PlotsTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlotInfo_PlotsLeft(t *testing.T) {
	tests := []struct {
		name  string
		input PlotInfo
		want  uint64
	}{
		{
			name: "1000 plots none reserved",
			input: PlotInfo{
				Disk: Disk{
					Free: SizeOfPlot*1000 + (SizeOfPlot - 1),
				},
				reserved: 0,
			},
			want: 1000,
		},
		{
			name: "1000 plots none reserved point",
			input: PlotInfo{
				Disk: Disk{
					Free: SizeOfPlot * 1000,
				},
				reserved: 0,
			},
			want: 1000,
		},
		{
			name: "1000 plots 200 reserved",
			input: PlotInfo{
				Disk: Disk{
					Free: SizeOfPlot*1000 + (SizeOfPlot - 1),
				},
				reserved: 200,
			},
			want: 1000 - 200,
		},
		{
			name: "2 plots 2 reserved",
			input: PlotInfo{
				Disk: Disk{
					Free: SizeOfPlot*2 + (SizeOfPlot - 1),
				},
				reserved: 2,
			},
			want: 0,
		},
		{
			name: "1 plot 2 reserved",
			input: PlotInfo{
				Disk: Disk{
					Free: SizeOfPlot + (SizeOfPlot - 1),
				},
				reserved: 2,
			},
			want: 0,
		},
		{
			name: "1 plot 50 reserved",
			input: PlotInfo{
				Disk: Disk{
					Free: SizeOfPlot + (SizeOfPlot - 1),
				},
				reserved: 50,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.PlotsLeft(); got != tt.want {
				t.Errorf("PlotsLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
