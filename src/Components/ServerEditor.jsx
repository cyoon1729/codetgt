import React, { useContext, useEffect, useState, useRef } from 'react';
import { useParams } from 'react-router';
import CodeMirror from 'codemirror';
import 'codemirror/lib/codemirror.css';
import 'codemirror/mode/python/python';
import * as CodeMirrorCollabExt from '@convergencelabs/codemirror-collab-ext';


const ServerEditor = (props) => {
  const [sessionParams, setSessionParams] = useState(null);
  const editorArea = useRef(null); 
  const ws = useRef(null);
	const editorId = props.eid;

	useEffect(() => {
		ws.current = new WebSocket("ws://localhost:3333/");
    ws.current.onopen = () => console.log("ws init");
		ws.current.onclose = () => console.log("ws close");
    
    const wsCurrent = ws.current;

		return () => wsCurrent.close();
	}, []);

  useEffect(() => {      
    const editor = CodeMirror.fromTextArea(
      editorArea.current,
      {
        lineNumbers: true,
        lineWrapping: true,
        scrollbarStyle: 'null',
        theme: 'default',
      }
    );
    editor.setSize("100%", "40vh");
    editor.setValue("// Hello World!");
   
    const contentManager = new CodeMirrorCollabExt.EditorContentManager({
      editor: editor
    }); 

    return () => {
      editor.toTextArea();
      contentManager.dispose();
    }
  }, [sessionParams]);

  return (
    <div>
      <textarea ref={editorArea} />
    </div>
  );
};


export default ServerEditor;
