import {Container} from 'react-bootstrap';
import './App.css';
import { ChakraProvider } from '@chakra-ui/react'

import CodeEditor from './CodeEditor';
import Home from './Home';
import JoinRoom from './JoinRoom';

function App() {
  return (
    <ChakraProvider>          
      <div className="App">
        <h1>
          codetgt
        </h1>
        <Home />
        <JoinRoom />
        <Container>
          <CodeEditor eid={1} />
        </Container>
      </div>
    </ChakraProvider>
  );
}

export default App;

