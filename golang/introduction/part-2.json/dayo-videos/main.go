package main

import "fmt"



type video struct {
	Id string  
	Title string  
	Description string 
	Imageurl string 
	Url string 
}



func main() {
	fmt.Println("Hello, world.")
	videos := getVideos()
	fmt.Println(videos)
}

func getVideos()(videos []video){
	//Get our videos here,

	video1 := video{
		Id : "eyvLwK5C2dw",
		Title : "Kubernetes on Azure",
		Imageurl : "https://i.ytimg.com/vi/eyvLwK5C2dw/mqdefault.jpg?sqp=CISC_PoF&rs=AOn4CLDo7kizrJozB0pxBhxL9JbyiW_EPw",
		Url : "https://youtu.be/eyvLwK5C2dw",
		Description : "",
	  }
	  
	  video2 := video{
		Id : "QThadS3Soig",
		Title : "Kubernetes on Amazon",
		Imageurl : "https://i.ytimg.com/vi/QThadS3Soig/sddefault.jpg",
		Url : "https://youtu.be/QThadS3Soig",
		Description : "",
	  }
	  
	  jkjk
	//and return them
	return []video{ video1, video2}
}
