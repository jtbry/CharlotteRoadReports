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
