docker build . -t face-api
docker run -i -t -p 8000:8000 \
  --rm --link FaceMysql:db --name FaceApi \
  -d --net mysql_default face-api
