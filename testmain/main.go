package main

import (
	"context"
	"fmt"
	"log"
)

type UserInfo struct {
	UserId   int
	UserName string
}

func main() {
	var ctx context.Context
	var userId = 10
	ui, ex := GetUserInfoFromCache(ctx, userId)
	if ex {
		return
	}
	log.Println(ui)
	ui = FindUserInfoFromDB(ctx, userId)
	log.Println(ui)
}

func GetUserInfoFromCache(ctx context.Context, userId int) (ui *UserInfo, ex bool) {
	fmt.Println("1111")
	fmt.Println("1111")
	fmt.Println("1111")
	fmt.Println("1111")
	return &UserInfo{
		UserId:   userId,
		UserName: "zhaolei",
	}, true
}

func FindUserInfoFromDB(ctx context.Context, userId int) (dui *UserInfo) {
	fmt.Println("1111")
	fmt.Println("1111")
	fmt.Println("1111")
	fmt.Println("1111")
	return &UserInfo{
		UserId:   userId,
		UserName: "zhaoleidb",
	}
}
