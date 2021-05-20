import { Container, FormControl, Grid, MenuItem, Paper, Select, Typography } from '@material-ui/core'
import React, { useEffect, useState } from 'react'
import moment from 'moment'
import 'react-dates/initialize'
import { DateRangePicker } from 'react-dates'
import 'react-dates/lib/css/_datepicker.css'
import '../assets/dates.css'
import axios from 'axios'
import Loading from '../components/Loading'
import { ReactComponent as ErrorSvg } from '../assets/undraw_error.svg'
import { Bar, Pie } from 'react-chartjs-2'

export default function Statistics () {
  const [startDate, setStartDate] = React.useState(moment())
  const [endDate, setEndDate] = React.useState(moment())
  const [focusDrp, setFocusDrp] = React.useState()
  const [data, setData] = useState(null)

  useEffect(() => {
    const RFC_3339 = 'YYYY-MM-DDTHH:mm:ssZ'
    const filter = {
      dateRangeStart: startDate.startOf('day').format(RFC_3339),
      dateRangeEnd: endDate.endOf('day').format(RFC_3339),
      activesOnly: false
    }
    axios.get('/api/incidents/search', { params: filter })
      .then(response => {
        setData(response.data)
      })
      .catch(error => {
        console.log(error)
        setData({ error: error })
      })
  }, [startDate, endDate])

  return (
    <Container>
      <Grid container spacing={2} style={{ marginTop: '1rem' }}>
        <Grid item xs={12}>
          <Typography variant='body1' color='textSecondary'>Use incidents from</Typography>
          <DateRangePicker
            startDate={startDate}
            startDateId='drp-start'
            endDate={endDate}
            endDateId='drp=end'
            numberOfMonths={1}
            showDefaultInputIcon
            minimumNights={0}
            isOutsideRange={(date) => date >= moment()}
            showClearDates
            onDatesChange={({ startDate, endDate }) => {
              // todo: there is an issue with setting endDate to null and causing an error
              // on use effect
              if (startDate) setStartDate(startDate)
              if (endDate) setEndDate(endDate)
              if (startDate == null && endDate == null) {
                // Allow clear dates functionality
                setStartDate(moment())
                setEndDate(moment())
              }
            }}
            focusedInput={focusDrp}
            onFocusChange={(focusedInput) => setFocusDrp(focusedInput)}
          />
        </Grid>

        <DisplayCards data={data} drstart={startDate} drend={endDate} />
      </Grid>
    </Container>
  )
}

function DisplayCards (props) {
  if (props.data === null) {
    return (<Grid item xs={12}><Loading pad /></Grid>)
  } else if (props.data.error) {
    return (
      <Grid item xs={12} style={{ textAlign: 'center' }}>
        <ErrorSvg style={{ marginTop: '1rem' }} width='35%' height='35%' />
        <Typography variant='h5'>Sorry</Typography>
        <Typography variant='h6'>The dashboard can not be loaded currently.</Typography>
        <Typography variant='h6'>Change your data range or try again later.</Typography>
      </Grid>
    )
  } else if (!props.data.length) {
    return (
      <Grid item xs={12} style={{ textAlign: 'center' }}>
        <Typography variant='h6'>There's not enough incidents in this date range.</Typography>
        <Typography variant='h6'>Change your data range or try again later.</Typography>
      </Grid>
    )
  } else {
    // todo: add a map that shows bubbles that have a correlation to the number of incidents
    // in that vicinity
    // similar to how google did their maps for the covid statistic maps
    // or some other way that will display event area frequency without cluttering the map
    // with a large number of markers
    return (
      <>
        <DisplayStatCards {...props} />
        <DisplayGraphMaker {...props} />
      </>
    )
  }
}

