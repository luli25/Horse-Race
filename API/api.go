package api

import (
	"log"
	"net/http"

	"horse-race/internal/race"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

func RaceWebsocketHandler(responseWrite http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(responseWrite, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	raceInstance := race.NewRace([]string{"A", "B", "C", "D"}, 50)

	updates := make(chan race.Update)

	go raceInstance.RunWithUpdates(updates)

	for update := range updates {
		err := conn.WriteJSON(update)
		if err != nil {
			log.Println("client disconnected")
			break
		}
	}
}
