package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
)

const baseSubsidy = uint64(1_776_465_535)

func main() {
	path := filepath.FromSlash("domain/consensus/processes/coinbasemanager/coinbasemanager.go")
	source, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	startMarker := []byte("var subsidyByDeflationaryMonthTable = []uint64{")
	start := bytes.Index(source, startMarker)
	if start < 0 {
		panic("subsidy table start was not found")
	}
	bodyStart := start + len(startMarker)
	endRelative := bytes.Index(source[bodyStart:], []byte("\n}"))
	if endRelative < 0 {
		panic("subsidy table end was not found")
	}
	end := bodyStart + endRelative + len("\n}")

	var table bytes.Buffer
	table.Write(startMarker)
	table.WriteByte('\n')
	for month := uint64(0); ; month++ {
		subsidy := uint64(float64(baseSubsidy) / math.Pow(1.4, float64(month)/12))
		if month%25 == 0 {
			table.WriteByte('\t')
		}
		fmt.Fprintf(&table, "%d,", subsidy)
		if month%25 == 24 || subsidy == 0 {
			table.WriteByte('\n')
		} else {
			table.WriteByte(' ')
		}
		if subsidy == 0 {
			break
		}
	}
	table.WriteByte('}')

	updated := make([]byte, 0, len(source)-end+start+table.Len())
	updated = append(updated, source[:start]...)
	updated = append(updated, table.Bytes()...)
	updated = append(updated, source[end:]...)
	if err := os.WriteFile(path, updated, 0o644); err != nil {
		panic(err)
	}
}
