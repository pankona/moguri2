import React from "react";
import { Character } from "./App";
import { styledTitleBackground } from "./Style";

type titleType = "initial" | "menu" | "continue";

const Title: React.FC<{
  onCharacterSelect: (character: Character) => void;
}> = ({ onCharacterSelect }) => {
  const [menu, setMenu] = React.useState<titleType>("initial");

  switch (menu) {
    case "initial":
      return (
        <div
          style={{ ...styledTitleBackground, cursor: "pointer" }}
          onClick={() => {
            setMenu("menu");
          }}
        >
          Click to start
        </div>
      );

    case "menu":
      return (
        <div style={styledTitleBackground}>
          <div
            style={{ cursor: "pointer" }}
            onClick={() => {
              /*not implemented yet*/
            }}
          >
            New Character
          </div>
          <div
            style={{ cursor: "pointer" }}
            onClick={() => {
              setMenu("continue");
            }}
          >
            Continue
          </div>
          <div
            style={{ cursor: "pointer" }}
            onClick={() => {
              setMenu("initial");
            }}
          >
            Back to title
          </div>
        </div>
      );

    case "continue":
      return (
        <div style={styledTitleBackground}>
          <div
            style={{ cursor: "pointer" }}
            onClick={() => {
              onCharacterSelect({
                cid: "test_character_1",
                name: "test character (1)",
              });
            }}
          >
            test character (1)
          </div>
          <div
            style={{ cursor: "pointer" }}
            onClick={() => {
              setMenu("menu");
            }}
          >
            Back
          </div>
        </div>
      );
  }
};

export default Title;
