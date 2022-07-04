import {Container} from 'react-bootstrap';
import './App.css';

import CodeEditor from './CodeEditor';

function App() {
  return (
    <div className="App">
      <h1>
        codetgt
      </h1>
      <Container>
        <CodeEditor eid={1} />
      </Container>
    </div>
  );
}

export default App;

