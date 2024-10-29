package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type Post struct {
	URL string
	res bool
}

func main() {

	var posts []Post
	var postNodes []*cdp.Node

	// init a controllable Chrome instance
	ctx, cancel := chromedp.NewContext(context.Background())
	// to release the browser resources when
	// it is no longer needed
	defer cancel()

	// SCROLLING LOGIC
	/*scrollingScript := `
		const scrolls = 8
		var scrollCount = 0

		const scrollInterval = setInterval(() => {
			window.scrollTo(0, document.body.scrollHeight)
			scrollCount++

			if (scrollCount == scrolls) {
				clearInterval(scrollInterval)
			}
		},500)
	`
	*/
	err := chromedp.Run(
		ctx,
		// visit target page

		chromedp.Navigate("https://www.instagram.com/webaitc/"),
		// wait 2 seconds

		chromedp.Sleep(2000*time.Millisecond),
		//chromedp.Evaluate(scrollingScript, nil),

		/// TODO ПОМЕНЯТЬ НОДУ КОГДА ПЕРЕХОДИШЬ НА ДРУГОЙ САЙТ
		chromedp.Nodes("div._aagv", &postNodes, chromedp.ByQueryAll),
	)
	if err != nil {
		log.Fatal("Error while performing the automation logic:", err)
	}
	fmt.Println("START")

	// Scraping LOGIC
	var url string
	var res bool

	var i = 0
	for _, node := range postNodes {
		if i >= 5 {
			break
		}
		err = chromedp.Run(ctx,
			// тут надо чтобы в начале были 2 слэша
			// чтобы получить такую строчку заходишь в консоль разраба и там выбираешь нужный элемент правая кнопка
			// copy -> XPATH - но потом лучше почитать про Xpath - то скопируешь можно будет сильно сократить
			// //img[@class='x5yr21d xu96u03 x10l6tqk x13vifvy x87ps6o xh8yej3']
			chromedp.AttributeValue(
				"//img[@class='x5yr21d xu96u03 x10l6tqk x13vifvy x87ps6o xh8yej3']",
				"src", &url, &res, chromedp.BySearch, chromedp.FromNode(node),
			),
			//chromedp.AttributeValue("img", "src", &name, &res, chromedp.BySearch, chromedp.FromNode(node)),

			//chromedp.Text("h5", &price, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		if err != nil {
			log.Fatal("Error:", err)
		}

		post := Post{}
		post.URL = url
		post.res = res
		posts = append(posts, post)
		i++
		fmt.Println("FIND POST:", post.URL)
	}
}
