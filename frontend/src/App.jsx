import UploadComponent from "./components/UploadComponent";
import TableComponent from "./components/TableComponent";
import React, { useState } from "react";

function App() {
  const [resultData, setResultData] = useState(undefined);

  return (
    <>
      <UploadComponent setResultData={setResultData} />
      <TableComponent resultData={resultData} />
    </>
  );
}

export default App;
