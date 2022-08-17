import { Character } from "./App";
import { styledTitleBackground } from "./Style";

const ChoiceNext: React.FC<{
  character: Character;
  onChoiceNext: (next: string) => void;
}> = ({ onChoiceNext }) => {
  return (
    <div style={styledTitleBackground}>
      <div>
        <div style={{ cursor: "pointer" }} onClick={() => onChoiceNext("A")}>
          A
        </div>
        <div style={{ cursor: "pointer" }} onClick={() => onChoiceNext("B")}>
          B
        </div>
        <div style={{ cursor: "pointer" }} onClick={() => onChoiceNext("C")}>
          C
        </div>
      </div>
      <div>どれを選ぶ？</div>
    </div>
  );
};

export default ChoiceNext;
