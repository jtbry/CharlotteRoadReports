# Scrape

This command will poll the CMPD incident API and update the database.

## Configuration

- **ENV**: If you are using `.env` files instead of traditional environment variables you need to declare **ENV** as either `production` or `development` so it knows which `.env` file to load. The default is `development`
- **DATABASE_URL**: The DSN for your MySQL database
- **SCHEDULED_SCRAPING**: Whether or not the process should handle scheduling internally. Defaults to `false`. If you are not using `cron` or a similar scheduler you should probably set this to `true`
