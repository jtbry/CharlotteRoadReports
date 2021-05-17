import React from 'react';
import { Paper, TableContainer, Table, TableCell, withStyles, TableHead, TableRow, TableBody, TablePagination } from '@material-ui/core';
import { Link } from 'react-router-dom';

export default function IncidentDataTable(props) {
    const [rowsPerPage, setRowsPerPage] = React.useState(10);
    const [page, setPage] = React.useState(0);
    const handleChangePage = (event, newPage) => {
        setPage(newPage);
      };
    
      const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
      };

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
        <Paper>
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
                    {props.data.slice(page * rowsPerPage, (page+1) * rowsPerPage).map(incident => {
                        const start_dt = new Date(incident.StartTimestamp);
                        return(
                            <TableRow key={incident.ID}>
                                <StyledTableCell>
                                    <Link to={`/incident/${incident.ID}`}>{incident.ID}</Link>
                                </StyledTableCell>
                                <StyledTableCell>{`${start_dt.getMonth()+1}/${start_dt.getDate()}, ${start_dt.toLocaleTimeString()}`}</StyledTableCell>
                                <StyledTableCell>{incident.TypeDesc}</StyledTableCell>
                                <StyledTableCell>{incident.Address}</StyledTableCell>
                            </TableRow>
                        );
                    })}
                </TableBody>
            </Table>
        </TableContainer>
        <TablePagination
            rowsPerPageOptions={[10, 25]}
            component="div"
            count={props.data.length}
            rowsPerPage={rowsPerPage}
            page={page}
            onChangeRowsPerPage={handleChangeRowsPerPage}
            onChangePage={handleChangePage}
        />
        </Paper>
    );
}
