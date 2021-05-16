import { Container, Grid, Typography } from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import { ReactComponent as ErrorSvg } from '../assets/undraw_error.svg';
import Loading from '../components/Loading';
import axios from 'axios';
import { Bar } from 'react-chartjs-2';
import IncidentDataTable from '../components/IncidentDataTable';

function DivisionDistributionChart(props) {
    const divisionLabels = [];
    const divisionValues = [];
    for(let i = 0; i < props.data.length; ++i) {
        if(!divisionLabels.includes(props.data[i].Division)) {
            divisionLabels.push(props.data[i].Division);
            divisionValues.push(props.data.filter(incident => incident.Division === props.data[i].Division).length);
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
        if(!typeLabels.includes(props.data[i].TypeCode)) {
            typeLabels.push(props.data[i].TypeCode);
            typeValues.push(props.data.filter(incident => incident.TypeCode === props.data[i].TypeCode).length);
            typeDescMap[props.data[i].TypeCode] = props.data[i].TypeDesc;
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

export default function Home(props) {
    const [data, setData] = useState(null);
    useEffect(() => {
        const getActiveIncidents = () => {
            axios.get('/api/incidents/active')
                .then((response) => {
                    setData( response.data )
                })
                .catch((error) => {
                    console.log(error);
                    setData({ error: error })
                });    
        }

        getActiveIncidents();
        const t = setTimeout(getActiveIncidents, 3000 * 60);
        return () => {
            clearInterval(t);
        }
    }, []);

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
                        <IncidentDataTable data={data} />
                    </Grid>
                </Grid>
            </Container>
        );
    }
}