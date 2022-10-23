package main

import (
	"fmt"

	"slice-search-create/pkg"
)

func main() {
	// 構造体をもとにスライスの作成
	companies := []pkg.Company{
		{
			ID:   1,
			Name: "A",
		},
		{
			ID:   2,
			Name: "b",
		},
	}

	users := []pkg.User{
		{
			ID:        1,
			CompanyID: 1,
			Name:      "taro",
			Age:       20,
		},
		{
			ID:        2,
			CompanyID: 1,
			Name:      "jiro",
			Age:       19,
		},
	}

	// 会社に所属するユーザーを入れる構造体
	var companyUsers []pkg.CompanyUsers

	for _, c := range companies {
		// 一時的なユーザーIDのスライス
		var userId []uint

		for _, u := range users {
			// 会社IDとユーザー.会社IDが一致した場合、ユーザーIDスライスに入れる
			if c.ID == u.CompanyID {
				userId = append(userId, u.ID)
			}
		}

		// 会社に所属するユーザーを入れたもの
		result := pkg.CompanyUsers{
			ID:     c.ID,
			Name:   c.Name,
			UserID: userId,
		}

		companyUsers = append(companyUsers, result)
	}

	fmt.Println(companyUsers)
}
