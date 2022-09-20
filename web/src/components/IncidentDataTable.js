import React from 'react'
import { Paper, TableContainer, Table, TableCell, withStyles, TableHead, TableRow, TableBody, TablePagination, Button, TableFooter, Dialog, DialogTitle, DialogActions, DialogContent, FormLabel, RadioGroup, FormControlLabel, Radio, Grid } from '@material-ui/core'
import { Link } from 'react-router-dom'

function ExportDialog (props) {
  const page = props.exportInfo.page
  const rowsPerPage = props.exportInfo.rowsPerPage

  const [dataSelection, setDataSelection] = React.useState('current')
  const [exportFormat, setExportFormat] = React.useState('csv')

  const exportData = () => {
    let data
    if (dataSelection === 'all') data = props.exportInfo.data
    if (dataSelection === 'current') data = props.exportInfo.data.slice(page * rowsPerPage, (page + 1) * rowsPerPage)

    let outputUrl
    if (exportFormat === 'json') {
      const json = JSON.stringify(data, null, 2)
      const blob = new Blob([json], { type: 'octet/stream' })
      outputUrl = window.URL.createObjectURL(blob)
    }
    if (exportFormat === 'csv') {
      const replacer = (key, value) => value === null ? '' : value
      const header = Object.keys(data[0])
      const csv = [
        header.join(','), // header row first
        ...data.map(row => header.map(fieldName => JSON.stringify(row[fieldName], replacer)).join(','))
      ].join('\r\n')

      const blob = new Blob([csv], { type: 'octet/stream' })
      outputUrl = window.URL.createObjectURL(blob)
    }

    const link = document.createElement('a')
    link.style = 'display: none'
    link.href = outputUrl
    link.download = `cltrr_${Date.now()}.${exportFormat}`
    document.body.appendChild(link)
    link.click()
    window.URL.revokeObjectURL(outputUrl)
    props.handleCloseExport()
  }

  return (
    <Dialog fullWidth open={props.openExport} onClose={props.handleCloseExport} aria-labelledby='export-data-dialog'>
      <DialogTitle id='form-dialog-title'>Export Data</DialogTitle>
      <DialogContent style={{ overflowY: 'hidden' }}>
        <Grid container spacing={3}>
          <Grid item>
            <FormLabel component='legend'>Data to Export</FormLabel>
            <RadioGroup aria-label='dataExportSelection' value={dataSelection} onChange={(e) => setDataSelection(e.target.value)}>
              <FormControlLabel value='all' control={<Radio />} label='All Data' />
              <FormControlLabel value='current' control={<Radio />} label='Current Page' />
            </RadioGroup>
          </Grid>
          <Grid item>
            <FormLabel component='legend'>Format to Export</FormLabel>
            <RadioGroup aria-label='exportFormatValue' value={exportFormat} onChange={(e) => setExportFormat(e.target.value)}>
              <FormControlLabel value='csv' control={<Radio />} label='CSV' />
              <FormControlLabel value='json' control={<Radio />} label='JSON' />
            </RadioGroup>
          </Grid>
        </Grid>
        <p>Exporting {dataSelection === 'all'
          ? props.exportInfo.data.length
          : (rowsPerPage >= props.exportInfo.data.length ? props.exportInfo.data.length : rowsPerPage)} rows as {exportFormat.toUpperCase()}
        </p>
      </DialogContent>
      <DialogActions>
        <Button onClick={props.handleCloseExport} color='secondary'>
          Cancel
        </Button>
        <Button onClick={exportData} color='primary'>
          Export
        </Button>
      </DialogActions>
    </Dialog>
  )
}

export default function IncidentDataTable (props) {
  const [openExport, setOpenExport] = React.useState(false)

  const [rowsPerPage, setRowsPerPage] = React.useState(10)
  const [page, setPage] = React.useState(0)
  const handleChangePage = (event, newPage) => {
    setPage(newPage)
  }
  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(parseInt(event.target.value, 10))
    setPage(0)
  }
  if (page >= (props.data.length / rowsPerPage)) setPage(0)

  const StyledTableCell = withStyles((theme) => ({
    head: {
      backgroundColor: theme.palette.primary.main,
      color: theme.palette.primary.contrastText,
      fontSize: 16
    },
    body: {
      fontSize: 14
    }
  }))(TableCell)
  return (
    <Paper>
      <ExportDialog openExport={openExport} handleCloseExport={() => setOpenExport(false)} exportInfo={{ data: props.data, page: page, rowsPerPage: rowsPerPage }} />
      <TableContainer>
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
            {props.data.slice(page * rowsPerPage, (page + 1) * rowsPerPage).map(incident => {
              const startDt = new Date(incident.StartTimestamp)
              return (
                <TableRow key={incident.ID}>
                  <StyledTableCell>
                    <Link to={`/incident/${incident.ID}`}>{incident.ID}</Link>
                  </StyledTableCell>
                  <StyledTableCell>{`${startDt.getMonth() + 1}/${startDt.getDate()}, ${startDt.toLocaleTimeString()}`}</StyledTableCell>
                  <StyledTableCell>{incident.TypeDesc}</StyledTableCell>
                  <StyledTableCell>{incident.Address}</StyledTableCell>
                </TableRow>
              )
            })}
          </TableBody>
        </Table>
      </TableContainer>
      <TableFooter>
        <TableRow>
          <TableCell>
            <Button disableElevation variant='contained' color='primary' onClick={() => setOpenExport(true)}>Export Data</Button>
          </TableCell>

          <TableCell>
            <TablePagination
              rowsPerPageOptions={[10, 25]}
              component='div'
              count={props.data.length}
              rowsPerPage={rowsPerPage}
              page={page}
              onChangeRowsPerPage={handleChangeRowsPerPage}
              onChangePage={handleChangePage}
            />
          </TableCell>
        </TableRow>
      </TableFooter>
    </Paper>
  )
}
