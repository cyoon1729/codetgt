import {Container} from 'react-bootstrap';
import './App.css';
import { ChakraProvider } from '@chakra-ui/react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import CodeEditor from './CodeEditor';
import Home from './Home';
import JoinRoom from './JoinRoom';
import Room from './Room';

function App() {
  return (
    <ChakraProvider>      
      <Router>                  
        <Routes>
          <Route path="/" element={<Room />} />     
          <Route path="/rooms/:roomId" element={<Room />} />
        </Routes>
      </Router>
    </ChakraProvider>
  );
}

export default App;

