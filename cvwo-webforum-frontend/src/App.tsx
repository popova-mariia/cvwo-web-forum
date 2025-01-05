//import { useState } from 'react';
import Login from './pages/Login';
import Register from './pages/Register';
import Home from './pages/Home';
import Replies from './pages/Replies';
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { blue, orange } from "@mui/material/colors";

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
                      <Route path="/" element={<Home />} />
                      <Route path="/thread/1" element={<Login />} />
                      <Route path="/thread/1/styled" element={<Register />} />
                      <Route path="/replies" element={<Replies />} />
                  </Routes>
              </BrowserRouter>
          </ThemeProvider>
      </div>
  );
};

export default App;

// const App = () => {
//   const [route, setRoute] = useState('/');

//   const renderPage = () => {
//     switch (route) {
//       case '/':
//         return <Login />;
//       case '/register':
//         return <Register />;
//       case '/dashboard':
//         return <Home />;
//       case '/:id/replies':
//         return <Replies />;
//       default:
//         return <Login />;
//     }
//   };

//   return (
//     <div>
//       <nav>
//         <button onClick={() => setRoute('/')}>Login</button>
//         <button onClick={() => setRoute('/register')}>Register</button>
//         <button onClick={() => setRoute('/dashboard')}>Home</button>
//         <button onClick={() => setRoute('/:id/replies')}>Replies</button>
//       </nav>
//       {renderPage()}
//     </div>
//   );
// };

// export default App;