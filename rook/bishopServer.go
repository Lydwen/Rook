package main

import(
	"fmt"
	"net/url"
	"errors"
)

/**
A Bishop is a server for Eat Hat AR
 */
type Bishop struct{
	Url url.URL
}

type Bishops struct{
	Servers []Bishop
}


/**
Create a Bishops DB
 */
func CreateBishops() *Bishops{
	bishops := Bishops{
		make([]Bishop, 0),
	}

	fmt.Printf("[DEBUG] Create a bishop")

	return &bishops
}

func (b Bishops) String() string{
	toReturn := "{servers : "
	for _,bishop := range b.Servers{
		toReturn+=bishop.Url.String() + ",\n"
	}
	toReturn += "}"
	return toReturn
}

/**
Add a Bishop server to the DB
 */
func (b *Bishops) AddServer(bishop Bishop){
	fmt.Printf("[DEBUG] Add a server : %s", bishop.Url.RequestURI())

	b.Servers = append(b.Servers, bishop)

	fmt.Println(b.String())
}

/**
Add a Bishop server with the corresponding URL to the DB
 */
func (b *Bishops) AddServerFromUrl(url url.URL){
	fmt.Printf("[DEBUG] Add a server from url: %s", url.String())

	b.Servers = append(b.Servers, Bishop{
		url,
	})

	fmt.Println(b.String())
}

/**
Remove a server from the DB
 */
func (b *Bishops) RemoveServer(bishopURL url.URL){
	fmt.Printf("[DEBUG] Remove a server : %s", bishopURL.RequestURI())

	//find the Bishop corresponding server
	serverPlace := -1
	for i,server := range b.Servers{
		if server.Url == bishopURL{
			serverPlace=i
			break
		}
	}

	fmt.Printf("[DEBUG] The place of the server is : %i", serverPlace)


	if serverPlace != -1{
		fmt.Printf("[DEBUG] SUCCESS! Find the server with the url %s", bishopURL.RequestURI())
		b.Servers = append(b.Servers[:serverPlace], b.Servers[serverPlace+1:]...)
	}else{
		fmt.Printf("[DEBUG] FAIL! Haven't find the server with the url %s", bishopURL.RequestURI())
	}
}

/**
Find a server in the DB
 */
func (b Bishops) FindBishop(url url.URL) (Bishop, error) {
	for _,server := range b.Servers{
		if server.Url == url{
			return server, nil
		}
	}

	return Bishop{}, errors.New("not any bishop server")
}