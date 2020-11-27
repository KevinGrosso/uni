package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"net/http"

	"github.com/holizz/terrapin"
)

func landing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
      <!doctype html>
      <title>Home</title>
      <h1>Home</h1>
      <a href="a/">Pagina a</a><br/>
      <a href="b/">Pagina b</a><br/>
      <a href="koch.html">Fiocco di koch</a><br/>
      `))
}

func pageA(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
      <!doctype html>
      <title>Pagina A</title>
      <h1>Pagina A</h1>
      <p>Questa è la pagina A</p>
      `))
}

func pageB(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
      <!doctype html>
      <title>Pagina B</title>
      <h1>Pagina B</h1>
      <p>Questa è la pagina B</p>
      `))
}

func koch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
      <!doctype html>
      <title>Fiocco di koch</title>
      <h1>Fiocco di koch</h1>
      <img src="/koch.png"></img>
      `))
}

func kochCurve(t *terrapin.Terrapin, lung float64, liv int) {
	if liv == 0 {
		t.Forward(lung)
		return
	}

	kochCurve(t, lung, liv-1)
	t.Left(math.Pi / 3.0)
	kochCurve(t, lung, liv-1)
	t.Right(math.Pi - math.Pi/3.0)
	kochCurve(t, lung, liv-1)
	t.Left(math.Pi / 3.0)
	kochCurve(t, lung, liv-1)
}

func kochSnowFlake(t *terrapin.Terrapin, lung float64, liv int) {
	kochCurve(t, lung, liv)
	t.Right(2.0 * math.Pi / 3)
	kochCurve(t, lung, liv)
	t.Right(2.0 * math.Pi / 3)
	kochCurve(t, lung, liv)
	t.Right(2.0 * math.Pi / 3)
}

func pageImg(w http.ResponseWriter, r *http.Request) {
	i := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	t := terrapin.NewTerrapin(i, terrapin.Position{280.0, 850.0})

	/*for i := 0; i < 4; i++ {
		t.Forward(50)
		t.Right(math.Pi / 2.0)
	}*/

	//kochCurve(t, 2.0, 4)

	kochSnowFlake(t, 1, 6)

	png.Encode(w, i)
}

func main() {
	http.HandleFunc("/", landing)
	http.HandleFunc("/a/", pageA)
	http.HandleFunc("/b/", pageB)
	http.HandleFunc("/koch.png", pageImg)
	http.HandleFunc("/koch.html", koch)

	fmt.Println("Running at http://localhost:4444/...")
	http.ListenAndServe("192.168.1.2:4444", nil)
}
