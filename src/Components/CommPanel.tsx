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
    <div className="CommPanel">
      <VStack width="full" height="100%" p={4} gap={4} align="left" justify="space-between" direction="column">
        <Flex width="full" height="40%" align="left">
          <Box>
            <Heading fontSize={14}> Actice Users </Heading>      
              <List py={2} px={4} spacing={2}>
                {listUsers}     
              </List>
          </Box>
        </Flex>
        <Flex height="10%"> <Box> </Box> </Flex>
        <Flex height="50%">
          <Chatbox />
        </Flex>
      </VStack>
    </div>
  );
};

export default CommPanel;
