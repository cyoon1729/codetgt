import React, { useEffect, useState, useRef } from 'react';
import CodeMirror from 'codemirror';
import 'codemirror/lib/codemirror.css';
import * as CodeMirrorCollabExt from '@convergencelabs/codemirror-collab-ext';
import { Flex, Box, Textarea, Text } from '@chakra-ui/react';

const textAreaStyle = {
    width: "100%",
    maxWidth: "100%",
    display: 'flex',
    flexDirection: 'column'
};

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
        width: '100%',
        height: '100%',
      }
    );
    editor.setSize("100%", "100%");
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
    <Flex
      w="full"
      direction="row"
      overflowY="scroll"
      flexDirection="column"
    >
      <Textarea w="100%" ref={editorArea} />
    </Flex>
  );
};


export default CodeEditor;
