package disk

import (
	"reflect"
	"syscall"
	"testing"
)

func TestSpace(t *testing.T) {
	tests := []struct {
		name     string
		s        Space
		wantByte uint64
		wantKiB  float64
		wantMiB  float64
		wantGiB  float64
		wantTiB  float64
	}{
		{
			name:     "one byte",
			s:        1,
			wantByte: 1,
			wantKiB:  1.0 / float64(KiB),
			wantMiB:  1.0 / float64(MiB),
			wantGiB:  1.0 / float64(GiB),
			wantTiB:  1.0 / float64(TiB),
		},
		{
			name:     "one kibibyte",
			s:        KiB,
			wantByte: KiB,
			wantKiB:  1.0,
			wantMiB:  1.0 / float64(KiB),
			wantGiB:  1.0 / float64(MiB),
			wantTiB:  1.0 / float64(GiB),
		},
		{
			name:     "one mebibyte",
			s:        MiB,
			wantByte: MiB,
			wantKiB:  KiB,
			wantMiB:  1.0,
			wantGiB:  1.0 / float64(KiB),
			wantTiB:  1.0 / float64(MiB),
		},
		{
			name:     "one Gibibyte",
			s:        GiB,
			wantByte: GiB,
			wantKiB:  MiB,
			wantMiB:  KiB,
			wantGiB:  1.0,
			wantTiB:  1.0 / float64(KiB),
		},
		{
			name:     "one Tebibyte",
			s:        TiB,
			wantByte: TiB,
			wantKiB:  GiB,
			wantMiB:  MiB,
			wantGiB:  KiB,
			wantTiB:  1.0,
		},
		{
			name:     "uint64 min",
			s:        0,
			wantByte: 0,
			wantKiB:  0,
			wantMiB:  0,
			wantGiB:  0,
			wantTiB:  0,
		},
		{
			name:     "uint64 max",
			s:        18446744073709551615,
			wantByte: 18446744073709551615,
			wantKiB:  18446744073709551615 / float64(KiB),
			wantMiB:  18446744073709551615 / float64(MiB),
			wantGiB:  18446744073709551615 / float64(GiB),
			wantTiB:  18446744073709551615 / float64(TiB),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Byte(); got != tt.wantByte {
				t.Errorf("Byte() = %v, want %v", got, tt.wantByte)
			}
			if got := tt.s.Kibibyte(); got != tt.wantKiB {
				t.Errorf("KiB() = %v, want %v", got, tt.wantKiB)
			}
			if got := tt.s.Mebibyte(); got != tt.wantMiB {
				t.Errorf("MiB() = %v, want %v", got, tt.wantMiB)
			}
			if got := tt.s.Gibibyte(); got != tt.wantGiB {
				t.Errorf("GiB() = %v, want %v", got, tt.wantGiB)
			}
			if got := tt.s.Tebibyte(); got != tt.wantTiB {
				t.Errorf("TiB() = %v, want %v", got, tt.wantTiB)
			}
		})
	}
}

func TestNewDisk(t *testing.T) {
	path := "/some/path"

	tests := []struct {
		name     string
		scenario syscall.Statfs_t
		want     Disk
		wantErr  bool
	}{
		{
			name: "mixed",
			scenario: syscall.Statfs_t{
				Bsize:  1024,
				Blocks: 7600,
				Bfree:  1432,
			},
			want: Disk{
				Path:  path,
				Total: 1024 * 7600,
				Free:  1432 * 1024,
				Used:  (1024 * 7600) - (1432 * 1024),
			},
		},
		{
			name: "no space left",
			scenario: syscall.Statfs_t{
				Bsize:  8,
				Blocks: 234324234,
				Bfree:  0,
			},
			want: Disk{
				Path:  path,
				Total: 8 * 234324234,
				Free:  0,
				Used:  8 * 234324234,
			},
		},
		{
			name: "all free",
			scenario: syscall.Statfs_t{
				Bsize:  123,
				Blocks: 234,
				Bfree:  234,
			},
			want: Disk{
				Path:  path,
				Total: 123 * 234,
				Free:  123 * 234,
				Used:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Statfs = func(path string, stat *syscall.Statfs_t) (err error) {
				stat.Bsize = tt.scenario.Bsize
				stat.Blocks = tt.scenario.Blocks
				stat.Bfree = tt.scenario.Bfree
				return nil
			}
			got, err := NewDisk(path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDisk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("NewDisk() got = %v, want %v", got, tt.want)
			}
		})
	}
}
