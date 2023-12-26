package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/samber/lo"
	"log"
	"strings"
	"time"
)

func main() {
	var response []SellerInfo
	var err error
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("start-fullscreen", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	ctx, cancel = chromedp.NewExecAllocator(ctx, opts...)
	ctx, cancel = chromedp.NewContext(ctx)

	defer cancel()

	err = loginTask(ctx)

	scrollCount := 5
	industry := All

	request := SearchRequest{
		ScrollCount: &scrollCount,
		Industry:    &industry,
	}

	if request.Industry != nil {
		task := chromedp.Tasks{
			chromedp.WaitVisible(IndustrySelector[*request.Industry]),
			chromedp.Evaluate(fmt.Sprintf(`document.querySelector("%s").click()`, IndustrySelector[*request.Industry]), nil),
			chromedp.Sleep(2 * time.Second),
		}
		err := chromedp.Run(ctx, task)
		if err != nil {
			log.Println(err)
			return
		}
	}

	for *request.ScrollCount > 0 {
		err = extractSellerInfo(ctx, &response, &request)
		*request.ScrollCount--
	}

	if err != nil {
		log.Println(err)
	}

	jsonData, err := json.Marshal(response)

	err = writeFile(response)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Print(string(jsonData))

	return
}

func loginTask(ctx context.Context) error {
	config, err := readConfigFromFile("./cookie.json")
	task := chromedp.Tasks{
		setCookies(config.Cookies),
		chromedp.Navigate("https://affiliate.tiktok.com/connection/creator?enter_from=seller_center_entry&shop_region=VN"),
		chromedp.Sleep(8 * time.Second),

		chromedp.WaitVisible(".arco-modal-wrapper.arco-modal-wrapper-align-center > div > div:nth-child(2) > span"),
		chromedp.Click(".arco-modal-wrapper.arco-modal-wrapper-align-center > div > div:nth-child(2) > span"),
		chromedp.Sleep(2 * time.Second),

		chromedp.WaitVisible("#___reactour > div:nth-child(4) > div > div.sc-kAyceB.KJlBq > div > span.sc-gEvEer.gnMrpw.sc-dcJsrY.esRJvK.reactour__close > button"),
		chromedp.Click("#___reactour > div:nth-child(4) > div > div.sc-kAyceB.KJlBq > div > span.sc-gEvEer.gnMrpw.sc-dcJsrY.esRJvK.reactour__close > button"),
		chromedp.Sleep(2 * time.Second),

		chromedp.WaitVisible("#arco-tabs-0-tab-1 > span > div > span"),
		chromedp.Evaluate(`document.querySelector("#arco-tabs-0-tab-0 > span").click()`, nil),
	}

	err = chromedp.Run(ctx, task)

	return err
}

func extractSellerInfo(ctx context.Context, response *[]SellerInfo, request *SearchRequest) error {
	task := chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			res, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				return err
			}
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
			if err != nil {
				return err
			}
			doc.Find(".arco-table-body > table > tbody > tr[class*='cursor-pointer']").
				Each(func(i int, selection *goquery.Selection) {
					userName := selection.Find("td:nth-child(1) > div > span > div > div.flex.flex-col.ml-16.text-overflow-single > div.flex.flex-col > div > span").Text()
					_, isExist := lo.Find(*response, func(sellerInfo SellerInfo) bool {
						return sellerInfo.UserName == userName
					})
					if isExist {
						return
					}

					follow := selection.Find("td:nth-child(2) > div > span > div").Text()
					viewAvg := selection.Find("td:nth-child(3) > div > span > div").Text()

					image := selection.Find(".m4b-avatar-image img").AttrOr("src", "")
					*response = append(*response, SellerInfo{
						Image:    image,
						UserName: userName,
						Url:      fmt.Sprintf("https://www.tiktok.com/@%s", userName),
						Follower: follow,
						ViewAvg:  viewAvg,
					})
				})

			return nil
		}),
	}

	if err := chromedp.Run(ctx, task); err != nil {
		return err
	}

	chromedp.Sleep(2 * time.Second)

	if *request.ScrollCount > 0 {
		chromedp.Run(ctx,
			chromedp.Evaluate(`document.querySelector(".arco-table-body > table > div:last-child").scrollIntoView()`, nil),
			chromedp.Sleep(5*time.Second))
	}

	return nil
}
