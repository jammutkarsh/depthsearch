import { createContext, useState } from "react";

export const GlobalContext = createContext("");
export default function GlobalState({ children }) {
  const [jsonData, setJsonData] = useState("");
  return (
    <GlobalContext.Provider value={{ jsonData, setJsonData }}>
      {children}
    </GlobalContext.Provider>
  );
} 
