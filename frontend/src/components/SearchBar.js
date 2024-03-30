import { useState } from "react";

const SearchBar = () => {
  const [searchTerm1, setSearchTerm1] = useState("");


  const handleSearchButtonClick = () => {
    // Add your logic here to generate dependency graph based on searchTerm2
    console.log("Generating dependency graph for:", searchTerm1);
  };

  return (
    <form className="w-full items-center justify-center flex flex-col">
      
      <div className="flex flex-row mt-44">
        <div className="flex flex-row items-center border border-solid border-black rounded-md mr-4 mt-6 px-2 py-1 w-96">
          <label htmlFor="search-term-1" className="text-sm text-gray-600 px-2">
            Repository URI
          </label>
          <input
            id="search-term-1"
            type="text"
            className="w-full px-2 py-2 border-0 "
            value={searchTerm1}
            onChange={(event) => setSearchTerm1(event.target.value)}
          />
        </div>
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded ml-2 mt-6 h-15"
          onClick={handleSearchButtonClick}
        >
          Resolve
        </button>
      </div>

      <div className="mt-4 text-start">
        <p className="text-gray-600 text-sm sm:text-base mt-5">Why?</p>
        <p className="text-base text-gray-700 mt-2 sm:text-lg">
          Lorem Ipsum is simply dummy text of the printing and typesetting
          industry.
        </p>
      </div>

      <div className="fixed bottom-0 left-0 w-full bg-slate-800 p-4">
        <div className="container mx-auto flex flex-row justify-between items-center">
          <div className="flex flex-col">
            <p className=" text-gray-200 text-base">
              Dev(s): Madhavi, Tejaswini, Utkarsh
            </p>
          </div>
          <a href="#" className="text-blue-500 hover:text-blue-700 underline">
            GitHub
          </a>
        </div>
      </div>
    </form>
  );
};

export default SearchBar;
