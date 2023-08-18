package main

import (
	"context"
	_ "embed"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"net/http"
	_ "net/http"

	"github.com/ServiceWeaver/weaver"
)

//go:embed index.html
var indexHtml string // index.html served on "/"

func main() {
	if err := weaver.Run(context.Background(), run); err != nil {
		panic(err)
	}
}

// app is the main component of our application.
type app struct {
	weaver.Implements[weaver.Main]
	// 02
	searcher weaver.Ref[Searcher]
	// 04
	listener weaver.Listener `weaver:"emojis"`
}

// run implements the application main.
func run(ctx context.Context, app *app) error {
	// 01
	fmt.Println("01: Hello, World!")
	// 02
	emojis, err := app.searcher.Get().Search(ctx, "pig")
	if err != nil {
		return err
	}
	fmt.Println("02: ", emojis)
	// 04
	app.Logger(ctx).Info("emojis listener available.", "addr", app.listener)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprint(w, indexHtml)
	})

	http.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {
		// Search for the list of matching emojis.
		query := request.URL.Query().Get("q")
		// 05 add logging
		app.Logger(ctx).Debug("Search", "query", query)
		emojis, err := app.searcher.Get().Search(request.Context(), query)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// JSON serialize the results.
		bytes, err := json.Marshal(emojis)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(writer, string(bytes))
	})
	return http.Serve(app.listener, nil)

	// return nil (for running 01, 02)
}
