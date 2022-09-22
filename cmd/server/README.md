# Scrape

This command will run the back-end server.

## Configuration

- **ENV**: If you are using `.env` files instead of traditional environment variables you need to declare **ENV** as either `production` or `development` so it knows which `.env` file to load. The default is `development`
- **DATABASE_URL**: The DSN for your MySQL database
- **PORT**: The port the web server should run on. Defaults to `8080`
- **SERVE_CLIENT**: Whether or not the server should also serve the client from `/web/build`. Defaults to `false`
