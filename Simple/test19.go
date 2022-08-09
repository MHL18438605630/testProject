package main

import "fmt"

type Hero struct {
	Name string
	Sex  string
	Age  int
}

// GetName get 函数
func (this *Hero) GetName() string {
	return this.Name
}

// SetName set 函数
func (this *Hero) SetName(name string) {
	this.Name = name
}

// Show 展示
func (this *Hero) Show() {
	fmt.Printf("Hero =%v\n", this)
}
func main() {

	hero := Hero{"MHL", "GIRL", 24}
	hero.Show()

	hero.SetName("MaHuiLi")
	hero.Show()

}
