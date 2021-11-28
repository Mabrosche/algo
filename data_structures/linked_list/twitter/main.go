package main

import (
	"fmt"
	"strings"
)

func main() {
	f := &Feed{}
	p1 := Post{
		body: "Lorem ipsum",
	}
	f.Append(&p1)

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)

	p2 := Post{
		body: "Dolor sit amet",
	}
	f.Append(&p2)

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)
	fmt.Printf("Second: %v\n", f.start.next)
	fmt.Println()
	fmt.Println(f)
}

type Post struct {
	body        string
	publishDate int64 // Unix timestamp
	next        *Post // link to the next Post
}

func (p Post) String() string {
	return fmt.Sprintf("Body: %s PubDate: %d", p.body, p.publishDate)
}

type Feed struct {
	length int // we'll use it later
	start  *Post
}

func (f Feed) String() string {
	sb := strings.Builder{}
	for i := 0; i < f.length; i++ {
		sb.WriteString(fmt.Sprintf("%s\n", f.))
	}
	return sb.String()
}

func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
	} else {
		currentPost := f.start
		for currentPost.next != nil {
			currentPost = currentPost.next
		}
		currentPost.next = newPost
	}
	f.length++
}