function DisplayStatCards (props) {
  let intervalStr = ''
  const daysDuration = Math.ceil((props.drend - props.drstart) / (1000 * 3600 * 24))
  if (daysDuration <= 1) {
    intervalStr = 'hour'
  } else if (daysDuration <= 14) {
    intervalStr = 'day'
  } else if (daysDuration <= 40) {
    intervalStr = 'week'
  } else {
    intervalStr = 'month'
  }

  return (
    <>
      <StatCard>
        <Typography variant='subtitle1'>Total incidents</Typography>
        <Typography variant='h6'>{props.data.length}</Typography>
      </StatCard>

      <StatCard>
        <Typography variant='subtitle1'>Average incidents per {intervalStr}</Typography>
        <Typography variant='h6'>{GetAverageIncidentPerTimeInterval(props.data, intervalStr, props.drstart)}</Typography>
      </StatCard>

      <StatCard>
        <Typography variant='subtitle1'>Average incident duration</Typography>
        <Typography variant='h6'>{GetAverageIncidentDuration(props.data)}</Typography>
      </StatCard>

      <StatCard>
        <Typography variant='subtitle1'>Most common type</Typography>
        <Typography variant='h6'>{GetMostFrequentFieldValue(props.data, 'TypeCode')}</Typography>
      </StatCard>

      <StatCard>
        <Typography variant='subtitle1'>Most common sub code</Typography>
        <Typography variant='h6'>{GetMostFrequentFieldValue(props.data, 'SubCode')}</Typography>
      </StatCard>

      <StatCard>
        <Typography variant='subtitle1'>Most common division</Typography>
        <Typography variant='h6'>{GetMostFrequentFieldValue(props.data, 'Division')}</Typography>
      </StatCard>
    </>
  )
}

function StatCard (props) {
  return (
    <Grid item xs={12} sm={4}>
      <Paper style={{ padding: '.5rem', textAlign: 'center' }} elevation={2}>
        {props.children}
      </Paper>
    </Grid>
  )
}

// Calculate average incidents per a given time interval
function GetAverageIncidentPerTimeInterval (data, interval, startRange) {
  const intervalsAndOccurences = {}
  for (let i = 0; i < data.length; i++) {
    const start = new Date(data[i].StartTimestamp)
    const daysPassed = Math.abs(Math.ceil((startRange._d - start) / (1000 * 3600 * 24)))
    let intervalKey
    if (interval === 'hour') return Math.round(data.length / new Date().getHours())
    else if (interval === 'day') intervalKey = `${start.getMonth()}/${start.getDate()}`
    // todo: we may want to make these more accurate
    // if someone selects within a range where a large number months/weeks don't
    // have data - it will make the average appear inaccurate because it will show the
    // average only of those weeks/months that do have data
    else if (interval === 'week') intervalKey = Math.floor(daysPassed / 7)
    else if (interval === 'month') intervalKey = `${start.getFullYear()}/${start.getMonth()}`

    if (intervalsAndOccurences[intervalKey]) intervalsAndOccurences[intervalKey]++
    else intervalsAndOccurences[intervalKey] = 1
  }

  let total = 0
  let keys = 0
  for (const key of Object.keys(intervalsAndOccurences)) {
    total += intervalsAndOccurences[key]
    keys++
  }
  return Math.round(total / keys)
}

// Calculate average duration of all incidents
function GetAverageIncidentDuration (data) {
  let totalDuration = 0
  for (let i = 0; i < data.length; i++) {
    // Don't calculate incidents that don't have an end time
    if (data[i].EndTimestamp === '0000-12-31T19:00:00-05:00') continue
    totalDuration += Math.abs(new Date(data[i].EndTimestamp) - new Date(data[i].StartTimestamp)) / 36e5
  }
  const avgDuration = totalDuration / data.length
  if (avgDuration < 1) return `${Math.ceil(avgDuration * 60)} minutes`
  else if (Math.ceil(avgDuration) === 1) return '1 hour'
  else return `${Math.ceil(avgDuration)} hours`
}

// Determine which value occurs the most frequently for a given field
function GetMostFrequentFieldValue (data, field) {
  const fieldsAndOccurences = {}
  for (let i = 0; i < data.length; i++) {
    if (fieldsAndOccurences[data[i][field]]) {
      fieldsAndOccurences[data[i][field]]++
    } else {
      fieldsAndOccurences[data[i][field]] = 1
    }
  }

  let max = ['N/A', 0]
  for (const key of Object.keys(fieldsAndOccurences)) {
    if (fieldsAndOccurences[key] > max[1]) max = [key, fieldsAndOccurences[key]]
  }
  return max[0]
}

