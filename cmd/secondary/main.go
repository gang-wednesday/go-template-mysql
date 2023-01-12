package main

import (
	"bytes"
	"fmt"
	"go-template/cmd/secondary/resize"
	"image"
	"image/color"
	"io"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "called the second api")
	})
	e.POST("/img", func(c echo.Context) error {
		req := c.Request()
		body := req.Body
		defer body.Close()
		buf, _ := io.ReadAll(body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		defer rdr1.Close()
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))
		defer rdr2.Close()
		rd3 := io.NopCloser(bytes.NewBuffer(buf))
		defer rd3.Close()
		i, _, err := image.Decode(rdr1)
		if err != nil {
			return err
		}

		j, _, err := image.Decode(rdr2)
		if err != nil {
			return err
		}
		k, _, err := image.Decode(rd3)
		if err != nil {
			return err
		}
		imgCh := make(chan image.Image)
		sizes := []int{40, 80, 120}
		images := []image.Image{i, j, k}
		for x, val := range sizes {
			go resize.Resize(images[x], val, imgCh)

		}
		dst := imaging.New(240, 240, color.NRGBA{0, 0, 0, 0})
		widthStart := 0
		for x := 0; x < 3; x++ {
			a := <-imgCh
			dst = imaging.Paste(dst, a, image.Pt(widthStart, 0))
			widthStart = widthStart + a.Bounds().Size().X
		}

		err = imaging.Save(dst, "./lol.jpeg")
		if err != nil {
			return err
		}
		c.Response().Writer.WriteHeader(http.StatusOK)
		c.Response().Header().Set("Content-Type", "application/octet-stream")
		w := c.Response().Writer
		_, err = fmt.Fprintf(w, "dst: %v\n", dst)
		return err

	})

	e.Logger.Fatal(e.Start(":8888"))
}
