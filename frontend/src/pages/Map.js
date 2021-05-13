import React, { useEffect, useState } from 'react';
import { Marker, MapContainer, TileLayer, Popup, AttributionControl } from 'react-leaflet';
import axios from 'axios';
import { Link } from 'react-router-dom';
import { Container, Typography } from '@material-ui/core'
import Loading from '../components/Loading';
import { ReactComponent as ErrorSvg } from '../assets/undraw_error.svg';

export default function Map(props) {
    // ? should this poll every 3min like the home page?
    const [data, setData] = useState(null);
    useEffect(() => {
        axios.get('/api/incidents/active')
            .then((response) => {
                setData( response.data )
            })
            .catch((error) => {
                console.log(error);
                setData({ error: error })
            });
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
            <div style={{overflow: "hidden"}}>
                <MapContainer style={{height: "95vh", overflow: "hidden"}} center={[35.227085, -80.843124]} zoom={11}>
                    <TileLayer attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors' url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
                    <AttributionControl position="topright" />
                    {data.map(incident => {
                        return(
                            <Marker position={[incident.latitude, incident.longitude]}>
                                <Popup>
                                    <Link to={`/incident/${incident.eventNo}`}>{incident.eventNo}</Link>
                                    <br />
                                    {incident.typeDescription}
                                    <br />
                                    {new Date(incident.DateTime).toLocaleString()}
                                    <br />
                                    {incident.division}
                                </Popup>
                            </Marker>
                        )
                    })}
                </MapContainer>
            </div>
        );
    }
}
