docker build . -t face_api_image
docker run -i -t -p 8000:8000 \
  --rm --link face_mysql:db --name face_api \
  -d --net mysql_default face_api_image