function DisplayGraphMaker (props) {
  // in the future I would like to make it so users can make their own custom graph
  // selecting what type it is (bar, pie, etc) and what the different values / axes will be
  // for now they can just select from a handful of presets
  const [graphMetric, setGraphMetric] = React.useState('time')
  const [graphType, setGraphType] = React.useState('bar')

  return (
    <>
      <Grid item xs={6}>
        <Typography gutterBottom variant='body1' color='textSecondary'>Select graph display</Typography>
        <FormControl variant='standard' style={{ minWidth: 120 }}>
          <Select value={graphMetric} onChange={(e) => setGraphMetric(e.target.value)}>
            <MenuItem value='time'>Time Distributon</MenuItem>
            <MenuItem value='TypeCode'>Type Distributon</MenuItem>
            <MenuItem value='SubCode'>Sub Code Distributon</MenuItem>
            <MenuItem value='Division'>Division Distributon</MenuItem>
          </Select>
        </FormControl>
      </Grid>
      <Grid item xs={6} style={{ textAlign: 'right' }}>
        <Typography gutterBottom variant='body1' color='textSecondary'>Select graph type</Typography>
        <FormControl variant='standard' style={{ minWidth: 120 }}>
          <Select value={graphType} onChange={(e) => setGraphType(e.target.value)}>
            <MenuItem value='bar'>Bar</MenuItem>
            <MenuItem value='pie'>Pie</MenuItem>
          </Select>
        </FormControl>
      </Grid>
      <Grid item xs={12}>
        <Paper style={{ padding: '1rem' }}>
          <div style={{ maxWidth: '75%', textAlign: 'center', margin: '0 auto' }}>
            <CreateGraph metric={graphMetric} type={graphType} data={props.data} />
          </div>
        </Paper>
      </Grid>
    </>
  )
}

function CreateGraph (props) {
  // todo: average duration distribution graph
  const data = props.data
  const dataLabels = []
  const dataValues = []
  if (props.metric === 'time') {
    const timesAndValues = {}
    for (let i = 0; i < data.length; i++) {
      const time = new Date(data[i].StartTimestamp)
      if (timesAndValues[time.getHours()]) timesAndValues[time.getHours()]++
      else timesAndValues[time.getHours()] = 1
    }

    for (const key of Object.keys(timesAndValues)) {
      const period = key < 12 ? 'AM' : 'PM'
      const hour = key === '0' ? 12 : (key > 12 ? key - 12 : key)
      dataLabels.push(`${hour}${period}`)
      dataValues.push(timesAndValues[key])
    }
  } else {
    for (let i = 0; i < data.length; i++) {
      const dataIdx = dataLabels.indexOf(data[i][props.metric])
      if (dataIdx === -1) {
        dataLabels.push(data[i][props.metric])
        dataValues.push(1)
      } else {
        dataValues[dataIdx]++
      }
    }
  }

  // Make graph
  const graphOptions = {
    responsive: true,
    plugins: {
      title: {
        display: true,
        text: `Incident count distributed by ${props.metric}`
      }
    },
    scales: {
      y: {
        grid: {
          display: false
        },
        ticks: {
          stepSize: 5
        }
      }
    }
  }

  const graphData = {
    labels: dataLabels,
    datasets: [{
      label: 'Incidents',
      backgroundColor: props.type === 'bar' ? '#05331E' : ['#82998F', '#113D29', '#031C11', '#000503', '#031A0F', '#05301D', '#768F83', '#05331E'],
      data: dataValues
    }]
  }

  if (props.type === 'bar') {
    return (
      <Bar
        data={graphData}
        options={graphOptions}
      />
    )
  } else if (props.type === 'pie') {
    return (
      <Pie
        data={graphData}
        options={graphOptions}
      />
    )
  } else {
    return (<p>Sorry, that's not a valid graph type</p>)
  }
}
