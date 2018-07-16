package main

import (
	"goLangLearn/retriever/mock"
	"fmt"
	real2 "goLangLearn/retriever/real"
	"time"
)

const url = "https://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})

	return s.Get(url)
}

func download(r Retriever) string {
	return  r.Get("https://www.baidu.com")
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name": "ccmouse",
			"course": "golang",
		})
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type Switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:",  v.Contents)
	case *real2.Retriever2:
		fmt.Println("UserAgent:", v.UserAgent)
	}

	fmt.Println()
}

func main() {

	var r Retriever

	retriever := mock.Retriever{"this is a fake baidu.com"}
	r = &retriever
	inspect(r)

	r = &real2.Retriever2{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// Type assertion
	if realRetriever, ok := r.(*real2.Retriever2); ok {
		fmt.Println(realRetriever.UserAgent)
	} else {
		fmt.Println("Not a real retriever")
	}

	fmt.Println(session(&retriever))
}
