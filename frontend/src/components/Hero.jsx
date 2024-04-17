import { useContext, useState } from "react";
import { LuNetwork } from "react-icons/lu";
import { GlobalContext } from "../context";
import { Oval } from "react-loader-spinner";
import Resolve from "./Resolve";
import Toggles from "./Toggles";
export default function Hero() {
  const { jsonData, setJsonData } = useContext(GlobalContext);
  const [searchUrl, setsearchUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [showVersion, setShowVersion] = useState(false);
  const [showStdLib, setShowStdLib] = useState(true);
  const [addPath, setAddPath] = useState(false);
  const [path, setPath] = useState("");
  const [message, setmessage] = useState("");

  // handle repoURL
  const handleResolve = (event) => {
    if (searchUrl == "") {
      setmessage("This field should not be empty!");
      return;
    }
    event.preventDefault();
    setLoading(true);
    //TODO: change it to env variable
    fetch("http://localhost:8080/resolve", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        repoURL: searchUrl,
        options: {
          version: showVersion,
          stdlib: showStdLib,
          path: path,
          branch: "",
        },
      }),
    })
      .then((response) => {
        if (!response.ok) {
          // Returning a Promise.reject() to continue to the catch block
          return Promise.reject();
        }
        return response.json();
      })
      .then((data) => {
        setJsonData(data);
        setmessage("");
        setLoading(false);
      })
      .catch((error) => {
        setLoading(false);
        setmessage("Provided URL/path is not valid!");
      });
    setJsonData("");
  };

  //handle displaying version
  const enableVersion = () => {
    showVersion ? setShowVersion(false) : setShowVersion(true);
  };

  //handle displaying standard libraries
  const enableStdlib = () => {
    showStdLib ? setShowStdLib(false) : setShowStdLib(true);
  };

  //handle adding path
  const enablePath = () => {
    addPath ? setAddPath(false) : setAddPath(true);
  };

  return (
    <div className="font-customFont">
      <div className="text-4xl flex justify-center h-full mt-28">
        <LuNetwork className="text-md mx-1.5 my-0.5" />
        DepthSearch
      </div>
      <form className="w-full items-center justify-center flex flex-col">
        <div className="flex flex-row mt-16 w-1/3">
          <input
            id="search-term"
            type="text"
            placeholder="Your repository URL"
            className="w-full px-2 py-2 border-2 rounded-md flex flex-row items-center border-solid border-gray-200 mr-1"
            value={searchUrl}
            onChange={(event) => {
              setsearchUrl(event.target.value);
              setmessage("");
              setJsonData("");
            }}
          />
          <button
            type="button"
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold px-4 rounded ml-1"
            onClick={handleResolve}
          >
            Resolve
          </button>
        </div>
        <span className="mt-1.5 text-xs text-red-700 flex justify-start w-1/3">
          {message}
        </span>
        <div className="flex justify-center mt-4">
          <Toggles
            feature="Show Version"
            show={showVersion}
            enable={enableVersion}
          />
          <Toggles
            feature="Show Standard Libraries"
            show={!showStdLib}
            enable={enableStdlib}
          />
          <Toggles feature="Add Path" show={addPath} enable={enablePath} />
        </div>
        <div className="flex justify-center mt-4">
          {addPath ? (
            <input
              id="addPath"
              type="text"
              placeholder="Path to the package or main function"
              className="p-1 w-80 border-2 rounded-md flex flex-row items-center text-sm"
              value={path}
              onChange={(event) => setPath(event.target.value)}
            />
          ) : (
            ""
          )}
        </div>
        {loading && searchUrl && (
          <div className="flex flex-col items-center mt-14">
            <Oval
              visible={true}
              height="40"
              width="40"
              ariaLabel="magnifying-glass-loading"
              wrapperClass="magnifying-glass-wrapper"
              color="#000000"
              secondaryColor="#2563eb"
            />
            <p className="mt-2">Generating your project dependencies...</p>
          </div>
        )}
      </form>
      <div className="flex justify-center">
        {jsonData && (
          <Resolve
            jsonData={jsonData}
            searchUrl={searchUrl}
            showVersion={showVersion}
          />
        )}
      </div>
    </div>
  );
}
