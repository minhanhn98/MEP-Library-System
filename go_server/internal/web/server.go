package web

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var Client *firestore.Client

func StartServer(host string, port int) {

	// Before starting the server establish the connection to firestore
	opt := option.WithCredentialsFile("./.credentials/mep-lib-pk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logrus.Fatalf("error initializing app: %v\n", err)
	}

	Client, err := app.Firestore(context.Background())
	if err != nil {
		logrus.Fatalf("failed to create client: %v\n", err)
	}

	defer Client.Close()

	// * Testing
	items := Client.Collection("items")
	books := items.Doc("books")

	docsnap, err := books.Get(context.Background())
	if err != nil {
		// TODO: Handle error.
		logrus.Fatal("some error ? %v\n", err)
	}
	dataMap := docsnap.Data()
	fmt.Println(dataMap)

	// * We might the below later

	// serverAdress := fmt.Sprintf("%s:%d", host, port)
	// r := mux.NewRouter()

	// srv := &http.Server{
	// 	Handler:      r,
	// 	Addr:         serverAdress,
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// // Run the server
	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil {
	// 		log.Fatalf("could not start server: %v", err)
	// 	}
	// }()

	// logrus.Println("started server on:", srv.Addr)

	// // Make a channel, and send a value on that channel
	// // whenever we get an os.Interrupt signal
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)

	// // Block until we receive our signal.
	// <-c

	// // Create a deadline to wait for, and shutdown
	// ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*15)
	// defer cancel()
	// srv.Shutdown(ctx)
	// logrus.Println("server shutting down -- goodbye!")
}
