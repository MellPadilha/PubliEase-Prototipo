import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import '@mantine/core/styles.css';
import "@mantine/charts";
import "@mantine/core";
import "@mantine/dates";
import "@mantine/dropzone";
import '@mantine/dropzone/styles.css';
import "@mantine/form";
import "@mantine/hooks";
import "@mantine/modals";
import "@mantine/nprogress";
import { createTheme, MantineProvider } from '@mantine/core';
const cors = require('cors')

const theme = createTheme({
  /** Put your mantine theme override here */
});


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <MantineProvider theme={theme}>
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </MantineProvider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
