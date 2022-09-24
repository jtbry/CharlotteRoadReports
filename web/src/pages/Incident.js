import { Card, CardContent, Container, Grid, Typography, withWidth } from '@material-ui/core'
import React, { useEffect, useState } from 'react'
import Loading from '../components/Loading'
import { ReactComponent as ErrorSvg } from '../assets/undraw_error.svg'
import axios from 'axios'
import { Marker, MapContainer, TileLayer, Popup } from 'react-leaflet'
import Rotate90DegreesCcwIcon from '@material-ui/icons/Rotate90DegreesCcw'

function IncidentInfoCard (props) {
  const incident = props.incident
  return (
    <Card variant='outlined'>
      <CardContent>
        <Grid container spacing={3}>
          <Grid item xs={4}>
            <Typography variant='h6'>Event No.</Typography>
            <Typography variant='subtitle1' gutterBottom>{incident.ID}</Typography>
          </Grid>
          <Grid item xs={4}>
            <Typography variant='h6'>Type Code</Typography>
            <Typography variant='subtitle1' gutterBottom>{incident.TypeCode}</Typography>
          </Grid>
          <Grid item xs={4}>
            <Typography variant='h6'>Division</Typography>
            <Typography variant='subtitle1' gutterBottom>{incident.Division}</Typography>
          </Grid>
          <Grid item xs={4}>
            <Typography variant='h6'>Description</Typography>
            <Typography variant='subtitle1' gutterBottom>{incident.TypeDesc}</Typography>
          </Grid>
          <Grid item xs={4}>
            <Typography variant='h6'>Address</Typography>
            <Typography variant='subtitle1' gutterBottom>{incident.Address}</Typography>
          </Grid>
          <Grid item xs={4}>
            <Typography variant='h6'>Lat, Lon</Typography>
            <Typography variant='subtitle1' gutterBottom>{incident.Latitude}<br />{incident.Longitude}</Typography>
          </Grid>
          <Grid item xs={4}>
            <Typography variant='h6'>Start Time</Typography>
            <Typography variant='subtitle1' gutterBottom>{new Date(incident.StartTimestamp).toLocaleString()}</Typography>
          </Grid>
          <Grid item xs={4}>
            <Typography variant='h6'>Status</Typography>
            <Typography variant='subtitle1' gutterBottom>{incident.Active ? 'Active' : 'Finished'}</Typography>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  )
}

function IncidentMap (props) {
  const incident = props.incident
  // todo: fix: the map does not appear on xs screen sizes but it does on sm
  if (props.width === 'xs') {
    return (
      <div style={{ textAlign: 'center' }}>
        <Rotate90DegreesCcwIcon />
        <Typography variant='h6'>
          Rotate your device to view the map.
        </Typography>
      </div>
    )
  } else {
    const mapHeight = props.width === 'sm' ? '300px' : '100%'
    return (
      <MapContainer style={{ height: mapHeight, borderRadius: '4px' }} center={[incident.Latitude, incident.Longitude]} zoom={15}>
        <TileLayer attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors' url='https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png' />
        <Marker position={[incident.Latitude, incident.Longitude]}>
          <Popup>
            {incident.ID}
          </Popup>
        </Marker>
      </MapContainer>
    )
  }
}
const ResponsiveIncidentMap = withWidth()(IncidentMap)

export default function Incident (props) {
  const eventNo = props.match.params.eventNo
  const [incident, setIncident] = useState(null)

  useEffect(() => {
    axios.get('/api/incidents/' + eventNo)
      .then(response => {
        setIncident(response.data)
      })
      .catch(error => {
        console.log(error)
        setIncident({ error: error })
      })
  }, [eventNo])

  if (incident === null) {
    return (<Loading pad />)
  } else if (incident.error) {
    return (
      <Container style={{ textAlign: 'center', padding: '1rem' }}>
        <ErrorSvg style={{ marginTop: '1rem' }} width='35%' height='35%' />
        <Typography variant='h5'>Are you sure that incident exists?</Typography>
        <Typography variant='h6'>If you're lost, try going <a href='/'>home</a>.</Typography>
      </Container>
    )
  } else {
    return (
      <Container style={{ marginTop: '3rem' }}>
        <Grid container spacing={3}>
          <Grid item sm={12} md={6}>
            <IncidentInfoCard incident={incident} />
          </Grid>

          <Grid item xs={12} sm={12} md={6}>
            <ResponsiveIncidentMap incident={incident} />
          </Grid>
        </Grid>
      </Container>
    )
  }
}
