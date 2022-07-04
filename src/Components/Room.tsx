import CodeEditor from './CodeEditor';
import Home from './Home';
import JoinRoom from './JoinRoom';

const Room = () => {
  return (
    <div className="App">
      <h1>
        codetgt
      </h1>
      <Home />
      <JoinRoom />
      <CodeEditor eid={1} />
    </div>      
  );
}

export default Room;

