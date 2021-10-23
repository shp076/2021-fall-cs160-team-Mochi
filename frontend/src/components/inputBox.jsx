import "../css/forms.css";

function InputBox({ placeholder, onChange, label, textArea }) {
  var labelStr = (<></>);
  if (label != null) {
    labelStr = (       
      <label className="agenda small">
        {label }&nbsp; 
      </label>
    )
  }
  if(textArea != null) {
    return (
      <div className="d-flex flex-row align-items-center">
      {labelStr}
        <textarea
          type="text"
          className="agenda text-input-box"
          placeholder={placeholder}
          onChange={(e) => onChange(e.target.value)}
          rows = "3"
        />
      </div>
    );
  } else {
  return (
    <div className="d-flex flex-row align-items-center">
    {labelStr}
      <input
        type="text"
        className="agenda text-input-box"
        placeholder={placeholder}
        onChange={(e) => onChange(e.target.value)}
      />
    </div>
  );
  }
}

export default InputBox;
