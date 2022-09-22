# CharlotteRoadReports (cltrr)

### View the [demo](https://charlotteroadreport.uc.r.appspot.com/)

<br />

CharlotteRoadReports is an open source web application that tracks and displays active and past traffic incidents and roadway obstructions encountered by the Charlotte-Mecklenburg Police Department. This application polls the CMPD's public data API which is based on their own Computer Aided Dispatch (CAD) system. Incidents are only tracked if they are reported / encountered by CMPD through this system.

<br />

## Deployment

Requirements

- Go 1.15+
- PostgreSQL Server 12.2+
- Node >= 10.16 and npm >= 5.6

### 1. Clone the Repo

`git clone https://github.com/jtbry/CharlotteRoadReports`

### 2. Set up environment

Please see the provided `dev.env` file and change it as needed. Further configuration options can be seen in the `AppConfig` struct located in `./internal/app/config.go`

### 3. Run the project

#### **Docker**

From the root directory, run `docker compose -f ./deployments/docker-compose.yml -p "cltrr" up -d`

#### **Manual**

You can build two binaries from this project

- **server**: Runs the back-end HTTP API by default but can be configured to serve the front-end as well.
- **scrape**: Runs the scraping process, can be executed via cron or it can be configured to schedule itself.

The **web** directory contains the source for the React front-end. Make sure to build this react app in order to serve it from the `server` binary or upload the built files and proxy to the server there is an `nginx.conf` there for example. For development you may se the `proxy` field in the client's `project.json` or use NGINX to proxy between the client dev server and the back-end.

<br />

# Contributing

All feedback and contributions welcome, submit an issue to share feedback or ideas or to discuss major changes before pull requests. Commits messages should generally follow the [conventional commits specs.](https://www.conventionalcommits.org/en/v1.0.0/#summary)

<br />

# License

CharlotteRoadReports is [licensed](/LICENSE) with the [MIT License](https://spdx.org/licenses/MIT.html)
