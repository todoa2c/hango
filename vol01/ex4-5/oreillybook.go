package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Oreillyの書籍情報
type Book struct {
	Title        string   `json:"title"`
	PictureLarge string   `json:"picture_large"`
	Picture      string   `json:"picture"`
	PictureSmall string   `json:"picture_small"`
	Authors      []string `json:"authors"`
	Released     string   `json:"released"`
	Pages        int      `json:"pages"`
	Price        int      `json:"price"`
	EbookPrice   int      `json:"ebook_price"`
	Original     string   `json:"original"`
	OriginalUrl  string   `json:"original_url"`
	ISBN         string   `json:"isbn"`
}

// OreillyBookはisbnに該当するOreillyの書籍情報を返す
func OreillyBook(isbn string) (*Book, error) {
	url := fmt.Sprintf("http://www.oreilly.co.jp/books/%s/biblio.json", isbn)
	log.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	book := &Book{}
	if err := json.Unmarshal(body, book); err != nil {
		return nil, err
	}

	return book, nil
}

// 課題4バージョン
// func main() {
// 	isbn := os.Args[1]
// 	book := OreillyBook(isbn)
// 	log.Printf("%s : %d pages, \\%d\n", book.Title, book.Pages, book.Price)
// 	for _, author := range book.Authors {
// 		log.Println("  ", author)
// 	}
// }

// 課題5バージョン
func main() {
	isbnCodes := os.Args[1:]
	bookChan := make(chan *Book, len(isbnCodes))

	for _, isbn := range isbnCodes {
		// go を外すと直列実行
		go func(isbn string) {
			book, err := OreillyBook(isbn)
			if err != nil {
				log.Println("ISBN: ", isbn, " : ", err)
				book = &Book{} //0ページ0円のダミーを返す
			}
			bookChan <- book
		}(isbn)
	}

	var totalPages, totalPrice int
	// 指定した件数分の結果の受信待ち
	for i := 0; i < len(isbnCodes); i++ {
		book := <-bookChan
		totalPages += book.Pages
		totalPrice += book.Price
	}

	log.Printf("Total Pages: %d, Total Price: \\%d\n", totalPages, totalPrice)
}
