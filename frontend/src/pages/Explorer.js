import React, { useEffect } from 'react';
import { AppBar, Container, Grid, Paper, Typography, Switch, TableContainer, Table, TableCell, withStyles, TableHead, TableRow, TableBody } from '@material-ui/core';
import { Link } from 'react-router-dom';
import moment from 'moment';
import 'react-dates/initialize';
import { SingleDatePicker } from 'react-dates';
import 'react-dates/lib/css/_datepicker.css';
import '../assets/dates.css'
import lodash from 'lodash';
import axios from 'axios';
import Loading from '../components/Loading';
import { ReactComponent as ErrorSvg } from '../assets/undraw_error.svg';

function IncidentDataTable(props) {
    const StyledTableCell = withStyles((theme) => ({
        head: {
            backgroundColor: theme.palette.primary.main,
            color: theme.palette.primary.contrastText,
            fontSize: 16,
        },
        body: {
            fontSize: 14,
        },
    }))(TableCell);

    return(
        <TableContainer component={Paper}>
            <Table>
                <TableHead>
                    <TableRow>
                        <StyledTableCell>ID</StyledTableCell>
                        <StyledTableCell>Start Time</StyledTableCell>
                        <StyledTableCell>Description</StyledTableCell>
                        <StyledTableCell>Address</StyledTableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {props.data.map(incident => {
                        const dt = new Date(incident.DateTime);
                        return(
                            <TableRow key={incident.eventNo}>
                                <StyledTableCell>
                                    <Link to={`/incident/${incident.eventNo}`}>{incident.eventNo}</Link>
                                </StyledTableCell>
                                <StyledTableCell>{`${dt.getMonth()+1}/${dt.getDate()}, ${dt.toLocaleTimeString()}`}</StyledTableCell>
                                <StyledTableCell>{incident.typeDescription}</StyledTableCell>
                                <StyledTableCell>{incident.address}</StyledTableCell>
                            </TableRow>
                        );
                    })}
                </TableBody>
            </Table>
        </TableContainer>
    );
}

function IsActiveFilter(props) {
    const [checked, setChecked] = React.useState(props.default);
    return(
        <Grid item style={{textAlign: "center"}}>
            <Typography variant="subtitle1">Active</Typography>
            <Switch
                checked={checked}
                onChange={() => { setChecked(!checked); props.update( (!checked) ? 1 : 0, "activesOnly") }}
                color="primary"
                name="checkedB"
                inputProps={{ 'aria-label': 'Show active incidents only' }}
            />
        </Grid>
    );
}

function DateRangeFilter(props) {
    const [startDate, setStartDate] = React.useState(props.default.start);
    const [startFocused, setStartFocused] = React.useState();

    const [endDate, setEndDate] = React.useState(props.default.end);
    const [endFocused, setEndFocused] = React.useState();

    const sdpProps = { numberOfMonths: 1, showDefaultInputIcon: true };
    const updateDate = function(value, isStart) {        
        if(isStart) {
            value.startOf("day");
            setStartDate(value);
            props.update(value, "dateRange.start");
        }
        else {
            value.endOf("day");
            setEndDate(value);
            props.update(value, "dateRange.end");
        }
    }

    return(
        <>
            <Grid item>
                <Typography variant="subtitle1">Start of Date Range</Typography>
                <SingleDatePicker {...sdpProps} isOutsideRange={(date) => (date > endDate)} date={startDate} onDateChange={(date) => updateDate(date, true)} focused={startFocused} onFocusChange={({focused}) => setStartFocused(focused)} />
            </Grid>
            <Grid item>
                <Typography variant="subtitle1">End of Date Range</Typography>
                <SingleDatePicker {...sdpProps} isOutsideRange={(date) => (date < startDate)} date={endDate} onDateChange={(date) => updateDate(date, false)} focused={endFocused} onFocusChange={({focused}) => setEndFocused(focused)} />
            </Grid>
        </>
    );
}

export default function Explorer() {
    const [filters, setFilters] = React.useState({
        dateRange: {
            start: moment().startOf("day"),
            end: moment().endOf("day")
        },
        activesOnly: 0
    });
    const mergeFilters = (value, jsonPath) => {
        let updatedFilter = {};
        lodash.set(updatedFilter, jsonPath, value);
        setFilters(lodash.merge({}, filters, updatedFilter));
    }
    
    const [data, setData] = React.useState(null);
    useEffect(() => {
        const RFC_3339 = 'YYYY-MM-DDTHH:mm:ssZ';
        const filterParams = {
            dateRangeStart: filters.dateRange.start.format(RFC_3339),
            dateRangeEnd: filters.dateRange.end.format(RFC_3339),
            activesOnly: filters.activesOnly
        };
        axios.get("/api/incidents/search", { params: filterParams })
            .then(response => {
                setData(response.data);
            })
            .catch(error => {
                console.log(error);
                setData({ error: error });
            })
    }, [filters]);

    let results;
    if(data == null) {
        results = (<div style={{padding: "1rem"}}><Loading /></div>)
    } else if(data.error) {
        results = (
            <Paper style={{textAlign: "center", padding: ".6rem"}}>
                <ErrorSvg style={{marginTop: "1rem"}} width="35%" height="35%" />
                <Typography variant="h5">We couldn't load any events.</Typography>
                <Typography variant="h6">Change your filters or try again later.</Typography>
            </Paper>
        );
    } else if(!data.length) {
        results = (
            <Paper>
                <Typography style={{textAlign: "center", padding: "1rem"}} variant="h6">No results found. Try changing your filters.</Typography>
            </Paper>
        );
    } else {
        results = (<IncidentDataTable data={data} />)
    }

    return(
        <div style={{marginTop: "3rem"}}>
            {/* Filters panel */}
            <Container>
                <Paper>
                    <AppBar position="static" elevation={0}>
                            <Typography variant="h6" style={{margin: ".6rem"}}>
                                Filters
                            </Typography>
                    </AppBar>
                    
                    <Grid container spacing={3} style={{margin: ".6rem"}}>
                        <DateRangeFilter update={mergeFilters} default={filters.dateRange} />       
                        <IsActiveFilter update={mergeFilters} default={filters.activesOnly} />             
                    </Grid>
                </Paper>
            </Container>

            {/* Results display */}
            <Container>
                {results}
            </Container>
        </div>
    );
}
