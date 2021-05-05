# CharlotteRoadReports (cltrr)
### View the [demo](https://cltrr.herokuapp.com/)
*The demo only tracks the 10,000 most recent incidents.*

<br />

CharlotteRoadReports is an open source web application that tracks and displays active and past traffic incidents and roadway obstructions encountered by the Charlotte-Mecklenburg Police Department. This application polls the CMPD's public data API which is based on their own Computer Aided Dispatch (CAD) system. Incidents are only tracked if they are reported / encountered by CMPD.

<br />

## Deployment
Requirements
* Go 1.15+
* PostgreSQL Server 12.2+
* Node >= 10.16 and npm >= 5.6 

### 1. Clone the Repo
`git clone https://github.com/jtbry/CharlotteRoadReports`

### 2. Set up environment
You can either set environment variables traditionally through your OS or by creating a `.env` file in the root directory. See the `example.env` file to get an idea of what environment variables are needed. If you are planning on developing the front-end, make sure to change the `proxy` field in `./frontend/package.json` to match the port in your environment variable.

### 3. Run the project
If you are running the project in a production environment or you don't plan to change/develop anything on the frontend, you can run `npm run build` to build the React frontend. After the frontend is built you can then run `go build ./cmd/cltrr` and then run the resulting binary which should be named `cltrr`

If you wish to change or develop the front-end, you will serve the frontend with react-scripts (make sure you ran `npm install` if you need to) and you can run the backend using the provided VSCode `launch.json` (debug) or by running `go run ./cmd/cltrr`

<br />

# Contributing
All feedback and contributions welcome, submit an issue to share feedback or ideas or to discuss major changes before pull requests. Commits messages should generally follow the [conventional commits specs.](https://www.conventionalcommits.org/en/v1.0.0/#summary)

<br />

# License
CharlotteRoadReports is [licensed](/LICENSE) with the [MIT License](https://spdx.org/licenses/MIT.html)