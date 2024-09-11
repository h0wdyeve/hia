import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Benefits3 from './benefits3'; // Import the BookingFlight component
import Benefits from './benefits'; // Import the YourComponent component
import Benefits2 from './benefits2'; // Import the EditFlight component
// import Navigation from './navigation'; // Import the Navigation component

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Benefits/>} />
        <Route path="/Benefits2" element={<Benefits2/>} /> 
        <Route path="/Benefits3" element={<Benefits3/>} /> 
        </Routes>
    </Router>
  );
};

export default App;