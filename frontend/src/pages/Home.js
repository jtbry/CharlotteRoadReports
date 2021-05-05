import { Container, Grid, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography, withStyles } from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import { ReactComponent as ErrorSvg } from '../assets/undraw_error.svg';
import Loading from '../components/Loading';
import Axios from 'axios';
import { Bar } from 'react-chartjs-2';
import { Link } from 'react-router-dom';

function DivisionDistributionChart(props) {
    const divisionLabels = [];
    const divisionValues = [];
    for(let i = 0; i < props.data.length; ++i) {
        if(!divisionLabels.includes(props.data[i].division)) {
            divisionLabels.push(props.data[i].division);
            divisionValues.push(props.data.filter(incident => incident.division === props.data[i].division).length);
        }
    }

    const data = {
        labels: divisionLabels,
        datasets: [
            {
                label: "Active Incidents",
                backgroundColor: "#05331e",
                data: divisionValues
            }
        ]
    };

    const options = {
        responsive: true,
        plugins:{
            title: {
                display: true,
                text: "Active incident count by division"
            }
        },
        scales: {
            y: {
                grid: {
                    display: false
                },
                ticks: {
                    stepSize: 1,
                }
            }
        }
    };

    return(
        <Bar 
            data={data}
            options={options}
        />
    );
}

function TypeDistributionChart(props) {
    const typeLabels = [];
    const typeDescMap = {};
    const typeValues = [];
    for(let i = 0; i < props.data.length; ++i) {
        if(!typeLabels.includes(props.data[i].typeCode)) {
            typeLabels.push(props.data[i].typeCode);
            typeValues.push(props.data.filter(incident => incident.typeCode === props.data[i].typeCode).length);
            typeDescMap[props.data[i].typeCode] = props.data[i].typeDescription;
        }
    }

    const data = {
        labels: typeLabels,
        datasets: [
            {
                label: "Active Incidents",
                backgroundColor: "#05331e",
                data: typeValues
            }
        ]
    };

    const options = {
        responsive: true,
        plugins:{
            title: {
                display: true,
                text: "Active incident count by type"
            },
            tooltip: {
                callbacks: {
                    label: (ti, obj) => {
                        return `${typeDescMap[ti.label]}: ${ti.formattedValue}`;
                    }
                }
            }
        },
        scales: {
            y: {
                grid: {
                    display: false
                },
                ticks: {
                    stepSize: 1,
                }
            }
        }
    };

    return(
        <Bar 
            data={data}
            options={options}
        />
    );
}

function ActiveIncidentsTable(props) {
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
                        <StyledTableCell>Date Time</StyledTableCell>
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

export default function Home(props) {
    const [data, setData] = useState(null);
    useEffect(() => {
        const getActiveIncidents = () => {
            Axios.get('/api/incidents/active')
                .then((response) => {
                    setData( response.data )
                })
                .catch((error) => {
                    console.log(error);
                    setData({ error: true })
                });    
        }

        getActiveIncidents();
        const t = setTimeout(getActiveIncidents, 3000 * 60);
        return () => {
            clearInterval(t);
        }
    }, [])

    if(data === null) {
        return(<Loading />);
    } else if(data.error) {
        return(
            <Container style={{textAlign: "center", padding: "1rem"}}>
                <ErrorSvg style={{marginTop: "1rem"}} width="35%" height="35%" />
                <Typography variant="h4">Sorry!</Typography>
                <Typography variant="h5">We can't load this page.</Typography>
            </Container>
        );
    } else {
        return(
            <Container style={{marginTop: "3rem"}}>
                <Grid container spacing={3}>
                    <Grid item xs={12} sm={6}>
                        <DivisionDistributionChart data={data} />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <TypeDistributionChart data={data} />
                    </Grid>
                    <Grid item xs={12}>
                        <Typography variant="subtitle1" gutterBottom>
                            {data.length} active incidents
                        </Typography>
                        <ActiveIncidentsTable data={data} />
                    </Grid>
                </Grid>
            </Container>
        );
    }
}