docker build -t ascii-art-web-docker .
docker run -p 8080:8080 ascii-art-web-docker
xdg-open http://localhost:8080
