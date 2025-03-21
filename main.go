package main

import (
	"fmt"

	"github.com/YoavIsaacs/go_blog_aggregator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = c.SetUser("YoavIsaacs")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c, err = config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(c.DatabaseURL)
	fmt.Println(c.CurrentUserName)
}
