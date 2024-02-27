// import React from 'react'
import { IoSearch } from "react-icons/io5";

const SearchBar = () => {
  return (
    <form className='w-full items-center justify-center flex flex-col'>
      <div className='relative flex flex-col items-center justify-center w-11/12 md:w-8/12 xl:w-1/2 mt-40 p-2 rounded-full bg-slate-800'>
        <input type="search" placeholder='Repository URL' className='block w-full ml-4 px-4 py-1 pl-10 rounded-full  bg-slate-800 text-white placeholder-gray-400 focus:outline-none '/>
        <button className='absolute right-0 top-0 mt-2 mr-2 p-2 items-center justify-center rounded-full bg-slate-600'>
        <IoSearch className="text-white"/>
        
        </button>
        </div>
    </form>
  )
}

export default SearchBar
