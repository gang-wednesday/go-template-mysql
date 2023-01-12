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
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))
		rd3 := io.NopCloser(bytes.NewBuffer(buf))
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

		a := <-imgCh
		b := <-imgCh
		d := <-imgCh
		dst := imaging.New(240, 240, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, a, image.Pt(0, 0))
		dst = imaging.Paste(dst, b, image.Pt(a.Bounds().Size().X, 0))
		dst = imaging.Paste(dst, d, image.Pt(b.Bounds().Size().X+d.Bounds().Size().X, 0))
		err = imaging.Save(dst, "./lol.jpeg")
		if err != nil {
			return err
		}
		c.Response().Writer.WriteHeader(http.StatusOK)
		c.Response().Header().Set("Content-Type", "application/octet-stream")
		w := c.Response().Writer
		_, err = fmt.Fprintf(w, "dst: %v\n", i)
		return err

	})

	e.Logger.Fatal(e.Start(":8888"))
}
