<h1>Ascii-art-web</h1>

<h2>By Abdulrahman khaled (akhaled), Sameer Ahmed (sahmed) and Mohammed Alsammak (malsamma)</h2>

<h3>Usage</h3>

1. `go run .` in the ascii-art-web/server directory

2. On any web browser, type `localhost:8080/` 

<h3>Implementation details: Algorithms</h3>

The `net/http` standard golang library was used to create the http server in `Go`. navigating to 
`ascii-art-web/server/main.go` and opening it, you will find the main function that handles both 
`/` and `/ascii-art` along with `favicon.ico` (handled by `http.NotFoundHandler()` function). Then, when the server is run, it listens for any requests at port 8080, ready to serve.

The server mainly listens for the GET request to load the main webpage, and a POST request when everything is 
submitted to `/ascii-art`. the server works by parsing HTML templates presented in the templates directory 
(along with styles.css). the ascii art is appended to the webpage inside the result `div`.

<h3> Requests returned </h3>

1. `OK (200)`, if everything went without errors.
2. `404 Not Found`, if nothing is found, for example templates or banners.
3. `400 Bad Request`, for incorrect requests.
4. `500 Internal Server Error`, for any unhandled errors.


<h4> Look at the comments in the code for more info</h4>
<h5> Only standard golang libraries were used </h5>