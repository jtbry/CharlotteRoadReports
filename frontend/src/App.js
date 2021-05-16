import React from 'react';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import { CssBaseline, Tabs, Tab, AppBar, Container, withWidth, Typography } from '@material-ui/core';
import { BrowserRouter, Link, Switch, Route, useLocation } from 'react-router-dom'
import { ReactComponent as ErrorSvg } from './assets/undraw_error.svg';
import Home from './pages/Home';
import Incident from './pages/Incident';
import Map from './pages/Map';
import Explorer from './pages/Explorer';

const theme = createMuiTheme({
  palette: {
    primary: {
      light: '#325d45',
      main: '#05331e',
      dark: '#000f00',
      contrastText: '#ffffff',
    },
    secondary: {
      light: '#f5e19b',
      main: '#c1af6c',
      dark: '#8f803f',
      contrastText: '#000000',
    },
    background: {
        default: "#f7f9fc",
        paper: "#fff"
    }
  },
});


function Nav(props) {
    const location = useLocation().pathname
    const [value, setValue] = React.useState(location);
    const handleChange = (event, newValue) => {
        setValue(newValue);
    };
    
    const navVariant = props.width === "xs" ? "fullWidth" : "standard";
    return (
        <div>
            <AppBar elevation={1} position="relative">
                <Container>
                    <Tabs value={value} onChange={handleChange} variant={navVariant} centered>
                        <Tab wrapped component={Link} to="/" value="/" label="Active Incidents"/>
                        <Tab wrapped component={Link} to="/map" value="/map" label="Active Map"/>
                        <Tab wrapped component={Link} to="/explorer" value="/explorer" label="Incident Explorer"/>
                        {/* tab commented out while still in progress */}
                        {/* <Tab wrapped component={Link} to="/stats" value="/stats" label="Statistics Dashboard"/> */}
                    </Tabs>
                </Container>
            </AppBar>
        </div>
    );
}
const ResponsiveNav = withWidth()(Nav);

export default function App(props) {
    return(
        <BrowserRouter>
            <ThemeProvider theme={theme}>
                {/* Navbar */}
                <CssBaseline />
                <ResponsiveNav />

                {/* Page Content */}
                <Switch>
                    <Route exact path="/">
                        <Home />
                    </Route>
 
                    <Route path="/incident/:eventNo" component={Incident} />

                    <Route exact path="/map">
                        <Map />
                    </Route>

                    <Route exact path="/explorer">
                        <Explorer />
                    </Route>
 
                    <Route path="*">
                        <Container style={{textAlign: "center", padding: "1rem"}}>
                            <ErrorSvg style={{marginTop: "1rem"}} width="35%" height="35%" />
                            <Typography variant="h4">Oh no!</Typography>
                            <Typography variant="h5">That page isn't available.</Typography>
                            <Typography variant="h6">If you're lost, try going <a href="/">home</a>.</Typography>
                        </Container>
                    </Route>
                </Switch>

                {/* Footer, don't render on the map page */}
                <Switch>
                    <Route path="/map" />
                    <Route path="*">
                        <div style={{textAlign: 'center', margin: "1rem", overflow:"hidden"}}>
                            <Typography variant="subtitle1">CharlotteRoadReports, contribute on <a href="https://github.com/jtbry/CharlotteRoadReports" rel="noreferrer" target="_blank">Github</a></Typography>
                        </div>
                    </Route>
                </Switch>
            </ThemeProvider>
        </BrowserRouter>
    );
}