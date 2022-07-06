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

const CommPanel = () => {
  const users = [
    {name: "chris", color: "blue"},
    {name: "chris", color: "blue"},
  ];

  const listUsers = users.map(usr => (
    <ListItem>
      <HStack justify="space-between">    
        <Box>
          {usr.name}
        </Box>
        <Box>
          <Circle size="8px" bg={usr.color} color="blue" />
        </Box>
      </HStack>
    </ListItem> 
  ));
  

  return (
    <div className="CommPanel">
      <VStack p={4} gap={4} align="left" position="relative">
        <Flex height="40%" minHeight="40%" maxHeight="40%">
          <Box>
          <Heading fontSize={14}> Current Users </Heading>      
            <List py={2} px={4} spacing={2}>
              {listUsers}     
            </List>
          </Box>
        </Flex>
        <Box background="blue">
          Hi
        </Box>
      </VStack>
    </div>
  );
};

export default CommPanel;
