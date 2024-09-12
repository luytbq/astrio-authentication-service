docker run -d \
  --name container-name \
  -e POSTGRES_USER=pg_user \
  -e POSTGRES_PASSWORD=pg_password \
  -e POSTGRES_DB=mydatabase \
  -p 54320:5432 \
  postgres
