import React from "react";
import AttractionTemplate from "./Attraction";
import ChoiceNext from "./ChoiceNext";
import Title from "./Title";

export interface Character {
  cid: string;
  name: string;
}

export interface Attraction {
  name: string;
}

const App: React.FC<{}> = () => {
  const [character, setCharacter] = React.useState<Character | null>(null);
  const onCharacterSelect = (c: Character) => {
    setCharacter(c);
  };

  const [attraction, setAttraction] = React.useState<Attraction | null>(null);

  switch (character) {
    case null:
      return <Title onCharacterSelect={onCharacterSelect} />;
    default:
      return attraction === null ? (
        <ChoiceNext
          character={character}
          onChoiceNext={(next: string) => {
            setAttraction({
              name: next,
            });
          }}
        />
      ) : (
        <AttractionTemplate attraction={attraction} />
      );
  }
};

export default App;
