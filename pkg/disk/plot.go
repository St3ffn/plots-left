package disk

// SizeOfPlot in Bytes
// A k32 will take up 101.3 GiB of space once completed
// 101.3 * 1024 * 1024 * 1024 = 108770046771,2
const SizeOfPlot Space = 108770046772

// PlotInfo gives further plot related information of a Disk
type PlotInfo struct {
	*Disk
	// reserved represents the amount of plots to be reserved when evaluating how many plots are left
	Reserved uint64
}

// PlotsTotal calculates the total amount of plots which can be stored on the Disk.
// The amount of reserved plots is not included in the assessment
func (p PlotInfo) PlotsTotal() uint64 {
	total := uint64(float64(p.Total) / float64(SizeOfPlot))
	return total
}

// PlotsLeft calculates the amount of plots which can still be stored on the Disk.
// The amount of reserved plots is included in the assessment
func (p PlotInfo) PlotsLeft() uint64 {
	// don't round, simply cut
	left := uint64(float64(p.Free) / float64(SizeOfPlot))
	if left > p.Reserved {
		return left - p.Reserved
	}
	return 0
}
