# **ASCII-ART-WEB** 

##     Description

This is a HTTP server written in Go with HTML templates in the root directory. The main objective is to create a web GUI version of Ascii-Art that allows users to generate different banner styles. The server uses appropriate HTTP status codes and follows good coding practices.


## Usage

To run the server, you need to install Go and run the following command in the root directory:
```
go run cmd/web/main.go
```
Then, you can access the server by going to http://localhost:8080.

## Implementation Details

The server has two endpoints:
+ `GET /`: sends the HTML response for the main page using Go templates to receive and display data from the server.
+ `POST /ascii-art`: sends data to the server (text and a banner) using form and other types of tags to make the post request.

The server uses the following HTTP status code
+ `200 OK`: if everything goes without errors.
+ `404 Not Found`: if nothing is found, for example templates or banners.
+ `400 Bad Request`: for incorrect requests.
+ `500 Internal Server Error`: for unhandled errors.

## Instructions

To use the web GUI version of Ascii-Art:

1. Enter text in the text input.
2. Select a banner style using radio buttons, select object or anything else.
3. Click on the "Generate" button to send a POST request to `/ascii-art`.
4. The result of the POST will be displayed in the same page or in a separate page depending on the implementation chosen.

## Algorithm

The server receives a text string and banner style from the client in a POST request to `/ascii-art`. It then uses the Ascii-Art package to generate the corresponding banner and returns it to the client. The server uses Go templates to display the main page and Go's built-in HTTP server to handle the HTTP requests.You can download with the button "Download" result of ascii art, that you will get result in txt.By the way, you can reset page by the button "Reset".