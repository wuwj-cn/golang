package main

import "fmt"

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type ByteSize float64

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", float64(b/YB))
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", float64(b/ZB))
	case b >= EB:
		return fmt.Sprintf("%.2fEB", float64(b/EB))
	case b >= PB:
		return fmt.Sprintf("%.2fPB", float64(b/PB))
	case b >= TB:
		return fmt.Sprintf("%.2fTB", float64(b/TB))
	case b >= GB:
		return fmt.Sprintf("%.2fGB", float64(b/GB))
	case b >= MB:
		return fmt.Sprintf("%.2fMB", float64(b/MB))
	case b >= KB:
		return fmt.Sprintf("%.2fKB", float64(b/KB))
	}
	return fmt.Sprintf("%.2fB", float64(b))
}

func main() {
	fmt.Println(ByteSize(1e13))
}
