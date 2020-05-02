package translate

func intTo32(i int) *int32 {
	i32 := int32(i)
	return &i32
}
