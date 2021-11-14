package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/eavesmy/gear"
)

type Upload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	port     = ":8080"
	filePath = "./"
)

func init() {
	flag.StringVar(&port, "p", ":8080", "service port.")
	flag.StringVar(&filePath, "path", "./", "saved log file path.")
	flag.Parse()
}

// 保存用户提交信息
func main() {
	app := gear.New()

	router := gear.NewRouter()
	router.Post("/upload", handler)

	app.Use(func(ctx *gear.Context) error { // 需要鉴权写到这里
		fmt.Println(ctx.Req.Header)
		return nil
	})
	app.UseHandler(router)

	fmt.Println("Server start successful.")

	if err := app.Listen(port); err != nil {
		fmt.Println("Server start error:", err)
	}
}

func handler(ctx *gear.Context) error {

	req := &Upload{}
	if err := ctx.ParseBody(req); err != nil {
		return err
	}

	filename := filePath + req.Title
	if err := writeFile(filename, req.Content); err != nil {
		return err
	}

	return ctx.HTML(200, "ok")
}

func writeFile(filename, content string) (err error) {
	return ioutil.WriteFile(filename, []byte(content), 0777)
}
