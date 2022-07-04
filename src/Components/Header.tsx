import { Flex, HStack, Button, Stack, Heading, chakra, Divider } from '@chakra-ui/react';
import JoinRoomModal from './JoinRoomModal'; 

const Header = () => {
  return (
    <chakra.header id="header">
      <Flex
        w="100%"
        px="6"
        py="2"
        align="center"
        justify="space-between"
        boxShadow="md"
      >
        <HStack w="100%" align="center" justify="space-between">
          <Heading> codetgt {'>'} </Heading>
          <JoinRoomModal modalDefault={true}/>
        </HStack>
      </Flex>
      <Divider borderWidth="xl" />
    </chakra.header>
  );
}

export default Header;
