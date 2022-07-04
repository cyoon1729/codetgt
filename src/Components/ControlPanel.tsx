import React from "react";
import {
  Container,
  FormControl,
  FormLabel,
  Code,
  Link,
  FormErrorMessage,
  VStack,
} from "@chakra-ui/react";
import {
  Select,
  CreatableSelect,
  AsyncSelect,
  OptionBase,
  GroupBase
} from "chakra-react-select";

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

const ControlPanel = () => {
    return (
      <div>
        <FormControl p={4}>
          <FormLabel fontSize={16}>
            Set Language
          </FormLabel>
          <Select<ProgrammingLanguage, true, GroupBase<ProgrammingLanguage>>
            name="languages"
            options={languages}
            selectedOptionStyle="check"
            closeMenuOnSelect={true}
            size="sm"
          />
        </FormControl>
      </div>
    )
};

export default ControlPanel;
