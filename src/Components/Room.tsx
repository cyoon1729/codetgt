import { useParams } from 'react-router-dom';
import CodeEditor from './CodeEditor';
import Home from './Home';
import { Flex, HStack, Box, Container } from '@chakra-ui/react';
import ControlPanel from './ControlPanel';

interface RoomProps {
    modalOpenDefault: boolean,
}

const Room = (props: RoomProps) => {
  const params = useParams();
  const roomId = params.roomId;
  const openModal = props.modalOpenDefault;
  
  return (
    <div className="Room">
        <Flex
          width="full"
          height="90vh"
          boxShadow="md"
          py={2}
          gap={2}
          justify="center"
          direction="row"
        >
        <Box width="20%" height="90%">
            <ControlPanel />
        </Box>

        <CodeEditor eid={1} />

        <Box width="20%" backgroundColor="blue"> hi </Box>
      </Flex>
    </div>      
  );
}

export default Room;

