import React from "react";
import { SimpleTreeView } from "@mui/x-tree-view/SimpleTreeView";
import { TreeItem } from "@mui/x-tree-view/TreeItem";
const renderTree = (data, parentId, showVersion) => {
  if (!data) return null;
  return (
    <div>
      {Array.isArray(data.tree.deps) &&
        data.tree.deps.map((dep, index) => (
          <React.Fragment key={`${parentId}-${data.tree.name}-${index}`}>
            <TreeItem
              itemId={`${parentId}-${data.tree.name}-${index}`}
              label={
                dep.name +
                (showVersion && dep.version != ""
                  ? "@" + `${dep.version}`
                  : "")
              }
            >
              {dep.deps &&
                renderTree(
                  { tree: dep },
                  `${parentId}-${data.tree.name}-${index}`,
                  showVersion
                )}
            </TreeItem>
          </React.Fragment>
        ))}
    </div>
  );
};

function Resolve(props) {
  const urlArray = props.searchUrl.split("/");
  return (
    <div className="font-customFont mt-10 border-2 w-2/5 bg-gray-100 rounded-md p-2 mb-28">
      <div className="flex flex-col items-center">
        <div className="text-2xl m-4 border-b-2 border-black border-spacing-1">{`${
          urlArray[urlArray.length - 1]
        } : ${props.jsonData.tree.resolved ? "Resolved" : ""}`}</div>
        <div className="flex flex-col items-center mt-2 font-customFont">
          {props.jsonData && (
            <div>
              {props.jsonData.summary.internal > 0 && (
                <div>
                  Internal Dependencies: {props.jsonData.summary.internal}
                </div>
              )}
              {props.jsonData.summary.external > 0 && (
                <div>
                  External Dependencies: {props.jsonData.summary.external}
                </div>
              )}
              {props.jsonData.summary.testing > 0 && (
                <div>
                  Testing Dependencies: {props.jsonData.summary.testing}
                </div>
              )}
            </div>
          )}
        </div>
        <div className="flex justify-center mt-6 border-t p-3 border-black">
          <SimpleTreeView>
            {renderTree(props.jsonData, "root", props.showVersion)}
          </SimpleTreeView>
        </div>
      </div>
    </div>
  );
}

export default Resolve;
