import { CircularProgress } from '@material-ui/core';
import React from 'react';

export default function Loading() {
    return(
        <div style={{textAlign: "center", marginTop: "18%"}}>
            <CircularProgress />
        </div>
    )
}