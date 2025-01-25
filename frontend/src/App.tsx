//import { useState } from 'react';
import Login from './pages/Login';
import Register from './pages/Register';
import Home from './pages/Home';
import Thread from './pages/Thread';
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { blue, orange } from "@mui/material/colors";
//import Thread from './components/Thread';

const theme = createTheme({
  palette: {
      primary: blue,
      secondary: orange,
  },
});

const App: React.FC = () => {
  return (
      <div className="App">
          <ThemeProvider theme={theme}>
              <BrowserRouter>
                  <Routes>
                      <Route path="/" element={<Login />} />
                      <Route path="/register" element={<Register />} />
                      <Route path="/home" element={<Home />} />
                      <Route path="/threads" element={<Thread />} />
                  </Routes>
              </BrowserRouter>
          </ThemeProvider>
      </div>
  );
};

export default App;
