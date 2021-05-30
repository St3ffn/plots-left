package printer

import (
	"github.com/St3ffn/plots-left/pkg/disk"
	"io/ioutil"
	"log"
	"os"
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
				Disk:     &disk.Disk{
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
				Disk:     &disk.Disk{
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
				Disk:     &disk.Disk{
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
			tmpFile, err := ioutil.TempFile(os.TempDir(), "test-output")
			if err != nil {
				log.Fatal("Cannot create temporary file", err)
			}
			Output = tmpFile

			DefaultPrinter{}.Print(&tt.info)

			// Close the file
			if err := tmpFile.Close(); err != nil {
				log.Fatal(err)
			}

			content, err := ioutil.ReadFile(tmpFile.Name())
			if err != nil {
				log.Fatal(err)
			}

			// Convert []byte to string and print to screen
			got := string(content)

			// cleanup
			_ = os.Remove(tmpFile.Name())

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
			want: "Path\tTotal\tReserved\tLeft\n/tmp\t0\t0\t\t0\n",
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
			want: "Path\tTotal\tReserved\tLeft\n/tmp\t5\t2\t\t1\n",
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
			want: "Path\tTotal\tReserved\tLeft\n/tmp\t7\t1\t\t5\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := ioutil.TempFile(os.TempDir(), "test-output")
			if err != nil {
				log.Fatal("Cannot create temporary file", err)
			}
			Output = tmpFile

			VerbosePrinter{}.Print(&tt.info)

			// Close the file
			if err := tmpFile.Close(); err != nil {
				log.Fatal(err)
			}

			content, err := ioutil.ReadFile(tmpFile.Name())
			if err != nil {
				log.Fatal(err)
			}

			// Convert []byte to string and print to screen
			got := string(content)

			// cleanup
			_ = os.Remove(tmpFile.Name())

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Print() got = %v, want %v", got, tt.want)
			}
		})
	}
}
