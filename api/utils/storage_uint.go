package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// StorageUint TODO
type StorageUint uint64

const (
	// B TODO
	B StorageUint = 1 << (10 * iota)
	// KB TODO
	KB
	// MB TODO
	MB
	// GB TODO
	GB
	// TB TODO
	TB
	// PB TODO
	PB
	// EB TODO
	EB
)

var (
	// StorageUint_name TODO
	StorageUint_name = map[StorageUint]string{
		B:  "B",
		KB: "KB",
		MB: "MB",
		GB: "GB",
		TB: "TB",
		PB: "PB",
		EB: "EB",
	}
	// StorageUint_value TODO
	StorageUint_value = map[string]StorageUint{
		"B":  B,
		"KB": KB,
		"MB": MB,
		"GB": GB,
		"TB": TB,
		"PB": PB,
		"EB": EB,
	}
)

// String TODO
func (su StorageUint) String() string {
	return StorageUint_name[su]
}

// Value TODO
func (su StorageUint) Value() uint64 {
	return uint64(su)
}

func (StorageUint) fromString(uintString string) StorageUint {
	return StorageUint_value[uintString]
}

// BytesToHuman TODO
func BytesToHuman(length uint64) string {
	if length == 0 {
		return "0 B"
	}
	var u StorageUint
	if length >= EB.Value() {
		u = EB
	} else if length >= PB.Value() {
		u = PB
	} else if length >= TB.Value() {
		u = TB
	} else if length >= GB.Value() {
		u = GB
	} else if length >= MB.Value() {
		u = MB
	} else if length >= KB.Value() {
		u = KB
	} else if length >= B.Value() {
		u = B
	}
	return fmt.Sprintf("%.2f "+u.String(), float64(length)/float64(u.Value())) // 保留2位小数
}

// StorageSize TODO
type StorageSize struct {
	Uint  StorageUint
	Value uint64
}

func (s *StorageSize) conv(storage string, uint StorageUint) error {
	var err error
	s.Uint = uint
	s.Value, err = strconv.ParseUint(storage[0:strings.LastIndex(storage, uint.String())], 10, 64)
	return err
}

// Parse TODO
func (s *StorageSize) Parse(storage string) error {
	var err error
	if strings.HasSuffix(storage, EB.String()) {
		err = s.conv(storage, EB)
	} else if strings.HasSuffix(storage, TB.String()) {
		err = s.conv(storage, TB)
	} else if strings.HasSuffix(storage, GB.String()) {
		err = s.conv(storage, GB)
	} else if strings.HasSuffix(storage, MB.String()) {
		err = s.conv(storage, MB)
	} else if strings.HasSuffix(storage, KB.String()) {
		err = s.conv(storage, KB)
	} else if strings.HasSuffix(storage, B.String()) {
		err = s.conv(storage, B)
	} else {
		err = errors.New("Unknown storage uint " + storage)
	}
	return err
}

// String TODO
func (s *StorageSize) String() string {
	return strconv.FormatUint(s.Value, 10) + s.Uint.String()
}

// Size TODO
func (s *StorageSize) Size() uint64 {
	return uint64(s.Uint) * s.Value
}
