package command

import (
	"file-stat/controller"
	"file-stat/handler"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

type ApiCommand struct{}

var (
	port   int
	apiCmd = &cobra.Command{
		Use:   "serve",
		Short: "REST API to list the files of a given directory",
		Run:   startApi,
	}
)

func (cc ApiCommand) GetCommand() *cobra.Command {

	apiCmd.PersistentFlags().IntVar(&port, "port", 5150, "Port on which the API should listen (default is 5150)")

	return apiCmd
}

func initRoutes(router *mux.Router) {
	// Override of the original error, because it's was HTML returned
	router.NotFoundHandler = handler.NotFoundHandler()
	router.MethodNotAllowedHandler = handler.MethodNotAllowedHandler()

	router.HandleFunc("/", (controller.HomeController{}).GetInfo).Methods("GET")
	router.HandleFunc("/files", (controller.FileController{}).GetFileList).Methods("GET")
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("CORS ok")

		next.ServeHTTP(w, r)
		return
	})
}

func startApi(cmd *cobra.Command, args []string) {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(cors)

	initRoutes(router)

	server := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 30,
	}

	log.Fatal(server.ListenAndServe())
}
