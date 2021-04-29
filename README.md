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

#### 1. Clone the Repo
`git clone https://github.com/jtbry/CharlotteRoadReports`

#### 2. Set up environment
See the example.env file for an example of the required environment variables. You can set these in your own .env file or by setting your environment variables appropriately.

#### 3. Run the project
You can use `go run .` or `go build` and run the resulting binary

<br />

# Contributing
All feedback and contributions welcome, submit an issue to share feedback or ideas or to discuss major changes before pull requests. Commits messages should generally follow the [conventional commits specs.](https://www.conventionalcommits.org/en/v1.0.0/#summary)

<br />

# License
CharlotteRoadReports is [licensed](/LICENSE) with the [MIT License](https://spdx.org/licenses/MIT.html)