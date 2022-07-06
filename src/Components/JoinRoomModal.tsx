import React from 'react';
import {v4 as uuidv4} from 'uuid';
import { nanoid } from 'nanoid'

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
import { useNavigate } from "react-router-dom";


interface JoinRoomProps {
  modalDefault: boolean,
}

const JoinRoomModal = (props: JoinRoomProps) => {
  let navigate = useNavigate();
    
  const generateRoomId = () => {
    return nanoid(8);
  }

  const joinRoom = (roomId: string) => {
    navigate(`/rooms/${roomId}`);
  }

  const [modalIsOpen, setIsOpen] = React.useState(props.modalDefault);

  const openModal = () => {
    setIsOpen(true);
  }
  
  const closeModal = () => {
    setIsOpen(false);
  }
  
  const [roomId, setRoomId] = React.useState('')
    
  const handleRoomIdInput = (event: any) => {
      setRoomId(event.target.value)
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
          <ModalCloseButton />
          <ModalBody py={8} mt={8}>
            <VStack spacing={4} align="center">  
              <Box width="100%">
                  <Button
                    width="100%"
                    onClick={() => {
                      joinRoom(generateRoomId());
                      closeModal()
                    }}
                  >
                    Create New Room
                  </Button>
              </Box>
              <Box width="100%" fontSize="md" fontWeight="semibold">
                <Center> Or </Center>
              </Box>
              <Box width="100%">
                <InputGroup>
                  <Input
                    placeholder='Room ID'
                    value={roomId}
                    size='md'
                    width="80%"
                    onChange={handleRoomIdInput}
                  />
                  <Button
                    width="20%"
                    ml={2}
                    onClick={() => {
                      joinRoom(roomId)
                      closeModal()
                    }}
                  >
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

export default JoinRoomModal;
