import { useParams } from 'react-router-dom';
import CodeEditor from './CodeEditor';
import Home from './Home';
import { Flex, HStack, Box, Container } from '@chakra-ui/react';
import ControlPanel from './ControlPanel';
import CommPanel from './CommPanel';

interface RoomProps {
    modalOpenDefault: boolean,
}

const Room = (props: RoomProps) => {
  const params = useParams();
  const roomId = params.roomId;
  const openModal = props.modalOpenDefault;
  
  return (
    <Flex className="Room" h="90vh" maxH="90vh" borderBottom={2}>
      <Flex
        w="full"
        h="90vh"
        boxShadow="md"
        py={2}
        gap={2}
        justify="center"
        direction="row"
        overflowY="hidden"
      >
        <Box w="25%" h="100%">
          <ControlPanel roomId={roomId}/>
        </Box>

        <CodeEditor eid={1} />

        <Box
          w="25%"
          h="100%"
          py={2}
          borderLeft="2px"
          borderLeftColor="#E1E1E1"
        >
          <CommPanel />
        </Box>
      </Flex>
    </Flex>      
  );
}

export default Room;

