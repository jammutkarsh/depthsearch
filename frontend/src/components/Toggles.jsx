import Switch from "@mui/material/Switch";
export default function Toggles(props) {
  return (
    <div
      className={`text-sm rounded-md border-2 p-1 m-1 ${
        props.show ? "bg-gray-200" : "border-gray-200"
      }`}
    >
      <span className="p-1">{props.feature}</span>
      <Switch
        checked={props.show}
        onChange={() => {
          props.enable();
        }}
        className="inline-flex h-6 w-11 items-center rounded-full text-sm"
      />
    </div>
  );
}
