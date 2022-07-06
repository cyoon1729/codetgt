import { ChakraProvider } from '@chakra-ui/react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import CodeEditor from './CodeEditor';
import Room from './Room';
import Header from './Header';
function App() {
  return (
    <ChakraProvider>           
      <Router>            
        <div className="App">
          <Header />
        </div>
        <Routes>
          <Route
            path="/"
            element={<Room modalOpenDefault={true}/>} />     
          <Route
            path="/rooms/:roomId"
            element={<Room  modalOpenDefault={false}/>} />
        </Routes>
      </Router>
    </ChakraProvider>
  );
}

export default App;

