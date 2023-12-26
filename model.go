package main

import "github.com/chromedp/cdproto/network"

type Config struct {
	URL     string                 `json:"url"`
	Cookies []*network.CookieParam `json:"cookies"`
}

type SellerInfo struct {
	UserName string `json:"userName"`
	Image    string `json:"image"`
	Url      string `json:"url"`
	Follower string `json:"follower"`
	ViewAvg  string `json:"viewAvg"`
}

type Industry int

const (
	All                   Industry = 0
	Beauty                Industry = 1
	Electronic            Industry = 2
	Fashion               Industry = 3
	Food                  Industry = 4
	HomeAndLife           Industry = 5
	MotherAndBaby         Industry = 6
	PersonalCareAndHealth Industry = 7
)

type SearchRequest struct {
	ScrollCount *int      `json:"scrollCount"`
	Industry    *Industry `json:"industry"`
}

const (
	AllSelector                   = "#content-container > main > div > div > div > div > div.mb-36.rounded-8 > div > div > div > div.arco-spin.w-full > div > div.flex.items-start.mt-8 > div.flex.items-center.flex-wrap > div.px-16.py-6.rounded-8.text-body-m-medium.cursor-pointer.mr-12.mt-12.hover\\:text-brand-normal.text-brand-normal"
	BeautySelector                = "#content-container > main > div > div > div > div > div.mb-36.rounded-8 > div > div > div > div.arco-spin.w-full > div > div.flex.items-start.mt-8 > div.flex.items-center.flex-wrap > div:nth-child(2) > div > div > div > div"
	ElectronicSelector            = "#content-container > main > div > div > div > div > div.mb-36.rounded-8 > div > div > div > div.arco-spin.w-full > div > div.flex.items-start.mt-8 > div.flex.items-center.flex-wrap > div:nth-child(3) > div > div > div > div"
	FashionSelector               = "#content-container > main > div > div > div > div > div.mb-36.rounded-8 > div > div > div > div.arco-spin.w-full > div > div.flex.items-start.mt-8 > div.flex.items-center.flex-wrap > div:nth-child(4) > div > div > div > div"
	FoodSelector                  = "#content-container > main > div > div > div > div > div.mb-36.rounded-8 > div > div > div > div.arco-spin.w-full > div > div.flex.items-start.mt-8 > div.flex.items-center.flex-wrap > div:nth-child(5) > div > div > div > div"
	HomeAndLifeSelector           = "#content-container > main > div > div > div > div > div.mb-36.rounded-8 > div > div > div > div.arco-spin.w-full > div > div.flex.items-start.mt-8 > div.flex.items-center.flex-wrap > div:nth-child(6) > div > div > div > div"
	MotherAndBabySelector         = "#content-container > main > div > div > div > div > div.mb-36.rounded-8 > div > div > div > div.arco-spin.w-full > div > div.flex.items-start.mt-8 > div.flex.items-center.flex-wrap > div:nth-child(7) > div > div > div > div"
	PersonalCareAndHealthSelector = "#content-container > main > div > div > div > div > div.mb-36.rounded-8 > div > div > div > div.arco-spin.w-full > div > div.flex.items-start.mt-8 > div.flex.items-center.flex-wrap > div:nth-child(8) > div > div > div > div"
)

var IndustrySelector = map[Industry]string{
	All:                   AllSelector,
	Beauty:                BeautySelector,
	Electronic:            ElectronicSelector,
	Fashion:               FashionSelector,
	Food:                  FoodSelector,
	HomeAndLife:           HomeAndLifeSelector,
	MotherAndBaby:         MotherAndBabySelector,
	PersonalCareAndHealth: PersonalCareAndHealthSelector,
}
