package main
import ("net/http"
"fmt"
"time"
)

func main(){
	links:=[]string{
		"https://google.com",
		"https://amazon.com",
		"https://golang.org",
		"https://swiggy.in",
		"https://stackoverflow.com",
	}

	c := make(chan string)
	for _,url :=range links{
		go checkLink(url,c)
	}

	for l :=range c{
		go func(link string){
			time.Sleep(5 * time.Second)
			checkLink(link,c)
		}(l)
		
	}

	// for i :=0;i<len(links);i++{
	// 	fmt.Println(<-c)
	// }

}

func checkLink(link string,c chan string){
	_,err :=http.Get(link)

	if(err!=nil){
		fmt.Println(link, "might be down")
		c<-link
		return 
	}
	fmt.Println(link, "is up")
	// c<-"its up"
		c<-link
}