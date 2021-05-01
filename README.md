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
You can either set environment variables traditionally through your OS or by creating a `.env` file in the root directory. See the `example.env` file to get an idea of what environment variables are needed.

#### 3. Run the project
Make sure your current working directory is the project's root folder. From there you can use `go run ./cmd/cltrr` or `go build ./cmd/cltrr` and run the resulting binary.

<br />

# Contributing
All feedback and contributions welcome, submit an issue to share feedback or ideas or to discuss major changes before pull requests. Commits messages should generally follow the [conventional commits specs.](https://www.conventionalcommits.org/en/v1.0.0/#summary)

<br />

# License
CharlotteRoadReports is [licensed](/LICENSE) with the [MIT License](https://spdx.org/licenses/MIT.html)