package main

type MemoryCard struct {
	data []int
}

func (mc *MemoryCard) CopyData() []int {
	return mc.data
}
func (mc *MemoryCard) StoreData(data []int) {
	mc.data = data
}
func (mc *MemoryCard) DeleteData() {
	mc.data = []int{}
}
func (mc *MemoryCard) AddData(data []int) {
	mc.data = append(mc.data, data...)
}
