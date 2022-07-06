import React from 'react';
import { useEffect, useRef } from 'react';
import {
  Flex,
  Box,
  VStack,  
  Text,
  Button,
  Input,
  InputGroup,
  InputRightElement,
  Heading,
  Divider,
} from '@chakra-ui/react';
import { ArrowUpIcon } from '@chakra-ui/icons';

interface Message {
  name: string;
  msg: string;
};

const messages = [
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
  {name: "chris", msg: "hello"},
];



const Chatbox = () => {  
  const msgs = messages.map(msg => (<Text> msg.msg </Text>));
  
  const messagesEndRef = useRef<HTMLDivElement>(null)
  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ block: 'end', behavior: 'smooth' });
  })
  
    const sendMessageEnter = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      console.log("!!");
    }
  };

  const sendMessageButton = (e: React.MouseEvent<HTMLButtonElement>) => {
    console.log("!!");
  };
  
  return (
    <VStack width="100%" spacing={0} align="left">
      <Heading fontSize={14} py={2}> Chat </Heading>
      <Divider paddingBottom={1} />
      <Flex paddingTop={1} maxHeight="70%" width="100%" overflowY="scroll" flexDirection="column">
          {msgs}
        <div ref={messagesEndRef} />
      </Flex>
      <Box width="100%">
        <InputGroup size='sm' borderRadius={1}>  
          <Input onKeyDown={sendMessageEnter} />
          <InputRightElement ml={2}>
          <Button onClick={sendMessageButton} size="sm" borderRadius={1}>
            <ArrowUpIcon />
          </Button>
          </InputRightElement>
        </InputGroup>
       </Box>
    </VStack>
  );
};

export default Chatbox;
