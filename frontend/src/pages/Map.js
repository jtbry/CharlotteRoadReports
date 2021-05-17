import React, { useEffect, useState } from 'react';
import { Marker, MapContainer, TileLayer, Popup, AttributionControl, LayersControl, LayerGroup } from 'react-leaflet';
import axios from 'axios';
import { Link } from 'react-router-dom';
import { Container, Typography } from '@material-ui/core'
import Loading from '../components/Loading';
import { ReactComponent as ErrorSvg } from '../assets/undraw_error.svg';
import { divIcon } from 'leaflet';

function randomColor(number) {
    const hue = number * 137.508; // use golden angle approximation
    return `hsl(${hue},50%,55%)`;
}

function generateLayersFromData(data) {
    let layers = [];
    for(let i = 0; i < data.length; i++) {
        let idx = layers.findIndex(layer => layer.name === data[i].TypeCode)
        if(idx === -1) {
            // Create new layer and icon
            const layerColor = randomColor(layers.length);
            const iconHtmlStyle = `
                background-color: ${layerColor};
                width: 2rem;
                height: 2rem;
                display: block;
                left: -1.5rem;
                top: -1.5rem;
                position: relative;
                border-radius: 3rem 3rem 0;
                transform: rotate(45deg);
                border: 1px solid #FFFFFF`;
            layers.push({
                name: data[i].TypeCode,
                color: layerColor,
                icon: divIcon({
                    className: '',
                    iconAnchor: [0, 24],
                    labelAnchor: [-6, 0],
                    popupAnchor: [0, -36],
                    html: `<span style="${iconHtmlStyle}"/>`
                }),
                data: [data[i]]
            })
        } else {
            // Push object to existing layer
            layers[idx].data.push(data[i])
        }
    }
    return layers;
}

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
        return(<Loading pad />);
    } else if(data.error) {
        return(
            <Container style={{textAlign: "center", padding: "1rem"}}>
                <ErrorSvg style={{marginTop: "1rem"}} width="35%" height="35%" />
                <Typography variant="h4">Sorry!</Typography>
                <Typography variant="h5">We can't load this page.</Typography>
            </Container>
        );
    } else {
        // Create an array for each type of incident
        let layers = generateLayersFromData(data);

        return(
            <div style={{overflow: "hidden"}}>
                <MapContainer style={{height: "95vh", overflow: "hidden"}} center={[35.227085, -80.843124]} zoom={11}>
                    <TileLayer attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors' url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
                    <AttributionControl position="topright" />
                    <LayersControl position="topright">
                        {layers.map(layer => {
                            const layerControlText = `<div style="display: -webkit-inline-box">
                                    ${layer.name}
                                    <div style="margin-left: .5rem; width: 1rem; height: 1rem; border: 1px solid rgba(0, 0, 0, .2); background-color: ${layer.color}"></div>
                                </div>`
                            return(<LayersControl.Overlay checked name={layerControlText} key={layer.name}>
                                <LayerGroup>
                                {layer.data.map(incident => {
                                    return(
                                        <Marker key={incident.ID} position={[incident.Latitude, incident.Longitude]} icon={layer.icon}>
                                            <Popup>
                                                <Link to={`/incident/${incident.ID}`}>{incident.ID}</Link>
                                                <br />
                                                {incident.TypeDesc}
                                                <br />
                                                {new Date(incident.StartTimestamp).toLocaleString()}
                                                <br />
                                                {incident.Division}
                                            </Popup>
                                        </Marker>
                                    )
                                })}
                                </LayerGroup>
                            </LayersControl.Overlay>);
                        })}
                    </LayersControl>
                </MapContainer>
            </div>
        );
    }
}
