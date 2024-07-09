package main

type Card interface {
	CopyData() []int
	StoreData(data []int)
	DeleteData()
	AddData(data []int)
}

type CardReader struct {
	Storage Card
}

func (cr CardReader) CopyDataOfStorage() []int {
	return cr.Storage.CopyData()
}
func (cr CardReader) StoreDataToStorage(data []int) {
	cr.Storage.StoreData(data)
}
func (cr CardReader) DeleteDataFromStorage() {
	cr.Storage.DeleteData()
}
func (cr CardReader) AddDataToStorage(data []int) {
	cr.Storage.AddData(data)
}
