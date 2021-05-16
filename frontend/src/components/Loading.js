import { CircularProgress } from '@material-ui/core';
import React from 'react';

export default function Loading(props) {
    return(
        <div style={{textAlign: "center", marginTop: (props.pad ? "18%" : "")}}>
            <CircularProgress />
        </div>
    )
}