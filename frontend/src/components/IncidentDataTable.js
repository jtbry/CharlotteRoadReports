import React from 'react';
import { Paper, TableContainer, Table, TableCell, withStyles, TableHead, TableRow, TableBody } from '@material-ui/core';
import { Link } from 'react-router-dom';

export default function IncidentDataTable(props) {
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
    );
}
