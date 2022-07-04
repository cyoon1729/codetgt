import React, { useEffect, useState, useRef } from 'react';
import CodeMirror from 'codemirror';
import 'codemirror/lib/codemirror.css';
import * as CodeMirrorCollabExt from '@convergencelabs/codemirror-collab-ext';


const CodeEditor = (props) => {
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
   
    editor.on('change', (instance, { origin }) => {
      if (origin !== 'setValue') {
        console.log(instance.getValue());
      }
    });
   
   const contentManager = new CodeMirrorCollabExt.EditorContentManager({
      editor: editor,
      onInsert(index, text) {
        console.log("Insert", index, text);
				ws.current.send(`editor: ${editorId} | insert!`);
      },
      onReplace(index, length, text) {
        console.log("Replace", index, length, text);
				ws.current.send(`editor: ${editorId} | replace!`);
      },
      onDelete(index, length) {
        console.log("Delete", index, length);
				ws.current.send(`editor: ${editorId} | delete!`);
      }
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


export default CodeEditor;
