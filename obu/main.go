package main

import (
	"github.com/gorilla/websocket"
	"log"
	"math"
	"math/rand"
	"time"
	"tolling/types"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		obuIds := generateOBUID(20)
		for i := 0; i < len(obuIds); i++ {
			data := types.OBUData{
				OBUID: obuIds[i],
				Long:  generateCoords(),
				Lat:   generateCoords(),
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(time.Second)
	}
}

func generateCoords() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func generateLatLong() (float64, float64) {
	return generateCoords(), generateCoords()
}

func generateOBUID(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}
