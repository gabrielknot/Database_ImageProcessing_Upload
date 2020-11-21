import React,{useState, useRef}  from 'react';
import Dropzone from 'react-dropzone'
import './styles/dropContainer.css'

export default function ImageUpload(props){
    const canvasRef = useRef()
    const currentDatabase = props.currentDatabase

    const getBase64Image=(img)=> {
        const canvas = <canvas ref={canvasRef} {...props}/>;
        canvas.width = img.width;
        canvas.height = img.height;
        const ctx = canvasRef.current.getContext("2d");
        ctx.drawImage(img, 0, 0);
        const dataURL = canvas.toDataURL("image/png");
        return dataURL.replace(/^data:image\/(png|jpg);base64,/, "");
      }
    
    
    
    return (
        <div className="imageUpload">
            <Dropzone accept="images/*" onDropAccepted={Element=>{
                props.handleAddImageInDataBase(currentDatabase,

                    getBase64Image( <img src={Element}/>))
                }
            }>
                {({ getRootProps, getInputProps, isDragActive, isDradReject })=>{
                    <div className={
                        "dropContainer"+ isDragActive ?
                        " isDragActive": isDradReject ?
                        " isDragReject":""
                    }
                        {... getRootProps()}
                        isDragActive={isDragActive}
                        isDradReject={isDradReject}>
                            <input { ... getInputProps()}>
                                <span className="dragMensage success"
                                type="success">
                                Your files acepted!
                                </span>:
                                    isDradReject?
                                <span className="dragMensage error"
                                        type="error">
                                    Your files aren't acepted, please try again with other type
                                </span>: 
                                    !isDradReject && !isDragActive?
                                <span className="dragMensage">
                                    Drag your files here!
                                </span>:''
                        </input>
                    </div>
                }}
            </Dropzone>     
        </div>
    );
    }

    