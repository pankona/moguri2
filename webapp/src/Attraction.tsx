import React from "react";
import { Attraction } from "./App";
import { styledTitleBackground } from "./Style";

type uiModeType = "menu" | "prepare" | "action";

const AttractionTemplate: React.FC<{ attraction: Attraction }> = ({
  attraction,
}) => {
  const [uiMode, setUIMode] = React.useState<uiModeType>("menu");

  switch (uiMode) {
    case "menu":
      return (
        <div style={styledTitleBackground}>
          <div>{attraction.name}が現れた</div>
          <div
            style={{ cursor: "pointer" }}
            onClick={() => setUIMode("action")}
          >
            戦う
          </div>
          <div
            style={{ cursor: "pointer" }}
            onClick={() => setUIMode("prepare")}
          >
            準備する
          </div>
        </div>
      );
    case "prepare":
      return (
        <div style={styledTitleBackground}>
          <div>準備する</div>
          <div>(implement build UI)</div>
          <div style={{ cursor: "pointer" }} onClick={() => setUIMode("menu")}>
            戻る
          </div>
        </div>
      );
    case "action":
      return <div style={styledTitleBackground}>戦う</div>;
  }
};

export default AttractionTemplate;
