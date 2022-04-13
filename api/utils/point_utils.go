package utils

// PointInt32 int32转指针
func PointInt32(i int32) *int32 {
	return &i
}

// PointInt64 int64转指针
func PointInt64(i int64) *int64 {
	return &i
}

// PointUint32 uint32转指针
func PointUint32(i uint32) *uint32 {
	return &i
}

// PointUint64 uint64转指针
func PointUint64(i uint64) *uint64 {
	return &i
}

// PointString string转指针
func PointString(s string) *string {
	return &s
}

// PointBool bool转指针
func PointBool(b bool) *bool {
	return &b
}
