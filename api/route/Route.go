package route

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

func StartRouting(db *gorm.DB) {
	r := chi.NewRouter()

	r.Get("/dev", func(writer http.ResponseWriter, request *http.Request) {
		//writer.Write(byte[])
		html := `
	<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Response Page</title>
				<style>
					body {
						font-family: Arial, sans-serif;
						background-color: #f0f0f0;
						margin: 0;
						padding: 20px;
					}
					h1 {
						color: #333;
					}
					p {
						color: #666;
					}
				</style>
			</head>
			<body>
				<h1>Hello, client!</h1>
				<p>This is a response from your Go server.</p>
			</body>
			</html>
    `
		_, err := writer.Write([]byte(html))
		if err != nil {
			http.Error(writer, "Error writing response", http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":3000", r)
}
