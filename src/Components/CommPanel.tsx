import {
  VStack,
  HStack,
  Box,
  Circle,
  Heading,
  List,
  ListItem,
  Flex,
} from "@chakra-ui/react"
import Chatbox from "./Chatbox";

const CommPanel = () => {
  const users = [
    {name: "chris", color: "blue"},
    {name: "chris", color: "blue"},
    {name: "chris", color: "blue"},
  ];

  const listUsers = users.map(usr => (
    <ListItem>
      <HStack width="full" justify="space-between" direction="column"> 
        <Box minWidth="60">
          {usr.name}
        </Box>
        <Flex width="20"></Flex>
        <Box>
          <Circle size="8px" bg={usr.color} color="blue" />
        </Box>
      </HStack>
    </ListItem> 
  ));
  

  return (
      <VStack w="full" h="100%" p={4} gap={4} align="left">
        <Flex w="full" h={60} align="left" >
          <Box w="full" h="full">
            <Heading fontSize={14}> Actice Users </Heading>      
              <List py={2} px={4} spacing={2}>
                {listUsers}     
              </List>
          </Box>
        </Flex>
        <Flex h="full">
          <Chatbox />
        </Flex>
      </VStack>
  );
};

export default CommPanel;
