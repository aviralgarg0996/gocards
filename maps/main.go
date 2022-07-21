package main
import "fmt"

func main(){
	// var colors map[string]string
	// colors :=map[string]string{
	// 	"red":"#ff",
	// 	"black":"#000",
	// }

colors:=make(map[string]string)
	fmt.Println(colors)
}

func printMap(c map[string]string){
	for color,hex:=range c{
		fmt.Println(color +" : "+hex)
	}
}