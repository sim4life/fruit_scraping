package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"

	"golang.org/x/net/publicsuffix"
)

func Test_extractFloat321(t *testing.T) {
	emptyStr := ""
	nonFloatStr := "some"
	floatStr := "3.99tt"

	exp := Number(0.0)
	act := extractFloat32(emptyStr)
	if exp != act {
		t.Fatalf("Expected %s but got %s", exp, act)
	}

	exp = Number(0.0)
	act = extractFloat32(nonFloatStr)
	if exp != act {
		t.Fatalf("Expected %s but got %s", exp, act)
	}

	exp = Number(3.99)
	act = extractFloat32(floatStr)
	if exp != act {
		t.Fatalf("Expected %s but got %s", exp, act)
	}

}

// Test case without Mocks
func Test_fruitInitScrape(t *testing.T) {
	// client *http.Client, uri string, fruitInQueue chan *FruitItem
	uri := "http://www.sainsburys.co.uk/webapp/wcs/stores/servlet" +
		"/CategoryDisplay?listView=true&orderBy=FAVOURITES_FIRST&" +
		"parent_category_rn=12518&top_category=12518&langId=44&" +
		"beginIndex=0&pageSize=20&catalogId=10137&searchTerm=&" +
		"categoryId=185749&listId=&storeId=10151&promotionId=#" +
		"langId=44&storeId=10151&catalogId=10137&categoryId=185749&" +
		"parent_category_rn=12518&top_category=12518&pageSize=20&" +
		"orderBy=FAVOURITES_FIRST&searchTerm=&beginIndex=0&" +
		"hideFilters=true"

	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	checkErr(err)
	client := http.Client{Jar: jar}

	fruitInQueue := make(chan *FruitItem, 20)

	go fruitInitScrape(&client, uri, fruitInQueue)

	iter := 0
	for fruitItem := range fruitInQueue {
		fruitItem.Description = ""
		iter++
	}
	exp := 18
	act := iter
	if exp != act {
		t.Fatalf("Expected %d but got %d", exp, act)
	}

}

// Test case without Mocks
func Test_fruitFinishScrape(t *testing.T) {
	// client *http.Client, uri string, fruitInQueue chan *FruitItem
	uri := "http://www.sainsburys.co.uk/webapp/wcs/stores/servlet" +
		"/CategoryDisplay?listView=true&orderBy=FAVOURITES_FIRST&" +
		"parent_category_rn=12518&top_category=12518&langId=44&" +
		"beginIndex=0&pageSize=20&catalogId=10137&searchTerm=&" +
		"categoryId=185749&listId=&storeId=10151&promotionId=#" +
		"langId=44&storeId=10151&catalogId=10137&categoryId=185749&" +
		"parent_category_rn=12518&top_category=12518&pageSize=20&" +
		"orderBy=FAVOURITES_FIRST&searchTerm=&beginIndex=0&" +
		"hideFilters=true"

	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	checkErr(err)
	client := http.Client{Jar: jar}

	fruitInQueue := make(chan *FruitItem, 20)
	fruitOutQueue := make(chan *FruitItem)

	go fruitInitScrape(&client, uri, fruitInQueue)
	go fruitFinishScrape(&client, fruitInQueue, fruitOutQueue)

	iter := 0
	for fruitItem := range fruitOutQueue {
		fruitItem.Description = ""
		iter++
	}
	exp := 18
	act := iter
	if exp != act {
		t.Fatalf("Expected %d but got %d", exp, act)
	}

}

// Test case without Mocks
func Test_getFruitsJSON(t *testing.T) {
	// client *http.Client, uri string, fruitInQueue chan *FruitItem
	uri := "http://www.sainsburys.co.uk/webapp/wcs/stores/servlet" +
		"/CategoryDisplay?listView=true&orderBy=FAVOURITES_FIRST&" +
		"parent_category_rn=12518&top_category=12518&langId=44&" +
		"beginIndex=0&pageSize=20&catalogId=10137&searchTerm=&" +
		"categoryId=185749&listId=&storeId=10151&promotionId=#" +
		"langId=44&storeId=10151&catalogId=10137&categoryId=185749&" +
		"parent_category_rn=12518&top_category=12518&pageSize=20&" +
		"orderBy=FAVOURITES_FIRST&searchTerm=&beginIndex=0&" +
		"hideFilters=true"

	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	checkErr(err)
	client := http.Client{Jar: jar}

	fruitInQueue := make(chan *FruitItem, 20)
	fruitOutQueue := make(chan *FruitItem)
	type Fruits struct {
		Results    []*FruitItem `json:"results"`
		TotalPrice Number       `json:"total"`
	}

	go fruitInitScrape(&client, uri, fruitInQueue)
	go fruitFinishScrape(&client, fruitInQueue, fruitOutQueue)

	fruitsJSON := getFruitsJSON(fruitOutQueue)
	var fruits Fruits

	err = json.Unmarshal(fruitsJSON, &fruits)
	checkErr(err)

	totalPrice := Number(0)
	for _, fruitItem := range fruits.Results {

		fmt.Println("Title: ", fruitItem.Title, " UnitPrice: ", fruitItem.UnitPrice)
		totalPrice += fruitItem.UnitPrice
	}
	fmt.Println("totalPrice is: ", totalPrice)

	exp := 18
	act := len(fruits.Results)
	if exp != act {
		t.Fatalf("Expected %d but got %d", exp, act)
	}

}
