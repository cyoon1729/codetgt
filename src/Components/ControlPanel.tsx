import React from "react";
import { useState } from "react";
import {
  Container,
  FormControl,
  FormLabel,
  Code,
  Link,
  FormErrorMessage,
  VStack,
  Heading,
  Input,
  InputGroup,
  InputLeftAddon,
  InputRightElement,
  Button,
  useClipboard,
} from "@chakra-ui/react";
import {
  Select,
  CreatableSelect,
  AsyncSelect,
  OptionBase,
  GroupBase
} from "chakra-react-select";
import { CopyIcon } from "@chakra-ui/icons"

const languages = [
  { value: "Python", label: "Python"},
  { value: "Go", label: "Go"},
  { value: "C", label: "C"},
  { value: "Haskell", label: "Haskell"},
];

interface ProgrammingLanguage extends OptionBase {
  label: string;
  value: string;
}

interface ControlProps {
  roomId?: string;
}

const ControlPanel = (props: ControlProps) => {
    const roomId = props.roomId;
    const [selectedLang, setLang] = useState<ProgrammingLanguage>();
    const selectChange = (option: any) => {
      const value = option;
      setLang(value)
      console.log(value)
    };

    const {hasCopied, onCopy} = useClipboard("https://codetgt.io/" + roomId)
    
    return (
      <div>
        <VStack p={4} gap={4} align="left">
          <FormControl> 
            <FormLabel fontSize={14}>
              <Heading fontSize={14}> Shareable Link: </Heading> 
            </FormLabel>
            <InputGroup size='sm' borderRadius={1}>  
              <InputLeftAddon children='codetgt.io/' />
              <Input isReadOnly value={roomId} />
              <InputRightElement ml={2}>
                <Button onClick={onCopy} size="sm" borderRadius={1}>
                  <CopyIcon />
                </Button>
              </InputRightElement>
            </InputGroup>
          </FormControl>

          <FormControl>
            <FormLabel fontSize={14}>
                <Heading fontSize={14}> Set Language </Heading>
            </FormLabel>
            <Select<ProgrammingLanguage, true, GroupBase<ProgrammingLanguage>>
              name="languages"
              options={languages}
              defaultValue={languages[0]}
              selectedOptionStyle="check"
              closeMenuOnSelect={true}
              size="sm"
              onChange={(option) => {selectChange(option)}}
            />
          </FormControl>
        </VStack>
      </div>
    )
};

export default ControlPanel;
