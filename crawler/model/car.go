/*
  author='du'
  date='2020/1/23 21:19'
*/
package model

type Car struct {
	Name         string
	Price        float64
	ImageURL     string
	Size         string
	Fuel         float64 //燃料
	Transmission string  //变速器
	Engine       string
	Displacement float64 //排量
	MaxSpeed     float64
	Acceleration float64 //加速度
}
