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
There are three possible binaries from this project.
- **Combined**: run the web app and scraper in the same process
- **Scrape**: run just the incident scraper
- **Web**: run the just the web app

If you build and run an option that runs the web app you need to make sure you've built the ReactJS front-end with `npm run build` first or that you are running the front-end in development mode with react-scripts with a proxy. You also need to make sure you provide a `PORT` environment variable when running the web app.

<br />

# Contributing
All feedback and contributions welcome, submit an issue to share feedback or ideas or to discuss major changes before pull requests. Commits messages should generally follow the [conventional commits specs.](https://www.conventionalcommits.org/en/v1.0.0/#summary)

<br />

# License
CharlotteRoadReports is [licensed](/LICENSE) with the [MIT License](https://spdx.org/licenses/MIT.html)