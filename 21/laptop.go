package main

import "fmt"

type Reader interface {
	CopyDataOfStorage() []int
	StoreDataToStorage(data []int)
	DeleteDataFromStorage()
	AddDataToStorage(data []int)
}

type Mouse struct {
}
type Keyboard struct {
}
type HeadPhones struct {
}
type UsbConnector struct {
	Reader Reader
}

type Laptop struct {
	Keyboard     Keyboard
	Mouse        Mouse
	HeadPhones   HeadPhones
	UsbConnector UsbConnector
}

// CardReader является адаптером для Laptop и MemoryCard
func main() {
	data := []int{1, 2, 3}
	mc := MemoryCard{data: data}

	cr := CardReader{Storage: &mc}

	laptop := Laptop{
		Keyboard:     Keyboard{},
		Mouse:        Mouse{},
		HeadPhones:   HeadPhones{},
		UsbConnector: UsbConnector{Reader: cr},
	}

	fmt.Println("Data of storage:", laptop.UsbConnector.Reader.CopyDataOfStorage())

	newData := []int{4, 5, 6}
	laptop.UsbConnector.Reader.AddDataToStorage(newData)
	fmt.Println("New data of storage:", laptop.UsbConnector.Reader.CopyDataOfStorage())

}
