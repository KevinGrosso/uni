package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"net/http"
	"strconv"

	"github.com/holizz/terrapin"
)

// Canvas condiviso tra tutti i client
var canvas *image.RGBA

// Tartaruga condivisa tra tutti i client
var tarta *terrapin.Terrapin

func pageCanvas(w http.ResponseWriter, r *http.Request) {
	png.Encode(w, canvas)
}

func ruotaTarta(rotazione string) {
	switch rotazione {
	case "45":
		tarta.Right(math.Pi / 4.0)
	case "90":
		tarta.Right(math.Pi / 2.0)
	case "180":
		tarta.Right(math.Pi)
	}
}

func muoviTarta(distanza string) {
	d, err := strconv.Atoi(distanza)
	if err == nil {
		tarta.Forward(float64(d))
	}
}

func pageMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r) // DEBUG

	// Se richiesto, aggiorno il Canvas
	query := r.URL.Query()

	rot := query.Get("rot")
	if rot != "" {
		ruotaTarta(rot)
	}

	d := query.Get("for")
	if d != "" {
		muoviTarta(d)
	}

	fmt.Fprintf(w, `
		<!doctype html>
		<title>TurtleDrawer</title>
		<script>
		function ruota(angolo) {
  			window.location.href = '?rot=' + angolo;
		}
		function avanti() {
			var distanza = document.getElementById("dist").value;
  			window.location.href = '?for=' + distanza;
		}
		</script>
		<img src="canvas.png"></img><br />
		<button onclick="ruota(45);">Ruota di 45°</button>
		<button onclick="ruota(90);">Ruota di 90°</button>
		<button onclick="ruota(180);">Ruota di 180°</button>
		<br /><input type="text" id="dist">
		<button onclick="avanti();">
		Avanti</button>
		`)
}

func main() {
	// Inizializzo il Canvas
	canvas = image.NewRGBA(image.Rect(0, 0, 900, 600))
	// Inizializzo la tartaruga
	tarta = terrapin.NewTerrapin(canvas, terrapin.Position{450, 500})

	http.HandleFunc("/canvas.png", pageCanvas)
	http.HandleFunc("/", pageMain)

	fmt.Println("Server online @ 127.0.0.1:4444")
	http.ListenAndServe("127.0.0.1:4444", nil)
}
