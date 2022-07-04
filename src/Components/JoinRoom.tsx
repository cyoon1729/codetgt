import React from 'react';
import {v4 as uuidv4} from 'uuid';
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalCloseButton,
  Button,
  VStack,
  Box,
  Input,
  InputGroup,
  Center,
} from '@chakra-ui/react'

const generateRoomId = () => {
    let newRoomId = uuidv4();
    console.log(newRoomId);
}

const JoinRoom = () => {
  const [modalIsOpen, setIsOpen] = React.useState(true);

  const openModal = () => {
    setIsOpen(true);
  }
  
  const closeModal = () => {
    setIsOpen(false);
  }

  return (
    <div>
      <Button onClick={openModal}>Create or Join Room</Button>

      <Modal isOpen={modalIsOpen} onClose={closeModal} size='lg'>
        <ModalOverlay 
          opacity="0.5"
          filter="blur(20px)"
        />
        <ModalContent height="30%">
          <ModalHeader> 
            code-tgt {'>'} &nbsp; Create or Join New Room
          </ModalHeader>
          <ModalCloseButton />
          <ModalBody mt={8}>
            <VStack spacing={4} align="center">  
              <Box width="100%">
                <Button width="100%">
                  Create New Room      
                </Button>
              </Box>
              <Box width="100%" fontSize="md" fontWeight="semibold">
                <Center> Or </Center>
              </Box>
              <Box width="100%">
                <InputGroup>
                  <Input placeholder='Room ID' size='md' width="80%"/>
                  <Button width="20%" ml={2}>
                    Join
                  </Button>
                </InputGroup>
              </Box>
            </VStack>
          </ModalBody>
        </ModalContent>
      </Modal>
    </div>
  );
}

export default JoinRoom;
