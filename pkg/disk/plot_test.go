package disk

import (
	"testing"
)

func TestPlotInfo_PlotsTotal(t *testing.T) {
	tests := []struct {
		name  string
		input PlotInfo
		want  uint64
	}{
		{
			name: "1000 plots",
			input: PlotInfo{
				Disk: &Disk{
					Total: SizeOfPlot * 1000,
				},
				Reserved: 0,
			},
			want: 1000,
		},
		{
			name: "2 plots",
			input: PlotInfo{
				Disk: &Disk{
					Total: SizeOfPlot*2 + (SizeOfPlot - 1),
				},
				Reserved: 2234,
			},
			want: 2,
		},
		{
			name: "1 plot",
			input: PlotInfo{
				Disk: &Disk{
					Total: SizeOfPlot + 23,
				},
				Reserved: 2,
			},
			want: 1,
		},
		{
			name: "1 plot another",
			input: PlotInfo{
				Disk: &Disk{
					Total: SizeOfPlot,
				},
				Reserved: 2,
			},
			want: 1,
		},
		{
			name: "0 plots",
			input: PlotInfo{
				Disk: &Disk{
					Total: SizeOfPlot - 1,
				},
				Reserved: 50,
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
				Disk: &Disk{
					Free: SizeOfPlot*1000 + (SizeOfPlot - 1),
				},
				Reserved: 0,
			},
			want: 1000,
		},
		{
			name: "1000 plots none reserved point",
			input: PlotInfo{
				Disk: &Disk{
					Free: SizeOfPlot * 1000,
				},
				Reserved: 0,
			},
			want: 1000,
		},
		{
			name: "1000 plots 200 reserved",
			input: PlotInfo{
				Disk: &Disk{
					Free: SizeOfPlot*1000 + (SizeOfPlot - 1),
				},
				Reserved: 200,
			},
			want: 1000 - 200,
		},
		{
			name: "2 plots 2 reserved",
			input: PlotInfo{
				Disk: &Disk{
					Free: SizeOfPlot*2 + (SizeOfPlot - 1),
				},
				Reserved: 2,
			},
			want: 0,
		},
		{
			name: "1 plot 2 reserved",
			input: PlotInfo{
				Disk: &Disk{
					Free: SizeOfPlot + (SizeOfPlot - 1),
				},
				Reserved: 2,
			},
			want: 0,
		},
		{
			name: "1 plot 50 reserved",
			input: PlotInfo{
				Disk: &Disk{
					Free: SizeOfPlot + (SizeOfPlot - 1),
				},
				Reserved: 50,
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

func TestPlotInfo_PlotsStored(t *testing.T) {
	tests := []struct {
		name  string
		input PlotInfo
		want  uint64
	}{
		{
			name: "1000 plots",
			input: PlotInfo{
				Disk: &Disk{
					Used: SizeOfPlot * 1000,
				},
				Reserved: 0,
			},
			want: 1000,
		},
		{
			name: "2 plots",
			input: PlotInfo{
				Disk: &Disk{
					Used: SizeOfPlot*2 + (SizeOfPlot - 1),
				},
				Reserved: 2234,
			},
			want: 2,
		},
		{
			name: "1 plot",
			input: PlotInfo{
				Disk: &Disk{
					Used: SizeOfPlot + 23,
				},
				Reserved: 2,
			},
			want: 1,
		},
		{
			name: "1 plot another",
			input: PlotInfo{
				Disk: &Disk{
					Used: SizeOfPlot,
				},
				Reserved: 2,
			},
			want: 1,
		},
		{
			name: "0 plots",
			input: PlotInfo{
				Disk: &Disk{
					Used: SizeOfPlot - 1,
				},
				Reserved: 50,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.PlotsStored(); got != tt.want {
				t.Errorf("PlotsStored() = %v, want %v", got, tt.want)
			}
		})
	}
}
