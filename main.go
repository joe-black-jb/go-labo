package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Revision string

type MoveCategory struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Moves []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"moves"`
}

func main() {
	start := time.Now()

	fmt.Println("Revision ⭐️: ", Revision)

	// var a []int
	// fmt.Println("a: ", a)
	// fmt.Println("a[2]: ", a[2])
	// fmt.Println("len(a): ", len(a))
	// fmt.Println("cap(a): ", cap(a))

	// ids := []int {1, 2, 3, 4}

	// ch := make(chan MoveCategory)
	// done := make (chan interface{})

	// for id := range ids {
	// 	go getPoke(done, id, ch)
	// }

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	close(ch)
	// }()

	// for result := range ch {
	// 		// fmt.Println("result: ", result)
	// 		fmt.Println("result.Name: ", result.Name)
	// 		// fmt.Println("result.Country: ", result.Country)
	// }

	//////// 正規表現 /////////
	// s := "10,897,603"
	// s := "※1,※2 10,897,603"
	s := "※１、※２ △10,897,603"
	result, err := FormatValueText(s)
	if err != nil {
		fmt.Println("変換エラー: ", err)
	}
	fmt.Println(result)

	///// 日付のループ /////
	// loc, err := time.LoadLocation("Asia/Tokyo")
	// if err != nil {
	// 	fmt.Println("load location error")
	// 	return
	// }
	// // 集計開始日付
	// date := time.Date(2024, time.January, 1, 1, 0, 0, 0, loc)
	// // 集計終了日付
	// // endDate := time.Date(2024, time.December, 31, 1, 0, 0, 0, loc)
	// endDate := time.Now()

	// for date.Before(endDate) || date.Equal(endDate) {
	// 	fmt.Println(date.Format("2006-01-02"))
	// 	date = date.AddDate(0, 0, 1) // 日付を1日進める
	// }

	fmt.Println("所要時間: ", time.Since(start))
}

// /////// 財務諸表のデータ加工 /////////////
var OnlyNumRe *regexp.Regexp = regexp.MustCompile(`\d+`)

// var OnlyNumRe *regexp.Regexp = regexp.MustCompile(`(?!※)\d+`)

// ※1 などを除外するためのパターン
var AsteriskAndHalfWidthNumRe *regexp.Regexp = regexp.MustCompile(`※\d+`)

func FormatValueText(text string) (int, error) {
	isMinus := false

	fmt.Println("================")
	// , を削除
	text = strings.ReplaceAll(text, ",", "")
	// ※1 などを削除
	asteriskAndHalfWidthNums := AsteriskAndHalfWidthNumRe.FindAllString(text, -1)
	for _, asteriskAndHalfWidthNum := range asteriskAndHalfWidthNums {
		text = strings.ReplaceAll(text, asteriskAndHalfWidthNum, "")
	}
	// マイナスチェック
	if strings.Contains(text, "△") {
		isMinus = true
	}
	// 数字部分のみ抜き出す
	text = OnlyNumRe.FindString(text)
	// スペースを削除
	text = strings.TrimSpace(text)
	// マイナスの場合、 - を先頭に追加する
	if isMinus {
		// previousMatch = "-" + previousMatch
		text = "-" + text
	}
	// previousInt, prevErr := strconv.Atoi(previousMatch)
	intValue, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

//////////////////////////////////////////

func getPoke(done chan interface{}, id int, ch chan MoveCategory) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/move-category/%v", id)
	res, err := http.Get(url)
	if err != nil {
		// return nil
		fmt.Printf("エラー内容 ❗️: ", err)
	}
	// var result interface{}
	var moveCategory MoveCategory
	body, _ := io.ReadAll(res.Body)
	// fmt.Println("body: ", body)
	json.Unmarshal(body, &moveCategory)
	// fmt.Println("result: ", result)
	// fmt.Printf("%+v\n", moveCategory)

	ch <- moveCategory
}

func getDogs(ch chan string, wg *sync.WaitGroup, urls []string) {
	fmt.Println("getDogs 開始")
	defer wg.Done()
	defer close(ch)
	for i := 0; i < len(urls); i++ {
		// url := urls[i]
		time.Sleep(2000 * time.Millisecond)

		dog := fmt.Sprintf("dog No.%v", i+1)
		fmt.Println("dog: ", dog)
		// dogs = append(dogs, dog)
		ch <- dog
	}
}

func getDog(wg *sync.WaitGroup, ch chan string, url string) {
	fmt.Println("getDog 開始")
	defer wg.Done()

	time.Sleep(2000 * time.Millisecond)
	// res, err := http.Get("https://dog.ceo/api/breeds/image/random")
	// fmt.Println("res: ", res)
	// if err != nil {
	// 	ch <- "Error"
	// }
	// defer res.Body.Close()
	// // return res.Body
	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	ch <- "Error"
	// }
	// fmt.Println("body", string(body))
	// ch <- string(body)

	msg := fmt.Sprintf("Done at %v", time.Now())
	fmt.Println("msg: ", msg)
	ch <- msg
	// v, ok := <- ch
	// fmt.Println("v: ", v)
	// fmt.Println("ok: ", ok)
}

func addWg(wg *sync.WaitGroup, num int) {
	time.Sleep(1000 * time.Millisecond)
	wg.Add(num)
}

// 他の goroutine で閉じている ch に書き込もうとして panic
func square(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		fmt.Println("i: ", i)
		go func(i int) {
			defer wg.Done()
			time.Sleep(2000 * time.Millisecond)
			fmt.Println("go i: ", i)
			result := i * i
			ch <- result
		}(i)
	}
	fmt.Println("closing")
}

func server(ch chan string) {
	defer close(ch)
	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func goSayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Hello!")
}
func goSayGoodBye(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Good Bye!")
}
func goSaySorry(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Sorry!")
}
func sayHello() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Hello!")
}
func sayGoodbye() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Goodbye!")
}
func saySorry() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Sorry!")
}

func getDogWithoutWg(ch chan string, url string) {
	fmt.Println("getDogWithoutWg 開始")

	time.Sleep(2000 * time.Millisecond)
	res, err := http.Get("https://dog.ceo/api/breeds/image/random")
	// fmt.Println("res: ", res)
	if err != nil {
		ch <- "Error"
	}
	defer res.Body.Close()
	// return res.Body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		ch <- "Error"
	}
	fmt.Println("body", string(body))
	ch <- string(body)
}
