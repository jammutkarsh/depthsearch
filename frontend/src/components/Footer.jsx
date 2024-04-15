import { VscGithubInverted } from "react-icons/vsc";

export default function Footer() {
  return (
    <div className="fixed bottom-0 left-0 w-full bg-slate-800 p-4 font-customFont">
      <div className="container mx-auto flex flex-row justify-between items-center">
        <div className="flex flex-col">
          <p className=" text-gray-200 text-base">
            <span>
              Dev(s):{" "}
              <a
                href="https://github.com/JammUtkarsh"
                target="_blank"
                className="m-1 text-blue-300 hover:text-blue-500"
              >
                Utkarsh
              </a>{" "}
              |
              <a
                href="https://github.com/Tejaswini-Tiwari"
                target="_blank"
                className="m-1 text-blue-300 hover:text-blue-500"
              >
                Tejaswini
              </a>{" "}
              |
              <a
                href="https://github.com/MadhaviGupta"
                target="_blank"
                className="m-1 text-blue-300 hover:text-blue-500"
              >
                Madhavi
              </a>
            </span>
          </p>
        </div>
        <a
          href="https://github.com/jammutkarsh/depthsearch"
          target="_blank"
          className="text-blue-300 hover:text-blue-500 flex"
        >
          <VscGithubInverted className="text-lg mx-1.5 my-0.5" />
          GitHub
        </a>
      </div>
    </div>
  );
}
