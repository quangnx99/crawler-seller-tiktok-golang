package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"os"
)

func writeFile(sellerInfos []SellerInfo) error {

	dirPath := "."
	fileName := "seller-infos.json"
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			fmt.Println("Unable to create working directory: ", err)
			return err
		}
	}
	filePath := fmt.Sprintf("%s/%s", dirPath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(sellerInfos)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func readConfigFromFile(filename string) (*Config, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func setCookies(cookies []*network.CookieParam) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		for _, cookie := range cookies {
			// SetCookieAction sets a cookie with the provided parameters.
			// Note: chromedp internally uses the `network.setCookie` method.
			if cookie.SameSite == "no_restriction" {
				cookie.SameSite = network.CookieSameSiteNone
			}
			setCookieAction := network.SetCookie(cookie.Name, cookie.Value).
				WithExpires(cookie.Expires).
				WithDomain(cookie.Domain).
				WithPath(cookie.Path).
				WithHTTPOnly(cookie.HTTPOnly).
				WithSecure(cookie.Secure)

			if err := setCookieAction.Do(ctx); err != nil {
				return err
			}
		}
		return nil
	}
}
